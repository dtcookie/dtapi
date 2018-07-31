package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// GetApplicationsBuilder TODO: documentation
type GetApplicationsBuilder struct {
	client *environment.APIClient
	opts   environment.GetApplicationsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder GetApplicationsBuilder) WithStartTimeStamp(startTimeStamp int64) GetApplicationsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder GetApplicationsBuilder) WithEndTimeDtamp(endTimeStamp int64) GetApplicationsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder GetApplicationsBuilder) WithRelativeTime(relativeTime string) GetApplicationsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder GetApplicationsBuilder) WithTags(tags []string) GetApplicationsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder GetApplicationsBuilder) WithEntityIDs(entityIDs []string) GetApplicationsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// Fetch TODO: documentation
func (builder GetApplicationsBuilder) Fetch() ([]environment.Application, error) {
	result, _, err := builder.client.TopologySmartscapeApplicationApi.GetApplications(nil, &builder.opts)
	return result, err
}
