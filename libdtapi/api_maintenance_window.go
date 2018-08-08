package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// MaintenanceWindowAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// configuring Maintenance Windows
type MaintenanceWindowAPI envService

// MaintenanceWindowQueryBuilder allows for conveniently
// specifying optional arguments when querying for
// Maintenance Windows.
type MaintenanceWindowQueryBuilder interface {
	// Get retrieves the configuration of all currently configured
	// Maintenance Windows
	Get() ([]dtapi.MaintenanceWindow, error)
	// WithStartTimeStamp is the start timestamp of the inquiry timeframe, in UTC milliseconds.
	// If 0 or missing, the current time is used.
	WithStartTimeStamp(timeStamp int64) MaintenanceWindowQueryBuilder
	// WithEndTimeStamp is the end timestamp of the inquiry timeframe, in UTC milliseconds.
	// If 0 or missing, all maintenance windows beginning after the `from` timestamp will be returned.
	WithEndTimeStamp(timeStamp int64) MaintenanceWindowQueryBuilder
	// WithType is the type of the maintenance window to return.
	// If `Unknown` or missing, all maintenance windows are returned.
	WithType(windowType string) MaintenanceWindowQueryBuilder
}

// NewQuery produces a builder for a parameterizable query for Maintenance Windows
func (api MaintenanceWindowAPI) NewQuery() MaintenanceWindowQueryBuilder {
	return newMaintenanceWindowQueryBuilder(api.client)
}

// All retrieves the configurations of all currently configured Maintenance Windows
func (api MaintenanceWindowAPI) All() ([]dtapi.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetAllMaintenanceWindowConfigs(nil, &dtapi.GetAllMaintenanceWindowConfigsOpts{})
	return result, err
}

// Get retrieves the configuration of a specific Maintenance Window
func (api MaintenanceWindowAPI) Get(uid string) (dtapi.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetMaintenanceWindowConfig(nil, uid)
	return result, err
}

// Delete deletes a specified maintenance window.
// Deletion cannot be undone.
func (api MaintenanceWindowAPI) Delete(uid string) error {
	_, err := api.client.MaintenanceWindowApi.RemoveMaintenanceWindowConfig(nil, uid)
	return err
}

// Store creates a new or updates an existing maintenance window.
func (api MaintenanceWindowAPI) Store(maintenanceWindow dtapi.MaintenanceWindow) error {
	_, err := api.client.MaintenanceWindowApi.StoreMaintenanceWindowConfig(nil, &dtapi.StoreMaintenanceWindowConfigOpts{
		MaintenanceWindow: optional.NewInterface(maintenanceWindow),
	})
	return err
}

type maintenanceWindowQueryBuilder struct {
	MaintenanceWindowQueryBuilder
	client *dtapi.APIClient
	opts   dtapi.GetAllMaintenanceWindowConfigsOpts
}

func newMaintenanceWindowQueryBuilder(client *dtapi.APIClient) MaintenanceWindowQueryBuilder {
	builder := maintenanceWindowQueryBuilder{}
	builder.client = client
	return &builder
}

func (builder *maintenanceWindowQueryBuilder) Get() ([]dtapi.MaintenanceWindow, error) {
	result, _, err := builder.client.MaintenanceWindowApi.GetAllMaintenanceWindowConfigs(nil, &builder.opts)
	builder.opts = dtapi.GetAllMaintenanceWindowConfigsOpts{}
	return result, err
}

func (builder *maintenanceWindowQueryBuilder) WithStartTimeStamp(timeStamp int64) MaintenanceWindowQueryBuilder {
	builder.opts.From = optional.NewInt64(timeStamp)
	return builder
}

func (builder *maintenanceWindowQueryBuilder) WithEndTimeStamp(timeStamp int64) MaintenanceWindowQueryBuilder {
	builder.opts.To = optional.NewInt64(timeStamp)
	return builder
}

func (builder *maintenanceWindowQueryBuilder) WithType(windowType string) MaintenanceWindowQueryBuilder {
	builder.opts.Type_ = optional.NewString(windowType)
	return builder
}
