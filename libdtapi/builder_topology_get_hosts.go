package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetHostsBuilder TODO: documentation
type GetHostsBuilder struct {
	client *dtapi.APIClient
	opts   dtapi.GetHostsOpts
}

// WithStartTimeStamp TODO: documentation
func (builder *GetHostsBuilder) WithStartTimeStamp(startTimeStamp int64) *GetHostsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

// WithEndTimeDtamp TODO: documentation
func (builder *GetHostsBuilder) WithEndTimeDtamp(endTimeStamp int64) *GetHostsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

// WithRelativeTime TODO: documentation
func (builder *GetHostsBuilder) WithRelativeTime(relativeTime string) *GetHostsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

// WithTags TODO: documentation
func (builder *GetHostsBuilder) WithTags(tags []string) *GetHostsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

// WithEntityIDs TODO: documentation
func (builder *GetHostsBuilder) WithEntityIDs(entityIDs []string) *GetHostsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

// WithShowMonitoringCandidates TODO: documentation
func (builder *GetHostsBuilder) WithShowMonitoringCandidates(showMonitoringCandidates bool) *GetHostsBuilder {
	builder.opts.ShowMonitoringCandidates = optional.NewBool(showMonitoringCandidates)
	return builder
}

// GET TODO: documentation
func (builder *GetHostsBuilder) GET() ([]dtapi.Host, error) {
	result, _, err := builder.client.TopologySmartscapeHostApi.GetHosts(nil, &builder.opts)
	return result, err
}