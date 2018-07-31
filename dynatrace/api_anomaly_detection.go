package dynatrace

import (
	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/config"
)

// AnomalyDetectionAPI TODO: documentation
type AnomalyDetectionAPI configService

// AwsAnomalyDetectionAPI TODO: documentation
type AwsAnomalyDetectionAPI configService

// DiskAnomalyDetectionAPI TODO: documentation
type DiskAnomalyDetectionAPI configService

// HostAnomalyDetectionAPI TODO: documentation
type HostAnomalyDetectionAPI configService

// ServiceAnomalyDetectionAPI TODO: documentation
type ServiceAnomalyDetectionAPI configService

// VMwareAnomalyDetectionAPI TODO: documentation
type VMwareAnomalyDetectionAPI configService

/*
Aws TODO: document
*/
func (api AnomalyDetectionAPI) Aws() AwsAnomalyDetectionAPI {
	return AwsAnomalyDetectionAPI{
		client: api.client,
	}
}

/*
VMware TODO: document
*/
func (api AnomalyDetectionAPI) VMware() VMwareAnomalyDetectionAPI {
	return VMwareAnomalyDetectionAPI{
		client: api.client,
	}
}

/*
Disks TODO: document
*/
func (api AnomalyDetectionAPI) Disks() DiskAnomalyDetectionAPI {
	return DiskAnomalyDetectionAPI{
		client: api.client,
	}
}

/*
Hosts TODO: document
*/
func (api AnomalyDetectionAPI) Hosts() HostAnomalyDetectionAPI {
	return HostAnomalyDetectionAPI{
		client: api.client,
	}
}

/*
Services TODO: document
*/
func (api AnomalyDetectionAPI) Services() ServiceAnomalyDetectionAPI {
	return ServiceAnomalyDetectionAPI{
		client: api.client,
	}
}

/*
Get Get configuration of anomaly detection for AWS.
@return AwsAnomalyDetectionConfig
*/
func (api AwsAnomalyDetectionAPI) Get() (config.AwsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig1(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Update configuration of anomaly detection for AWS.
 * @param optional nil or *UpdateAwsAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "AwsAnomalyDetectionConfig" (optional.Interface of AwsAnomalyDetectionConfig) -
*/
func (api AwsAnomalyDetectionAPI) Update(c config.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.UpdateAwsAnomalyDetectionConfig1(nil, &config.UpdateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Validate Validate configuration of anomaly detection for AWS (for PUT request).
 * @param optional nil or *ValidateAwsAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "AwsAnomalyDetectionConfig" (optional.Interface of AwsAnomalyDetectionConfig) -
*/
func (api AwsAnomalyDetectionAPI) Validate(c config.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.ValidateAwsAnomalyDetectionConfig1(nil, &config.ValidateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Create Create a new disk event rule.
 * @param optional nil or *CreateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
@return DiskEventAnomalyDetectionConfig
*/
func (api DiskAnomalyDetectionAPI) Create(c config.DiskEventAnomalyDetectionConfig) (config.DiskEventAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.CreateDiskEventConfig1(nil, &config.CreateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Delete Delete the disk alert rule with the given id.
 * @param id
*/
func (api DiskAnomalyDetectionAPI) Delete(id string) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.DeleteDiskEventConfig1(nil, id)
	return err
}

/*
Get Get the disk event rule with the given id.
 * @param id
@return DiskEventAnomalyDetectionConfig
*/
func (api DiskAnomalyDetectionAPI) Get(id string) (config.DiskEventAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.GetDiskEventConfig1(nil, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
List Get the list of disk event rules.
@return DiskEventAnomalyDetectionConfigList
*/
func (api DiskAnomalyDetectionAPI) List() (config.DiskEventAnomalyDetectionConfigList, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.ListDiskEventConfigs1(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Update the disk event rule with the given id.
 * @param id
 * @param optional nil or *UpdateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
*/
func (api DiskAnomalyDetectionAPI) Update(id string, c config.DiskEventAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.UpdateDiskEventConfig1(nil, id, &config.UpdateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Validate Validate configuration of disk event rule.
 * @param optional nil or *ValidateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
*/
func (api DiskAnomalyDetectionAPI) Validate(c config.DiskEventAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.ValidateDiskEventConfig1(nil, &config.ValidateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Get Get configuration of anomaly detection for hosts.
 * @param  - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return HostsAnomalyDetectionConfig
*/
func (api HostAnomalyDetectionAPI) Get() (config.HostsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionHostsApi.GetHostEventsConfig1(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Update configuration of anomaly detection for hosts.
 * @param optional nil or *UpdateHostEventsConfig1Opts - Optional Parameters:
 * @param "HostsAnomalyDetectionConfig" (optional.Interface of HostsAnomalyDetectionConfig) -
*/
func (api HostAnomalyDetectionAPI) Update(c config.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.UpdateHostEventsConfig1(nil, &config.UpdateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Validate Validate configuration of anomaly detection for hosts (for PUT request).
 * @param optional nil or *ValidateHostEventsConfig1Opts - Optional Parameters:
 * @param "HostsAnomalyDetectionConfig" (optional.Interface of HostsAnomalyDetectionConfig) -
*/
func (api HostAnomalyDetectionAPI) Validate(c config.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.ValidateHostEventsConfig1(nil, &config.ValidateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Get Gets the service anomaly detection configuration
@return ServiceAnomalyDetectionConfig
*/
func (api ServiceAnomalyDetectionAPI) Get() (config.ServiceAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionServicesApi.GetConfiguration1(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Updates the service anomaly detection configuration
 * @param optional nil or *UpdateConfiguration1Opts - Optional Parameters:
 * @param "ServiceAnomalyDetectionConfig" (optional.Interface of ServiceAnomalyDetectionConfig) -
@return ServiceAnomalyDetectionConfig
*/
func (api ServiceAnomalyDetectionAPI) Update(c config.ServiceAnomalyDetectionConfig) (config.ServiceAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionServicesApi.UpdateConfiguration1(nil, &config.UpdateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Validate Validates the service anomaly detection configuration
 * @param optional nil or *ValidateConfiguration1Opts - Optional Parameters:
 * @param "ServiceAnomalyDetectionConfig" (optional.Interface of ServiceAnomalyDetectionConfig) -
*/
func (api ServiceAnomalyDetectionAPI) Validate(c config.ServiceAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionServicesApi.ValidateConfiguration1(nil, &config.ValidateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Get Get configuration of anomaly detection for VMware.
@return VMwareAnomalyDetectionConfig
*/
func (api VMwareAnomalyDetectionAPI) Get() (config.VMwareAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig1(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
Update Update configuration of anomaly detection for VMware.
 * @param optional nil or *UpdateVMwareAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "VMwareAnomalyDetectionConfig" (optional.Interface of VMwareAnomalyDetectionConfig) -
*/
func (api VMwareAnomalyDetectionAPI) Update(c config.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.UpdateVMwareAnomalyDetectionConfig1(nil, &config.UpdateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
Validate Validate configuration of anomaly detection for VMware (for PUT request).
 * @param c config.VMwareAnomalyDetectionConfig
*/
func (api VMwareAnomalyDetectionAPI) Validate(c config.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.ValidateVMwareAnomalyDetectionConfig1(nil, &config.ValidateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}
