# MonitoringSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchRequests** | **bool** | Capture fetch() request enabled/disabled | 
**XmlHttpRequest** | **bool** | Support for XmlHttpRequest enabled/disabled | 
**JavaScriptFrameworkSupport** | [**JavaScriptFrameworkSupport**](JavaScriptFrameworkSupport.md) |  | 
**ContentCapture** | [**ContentCapture**](ContentCapture.md) |  | 
**ExcludeXhrRegex** | **string** | A regular expression to match all URLs that should be excluded from becoming XHR actions.  An empty value disables the feature. | 
**InjectionMode** | **string** | JavaScript injection mode. | 
**LibraryFileLocation** | **string** | The source path for placement of your application’s custom JavaScript library file.  An empty value sets it to the root directory of your web server. | 
**MonitoringDataPath** | **string** | The path where the JavaScript tag should send monitoring data.  Specify either a relative or an absolute URL. If you enter an absolute URL, data will be sent using CORS. | 
**CustomConfigurationProperties** | **string** | Additional JavaScript tag properties that are specific to your application. To do this, type key-value pairs defined using (&#x3D;) and separated using a (|) symbol. | 
**ServerRequestPathId** | **string** | Path that is to be used to identify the server’s request ID. | 
**SecureCookieAttribute** | **bool** | Use the secure cookie attribute for cookies set by Dynatrace enabled/disabled. | 
**CookiePlacementDomain** | **string** | Domain to be used for cookie placement. | 
**CacheControlHeaderOptimizations** | **bool** | Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled. | 
**AdvancedJavaScriptTagSettings** | [**AdvancedJavaScriptTagSettings**](AdvancedJavaScriptTagSettings.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


