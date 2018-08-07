package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// AnomalyDetectionAPI TODO: documentation
type AnomalyDetectionAPI confService

// AwsAnomalyDetectionAPI TODO: documentation
type AwsAnomalyDetectionAPI confService

// DiskAnomalyDetectionAPI TODO: documentation
type DiskAnomalyDetectionAPI confService

// HostAnomalyDetectionAPI TODO: documentation
type HostAnomalyDetectionAPI confService

// ServiceAnomalyDetectionAPI TODO: documentation
type ServiceAnomalyDetectionAPI confService

// VMwareAnomalyDetectionAPI TODO: documentation
type VMwareAnomalyDetectionAPI confService

/*
AWS TODO: document
*/
func (api AnomalyDetectionAPI) AWS() AwsAnomalyDetectionAPI {
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
GET Get configuration of anomaly detection for AWS.
@return AwsAnomalyDetectionConfig
*/
func (api AwsAnomalyDetectionAPI) GET() (dtapi.AwsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig1(nil)
	return result, err
}

/*
UPDATE Updates (HTTP PUT) configuration of anomaly detection for AWS.
 * @param optional nil or *UpdateAwsAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "AwsAnomalyDetectionConfig" (optional.Interface of AwsAnomalyDetectionConfig) -
*/
func (api AwsAnomalyDetectionAPI) UPDATE(c dtapi.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.UpdateAwsAnomalyDetectionConfig1(nil, &dtapi.UpdateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validate configuration of anomaly detection for AWS (for PUT request).
 * @param optional nil or *ValidateAwsAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "AwsAnomalyDetectionConfig" (optional.Interface of AwsAnomalyDetectionConfig) -
*/
func (api AwsAnomalyDetectionAPI) VALIDATE(c dtapi.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.ValidateAwsAnomalyDetectionConfig1(nil, &dtapi.ValidateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
CREATE Creates (HTTP POST) a new disk event rule.
 * @param optional nil or *CreateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
@return DiskEventAnomalyDetectionConfig
*/
func (api DiskAnomalyDetectionAPI) CREATE(c dtapi.DiskEventAnomalyDetectionConfig) (dtapi.DiskEventAnomalyDetectionConfigStub, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.CreateDiskEventConfig1(nil, &dtapi.CreateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return result, err
}

/*
DELETE Deletes the disk alert rule with the given id.
 * @param id
*/
func (api DiskAnomalyDetectionAPI) DELETE(ID string) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.DeleteDiskEventConfig1(nil, ID)
	return err
}

/*
GET Gets the disk event rule with the given id.
 * @param id
@return DiskEventAnomalyDetectionConfig
*/
func (api DiskAnomalyDetectionAPI) GET(ID string) (dtapi.DiskEventAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.GetDiskEventConfig1(nil, ID)
	return result, err
}

/*
GETALL Get the list of disk event rules.
@return DiskEventAnomalyDetectionConfigList
*/
func (api DiskAnomalyDetectionAPI) GETALL() ([]dtapi.DiskEventAnomalyDetectionConfigStub, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.ListDiskEventConfigs1(nil)
	return result.Values, err
}

/*
UPDATE Updates (HTTP PUT) the disk event rule with the given id.
 * @param id
 * @param optional nil or *UpdateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
*/
func (api DiskAnomalyDetectionAPI) UPDATE(id string, c dtapi.DiskEventAnomalyDetectionConfigStub) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.UpdateOrCreateDiskEventConfig1(nil, id, &dtapi.UpdateOrCreateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validates configuration of disk event rule.
 * @param optional nil or *ValidateDiskEventConfig1Opts - Optional Parameters:
 * @param "DiskEventAnomalyDetectionConfig" (optional.Interface of DiskEventAnomalyDetectionConfig) -
*/
func (api DiskAnomalyDetectionAPI) VALIDATE(c dtapi.DiskEventAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.ValidateDiskEventConfig1(nil, &dtapi.ValidateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
GET Gets configuration of anomaly detection for hosts.
 * @param  - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return HostsAnomalyDetectionConfig
*/
func (api HostAnomalyDetectionAPI) GET() (dtapi.HostsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionHostsApi.GetHostEventsConfig1(nil)
	return result, err
}

/*
UPDATE Updates configuration of anomaly detection for hosts.
 * @param optional nil or *UpdateHostEventsConfig1Opts - Optional Parameters:
 * @param "HostsAnomalyDetectionConfig" (optional.Interface of HostsAnomalyDetectionConfig) -
*/
func (api HostAnomalyDetectionAPI) UPDATE(c dtapi.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.UpdateHostEventsConfig1(nil, &dtapi.UpdateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validates configuration of anomaly detection for hosts (for PUT request).
 * @param optional nil or *ValidateHostEventsConfig1Opts - Optional Parameters:
 * @param "HostsAnomalyDetectionConfig" (optional.Interface of HostsAnomalyDetectionConfig) -
*/
func (api HostAnomalyDetectionAPI) VALIDATE(c dtapi.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.ValidateHostEventsConfig1(nil, &dtapi.ValidateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
GET Gets the service anomaly detection configuration
@return ServiceAnomalyDetectionConfig
*/
func (api ServiceAnomalyDetectionAPI) GET() (dtapi.ServiceAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionServicesApi.GetConfiguration1(nil)
	return result, err
}

/*
UPDATE Updates (HTTP PUT) the service anomaly detection configuration
 * @param optional nil or *UpdateConfiguration1Opts - Optional Parameters:
 * @param "ServiceAnomalyDetectionConfig" (optional.Interface of ServiceAnomalyDetectionConfig) -
@return ServiceAnomalyDetectionConfig
*/
func (api ServiceAnomalyDetectionAPI) UPDATE(c dtapi.ServiceAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionServicesApi.UpdateConfiguration1(nil, &dtapi.UpdateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validates the service anomaly detection configuration
 * @param optional nil or *ValidateConfiguration1Opts - Optional Parameters:
 * @param "ServiceAnomalyDetectionConfig" (optional.Interface of ServiceAnomalyDetectionConfig) -
*/
func (api ServiceAnomalyDetectionAPI) VALIDATE(c dtapi.ServiceAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionServicesApi.ValidateConfiguration1(nil, &dtapi.ValidateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
GET Gets configuration of anomaly detection for VMware.
@return VMwareAnomalyDetectionConfig
*/
func (api VMwareAnomalyDetectionAPI) GET() (dtapi.VMwareAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig1(nil)
	return result, err
}

/*
UPDATE Updates (HTTP PUT) configuration of anomaly detection for VMware.
 * @param optional nil or *UpdateVMwareAnomalyDetectionConfig1Opts - Optional Parameters:
 * @param "VMwareAnomalyDetectionConfig" (optional.Interface of VMwareAnomalyDetectionConfig) -
*/
func (api VMwareAnomalyDetectionAPI) UPDATE(c dtapi.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.UpdateVMwareAnomalyDetectionConfig1(nil, &dtapi.UpdateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

/*
VALIDATE Validates configuration of anomaly detection for VMware (for PUT request).
 * @param c config.VMwareAnomalyDetectionConfig
*/
func (api VMwareAnomalyDetectionAPI) VALIDATE(c dtapi.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.ValidateVMwareAnomalyDetectionConfig1(nil, &dtapi.ValidateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}
