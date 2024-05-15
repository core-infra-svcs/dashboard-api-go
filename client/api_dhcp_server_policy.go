/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
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


// DhcpServerPolicyApiService DhcpServerPolicyApi service
type DhcpServerPolicyApiService service

type DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject132
}

func (r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject132) DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*InlineResponse200143, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network

Add a server to be trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string) DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200143
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*InlineResponse200143, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200143
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
}

func (r DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network

Remove a server from being trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterToString(r.trustedServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest) Execute() (*InlineResponse200142, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicy Return the DHCP server settings

Return the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicy(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200142
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest) (*InlineResponse200142, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200142
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) PerPage(perPage int32) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) Execute() ([]InlineResponse200143, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network

Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as allow listed snoop entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200143
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) ([]InlineResponse200143, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200143
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) PerPage(perPage int32) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) Execute() ([]InlineResponse200144, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings

Return the devices that have a Dynamic ARP Inspection warning and their warnings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200144
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) ([]InlineResponse200144, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200144
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	updateNetworkSwitchDhcpServerPolicy *InlineObject131
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy InlineObject131) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	r.updateNetworkSwitchDhcpServerPolicy = &updateNetworkSwitchDhcpServerPolicy
	return r
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) Execute() (*InlineResponse200142, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicy Update the DHCP server settings

Update the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicy(ctx context.Context, networkId string) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200142
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) (*InlineResponse200142, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200142
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject133
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject133) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*InlineResponse200143, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network

Update a server that is trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
//  @return InlineResponse200143
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*InlineResponse200143, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200143
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterToString(r.trustedServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
