package dynatrace

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// UserSessionQLAPI TODO: documentation
type UserSessionQLAPI service

/*
DTAQLResultBuilder TODO: documentation
*/
type DTAQLResultBuilder struct {
	client *environment.APIClient
	query  string
	opts   environment.GetDTAQLResultAsTreeOpts
	f      func(context.Context, string, *environment.GetDTAQLResultAsTreeOpts) (environment.DtaqlResultAsTable, *http.Response, error)
}

/*
AsTable TODO: documentation
*/
func (api UserSessionQLAPI) AsTable(query string) DTAQLResultBuilder {
	return DTAQLResultBuilder{
		client: api.client,
		query:  query,
		f:      api.client.UserSessionQueryLanguageApi.GetDTAQLResultAsTree,
	}
}

/*
AsTree TODO: documentation
*/
func (api UserSessionQLAPI) AsTree(query string) DTAQLResultBuilder {
	return DTAQLResultBuilder{
		client: api.client,
		query:  query,
	}
}

/*
Fetch TODO: documentation
*/
func (builder DTAQLResultBuilder) Fetch() (environment.DtaqlResultAsTable, error) {
	result, _, err := builder.f(nil, builder.query, &builder.opts)
	return result, err
}

/*
WithStartTimestamp TODO: documentation
*/
func (builder DTAQLResultBuilder) WithStartTimestamp(startTimeStamp int64) DTAQLResultBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

/*
WithEndTimestamp TODO: documentation
*/
func (builder DTAQLResultBuilder) WithEndTimestamp(endTimeStamp int64) DTAQLResultBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}
