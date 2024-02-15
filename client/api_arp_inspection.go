/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
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


// ArpInspectionApiService ArpInspectionApi service
type ArpInspectionApiService service

type ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *ArpInspectionApiService
	networkId string
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject129
}

func (r ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject129) ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*InlineResponse20087, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network

Add a server to be trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *ArpInspectionApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string) ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20087
func (a *ArpInspectionApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r ArpInspectionApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*InlineResponse20087, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20087
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArpInspectionApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
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

type ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *ArpInspectionApiService
	networkId string
	trustedServerId string
}

func (r ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network

Remove a server from being trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *ArpInspectionApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
func (a *ArpInspectionApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r ArpInspectionApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArpInspectionApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
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

type ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct {
	ctx context.Context
	ApiService *ArpInspectionApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) PerPage(perPage int32) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) StartingAfter(startingAfter string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) EndingBefore(endingBefore string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) Execute() ([]InlineResponse20087, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network

Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as allow listed snoop entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest
*/
func (a *ArpInspectionApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx context.Context, networkId string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	return ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20087
func (a *ArpInspectionApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) ([]InlineResponse20087, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20087
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArpInspectionApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
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

type ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct {
	ctx context.Context
	ApiService *ArpInspectionApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) PerPage(perPage int32) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) StartingAfter(startingAfter string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) EndingBefore(endingBefore string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) Execute() ([]InlineResponse20088, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings

Return the devices that have a Dynamic ARP Inspection warning and their warnings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest
*/
func (a *ArpInspectionApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx context.Context, networkId string) ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	return ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20088
func (a *ArpInspectionApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r ArpInspectionApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) ([]InlineResponse20088, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20088
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArpInspectionApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
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

type ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *ArpInspectionApiService
	networkId string
	trustedServerId string
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer *InlineObject130
}

func (r ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer InlineObject130) ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer = &updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer
	return r
}

func (r ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*InlineResponse20087, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network

Update a server that is trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *ArpInspectionApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
//  @return InlineResponse20087
func (a *ArpInspectionApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r ArpInspectionApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*InlineResponse20087, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20087
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArpInspectionApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
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
