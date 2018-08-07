package dtapi

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// UserSessionQLAPI TODO: documentation
type UserSessionQLAPI envService

/*
DTAQLResultBuilder TODO: documentation
*/
type DTAQLResultBuilder struct {
	client *dtapi.APIClient
	query  string
	opts   dtapi.GetDTAQLResultAsTreeOpts
	f      func(context.Context, string, *dtapi.GetDTAQLResultAsTreeOpts) (dtapi.DtaqlResultAsTable, *http.Response, error)
}

/*
AsTable TODO: documentation
*/
func (api UserSessionQLAPI) AsTable(query string) *DTAQLResultBuilder {
	return &DTAQLResultBuilder{
		client: api.client,
		query:  query,
		f:      api.client.UserSessionQueryLanguageApi.GetDTAQLResultAsTree,
	}
}

/*
AsTree TODO: documentation
*/
func (api UserSessionQLAPI) AsTree(query string) *DTAQLResultBuilder {
	return &DTAQLResultBuilder{
		client: api.client,
		query:  query,
	}
}

/*
GET TODO: documentation
*/
func (builder *DTAQLResultBuilder) GET() (dtapi.DtaqlResultAsTable, error) {
	result, _, err := builder.f(nil, builder.query, &builder.opts)
	return result, err
}

/*
WithStartTimestamp TODO: documentation
*/
func (builder *DTAQLResultBuilder) WithStartTimestamp(startTimeStamp int64) *DTAQLResultBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

/*
WithEndTimestamp TODO: documentation
*/
func (builder *DTAQLResultBuilder) WithEndTimestamp(endTimeStamp int64) *DTAQLResultBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}
