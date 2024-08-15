/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ReadingsApiService ReadingsApi service
type ReadingsApiService service

type ReadingsApiGetOrganizationSensorReadingsHistoryRequest struct {
	ctx context.Context
	ApiService *ReadingsApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	networkIds *[]string
	serials *[]string
	metrics *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) PerPage(perPage int32) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) StartingAfter(startingAfter string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) EndingBefore(endingBefore string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) T0(t0 string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) T1(t1 string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) Timespan(timespan float32) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.timespan = &timespan
	return r
}

// Optional parameter to filter readings by network.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) NetworkIds(networkIds []string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter readings by sensor.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) Serials(serials []string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.serials = &serials
	return r
}

// Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water.
func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) Metrics(metrics []string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	r.metrics = &metrics
	return r
}

func (r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) Execute() ([]InlineResponse200269, *http.Response, error) {
	return r.ApiService.GetOrganizationSensorReadingsHistoryExecute(r)
}

/*
GetOrganizationSensorReadingsHistory Return all reported readings from sensors in a given timespan, sorted by timestamp

Return all reported readings from sensors in a given timespan, sorted by timestamp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ReadingsApiGetOrganizationSensorReadingsHistoryRequest
*/
func (a *ReadingsApiService) GetOrganizationSensorReadingsHistory(ctx context.Context, organizationId string) ReadingsApiGetOrganizationSensorReadingsHistoryRequest {
	return ReadingsApiGetOrganizationSensorReadingsHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200269
func (a *ReadingsApiService) GetOrganizationSensorReadingsHistoryExecute(r ReadingsApiGetOrganizationSensorReadingsHistoryRequest) ([]InlineResponse200269, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200269
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadingsApiService.GetOrganizationSensorReadingsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sensor/readings/history"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.metrics != nil {
		localVarQueryParams.Add("metrics", parameterToString(*r.metrics, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ReadingsApiGetOrganizationSensorReadingsLatestRequest struct {
	ctx context.Context
	ApiService *ReadingsApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	metrics *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) PerPage(perPage int32) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) StartingAfter(startingAfter string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) EndingBefore(endingBefore string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter readings by network.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) NetworkIds(networkIds []string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter readings by sensor.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) Serials(serials []string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.serials = &serials
	return r
}

// Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are apparentPower, battery, button, co2, current, door, downstreamPower, frequency, humidity, indoorAirQuality, noise, pm25, powerFactor, realPower, remoteLockoutSwitch, temperature, tvoc, voltage, and water.
func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) Metrics(metrics []string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	r.metrics = &metrics
	return r
}

func (r ReadingsApiGetOrganizationSensorReadingsLatestRequest) Execute() ([]InlineResponse200270, *http.Response, error) {
	return r.ApiService.GetOrganizationSensorReadingsLatestExecute(r)
}

/*
GetOrganizationSensorReadingsLatest Return the latest available reading for each metric from each sensor, sorted by sensor serial

Return the latest available reading for each metric from each sensor, sorted by sensor serial

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ReadingsApiGetOrganizationSensorReadingsLatestRequest
*/
func (a *ReadingsApiService) GetOrganizationSensorReadingsLatest(ctx context.Context, organizationId string) ReadingsApiGetOrganizationSensorReadingsLatestRequest {
	return ReadingsApiGetOrganizationSensorReadingsLatestRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200270
func (a *ReadingsApiService) GetOrganizationSensorReadingsLatestExecute(r ReadingsApiGetOrganizationSensorReadingsLatestRequest) ([]InlineResponse200270, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200270
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadingsApiService.GetOrganizationSensorReadingsLatest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sensor/readings/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.metrics != nil {
		localVarQueryParams.Add("metrics", parameterToString(*r.metrics, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["bearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
