/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
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


// LatencyHistoryApiService LatencyHistoryApi service
type LatencyHistoryApiService service

type LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest struct {
	ctx context.Context
	ApiService *LatencyHistoryApiService
	networkId string
	clientId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 791 days from today.
func (r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) T0(t0 string) LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 791 days after t0.
func (r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) T1(t1 string) LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day.
func (r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) Timespan(timespan float32) LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400.
func (r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) Resolution(resolution int32) LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest {
	r.resolution = &resolution
	return r
}

func (r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientLatencyHistoryExecute(r)
}

/*
GetNetworkWirelessClientLatencyHistory Return the latency history for a client

Return the latency history for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP. The latency data is from a sample of 2% of packets and is grouped into 4 traffic categories: background, best effort, video, voice. Within these categories the sampled packet counters are bucketed by latency in milliseconds.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest
*/
func (a *LatencyHistoryApiService) GetNetworkWirelessClientLatencyHistory(ctx context.Context, networkId string, clientId string) LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest {
	return LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LatencyHistoryApiService) GetNetworkWirelessClientLatencyHistoryExecute(r LatencyHistoryApiGetNetworkWirelessClientLatencyHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyHistoryApiService.GetNetworkWirelessClientLatencyHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/latencyHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

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
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
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

type LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest struct {
	ctx context.Context
	ApiService *LatencyHistoryApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	autoResolution *bool
	clientId *string
	deviceSerial *string
	apTag *string
	band *string
	ssid *int32
	accessCategory *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) T0(t0 string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) T1(t1 string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) Timespan(timespan float32) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) Resolution(resolution int32) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.resolution = &resolution
	return r
}

// Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) AutoResolution(autoResolution bool) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.autoResolution = &autoResolution
	return r
}

// Filter results by network client.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) ClientId(clientId string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.clientId = &clientId
	return r
}

// Filter results by device.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) DeviceSerial(deviceSerial string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.deviceSerial = &deviceSerial
	return r
}

// Filter results by AP tag.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) ApTag(apTag string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.apTag = &apTag
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;).
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) Band(band string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.band = &band
	return r
}

// Filter results by SSID number.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) Ssid(ssid int32) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.ssid = &ssid
	return r
}

// Filter by access category.
func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) AccessCategory(accessCategory string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	r.accessCategory = &accessCategory
	return r
}

func (r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) Execute() ([]InlineResponse200109, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessLatencyHistoryExecute(r)
}

/*
GetNetworkWirelessLatencyHistory Return average wireless latency over time for a network, device, or network client

Return average wireless latency over time for a network, device, or network client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest
*/
func (a *LatencyHistoryApiService) GetNetworkWirelessLatencyHistory(ctx context.Context, networkId string) LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest {
	return LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200109
func (a *LatencyHistoryApiService) GetNetworkWirelessLatencyHistoryExecute(r LatencyHistoryApiGetNetworkWirelessLatencyHistoryRequest) ([]InlineResponse200109, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200109
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyHistoryApiService.GetNetworkWirelessLatencyHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/latencyHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
	}
	if r.autoResolution != nil {
		localVarQueryParams.Add("autoResolution", parameterToString(*r.autoResolution, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("clientId", parameterToString(*r.clientId, ""))
	}
	if r.deviceSerial != nil {
		localVarQueryParams.Add("deviceSerial", parameterToString(*r.deviceSerial, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.accessCategory != nil {
		localVarQueryParams.Add("accessCategory", parameterToString(*r.accessCategory, ""))
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
