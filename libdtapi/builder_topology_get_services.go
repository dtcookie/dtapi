package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetServiceBuilder TODO: documentation
type GetServiceBuilder interface {
	// WithStartTimeStamp TODO: documentation
	WithStartTimeStamp(startTimeStamp int64) GetServiceBuilder
	// WithEndTimeDtamp TODO: documentation
	WithEndTimeDtamp(endTimeStamp int64) GetServiceBuilder
	// WithRelativeTime TODO: documentation
	WithRelativeTime(relativeTime string) GetServiceBuilder
	// WithTags TODO: documentation
	WithTags(tags []string) GetServiceBuilder
	// WithEntityIDs TODO: documentation
	WithEntityIDs(entityIDs []string) GetServiceBuilder
	// Get TODO: documentation
	Get() ([]dtapi.Service, error)
}

type getServiceBuilder struct {
	GetServiceBuilder
	client *dtapi.APIClient
	opts   dtapi.GetServicesOpts
}

func (builder *getServiceBuilder) WithStartTimeStamp(startTimeStamp int64) GetServiceBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

func (builder *getServiceBuilder) WithEndTimeDtamp(endTimeStamp int64) GetServiceBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

func (builder *getServiceBuilder) WithRelativeTime(relativeTime string) GetServiceBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

func (builder *getServiceBuilder) WithTags(tags []string) GetServiceBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

func (builder *getServiceBuilder) WithEntityIDs(entityIDs []string) GetServiceBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

func (builder *getServiceBuilder) Get() ([]dtapi.Service, error) {
	result, _, err := builder.client.TopologySmartscapeServiceApi.GetServices(nil, &builder.opts)
	return result, err
}
