/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
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


// L2ApiService L2Api service
type L2ApiService service

type L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest struct {
	ctx context.Context
	ApiService *L2ApiService
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
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) Serials(serials []string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.serials = &serials
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) T0(t0 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) T1(t1 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) Timespan(timespan float32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) PerPage(perPage int32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) StartingAfter(startingAfter string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) EndingBefore(endingBefore string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) Execute() (*InlineResponse200332, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceExecute(r)
}

/*
GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice List wireless LAN controller layer 2 interfaces in an organization

List wireless LAN controller layer 2 interfaces in an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest
*/
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice(ctx context.Context, organizationId string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest {
	return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200332
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceExecute(r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2ByDeviceRequest) (*InlineResponse200332, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200332
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L2ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2ByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wirelessController/devices/interfaces/l2/byDevice"
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

type L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest struct {
	ctx context.Context
	ApiService *L2ApiService
	organizationId string
	serials *[]string
	includeInterfacesWithoutChanges *bool
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) Serials(serials []string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.serials = &serials
	return r
}

// By default, interfaces without changes are omitted from the response for brevity. If you want to include the interfaces even if they have no changes, set to true. (default: false)
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) IncludeInterfacesWithoutChanges(includeInterfacesWithoutChanges bool) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.includeInterfacesWithoutChanges = &includeInterfacesWithoutChanges
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) T0(t0 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) T1(t1 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) Timespan(timespan float32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) PerPage(perPage int32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) StartingAfter(startingAfter string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) EndingBefore(endingBefore string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) Execute() (*InlineResponse200333, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceExecute(r)
}

/*
GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice List wireless LAN controller layer 2 interfaces history status in an organization

List wireless LAN controller layer 2 interfaces history status in an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest
*/
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice(ctx context.Context, organizationId string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest {
	return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200333
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceExecute(r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceRequest) (*InlineResponse200333, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200333
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L2ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wirelessController/devices/interfaces/l2/statuses/changeHistory/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.includeInterfacesWithoutChanges != nil {
		localVarQueryParams.Add("includeInterfacesWithoutChanges", parameterToString(*r.includeInterfacesWithoutChanges, ""))
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

type L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest struct {
	ctx context.Context
	ApiService *L2ApiService
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
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) Serials(serials []string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.serials = &serials
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) T0(t0 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) T1(t1 string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) Timespan(timespan float32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) PerPage(perPage int32) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) StartingAfter(startingAfter string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) EndingBefore(endingBefore string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) Execute() (*InlineResponse200334, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalExecute(r)
}

/*
GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval List wireless LAN controller layer 2 interfaces history usage in an organization

List wireless LAN controller layer 2 interfaces history usage in an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest
*/
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(ctx context.Context, organizationId string) L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest {
	return L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200334
func (a *L2ApiService) GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalExecute(r L2ApiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest) (*InlineResponse200334, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200334
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "L2ApiService.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wirelessController/devices/interfaces/l2/usage/history/byInterval"
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
