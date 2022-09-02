package uda

import (
	"context"
	"fmt"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var _ executor = (*aliyun)(nil)

type aliyun struct {
	clients map[string]*alidns.Client
}

func (a *aliyun) resolve(
	ctx context.Context,
	domain string, rr string, value string,
	options *options,
) (result *Result, err error) {
	if record, getErr := a.get(ctx, domain, rr, options); nil != getErr {
		err = getErr
	} else if nil != record {
		err = a.update(ctx, record, value, options)
	} else {
		err = a.add(ctx, domain, rr, value, options)
	}

	return
}

func (a *aliyun) update(_ context.Context, record *Record, value string, options *options) (err error) {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.RecordId = record.Id
	req.RR = record.Subdomain
	req.Type = options.typ
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Id, options.secret.Key); nil != clientErr {
		err = clientErr
	} else if rsp, updateErr := client.UpdateDomainRecord(req); nil != updateErr {
		err = updateErr
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`更新域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *aliyun) add(_ context.Context, domain string, rr string, value string, options *options) (err error) {
	req := alidns.CreateAddDomainRecordRequest()
	req.DomainName = domain
	req.RR = rr
	req.Type = options.typ
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Id, options.secret.Key); nil != clientErr {
		err = clientErr
	} else if rsp, addErr := client.AddDomainRecord(req); nil != addErr {
		err = addErr
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`添加域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *aliyun) get(_ context.Context, domain string, rr string, options *options) (record *Record, err error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.DomainName = domain
	req.RRKeyWord = rr
	req.TypeKeyWord = options.typ

	if client, clientErr := a.getClient(options.secret.Id, options.secret.Key); nil != clientErr {
		err = clientErr
	} else if rsp, getErr := client.DescribeDomainRecords(req); nil != getErr {
		err = getErr
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`获取域名解析记录出错`, field.String(`domain`, fmt.Sprintf(`%s.%s`, domain, rr)))
	} else {
		for _, _record := range rsp.DomainRecords.Record {
			if domain == _record.DomainName && options.typ == _record.Type && rr == _record.RR {
				record = new(Record)
				record.Id = _record.RecordId
				record.Name = _record.DomainName
				record.Subdomain = _record.RR
				record.Value = _record.Value
			}
		}
	}

	return
}

func (a *aliyun) getClient(key string, secret string) (client *alidns.Client, err error) {
	if cacheClient, ok := a.clients[key]; !ok {
		a.clients[key], err = alidns.NewClientWithAccessKey("cn-hangzhou", key, secret)
	} else {
		client = cacheClient
	}

	return
}
