package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetProcessGroupsBuilder TODO: documentation
type GetProcessGroupsBuilder interface {
	// WithStartTimeStamp TODO: documentation
	WithStartTimeStamp(startTimeStamp int64) GetProcessGroupsBuilder
	// WithEndTimeDtamp TODO: documentation
	WithEndTimeDtamp(endTimeStamp int64) GetProcessGroupsBuilder
	// WithRelativeTime TODO: documentation
	WithRelativeTime(relativeTime string) GetProcessGroupsBuilder
	// WithTags TODO: documentation
	WithTags(tags []string) GetProcessGroupsBuilder
	// WithEntityIDs TODO: documentation
	WithEntityIDs(entityIDs []string) GetProcessGroupsBuilder
	// Get TODO: documentation
	Get() ([]dtapi.ProcessGroup, error)
}

type getProcessGroupsBuilder struct {
	GetProcessGroupsBuilder
	client *dtapi.APIClient
	opts   dtapi.GetProcessGroupsOpts
}

func (builder *getProcessGroupsBuilder) WithStartTimeStamp(startTimeStamp int64) GetProcessGroupsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

func (builder *getProcessGroupsBuilder) WithEndTimeDtamp(endTimeStamp int64) GetProcessGroupsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

func (builder *getProcessGroupsBuilder) WithRelativeTime(relativeTime string) GetProcessGroupsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

func (builder *getProcessGroupsBuilder) WithTags(tags []string) GetProcessGroupsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

func (builder *getProcessGroupsBuilder) WithEntityIDs(entityIDs []string) GetProcessGroupsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

func (builder *getProcessGroupsBuilder) Get() ([]dtapi.ProcessGroup, error) {
	result, _, err := builder.client.TopologySmartscapeProcessGroupApi.GetProcessGroups(nil, &builder.opts)
	return result, err
}
