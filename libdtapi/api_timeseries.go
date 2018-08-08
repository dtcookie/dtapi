package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TimeseriesAPI is a convenience wrapper around
// the Services offered via
// "github.com/dtcookie/dtapi/libdtapienv" related to
// configuring Timeseries and Custom Metrics
type TimeseriesAPI envService

// Create creates a new custom metric.
func (api TimeseriesAPI) Create(ID string, c *dtapi.TimeseriesRegistrationMessage) (dtapi.TimeseriesDefinition, error) {
	result, _, err := api.client.TimeseriesApi.CreateCustomTimeseries(nil, ID, &dtapi.CreateCustomTimeseriesOpts{
		TimeseriesRegistrationMessage: optional.NewInterface(c),
	})
	return result, err
}

// Delete deletes a custom metric.
func (api TimeseriesAPI) Delete(ID string) error {
	_, err := api.client.TimeseriesApi.DeleteCustomTimeseries(nil, ID)
	return err
}

// GetAll Lists all metric definitions, along with parameters of each metric.
// You can specify filtering paramters to return only matched metrics.
// If no parameters are specified, the call will list all the defined and exposed metrics.
//
// Paramters:
//	- source			Metric type. Allowed values are `BUILTIN`, `PLUGIN`, and `CUSTOM`.
//						Specify "" for "don't care"
//	- detailedSource	The feature, where metrics originate, such as Synthetic or RUM.
//						Specify "" for "don't care"
//
func (api TimeseriesAPI) GetAll(source string, detailedSource string) ([]dtapi.TimeseriesDefinition, error) {
	result, _, err := api.client.TimeseriesApi.GetAllTimeseriesDefinitions(nil, &dtapi.GetAllTimeseriesDefinitionsOpts{
		Source:         optional.NewString(source),
		DetailedSource: optional.NewString(detailedSource),
	})
	return result, err
}

// Query Lists all available metric data points, matching specified parameters.
func (api TimeseriesAPI) Query(ID string, query dtapi.TimeseriesQueryMessage) (dtapi.TimeseriesDataPointQueryResult, error) {
	result, _, err := api.client.TimeseriesApi.ReadTimeseriesComplex(nil, ID, &dtapi.ReadTimeseriesComplexOpts{
		TimeseriesQueryMessage: optional.NewInterface(query),
	})
	return result.Result, err
}

// /*
// ReadTimeseriesData Gets the parameters of the specified metric and optionally data points.
// To obtain data points, set **includeData** to &#x60;true&#x60;.   You can obtain either data points or the scalar result of the specified timeseries, depending on the **queryMode**.   To obtain data points you must specify the timeframe, either as **relativeTime** or as combination of **startTimestamp** and **endTimestamp**. You must also provide **aggregationType**, supported by the metric.
//  * @param timeseriesIdentifier Case-sensitive identifier of the timeseries, where you want to read parameters and data points.
//  * @param optional nil or *ReadTimeseriesDataOpts - Optional Parameters:
//  * @param "IncludeData" (optional.Bool) -  Flag to include data points to the response. \\n\\n To obtain data points you must specify the timeframe and aggregation type.
//  * @param "AggregationType" (optional.String) -  The aggregation type for the resulting data points. \\n\\n\" +      \"If the requested metric doesn't support the specified aggregation, the request will result in an error.
//  * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC). The start time must be earlier than the end time.
//  * @param "EndTimestamp" (optional.Int64) -  End timestamp of the requested timeframe, in milliseconds (UTC). End time must be later than the start time. \\n\\n \" +      \"If later than the current time, Dynatrace automatically uses current time instead.
//  * @param "Predict" (optional.Bool) -  Used to predict future data points.
//  * @param "RelativeTime" (optional.String) -  Relative timeframe, back from the current time.
//  * @param "QueryMode" (optional.String) -  The type of result that the call should return. Valid result modes are: \\n\" +         \"`series`: returns all the data points of the timeseries in specified timeframe. \\n `total`: returns one scalar value for the specified timeframe. \\n\\n\"+         \"By default, the `series` mode is used.
//  * @param "Entity" (optional.Interface of []string) -  Filters requested data points by entity which should deliver them. \\n\\n\" +         \"Allowed values are Dynatrace entity IDs. You can find them in the URL of the corresponding Dynatrace entity page, for example, `HOST-007`. \\n\\n\" +         \"If the selected entity doesn't support the requested timeseries, the request will result in an error.
//  * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of applications by the specified tag. \\n\\n\" +         \"Use multiple tag parameters to combine multiple tag filters using the logical operator AND. \\n\\n\" +         \"In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: `[context]key:value`.
//  * @param "Percentile" (optional.Int32) -  In case of the percentile aggregation type, this parameter specifies which percentile of the selected response time metric should be delivered. \" +         \"Valid values for percentile are between 1 and 99. \\n\\n\" +         \"Please keep in mind that percentile export is only possible for response-time based metrics such as application and service response times.
// @return []TimeseriesQueryResult
// */
// func (api TimeseriesAPI) ReadTimeseriesData(timeseriesIdentifier string, opts *dtapi.ReadTimeseriesDataOpts) ([]dtapi.TimeseriesQueryResult, error) {
// 	result, _, err := api.client.TimeseriesApi.ReadTimeseriesData(nil, timeseriesIdentifier, opts)
// 	if err != nil {
// 		return result, err
// 	}
// 	return result, nil
// }
