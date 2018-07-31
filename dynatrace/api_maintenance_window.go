package dynatrace

import (
	"github.com/dtcookie/dtapi/environment"
)

// MaintenanceWindowAPI TODO: documentation
type MaintenanceWindowAPI service

// GetAllMaintenanceWindowConfigs TODO: documentation
func (api MaintenanceWindowAPI) GetAllMaintenanceWindowConfigs(opts *environment.GetAllMaintenanceWindowConfigsOpts) ([]environment.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetAllMaintenanceWindowConfigs(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetMaintenanceWindowConfig TODO: documentation
func (api MaintenanceWindowAPI) GetMaintenanceWindowConfig(uid string) (environment.MaintenanceWindow, error) {
	result, _, err := api.client.MaintenanceWindowApi.GetMaintenanceWindowConfig(nil, uid)
	if err != nil {
		return result, err
	}
	return result, nil
}

// RemoveMaintenanceWindowConfig TODO: documentation
func (api MaintenanceWindowAPI) RemoveMaintenanceWindowConfig(uid string) error {
	_, err := api.client.MaintenanceWindowApi.RemoveMaintenanceWindowConfig(nil, uid)
	if err != nil {
		return err
	}
	return nil
}

// StoreMaintenanceWindowConfig TODO: documentation
func (api MaintenanceWindowAPI) StoreMaintenanceWindowConfig(opts *environment.StoreMaintenanceWindowConfigOpts) error {
	_, err := api.client.MaintenanceWindowApi.StoreMaintenanceWindowConfig(nil, opts)
	if err != nil {
		return err
	}
	return nil
}
