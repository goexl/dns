package dns

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

var _ executor = (*tencentCloud)(nil)

type tencentCloud struct {
	clients map[string]*dnspod.Client
}

func newTencentCloud() *tencentCloud {
	return &tencentCloud{
		clients: make(map[string]*dnspod.Client),
	}
}

func (a *tencentCloud) update(_ context.Context, record *Record, value string, options *options) (err error) {
	req := alidns.CreateUpdateDomainRecordRequest()
	req.RecordId = record.Id
	req.RR = record.Subdomain
	req.Type = string(options.typ)
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, updateErr := client.UpdateDomainRecord(req); nil != updateErr {
		err = updateErr
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`更新域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *tencentCloud) add(_ context.Context, domain string, rr string, value string, options *options) (err error) {
	req := alidns.CreateAddDomainRecordRequest()
	req.DomainName = domain
	req.RR = rr
	req.Type = string(options.typ)
	req.Value = value
	req.TTL = requests.NewInteger(int(options.ttl.Seconds()))

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, addErr := client.AddDomainRecord(req); nil != addErr {
		err = addErr
	} else if nil != rsp && !rsp.IsSuccess() {
		err = exc.NewFields(`添加域名解析记录出错`, field.String(`value`, value))
	}

	return
}

func (a *tencentCloud) get(_ context.Context, domain string, rr string, options *options) (record *Record, err error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.DomainName = domain
	req.RRKeyWord = rr
	req.TypeKeyWord = string(options.typ)

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if rsp, getErr := client.DescribeDomainRecords(req); nil != getErr {
		err = getErr
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

func (a *tencentCloud) delete(_ context.Context, record *Record, options *options) (err error) {
	req := dnspod.NewDeleteRecordRequest()
	req.Domain = &record.Name

	if id, parseErr := strconv.ParseUint(record.Id, 10, 64); nil != parseErr {
		err = parseErr
	} else {
		req.RecordId = &id
	}
	if nil != err {
		return
	}

	if client, clientErr := a.getClient(options.secret.Ak, options.secret.Sk); nil != clientErr {
		err = clientErr
	} else if _, de := client.DeleteRecord(req); nil != de {
		err = de
	}

	return
}

func (a *tencentCloud) getClient(ak string, sk string) (client *dnspod.Client, err error) {
	if cacheClient, ok := a.clients[ak]; !ok {
		credential := common.NewCredential(ak, sk)
		if client, err = dnspod.NewClient(credential, regions.Guangzhou, profile.NewClientProfile()); nil != err {
			return
		}
		a.clients[ak] = client
	} else {
		client = cacheClient
	}

	return
}
