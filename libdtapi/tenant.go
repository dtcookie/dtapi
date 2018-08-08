package dtapi

// Tenant TODO: comment
type Tenant struct {
	Cluster              *ClusterAPI
	Anonymization        *AnonymizationAPI
	Events               *EventAPI
	MaintenanceWindows   *MaintenanceWindowAPI
	Problems             *ProblemAPI
	Synthetic            *SyntheticAPI
	Thresholds           *ThresholdAPI
	Timeseries           *TimeseriesAPI
	Topology             *TopologyAPI
	AnomalyDetection     *AnomalyDetectionAPI
	ApplicationDetection *ApplicationDetectionConfigAPI
	CustomServices       *CustomServicesAPI
	DataPrivacy          *DataPrivacyAPI
	RequestAttributes    *RequestAttributesAPI
	UserSessionQL        *UserSessionQLAPI
	WebApplications      *WebApplicationConfigAPI
}

// NewTenant TODO: comment
func NewTenant(environment string, apiToken string) *Tenant {
	envClient := newSaasEnvClient(environment, apiToken)
	confClient := newSaasConfigClient(environment, apiToken)
	// envService := service{
	// 	client: envClient,
	// }
	webAppConfigAPI := WebApplicationConfigAPI{}
	webAppConfigAPI.client = confClient
	defaultWebAppConfigAPI := DefaultWebApplicationConfigAPI{}
	defaultWebAppConfigAPI.client = confClient
	defaultWebAppDataPrivacySettingsAPI := DefaultWebApplicationDataPrivacySettingsAPI{}
	defaultWebAppDataPrivacySettingsAPI.client = confClient
	defaultWebAppConfigAPI.DataPrivacySettings = defaultWebAppDataPrivacySettingsAPI
	webAppConfigAPI.Default = defaultWebAppConfigAPI

	anonymization := AnonymizationAPI{}
	anonymization.client = envClient

	return &Tenant{
		Cluster: &ClusterAPI{
			client: envClient,
		},
		Anonymization: &anonymization,
		Events: &EventAPI{
			client: envClient,
		},
		MaintenanceWindows: &MaintenanceWindowAPI{
			client: envClient,
		},
		Problems: &ProblemAPI{
			client: envClient,
		},
		Synthetic: &SyntheticAPI{
			client: envClient,
		},
		Thresholds: &ThresholdAPI{
			client: envClient,
		},
		Timeseries: &TimeseriesAPI{
			client: envClient,
		},
		Topology: newTopologyAPI(envClient),
		AnomalyDetection: &AnomalyDetectionAPI{
			client: confClient,
		},
		ApplicationDetection: &ApplicationDetectionConfigAPI{
			client: confClient,
		},
		CustomServices: newCustomServicesAPI(confClient),
		DataPrivacy: &DataPrivacyAPI{
			client: confClient,
		},
		RequestAttributes: &RequestAttributesAPI{
			client: confClient,
		},
		UserSessionQL: &UserSessionQLAPI{
			client: envClient,
		},
		WebApplications: &webAppConfigAPI,
	}
}
