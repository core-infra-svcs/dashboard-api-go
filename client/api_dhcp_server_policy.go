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

// DhcpServerPolicyApiService DhcpServerPolicyApi service
type DhcpServerPolicyApiService service

type DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject114
}

func (r DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject114) DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (InlineResponse20048, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network

Add a server to be trusted by Dynamic ARP Inspection on this network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx _context.Context, networkId string) DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20048
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20048
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
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

type DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
}


func (r DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network

Remove a server from being trusted by Dynamic ARP Inspection on this network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param trustedServerId
 @return DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx _context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", _neturl.PathEscape(parameterToString(r.trustedServerId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
}


func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicy Return the DHCP server settings

Return the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicy(ctx _context.Context, networkId string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
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

type DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) PerPage(perPage int32) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) Execute() ([]InlineResponse20048, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network

Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as whitelisted snoop entries

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx _context.Context, networkId string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20048
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) ([]InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20048
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) PerPage(perPage int32) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) Execute() ([]InlineResponse20049, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings

Return the devices that have a Dynamic ARP Inspection warning and their warnings

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx _context.Context, networkId string) DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	return DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20049
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r DhcpServerPolicyApiApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) ([]InlineResponse20049, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse20049
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	updateNetworkSwitchDhcpServerPolicy *InlineObject113
}

func (r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest) UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy InlineObject113) DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	r.updateNetworkSwitchDhcpServerPolicy = &updateNetworkSwitchDhcpServerPolicy
	return r
}

func (r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicy Update the DHCP server settings

Update the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicy(ctx _context.Context, networkId string) DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateNetworkSwitchDhcpServerPolicy
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

type DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx _context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject115
}

func (r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject115) DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (InlineResponse20048, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network

Update a server that is trusted by Dynamic ARP Inspection on this network

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param trustedServerId
 @return DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx _context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
//  @return InlineResponse20048
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (InlineResponse20048, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20048
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", _neturl.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", _neturl.PathEscape(parameterToString(r.trustedServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
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
