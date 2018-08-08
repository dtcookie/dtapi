package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapiconf"
)

// AnomalyDetectionAPI is a cleaned up wrapper around the
// Anomaly Detection functinality offered via
// "github.com/dtcookie/dtapi/libdtapiconf"
type AnomalyDetectionAPI confService

// AwsAnomalyDetectionAPI covers anomaly detection
// for AWS
type AwsAnomalyDetectionAPI confService

// DiskAnomalyDetectionAPI covers anomaly detection
// for Disks
type DiskAnomalyDetectionAPI confService

// HostAnomalyDetectionAPI covers anomaly detection
// for Hosts
type HostAnomalyDetectionAPI confService

// ServiceAnomalyDetectionAPI covers anomaly detection
// for Services
type ServiceAnomalyDetectionAPI confService

// VMwareAnomalyDetectionAPI covers anomaly detection for
// VMWare
type VMwareAnomalyDetectionAPI confService

// AWS returns an API covering Anomaly Detection for AWS
func (api *AnomalyDetectionAPI) AWS() *AwsAnomalyDetectionAPI {
	return &AwsAnomalyDetectionAPI{
		client: api.client,
	}
}

// VMware returns an API covering Anomaly Detection for VMware
func (api *AnomalyDetectionAPI) VMware() *VMwareAnomalyDetectionAPI {
	return &VMwareAnomalyDetectionAPI{
		client: api.client,
	}
}

// Disks returns an API covering Anomaly Detection for Disks
func (api *AnomalyDetectionAPI) Disks() *DiskAnomalyDetectionAPI {
	return &DiskAnomalyDetectionAPI{
		client: api.client,
	}
}

// Hosts returns an API covering Anomaly Detection for Hosts
func (api *AnomalyDetectionAPI) Hosts() *HostAnomalyDetectionAPI {
	return &HostAnomalyDetectionAPI{
		client: api.client,
	}
}

// Services returns an API covering Anomaly Detection for Services
func (api *AnomalyDetectionAPI) Services() *ServiceAnomalyDetectionAPI {
	return &ServiceAnomalyDetectionAPI{
		client: api.client,
	}
}

// Get retrieves the current settings for Anomaly Detection related to AWS
func (api *AwsAnomalyDetectionAPI) Get() (dtapi.AwsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionAWSApi.GetAwsAnomalyDetectionConfig1(nil)
	return result, err
}

// Update updates (HTTP PUT) configuration of anomaly detection for AWS.
func (api *AwsAnomalyDetectionAPI) Update(c dtapi.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.UpdateAwsAnomalyDetectionConfig1(nil, &dtapi.UpdateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating the configuration for Anomaly Detection related to AWS
// before the actual Update.
func (api *AwsAnomalyDetectionAPI) Validate(c dtapi.AwsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionAWSApi.ValidateAwsAnomalyDetectionConfig1(nil, &dtapi.ValidateAwsAnomalyDetectionConfig1Opts{
		AwsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Create creates (HTTP POST) a new rule for Disk Event Anomaly Detection
// Returns:
// 	- DiskEventAnomalyDetectionConfigStub	 a name/ID pair for the created disk event rule
func (api *DiskAnomalyDetectionAPI) Create(c dtapi.DiskEventAnomalyDetectionConfig) (dtapi.DiskEventAnomalyDetectionConfigStub, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.CreateDiskEventConfig1(nil, &dtapi.CreateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return result, err
}

// Delete removes a rule for Disk Event Anomaly Detection
//
// Parameters:
//	- ID	the unique identifier of the rule to remove
func (api *DiskAnomalyDetectionAPI) Delete(ID string) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.DeleteDiskEventConfig1(nil, ID)
	return err
}

// Get retrieves the detailed configuration for a Disk Event Rule
// for Anomaly Detection
//
// Parameters:
//	- ID	the unique identifier of the disk event rule
// Returns:
//	- DiskEventAnomalyDetectionConfig	detailed configuration of the
//										disk event rule
func (api *DiskAnomalyDetectionAPI) Get(ID string) (dtapi.DiskEventAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.GetDiskEventConfig1(nil, ID)
	return result, err
}

// GetAll retrieves name/ID pairs for all currently configured Disk Event Rules
// for Anomaly Detection
func (api *DiskAnomalyDetectionAPI) GetAll() ([]dtapi.DiskEventAnomalyDetectionConfigStub, error) {
	result, _, err := api.client.AnomalyDetectionDiskEventsApi.ListDiskEventConfigs1(nil)
	return result.Values, err
}

// Update updates (HTTP PUT) a Disk Event Rule for Anomaly Detection
//
// Note:
//	In case the specified DiskEventAnomalyDetectionConfig contains an
//	identifier for the Disk Event Rule, it MUST equal the specified 'ID'
//
// TODO:
//	Perform check for equality of IDs offline
//
// Parameters:
//	- ID	is the unique identifier of the Disk Event Rule to update
//	- c		is the detailed configuration of the Disk Event Rule to update.
func (api *DiskAnomalyDetectionAPI) Update(ID string, c dtapi.DiskEventAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.UpdateOrCreateDiskEventConfig1(nil, ID, &dtapi.UpdateOrCreateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating the given DiskEventAnomalyDetectionConfig before performing the
// actual Update or Create operation.
func (api *DiskAnomalyDetectionAPI) Validate(c dtapi.DiskEventAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionDiskEventsApi.ValidateDiskEventConfig1(nil, &dtapi.ValidateDiskEventConfig1Opts{
		DiskEventAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Get retrieves the detailed configuration for Host Anomaly Detection
func (api *HostAnomalyDetectionAPI) Get() (dtapi.HostsAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionHostsApi.GetHostEventsConfig1(nil)
	return result, err
}

// Update updates (HTTP PUT) the configuration for Host Anomaly Detection
func (api *HostAnomalyDetectionAPI) Update(c dtapi.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.UpdateHostEventsConfig1(nil, &dtapi.UpdateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating a configuration for Host Anomaly Detection
// before performing the actual Update.
func (api *HostAnomalyDetectionAPI) Validate(c dtapi.HostsAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionHostsApi.ValidateHostEventsConfig1(nil, &dtapi.ValidateHostEventsConfig1Opts{
		HostsAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Get retrieves the detailed configuration for Service Anomaly Detection
func (api *ServiceAnomalyDetectionAPI) Get() (dtapi.ServiceAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionServicesApi.GetConfiguration1(nil)
	return result, err
}

// Update updates (HTTP PUT) the configuration for Service Anomaly Detection
func (api *ServiceAnomalyDetectionAPI) Update(c dtapi.ServiceAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionServicesApi.UpdateConfiguration1(nil, &dtapi.UpdateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating a configuration for Service Anomaly Detection
// before performing the actual Update.
func (api *ServiceAnomalyDetectionAPI) Validate(c dtapi.ServiceAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionServicesApi.ValidateConfiguration1(nil, &dtapi.ValidateConfiguration1Opts{
		ServiceAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Get retrieves the detailed configuration for VMWare Anomaly Detection
func (api *VMwareAnomalyDetectionAPI) Get() (dtapi.VMwareAnomalyDetectionConfig, error) {
	result, _, err := api.client.AnomalyDetectionVMwareApi.GetVMwareAnomalyDetectionConfig1(nil)
	return result, err
}

// Update updates (HTTP PUT) the configuration for VMWare Anomaly Detection
func (api *VMwareAnomalyDetectionAPI) Update(c dtapi.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.UpdateVMwareAnomalyDetectionConfig1(nil, &dtapi.UpdateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}

// Validate allows for validating a configuration for VMWare Anomaly Detection
// before performing the actual Update.
func (api *VMwareAnomalyDetectionAPI) Validate(c dtapi.VMwareAnomalyDetectionConfig) error {
	_, err := api.client.AnomalyDetectionVMwareApi.ValidateVMwareAnomalyDetectionConfig1(nil, &dtapi.ValidateVMwareAnomalyDetectionConfig1Opts{
		VMwareAnomalyDetectionConfig: optional.NewInterface(c),
	})
	return err
}
