/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
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


// SystemApiService SystemApi service
type SystemApiService service

type SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
	networkIds *[]string
	serials *[]string
	productTypes *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) PerPage(perPage int32) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) StartingAfter(startingAfter string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) EndingBefore(endingBefore string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) T0(t0 string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) T1(t1 string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) Timespan(timespan float32) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) Interval(interval int32) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.interval = &interval
	return r
}

// Optional parameter to filter the result set by the included set of network IDs
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) NetworkIds(networkIds []string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities history by device serial numbers
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) Serials(serials []string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect.
func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) ProductTypes(productTypes []string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	r.productTypes = &productTypes
	return r
}

func (r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) Execute() (*InlineResponse200281, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesSystemMemoryUsageHistoryByIntervalExecute(r)
}

/*
GetOrganizationDevicesSystemMemoryUsageHistoryByInterval Return the memory utilization history in kB for devices in the organization.

Return the memory utilization history in kB for devices in the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest
*/
func (a *SystemApiService) GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(ctx context.Context, organizationId string) SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest {
	return SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200281
func (a *SystemApiService) GetOrganizationDevicesSystemMemoryUsageHistoryByIntervalExecute(r SystemApiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest) (*InlineResponse200281, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200281
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/system/memory/usage/history/byInterval"
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
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
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

type SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	organizationId string
	serials *[]string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) Serials(serials []string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.serials = &serials
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) T0(t0 string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) T1(t1 string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) Timespan(timespan float32) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) PerPage(perPage int32) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) StartingAfter(startingAfter string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) EndingBefore(endingBefore string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) Execute() (*InlineResponse200381, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalExecute(r)
}

/*
GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval List cpu utilization data of wireless LAN controllers in an organization

List cpu utilization data of wireless LAN controllers in an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest
*/
func (a *SystemApiService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval(ctx context.Context, organizationId string) SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest {
	return SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200381
func (a *SystemApiService) GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalExecute(r SystemApiGetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByIntervalRequest) (*InlineResponse200381, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200381
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.GetOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wirelessController/devices/system/utilization/history/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
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
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
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

type SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
}

// The beginning of the timespan for the data. The maximum lookback period is 1 day from today.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) T0(t0 string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 1 day after t0.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) T1(t1 string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 day.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) Timespan(timespan float32) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 20. Default is 10.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) PerPage(perPage int32) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) StartingAfter(startingAfter string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) EndingBefore(endingBefore string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter the result set by the included set of network IDs
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) NetworkIds(networkIds []string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities history by device serial numbers
func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) Serials(serials []string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	r.serials = &serials
	return r
}

func (r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) Execute() (*InlineResponse200360, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesSystemCpuLoadHistoryExecute(r)
}

/*
GetOrganizationWirelessDevicesSystemCpuLoadHistory Return the CPU Load history for a list of wireless devices in the organization.

Return the CPU Load history for a list of wireless devices in the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest
*/
func (a *SystemApiService) GetOrganizationWirelessDevicesSystemCpuLoadHistory(ctx context.Context, organizationId string) SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest {
	return SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200360
func (a *SystemApiService) GetOrganizationWirelessDevicesSystemCpuLoadHistoryExecute(r SystemApiGetOrganizationWirelessDevicesSystemCpuLoadHistoryRequest) (*InlineResponse200360, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200360
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.GetOrganizationWirelessDevicesSystemCpuLoadHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/system/cpu/load/history"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
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
