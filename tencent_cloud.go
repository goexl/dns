package dns

import (
	"context"
	"strconv"

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
	req := dnspod.NewModifyRecordRequest()
	req.SubDomain = &record.Subdomain
	req.Value = &value
	typ := string(options.typ)
	req.RecordType = &typ
	ttl := uint64(options.ttl.Seconds())
	req.TTL = &ttl
	if req.RecordId, err = record.TencentCloudId(); nil != err {
		return
	}

	if client, ce := a.getClient(options.secret.Ak, options.secret.Sk); nil != ce {
		err = ce
	} else if _, ue := client.ModifyRecord(req); nil != ue {
		err = ue
	}

	return
}

func (a *tencentCloud) add(
	_ context.Context,
	domain string, subdomain string, value string,
	options *options,
) (err error) {
	req := dnspod.NewCreateRecordRequest()
	req.Domain = &domain
	req.SubDomain = &subdomain
	typ := string(options.typ)
	req.RecordType = &typ
	req.Value = &value
	ttl := uint64(options.ttl.Seconds())
	req.TTL = &ttl

	if client, ce := a.getClient(options.secret.Ak, options.secret.Sk); nil != ce {
		err = ce
	} else if _, ae := client.CreateRecord(req); nil != ae {
		err = ae
	}

	return
}

func (a *tencentCloud) get(
	_ context.Context,
	domain string, subdomain string,
	options *options,
) (record *Record, err error) {
	req := dnspod.NewDescribeRecordListRequest()
	req.Domain = &domain
	req.Subdomain = &subdomain
	typ := string(options.typ)
	req.RecordType = &typ

	if client, ce := a.getClient(options.secret.Ak, options.secret.Sk); nil != ce {
		err = ce
	} else if rsp, ge := client.DescribeRecordList(req); nil != ge {
		err = ge
	} else {
		for _, _record := range rsp.Response.RecordList {
			if domain == *_record.Name && options.typ == Type(*_record.Type) {
				record = new(Record)
				record.Id = strconv.FormatUint(*_record.RecordId, 10)
				record.Name = *_record.Name
				record.Subdomain = subdomain
				record.Type = *_record.Type
				record.Value = *_record.Value
			}
		}
	}

	return
}

func (a *tencentCloud) delete(_ context.Context, record *Record, options *options) (err error) {
	req := dnspod.NewDeleteRecordRequest()
	req.Domain = &record.Name
	if req.RecordId, err = record.TencentCloudId(); nil != err {
		return
	}

	if client, ce := a.getClient(options.secret.Ak, options.secret.Sk); nil != ce {
		err = ce
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
