package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// MaintenanceWindowAPI TODO: documentation
type MaintenanceWindowAPI envService

// MaintenanceWindowQueryBuilder TODO: documentation
type MaintenanceWindowQueryBuilder struct {
	client *dtapi.APIClient
	opts   dtapi.GetAllMaintenanceWindowConfigsOpts
}

func newMaintenanceWindowQueryBuilder(client *dtapi.APIClient) *MaintenanceWindowQueryBuilder {
	builder := MaintenanceWindowQueryBuilder{}
	builder.client = client
	return &builder
}

// GET TODO: documentation
func (builder *MaintenanceWindowQueryBuilder) GET() ([]dtapi.MaintenanceWindow, error) {
	result, _, err := builder.client.MaintenanceWindowApi.GetAllMaintenanceWindowConfigs(nil, &builder.opts)
	return result, err
}

// WithStartTimeStamp TODO: documentation
func (builder *MaintenanceWindowQueryBuilder) WithStartTimeStamp(timeStamp int64) *MaintenanceWindowQueryBuilder {
	builder.opts.From = optional.NewInt64(timeStamp)
	return builder
}

// WithEndTimeStamp TODO: documentation
func (builder *MaintenanceWindowQueryBuilder) WithEndTimeStamp(timeStamp int64) *MaintenanceWindowQueryBuilder {
	builder.opts.To = optional.NewInt64(timeStamp)
	return builder
}

// WithType TODO: documentation
func (builder *MaintenanceWindowQueryBuilder) WithType(windowType string) *MaintenanceWindowQueryBuilder {
	builder.opts.Type_ = optional.NewString(windowType)
	return builder
}

// BuildQuery TODO: documentation
func (api MaintenanceWindowAPI) BuildQuery() *MaintenanceWindowQueryBuilder {
	return newMaintenanceWindowQueryBuilder(api.client)
}

// ALL TODO: documentation
func (api MaintenanceWindowAPI) ALL() ([]dtapi.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetAllMaintenanceWindowConfigs(nil, &dtapi.GetAllMaintenanceWindowConfigsOpts{})
	return result, err
}

// GET TODO: documentation
func (api MaintenanceWindowAPI) GET(uid string) (dtapi.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetMaintenanceWindowConfig(nil, uid)
	return result, err
}

// DELETE TODO: documentation
func (api MaintenanceWindowAPI) DELETE(uid string) error {
	_, err := api.client.MaintenanceWindowApi.RemoveMaintenanceWindowConfig(nil, uid)
	return err
}

// STORE TODO: documentation
func (api MaintenanceWindowAPI) STORE(maintenanceWindow dtapi.MaintenanceWindow) error {
	_, err := api.client.MaintenanceWindowApi.StoreMaintenanceWindowConfig(nil, &dtapi.StoreMaintenanceWindowConfigOpts{
		MaintenanceWindow: optional.NewInterface(maintenanceWindow),
	})
	return err
}
