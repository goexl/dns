package dns

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

func newAliyun() *aliyun {
	return &aliyun{
		clients: make(map[string]*alidns.Client),
	}
}

func (a *aliyun) update(_ context.Context, record *Record, value string, options *options) (err error) {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.RecordId = record.Id
	req.RR = record.Subdomain
	req.Type = string(options.typ)
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, ue := client.UpdateDomainRecord(req); nil != ue {
		err = ue
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`更新域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *aliyun) add(_ context.Context, domain string, rr string, value string, options *options) (err error) {
	req := alidns.CreateAddDomainRecordRequest()
	req.DomainName = domain
	req.RR = rr
	req.Type = string(options.typ)
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, ae := client.AddDomainRecord(req); nil != ae {
		err = ae
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`添加域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *aliyun) get(_ context.Context, domain string, rr string, options *options) (record *Record, err error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.DomainName = domain
	req.RRKeyWord = rr
	req.TypeKeyWord = string(options.typ)

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, ge := client.DescribeDomainRecords(req); nil != ge {
		err = ge
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`获取域名解析记录出错`, field.String(`domain`, fmt.Sprintf(`%s.%s`, domain, rr)))
	} else {
		for _, _record := range rsp.DomainRecords.Record {
			if domain == _record.DomainName && string(options.typ) == _record.Type && rr == _record.RR {
				record = new(Record)
				record.Id = _record.RecordId
				record.Name = _record.DomainName
				record.Subdomain = _record.RR
				record.Type = _record.Type
				record.Value = _record.Value
			}
		}
	}

	return
}

func (a *aliyun) delete(_ context.Context, record *Record, options *options) (err error) {
	req := alidns.CreateDeleteDomainRecordRequest()
	req.RecordId = record.Id

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, de := client.DeleteDomainRecord(req); nil != de {
		err = de
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`删除域名解析记录出错`, field.String(`domain`, record.Final()))
	}

	return
}

func (a *aliyun) getClient(ak string, sk string) (client *alidns.Client, err error) {
	if cacheClient, ok := a.clients[ak]; !ok {
		if client, err = alidns.NewClientWithAccessKey("cn-hangzhou", ak, sk); nil != err {
			return
		}
		a.clients[ak] = client
	} else {
		client = cacheClient
	}

	return
}
