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

// EventsApiService EventsApi service
type EventsApiService service

type EventsApiApiGetNetworkApplianceClientSecurityEventsRequest struct {
	ctx _context.Context
	ApiService *EventsApiService
	networkId string
	clientId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	sortOrder *string
}

// The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) T0(t0 string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 791 days after t0.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) T1(t1 string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) Timespan(timespan float32) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.timespan = &timespan
	return r
}
// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) PerPage(perPage int32) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) StartingAfter(startingAfter string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) EndingBefore(endingBefore string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.endingBefore = &endingBefore
	return r
}
// Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order.
func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) SortOrder(sortOrder string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkApplianceClientSecurityEventsExecute(r)
}

/*
GetNetworkApplianceClientSecurityEvents List the security events for a client

List the security events for a client. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param clientId
 @return EventsApiApiGetNetworkApplianceClientSecurityEventsRequest
*/
func (a *EventsApiService) GetNetworkApplianceClientSecurityEvents(ctx _context.Context, networkId string, clientId string) EventsApiApiGetNetworkApplianceClientSecurityEventsRequest {
	return EventsApiApiGetNetworkApplianceClientSecurityEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *EventsApiService) GetNetworkApplianceClientSecurityEventsExecute(r EventsApiApiGetNetworkApplianceClientSecurityEventsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetNetworkApplianceClientSecurityEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/clients/{clientId}/security/events"
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
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
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

type EventsApiApiGetNetworkApplianceSecurityEventsRequest struct {
	ctx _context.Context
	ApiService *EventsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	sortOrder *string
}

// The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) T0(t0 string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) T1(t1 string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) Timespan(timespan float32) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.timespan = &timespan
	return r
}
// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) PerPage(perPage int32) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) StartingAfter(startingAfter string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) EndingBefore(endingBefore string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.endingBefore = &endingBefore
	return r
}
// Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order.
func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) SortOrder(sortOrder string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r EventsApiApiGetNetworkApplianceSecurityEventsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkApplianceSecurityEventsExecute(r)
}

/*
GetNetworkApplianceSecurityEvents List the security events for a network

List the security events for a network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return EventsApiApiGetNetworkApplianceSecurityEventsRequest
*/
func (a *EventsApiService) GetNetworkApplianceSecurityEvents(ctx _context.Context, networkId string) EventsApiApiGetNetworkApplianceSecurityEventsRequest {
	return EventsApiApiGetNetworkApplianceSecurityEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *EventsApiService) GetNetworkApplianceSecurityEventsExecute(r EventsApiApiGetNetworkApplianceSecurityEventsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetNetworkApplianceSecurityEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/security/events"
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
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
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

type EventsApiApiGetNetworkEventsRequest struct {
	ctx _context.Context
	ApiService *EventsApiService
	networkId string
	productType *string
	includedEventTypes *[]string
	excludedEventTypes *[]string
	deviceMac *string
	deviceSerial *string
	deviceName *string
	clientIp *string
	clientMac *string
	clientName *string
	smDeviceMac *string
	smDeviceName *string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, and cellularGateway
func (r EventsApiApiGetNetworkEventsRequest) ProductType(productType string) EventsApiApiGetNetworkEventsRequest {
	r.productType = &productType
	return r
}
// A list of event types. The returned events will be filtered to only include events with these types.
func (r EventsApiApiGetNetworkEventsRequest) IncludedEventTypes(includedEventTypes []string) EventsApiApiGetNetworkEventsRequest {
	r.includedEventTypes = &includedEventTypes
	return r
}
// A list of event types. The returned events will be filtered to exclude events with these types.
func (r EventsApiApiGetNetworkEventsRequest) ExcludedEventTypes(excludedEventTypes []string) EventsApiApiGetNetworkEventsRequest {
	r.excludedEventTypes = &excludedEventTypes
	return r
}
// The MAC address of the Meraki device which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) DeviceMac(deviceMac string) EventsApiApiGetNetworkEventsRequest {
	r.deviceMac = &deviceMac
	return r
}
// The serial of the Meraki device which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) DeviceSerial(deviceSerial string) EventsApiApiGetNetworkEventsRequest {
	r.deviceSerial = &deviceSerial
	return r
}
// The name of the Meraki device which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) DeviceName(deviceName string) EventsApiApiGetNetworkEventsRequest {
	r.deviceName = &deviceName
	return r
}
// The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks.
func (r EventsApiApiGetNetworkEventsRequest) ClientIp(clientIp string) EventsApiApiGetNetworkEventsRequest {
	r.clientIp = &clientIp
	return r
}
// The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks.
func (r EventsApiApiGetNetworkEventsRequest) ClientMac(clientMac string) EventsApiApiGetNetworkEventsRequest {
	r.clientMac = &clientMac
	return r
}
// The name, or partial name, of the client which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) ClientName(clientName string) EventsApiApiGetNetworkEventsRequest {
	r.clientName = &clientName
	return r
}
// The MAC address of the Systems Manager device which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) SmDeviceMac(smDeviceMac string) EventsApiApiGetNetworkEventsRequest {
	r.smDeviceMac = &smDeviceMac
	return r
}
// The name of the Systems Manager device which the list of events will be filtered with
func (r EventsApiApiGetNetworkEventsRequest) SmDeviceName(smDeviceName string) EventsApiApiGetNetworkEventsRequest {
	r.smDeviceName = &smDeviceName
	return r
}
// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
func (r EventsApiApiGetNetworkEventsRequest) PerPage(perPage int32) EventsApiApiGetNetworkEventsRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkEventsRequest) StartingAfter(startingAfter string) EventsApiApiGetNetworkEventsRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetNetworkEventsRequest) EndingBefore(endingBefore string) EventsApiApiGetNetworkEventsRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r EventsApiApiGetNetworkEventsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkEventsExecute(r)
}

/*
GetNetworkEvents List the events for the network

List the events for the network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return EventsApiApiGetNetworkEventsRequest
*/
func (a *EventsApiService) GetNetworkEvents(ctx _context.Context, networkId string) EventsApiApiGetNetworkEventsRequest {
	return EventsApiApiGetNetworkEventsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *EventsApiService) GetNetworkEventsExecute(r EventsApiApiGetNetworkEventsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetNetworkEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.productType != nil {
		localVarQueryParams.Add("productType", parameterToString(*r.productType, ""))
	}
	if r.includedEventTypes != nil {
		localVarQueryParams.Add("includedEventTypes", parameterToString(*r.includedEventTypes, "csv"))
	}
	if r.excludedEventTypes != nil {
		localVarQueryParams.Add("excludedEventTypes", parameterToString(*r.excludedEventTypes, "csv"))
	}
	if r.deviceMac != nil {
		localVarQueryParams.Add("deviceMac", parameterToString(*r.deviceMac, ""))
	}
	if r.deviceSerial != nil {
		localVarQueryParams.Add("deviceSerial", parameterToString(*r.deviceSerial, ""))
	}
	if r.deviceName != nil {
		localVarQueryParams.Add("deviceName", parameterToString(*r.deviceName, ""))
	}
	if r.clientIp != nil {
		localVarQueryParams.Add("clientIp", parameterToString(*r.clientIp, ""))
	}
	if r.clientMac != nil {
		localVarQueryParams.Add("clientMac", parameterToString(*r.clientMac, ""))
	}
	if r.clientName != nil {
		localVarQueryParams.Add("clientName", parameterToString(*r.clientName, ""))
	}
	if r.smDeviceMac != nil {
		localVarQueryParams.Add("smDeviceMac", parameterToString(*r.smDeviceMac, ""))
	}
	if r.smDeviceName != nil {
		localVarQueryParams.Add("smDeviceName", parameterToString(*r.smDeviceName, ""))
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

type EventsApiApiGetNetworkEventsEventTypesRequest struct {
	ctx _context.Context
	ApiService *EventsApiService
	networkId string
}


func (r EventsApiApiGetNetworkEventsEventTypesRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkEventsEventTypesExecute(r)
}

/*
GetNetworkEventsEventTypes List the event type to human-readable description

List the event type to human-readable description

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return EventsApiApiGetNetworkEventsEventTypesRequest
*/
func (a *EventsApiService) GetNetworkEventsEventTypes(ctx _context.Context, networkId string) EventsApiApiGetNetworkEventsEventTypesRequest {
	return EventsApiApiGetNetworkEventsEventTypesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *EventsApiService) GetNetworkEventsEventTypesExecute(r EventsApiApiGetNetworkEventsEventTypesRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetNetworkEventsEventTypes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/events/eventTypes"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type EventsApiApiGetOrganizationApplianceSecurityEventsRequest struct {
	ctx _context.Context
	ApiService *EventsApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	sortOrder *string
}

// The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) T0(t0 string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 365 days after t0.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) T1(t1 string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) Timespan(timespan float32) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.timespan = &timespan
	return r
}
// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) PerPage(perPage int32) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) StartingAfter(startingAfter string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) EndingBefore(endingBefore string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.endingBefore = &endingBefore
	return r
}
// Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order.
func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) SortOrder(sortOrder string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationApplianceSecurityEventsExecute(r)
}

/*
GetOrganizationApplianceSecurityEvents List the security events for an organization

List the security events for an organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return EventsApiApiGetOrganizationApplianceSecurityEventsRequest
*/
func (a *EventsApiService) GetOrganizationApplianceSecurityEvents(ctx _context.Context, organizationId string) EventsApiApiGetOrganizationApplianceSecurityEventsRequest {
	return EventsApiApiGetOrganizationApplianceSecurityEventsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *EventsApiService) GetOrganizationApplianceSecurityEventsExecute(r EventsApiApiGetOrganizationApplianceSecurityEventsRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.GetOrganizationApplianceSecurityEvents")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/security/events"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
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
