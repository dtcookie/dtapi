package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// GetProcessesBuilder TODO: documentation
type GetProcessesBuilder struct {
	client *environment.APIClient
	opts   environment.GetProcessesOpts
}

// WithStartTimeStamp TODO: documentation
func (builder GetProcessesBuilder) WithStartTimeStamp(startTimeStamp int64) GetProcessesBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder GetProcessesBuilder) WithEndTimeDtamp(endTimeStamp int64) GetProcessesBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder GetProcessesBuilder) WithRelativeTime(relativeTime string) GetProcessesBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder GetProcessesBuilder) WithTags(tags []string) GetProcessesBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder GetProcessesBuilder) WithEntityIDs(entityIDs []string) GetProcessesBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// Fetch TODO: documentation
func (builder GetProcessesBuilder) Fetch() ([]environment.ProcessGroupInstance, error) {
	result, _, err := builder.client.TopologySmartscapeProcessApi.GetProcesses(nil, &builder.opts)
	return result, err
}
