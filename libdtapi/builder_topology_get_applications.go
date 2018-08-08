package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetApplicationsBuilder TODO: documentation
type GetApplicationsBuilder interface {
	// WithStartTimeStamp TODO: documentation
	WithStartTimeStamp(startTimeStamp int64) GetApplicationsBuilder
	// WithEndTimeDtamp TODO: documentation
	WithEndTimeDtamp(endTimeStamp int64) GetApplicationsBuilder
	// WithRelativeTime TODO: documentation
	WithRelativeTime(relativeTime string) GetApplicationsBuilder
	// WithTags TODO: documentation
	WithTags(tags []string) GetApplicationsBuilder
	// WithEntityIDs TODO: documentation
	WithEntityIDs(entityIDs []string) GetApplicationsBuilder
	// GET TODO: documentation
	Get() ([]dtapi.Application, error)
}

type getApplicationsBuilder struct {
	GetApplicationsBuilder
	client *dtapi.APIClient
	opts   dtapi.GetApplicationsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder *getApplicationsBuilder) WithStartTimeStamp(startTimeStamp int64) GetApplicationsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder *getApplicationsBuilder) WithEndTimeDtamp(endTimeStamp int64) GetApplicationsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder *getApplicationsBuilder) WithRelativeTime(relativeTime string) GetApplicationsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder *getApplicationsBuilder) WithTags(tags []string) GetApplicationsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder *getApplicationsBuilder) WithEntityIDs(entityIDs []string) GetApplicationsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// GET TODO: documentation
func (builder *getApplicationsBuilder) Get() ([]dtapi.Application, error) {
	result, _, err := builder.client.TopologySmartscapeApplicationApi.GetApplications(nil, &builder.opts)
	return result, err
}
