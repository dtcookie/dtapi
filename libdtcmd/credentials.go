package dtcmd

// Credentials hold the necessary information in order to access
// the REST API of a Dynatrace Tenant
type Credentials struct {
	// Environment specifies the unique identifier of a
	// Dynatrace Tenant.
	// For a SaaS Tenant, this is the first part of the Web URL
	// Example: https://<environment>.live.dynatrace.com
	Environment string
	// APIToken specifies the access token to be used to
	// authenticate against the Dynatrace REST API.
	APIToken string
}
