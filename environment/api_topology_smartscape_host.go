/*
 * Dynatrace Environment API
 *
 * Documentation of the Dynatrace REST API. Refer to the [help page](https://www.dynatrace.com/support/help/shortlink/section-api) to read about use-cases and examples.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package environment

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type TopologySmartscapeHostApiService service

/*
TopologySmartscapeHostApiService Lists all available hosts in your environment
You can narrow down the output by specifying filtering parameters of the request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetHostsOpts - Optional Parameters:
 * @param "StartTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used.
 * @param "EndTimestamp" (optional.Int64) -  Start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then now is used.
 * @param "RelativeTime" (optional.String) -  Relative timeframe, back from now.
 * @param "Tag" (optional.Interface of []string) -  Filters the resulting set of hosts by the specified tag.    A host has to match ALL specified tags.
 * @param "ShowMonitoringCandidates" (optional.Bool) -  Include/exclude monitoring canditate to the response.   Monitoring candidates are network entities, which are detected but not monitored.
 * @param "Entity" (optional.Interface of []string) -  Only return specified hosts.
@return []Host
*/

type GetHostsOpts struct {
    StartTimestamp optional.Int64
    EndTimestamp optional.Int64
    RelativeTime optional.String
    Tag optional.Interface
    ShowMonitoringCandidates optional.Bool
    Entity optional.Interface
}

func (a *TopologySmartscapeHostApiService) GetHosts(ctx context.Context, localVarOptionals *GetHostsOpts) ([]Host, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue []Host
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entity/infrastructure/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.StartTimestamp.IsSet() {
		localVarQueryParams.Add("startTimestamp", parameterToString(localVarOptionals.StartTimestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndTimestamp.IsSet() {
		localVarQueryParams.Add("endTimestamp", parameterToString(localVarOptionals.EndTimestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RelativeTime.IsSet() {
		localVarQueryParams.Add("relativeTime", parameterToString(localVarOptionals.RelativeTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tag.IsSet() {
		localVarQueryParams.Add("tag", parameterToString(localVarOptionals.Tag.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ShowMonitoringCandidates.IsSet() {
		localVarQueryParams.Add("showMonitoringCandidates", parameterToString(localVarOptionals.ShowMonitoringCandidates.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Entity.IsSet() {
		localVarQueryParams.Add("entity", parameterToString(localVarOptionals.Entity.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("Api-Token", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []Host
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
TopologySmartscapeHostApiService Gets parameters of the specified host
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param meIdentifier Dynatrace entity ID of the host you're inquiring.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
@return Host
*/
func (a *TopologySmartscapeHostApiService) GetSingleHost(ctx context.Context, meIdentifier string) (Host, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue Host
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entity/infrastructure/hosts/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", fmt.Sprintf("%v", meIdentifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("Api-Token", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Host
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
TopologySmartscapeHostApiService Updates properties of the specified host
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param meIdentifier Dynatrace entity ID of the host to be updated.   You can find it in the URL of the corresponding host page, for example, `HOST-007`.
 * @param optional nil or *UpdateHostOpts - Optional Parameters:
 * @param "UpdateEntity" (optional.Interface of UpdateEntity) - 
*/

type UpdateHostOpts struct {
    UpdateEntity optional.Interface
}

func (a *TopologySmartscapeHostApiService) UpdateHost(ctx context.Context, meIdentifier string, localVarOptionals *UpdateHostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entity/infrastructure/hosts/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", fmt.Sprintf("%v", meIdentifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.UpdateEntity.IsSet() {
		localVarOptionalUpdateEntity, localVarOptionalUpdateEntityok := localVarOptionals.UpdateEntity.Value().(UpdateEntity)
		if !localVarOptionalUpdateEntityok {
			return nil, reportError("updateEntity should be UpdateEntity")
		}
		localVarPostBody = &localVarOptionalUpdateEntity
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("Api-Token", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}