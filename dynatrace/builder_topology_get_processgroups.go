package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// GetProcessGroupsBuilder TODO: documentation
type GetProcessGroupsBuilder struct {
	client *environment.APIClient
	opts   environment.GetProcessGroupsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder GetProcessGroupsBuilder) WithStartTimeStamp(startTimeStamp int64) GetProcessGroupsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder GetProcessGroupsBuilder) WithEndTimeDtamp(endTimeStamp int64) GetProcessGroupsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder GetProcessGroupsBuilder) WithRelativeTime(relativeTime string) GetProcessGroupsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder GetProcessGroupsBuilder) WithTags(tags []string) GetProcessGroupsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder GetProcessGroupsBuilder) WithEntityIDs(entityIDs []string) GetProcessGroupsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// Fetch TODO: documentation
func (builder GetProcessGroupsBuilder) Fetch() ([]environment.ProcessGroup, error) {
	result, _, err := builder.client.TopologySmartscapeProcessGroupApi.GetProcessGroups(nil, &builder.opts)
	return result, err
}
