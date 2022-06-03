/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// NetworkHealthApiService NetworkHealthApi service
type NetworkHealthApiService service

type NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest struct {
	ctx _context.Context
	ApiService *NetworkHealthApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) T0(t0 string) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) T1(t1 string) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) Timespan(timespan float32) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.timespan = &timespan
	return r
}
// The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) Resolution(resolution int32) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.resolution = &resolution
	return r
}
// The number of entries per page returned. Acceptable range is 3 - 100. Default is 10.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) PerPage(perPage int32) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) StartingAfter(startingAfter string) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) EndingBefore(endingBefore string) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkNetworkHealthChannelUtilizationExecute(r)
}

/*
GetNetworkNetworkHealthChannelUtilization Get the channel utilization over each radio for all APs in a network.

Get the channel utilization over each radio for all APs in a network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest
*/
func (a *NetworkHealthApiService) GetNetworkNetworkHealthChannelUtilization(ctx _context.Context, networkId string) NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest {
	return NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *NetworkHealthApiService) GetNetworkNetworkHealthChannelUtilizationExecute(r NetworkHealthApiApiGetNetworkNetworkHealthChannelUtilizationRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkHealthApiService.GetNetworkNetworkHealthChannelUtilization")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/networkHealth/channelUtilization"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
