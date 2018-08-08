package dtapi

import dtapienv "github.com/dtcookie/dtapi/libdtapienv"

// ClusterAPI is a convenience wrapper around the
// Services offered for Cluster Time and Cluster Version
// via "github.com/dtcookie/dtapi/libdtapienv"
type ClusterAPI envService

// ClusterTime retrieves the current cluster time of the
// Tenant in UTC milliseconds.
func (api *ClusterAPI) ClusterTime() (int64, error) {
	var err error
	var clusterTime int64

	if clusterTime, _, err = api.client.ClusterTimeApi.GetCurrentClusterTime(nil); err != nil {
		return 0, err
	}
	return clusterTime, nil
}

// Version retrieves the current cluster version of the
// Tenant.
func (api *ClusterAPI) Version() (string, error) {
	var err error
	var version dtapienv.ClusterVersion

	if version, _, err = api.client.ClusterVersionApi.GetVersion(nil); err != nil {
		return "", err
	}
	return version.Version, nil
}
