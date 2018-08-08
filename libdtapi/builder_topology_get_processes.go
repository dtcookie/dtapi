package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetProcessesBuilder TODO: documentation
type GetProcessesBuilder interface {
	// WithStartTimeStamp TODO: documentation
	WithStartTimeStamp(startTimeStamp int64) GetProcessesBuilder
	// WithEndTimeDtamp TODO: documentation
	WithEndTimeDtamp(endTimeStamp int64) GetProcessesBuilder
	// WithRelativeTime TODO: documentation
	WithRelativeTime(relativeTime string) GetProcessesBuilder
	// WithTags TODO: documentation
	WithTags(tags []string) GetProcessesBuilder
	// WithEntityIDs TODO: documentation
	WithEntityIDs(entityIDs []string) GetProcessesBuilder
	// Get TODO: documentation
	Get() ([]dtapi.ProcessGroupInstance, error)
}

type getProcessesBuilder struct {
	GetProcessesBuilder
	client *dtapi.APIClient
	opts   dtapi.GetProcessesOpts
}

// WithStartTimeStamp TODO: documentation
func (builder *getProcessesBuilder) WithStartTimeStamp(startTimeStamp int64) GetProcessesBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder *getProcessesBuilder) WithEndTimeDtamp(endTimeStamp int64) GetProcessesBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder *getProcessesBuilder) WithRelativeTime(relativeTime string) GetProcessesBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder *getProcessesBuilder) WithTags(tags []string) GetProcessesBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder *getProcessesBuilder) WithEntityIDs(entityIDs []string) GetProcessesBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// GET TODO: documentation
func (builder *getProcessesBuilder) Get() ([]dtapi.ProcessGroupInstance, error) {
	result, _, err := builder.client.TopologySmartscapeProcessApi.GetProcesses(nil, &builder.opts)
	return result, err
}
