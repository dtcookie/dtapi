package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// GetHostsBuilder TODO: documentation
type GetHostsBuilder interface {
	// WithStartTimeStamp TODO: documentation
	WithStartTimeStamp(startTimeStamp int64) GetHostsBuilder
	// WithEndTimeDtamp TODO: documentation
	WithEndTimeDtamp(endTimeStamp int64) GetHostsBuilder
	// WithRelativeTime TODO: documentation
	WithRelativeTime(relativeTime string) GetHostsBuilder
	// WithTags TODO: documentation
	WithTags(tags []string) GetHostsBuilder
	// WithEntityIDs TODO: documentation
	WithEntityIDs(entityIDs []string) GetHostsBuilder
	// WithShowMonitoringCandidates TODO: documentation
	WithShowMonitoringCandidates(showMonitoringCandidates bool) GetHostsBuilder
	// Get TODO: documentation
	Get() ([]dtapi.Host, error)
}

type getHostsBuilder struct {
	GetHostsBuilder
	client *dtapi.APIClient
	opts   dtapi.GetHostsOpts
}

func (builder *getHostsBuilder) WithStartTimeStamp(startTimeStamp int64) GetHostsBuilder {
	builder.opts.StartTimestamp = optional.NewInt64(startTimeStamp)
	return builder
}

func (builder *getHostsBuilder) WithEndTimeDtamp(endTimeStamp int64) GetHostsBuilder {
	builder.opts.EndTimestamp = optional.NewInt64(endTimeStamp)
	return builder
}

func (builder *getHostsBuilder) WithRelativeTime(relativeTime string) GetHostsBuilder {
	builder.opts.RelativeTime = optional.NewString(relativeTime)
	return builder
}

func (builder *getHostsBuilder) WithTags(tags []string) GetHostsBuilder {
	builder.opts.Tag = optional.NewInterface(tags)
	return builder
}

func (builder *getHostsBuilder) WithEntityIDs(entityIDs []string) GetHostsBuilder {
	builder.opts.Entity = optional.NewInterface(entityIDs)
	return builder
}

func (builder *getHostsBuilder) WithShowMonitoringCandidates(showMonitoringCandidates bool) GetHostsBuilder {
	builder.opts.ShowMonitoringCandidates = optional.NewBool(showMonitoringCandidates)
	return builder
}

func (builder *getHostsBuilder) Get() ([]dtapi.Host, error) {
	result, _, err := builder.client.TopologySmartscapeHostApi.GetHosts(nil, &builder.opts)
	return result, err
}
