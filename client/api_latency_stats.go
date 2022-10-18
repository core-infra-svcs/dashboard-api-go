/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
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

// LatencyStatsApiService LatencyStatsApi service
type LatencyStatsApiService service

type LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest struct {
	ctx _context.Context
	ApiService *LatencyStatsApiService
	serial string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) T0(t0 string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) T1(t1 string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.timespan = &timespan
	return r
}
// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Band(band string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.band = &band
	return r
}
// Filter results by SSID
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.ssid = &ssid
	return r
}
// Filter results by VLAN
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.vlan = &vlan
	return r
}
// Filter results by AP Tag
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.apTag = &apTag
	return r
}
// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Fields(fields string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceWirelessLatencyStatsExecute(r)
}

/*
GetDeviceWirelessLatencyStats Aggregated latency info for a given AP on this network

Aggregated latency info for a given AP on this network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetDeviceWirelessLatencyStats(ctx _context.Context, serial string) LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest {
	return LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetDeviceWirelessLatencyStatsExecute(r LatencyStatsApiApiGetDeviceWirelessLatencyStatsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetDeviceWirelessLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/wireless/latencyStats"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)

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
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest struct {
	ctx _context.Context
	ApiService *LatencyStatsApiService
	networkId string
	clientId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) T0(t0 string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) T1(t1 string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.timespan = &timespan
	return r
}
// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Band(band string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.band = &band
	return r
}
// Filter results by SSID
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.ssid = &ssid
	return r
}
// Filter results by VLAN
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.vlan = &vlan
	return r
}
// Filter results by AP Tag
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.apTag = &apTag
	return r
}
// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Fields(fields string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessClientLatencyStatsExecute(r)
}

/*
GetNetworkWirelessClientLatencyStats Aggregated latency info for a given client on this network

Aggregated latency info for a given client on this network. Clients are identified by their MAC.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param clientId
 @return LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessClientLatencyStats(ctx _context.Context, networkId string, clientId string) LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest {
	return LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessClientLatencyStatsExecute(r LatencyStatsApiApiGetNetworkWirelessClientLatencyStatsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessClientLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/latencyStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", _neturl.PathEscape(parameterToString(r.clientId, "")), -1)

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
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest struct {
	ctx _context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) T0(t0 string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) T1(t1 string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.timespan = &timespan
	return r
}
// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Band(band string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.band = &band
	return r
}
// Filter results by SSID
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.ssid = &ssid
	return r
}
// Filter results by VLAN
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.vlan = &vlan
	return r
}
// Filter results by AP Tag
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.apTag = &apTag
	return r
}
// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Fields(fields string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessClientsLatencyStatsExecute(r)
}

/*
GetNetworkWirelessClientsLatencyStats Aggregated latency info for this network, grouped by clients

Aggregated latency info for this network, grouped by clients

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessClientsLatencyStats(ctx _context.Context, networkId string) LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest {
	return LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessClientsLatencyStatsExecute(r LatencyStatsApiApiGetNetworkWirelessClientsLatencyStatsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessClientsLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/latencyStats"
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
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest struct {
	ctx _context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) T0(t0 string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) T1(t1 string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.timespan = &timespan
	return r
}
// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Band(band string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.band = &band
	return r
}
// Filter results by SSID
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.ssid = &ssid
	return r
}
// Filter results by VLAN
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.vlan = &vlan
	return r
}
// Filter results by AP Tag
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.apTag = &apTag
	return r
}
// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Fields(fields string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessDevicesLatencyStatsExecute(r)
}

/*
GetNetworkWirelessDevicesLatencyStats Aggregated latency info for this network, grouped by node

Aggregated latency info for this network, grouped by node

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessDevicesLatencyStats(ctx _context.Context, networkId string) LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest {
	return LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessDevicesLatencyStatsExecute(r LatencyStatsApiApiGetNetworkWirelessDevicesLatencyStatsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessDevicesLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/devices/latencyStats"
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
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest struct {
	ctx _context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) T0(t0 string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) T1(t1 string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.timespan = &timespan
	return r
}
// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Band(band string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.band = &band
	return r
}
// Filter results by SSID
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.ssid = &ssid
	return r
}
// Filter results by VLAN
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.vlan = &vlan
	return r
}
// Filter results by AP Tag
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.apTag = &apTag
	return r
}
// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Fields(fields string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkWirelessLatencyStatsExecute(r)
}

/*
GetNetworkWirelessLatencyStats Aggregated latency info for this network

Aggregated latency info for this network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessLatencyStats(ctx _context.Context, networkId string) LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest {
	return LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessLatencyStatsExecute(r LatencyStatsApiApiGetNetworkWirelessLatencyStatsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/latencyStats"
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
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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
