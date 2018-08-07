package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetProcessGroupsBuilder TODO: documentation
type GetProcessGroupsBuilder struct {
	client *dtapi.APIClient
	opts   dtapi.GetProcessGroupsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder *GetProcessGroupsBuilder) WithStartTimeStamp(startTimeStamp int64) *GetProcessGroupsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder *GetProcessGroupsBuilder) WithEndTimeDtamp(endTimeStamp int64) *GetProcessGroupsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder *GetProcessGroupsBuilder) WithRelativeTime(relativeTime string) *GetProcessGroupsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder *GetProcessGroupsBuilder) WithTags(tags []string) *GetProcessGroupsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder *GetProcessGroupsBuilder) WithEntityIDs(entityIDs []string) *GetProcessGroupsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// GET TODO: documentation
func (builder *GetProcessGroupsBuilder) GET() ([]dtapi.ProcessGroup, error) {
	result, _, err := builder.client.TopologySmartscapeProcessGroupApi.GetProcessGroups(nil, &builder.opts)
	return result, err
}
