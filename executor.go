package dns

import (
	"context"
)

type executor interface {
	add(_ context.Context, domain string, rr string, value string, options *options) (err error)

	get(ctx context.Context, domain string, rr string, options *options) (record *Record, err error)

	update(_ context.Context, record *Record, value string, options *options) (err error)

	delete(_ context.Context, record *Record, options *options) (err error)
}
