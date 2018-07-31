package dynatrace

// ClusterAPI TODO: documentation
type ClusterAPI service

// ClusterTime TODO: documentation
func (api ClusterAPI) ClusterTime() (int64, error) {
	clusterTime, _, err := api.client.ClusterTimeApi.GetCurrentClusterTime(nil)
	if err != nil {
		return 0, err
	}
	return clusterTime, nil
}

// Version TODO: documentation
func (api ClusterAPI) Version() (string, error) {
	version, _, err := api.client.ClusterVersionApi.GetVersion(nil)
	if err != nil {
		return "", err
	}
	return version.Version, nil
}
