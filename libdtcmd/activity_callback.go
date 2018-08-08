package dtcmd

// ActivityCallback is being used by a Tenant for logging
// purposes.
type ActivityCallback interface {
	// OnCreateWebApp is getting called by a Tenant when it is
	// about to create a Web Application via Dynatrace REST API
	//
	// Purpose of this callback function is to customize log output
	// (e.g. print out technology specific stuff)
	OnCreateWebApp(name string)
}
