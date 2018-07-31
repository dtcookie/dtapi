package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// GetServiceBuilder TODO: documentation
type GetServiceBuilder struct {
	client *environment.APIClient
	opts   environment.GetServicesOpts
}

// WithStartTimeStamp TODO: documentation
func (builder GetServiceBuilder) WithStartTimeStamp(startTimeStamp int64) GetServiceBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder GetServiceBuilder) WithEndTimeDtamp(endTimeStamp int64) GetServiceBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder GetServiceBuilder) WithRelativeTime(relativeTime string) GetServiceBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder GetServiceBuilder) WithTags(tags []string) GetServiceBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder GetServiceBuilder) WithEntityIDs(entityIDs []string) GetServiceBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// Fetch TODO: documentation
func (builder GetServiceBuilder) Fetch() ([]environment.Service, error) {
	result, _, err := builder.client.TopologySmartscapeServiceApi.GetServices(nil, &builder.opts)
	return result, err
}
