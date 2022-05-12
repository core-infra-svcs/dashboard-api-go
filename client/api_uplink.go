/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
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

// UplinkApiService UplinkApi service
type UplinkApiService service

type UplinkApiApiGetNetworkCellularGatewayUplinkRequest struct {
	ctx _context.Context
	ApiService *UplinkApiService
	networkId string
}


func (r UplinkApiApiGetNetworkCellularGatewayUplinkRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkCellularGatewayUplinkExecute(r)
}

/*
GetNetworkCellularGatewayUplink Returns the uplink settings for your MG network.

Returns the uplink settings for your MG network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return UplinkApiApiGetNetworkCellularGatewayUplinkRequest
*/
func (a *UplinkApiService) GetNetworkCellularGatewayUplink(ctx _context.Context, networkId string) UplinkApiApiGetNetworkCellularGatewayUplinkRequest {
	return UplinkApiApiGetNetworkCellularGatewayUplinkRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *UplinkApiService) GetNetworkCellularGatewayUplinkExecute(r UplinkApiApiGetNetworkCellularGatewayUplinkRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.GetNetworkCellularGatewayUplink")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/uplink"
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

type UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest struct {
	ctx _context.Context
	ApiService *UplinkApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	iccids *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) PerPage(perPage int32) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) StartingAfter(startingAfter string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) EndingBefore(endingBefore string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.endingBefore = &endingBefore
	return r
}
// A list of network IDs. The returned devices will be filtered to only include these networks.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) NetworkIds(networkIds []string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.networkIds = &networkIds
	return r
}
// A list of serial numbers. The returned devices will be filtered to only include these serials.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) Serials(serials []string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.serials = &serials
	return r
}
// A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) Iccids(iccids []string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	r.iccids = &iccids
	return r
}

func (r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationApplianceUplinkStatusesExecute(r)
}

/*
GetOrganizationApplianceUplinkStatuses List the uplink status of every Meraki MX and Z series appliances in the organization

List the uplink status of every Meraki MX and Z series appliances in the organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest
*/
func (a *UplinkApiService) GetOrganizationApplianceUplinkStatuses(ctx _context.Context, organizationId string) UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest {
	return UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *UplinkApiService) GetOrganizationApplianceUplinkStatusesExecute(r UplinkApiApiGetOrganizationApplianceUplinkStatusesRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.GetOrganizationApplianceUplinkStatuses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/uplink/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.iccids != nil {
		localVarQueryParams.Add("iccids", parameterToString(*r.iccids, "csv"))
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

type UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest struct {
	ctx _context.Context
	ApiService *UplinkApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	iccids *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) PerPage(perPage int32) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.perPage = &perPage
	return r
}
// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) StartingAfter(startingAfter string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.startingAfter = &startingAfter
	return r
}
// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) EndingBefore(endingBefore string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.endingBefore = &endingBefore
	return r
}
// A list of network IDs. The returned devices will be filtered to only include these networks.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) NetworkIds(networkIds []string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.networkIds = &networkIds
	return r
}
// A list of serial numbers. The returned devices will be filtered to only include these serials.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) Serials(serials []string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.serials = &serials
	return r
}
// A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) Iccids(iccids []string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.iccids = &iccids
	return r
}

func (r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationCellularGatewayUplinkStatusesExecute(r)
}

/*
GetOrganizationCellularGatewayUplinkStatuses List the uplink status of every Meraki MG cellular gateway in the organization

List the uplink status of every Meraki MG cellular gateway in the organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest
*/
func (a *UplinkApiService) GetOrganizationCellularGatewayUplinkStatuses(ctx _context.Context, organizationId string) UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	return UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *UplinkApiService) GetOrganizationCellularGatewayUplinkStatusesExecute(r UplinkApiApiGetOrganizationCellularGatewayUplinkStatusesRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.GetOrganizationCellularGatewayUplinkStatuses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/cellularGateway/uplink/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.iccids != nil {
		localVarQueryParams.Add("iccids", parameterToString(*r.iccids, "csv"))
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

type UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest struct {
	ctx _context.Context
	ApiService *UplinkApiService
	networkId string
	updateNetworkCellularGatewayUplink *InlineObject61
}

func (r UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest) UpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink InlineObject61) UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest {
	r.updateNetworkCellularGatewayUplink = &updateNetworkCellularGatewayUplink
	return r
}

func (r UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkCellularGatewayUplinkExecute(r)
}

/*
UpdateNetworkCellularGatewayUplink Updates the uplink settings for your MG network.

Updates the uplink settings for your MG network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest
*/
func (a *UplinkApiService) UpdateNetworkCellularGatewayUplink(ctx _context.Context, networkId string) UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest {
	return UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *UplinkApiService) UpdateNetworkCellularGatewayUplinkExecute(r UplinkApiApiUpdateNetworkCellularGatewayUplinkRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.UpdateNetworkCellularGatewayUplink")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/uplink"
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
	localVarPostBody = r.updateNetworkCellularGatewayUplink
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