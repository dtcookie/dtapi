package dtapi

import (
	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// TimeseriesAPI TODO: documentation
type TimeseriesAPI envService

/*
CreateCustomTimeseries Creates a new custom metric.
 * @param timeseriesIdentifier The ID for the new metric. It must start with the `custom:` prefix.   If you use the ID of an existing metric the respective parameters will be updated.
 * @param optional nil or *CreateCustomTimeseriesOpts - Optional Parameters:
 * @param "TimeseriesRegistrationMessage" (optional.Interface of TimeseriesRegistrationMessage) -  The body of the request, containing metric parameters.
@return TimeseriesDefinition
*/
func (api TimeseriesAPI) CreateCustomTimeseries(timeseriesIdentifier string, c *dtapi.TimeseriesRegistrationMessage) (dtapi.TimeseriesDefinition, error) {
	result, _, err := api.client.TimeseriesApi.CreateCustomTimeseries(nil, timeseriesIdentifier, &dtapi.CreateCustomTimeseriesOpts{
		TimeseriesRegistrationMessage: optional.NewInterface(c),
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
DeleteCustomTimeseries Deletes a custom metric.
 * @param timeseriesIdentifier The ID of the metric to be deleted.
*/
func (api TimeseriesAPI) DeleteCustomTimeseries(timeseriesIdentifier string) error {
	_, err := api.client.TimeseriesApi.DeleteCustomTimeseries(nil, timeseriesIdentifier)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAllTimeseriesDefinitions Lists all metric definitions, along with parameters of each metric.
You can specify filtering paramters to return only matched metrics. If no parameters are specified, the call will list all the defined and exposed metrics.
 * @param optional nil or *GetAllTimeseriesDefinitionsOpts - Optional Parameters:
 * @param "Source" (optional.String) -  Metric type. Allowed values are `BUILTIN`, `PLUGIN`, and `CUSTOM`.
 * @param "DetailedSource" (optional.String) -  The feature, where metrics originate, such as Synthetic or RUM.
@return []TimeseriesDefinition
*/
func (api TimeseriesAPI) GetAllTimeseriesDefinitions(opts *dtapi.GetAllTimeseriesDefinitionsOpts) ([]dtapi.TimeseriesDefinition, error) {
	result, _, err := api.client.TimeseriesApi.GetAllTimeseriesDefinitions(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
ReadTimeseriesComplex Lists all available metric data points, matching specified parameters.
Provides advanced filtering possibilities, comparing to the &#x60;GET /timeseries/{metricIdentifier}&#x60; request.
 * @param timeseriesIdentifier Case-sensitive identifier of the timeseries, where you want to read parameters and data points.\"
 * @param optional nil or *ReadTimeseriesComplexOpts - Optional Parameters:
 * @param "TimeseriesQueryMessage" (optional.Interface of TimeseriesQueryMessage) -  JSON body of the request, containing parameters to identify the required data points.
@return TimeseriesQueryResultWrapper
*/
func (api TimeseriesAPI) ReadTimeseriesComplex(timeseriesIdentifier string, opts *dtapi.ReadTimeseriesComplexOpts) (dtapi.TimeseriesQueryResultWrapper, error) {
	result, _, err := api.client.TimeseriesApi.ReadTimeseriesComplex(nil, timeseriesIdentifier, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// /*
// ReadTimeseriesComplexWithoutPathParam TODO: comment
//  * @param optional nil or *ReadTimeseriesComplexWithoutPathParamOpts - Optional Parameters:
//  * @param "TimeseriesQueryMessage" (optional.Interface of TimeseriesQueryMessage) -
// */
// func (api TimeseriesAPI) ReadTimeseriesComplexWithoutPathParam(opts *dtapi.ReadTimeseriesComplexWithoutPathParamOpts) error {
// 	_, err := api.client.TimeseriesApi.ReadTimeseriesComplexWithoutPathParam(nil, opts)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

/*
ReadTimeseriesData Gets the parameters of the specified metric and optionally data points.
To obtain data points, set **includeData** to &#x60;true&#x60;.   You can obtain either data points or the scalar result of the specified timeseries, depending on the **queryMode**.   To obtain data points you must specify the timeframe, either as **relativeTime** or as combination of **startTimestamp** and **endTimestamp**. You must also provide **aggregationType**, supported by the metric.
 * @param timeseriesIdentifier Case-sensitive identifier of the timeseries, where you want to read parameters and data points.
 * @param optional nil or *ReadTimeseriesDataOpts - Optional Parameters:
 * @param "IncludeData" (optional.Bool) -  Flag to include data points to the response. \\n\\n To obtain data points you must specify the timeframe and aggregation type.
 * @param "AggregationType" (optional.String) -  The aggregation type for the resulting data points. \\n\\n\" +      \"If the requested metric doesn't support the specified aggregation, the request will result in an error.
 * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC). The start time must be earlier than the end time.
 * @param "EndTimestamp" (optional.Int64) -  End timestamp of the requested timeframe, in milliseconds (UTC). End time must be later than the start time. \\n\\n \" +      \"If later than the current time, Dynatrace automatically uses current time instead.
 * @param "Predict" (optional.Bool) -  Used to predict future data points.
 * @param "RelativeTime" (optional.String) -  Relative timeframe, back from the current time.
 * @param "QueryMode" (optional.String) -  The type of result that the call should return. Valid result modes are: \\n\" +         \"`series`: returns all the data points of the timeseries in specified timeframe. \\n `total`: returns one scalar value for the specified timeframe. \\n\\n\"+         \"By default, the `series` mode is used.
 * @param "Entity" (optional.Interface of []string) -  Filters requested data points by entity which should deliver them. \\n\\n\" +         \"Allowed values are Dynatrace entity IDs. You can find them in the URL of the corresponding Dynatrace entity page, for example, `HOST-007`. \\n\\n\" +         \"If the selected entity doesn't support the requested timeseries, the request will result in an error.
 * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of applications by the specified tag. \\n\\n\" +         \"Use multiple tag parameters to combine multiple tag filters using the logical operator AND. \\n\\n\" +         \"In case of key-value tags, such as imported AWS or CloudFoundry tags use following format: `[context]key:value`.
 * @param "Percentile" (optional.Int32) -  In case of the percentile aggregation type, this parameter specifies which percentile of the selected response time metric should be delivered. \" +         \"Valid values for percentile are between 1 and 99. \\n\\n\" +         \"Please keep in mind that percentile export is only possible for response-time based metrics such as application and service response times.
@return []TimeseriesQueryResult
*/
func (api TimeseriesAPI) ReadTimeseriesData(timeseriesIdentifier string, opts *dtapi.ReadTimeseriesDataOpts) ([]dtapi.TimeseriesQueryResult, error) {
	result, _, err := api.client.TimeseriesApi.ReadTimeseriesData(nil, timeseriesIdentifier, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
