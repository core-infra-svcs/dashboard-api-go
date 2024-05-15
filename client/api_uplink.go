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


// UplinkApiService UplinkApi service
type UplinkApiService service

type UplinkApiGetNetworkCellularGatewayUplinkRequest struct {
	ctx context.Context
	ApiService *UplinkApiService
	networkId string
}

func (r UplinkApiGetNetworkCellularGatewayUplinkRequest) Execute() (*InlineResponse20076, *http.Response, error) {
	return r.ApiService.GetNetworkCellularGatewayUplinkExecute(r)
}

/*
GetNetworkCellularGatewayUplink Returns the uplink settings for your MG network.

Returns the uplink settings for your MG network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return UplinkApiGetNetworkCellularGatewayUplinkRequest
*/
func (a *UplinkApiService) GetNetworkCellularGatewayUplink(ctx context.Context, networkId string) UplinkApiGetNetworkCellularGatewayUplinkRequest {
	return UplinkApiGetNetworkCellularGatewayUplinkRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20076
func (a *UplinkApiService) GetNetworkCellularGatewayUplinkExecute(r UplinkApiGetNetworkCellularGatewayUplinkRequest) (*InlineResponse20076, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20076
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.GetNetworkCellularGatewayUplink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/uplink"
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

type UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest struct {
	ctx context.Context
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
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) PerPage(perPage int32) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) StartingAfter(startingAfter string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) EndingBefore(endingBefore string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.endingBefore = &endingBefore
	return r
}

// A list of network IDs. The returned devices will be filtered to only include these networks.
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) NetworkIds(networkIds []string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.networkIds = &networkIds
	return r
}

// A list of serial numbers. The returned devices will be filtered to only include these serials.
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) Serials(serials []string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.serials = &serials
	return r
}

// A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.
func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) Iccids(iccids []string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	r.iccids = &iccids
	return r
}

func (r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) Execute() ([]InlineResponse200224, *http.Response, error) {
	return r.ApiService.GetOrganizationCellularGatewayUplinkStatusesExecute(r)
}

/*
GetOrganizationCellularGatewayUplinkStatuses List the uplink status of every Meraki MG cellular gateway in the organization

List the uplink status of every Meraki MG cellular gateway in the organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest
*/
func (a *UplinkApiService) GetOrganizationCellularGatewayUplinkStatuses(ctx context.Context, organizationId string) UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest {
	return UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200224
func (a *UplinkApiService) GetOrganizationCellularGatewayUplinkStatusesExecute(r UplinkApiGetOrganizationCellularGatewayUplinkStatusesRequest) ([]InlineResponse200224, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200224
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.GetOrganizationCellularGatewayUplinkStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/cellularGateway/uplink/statuses"
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

type UplinkApiUpdateNetworkCellularGatewayUplinkRequest struct {
	ctx context.Context
	ApiService *UplinkApiService
	networkId string
	updateNetworkCellularGatewayUplink *InlineObject84
}

func (r UplinkApiUpdateNetworkCellularGatewayUplinkRequest) UpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink InlineObject84) UplinkApiUpdateNetworkCellularGatewayUplinkRequest {
	r.updateNetworkCellularGatewayUplink = &updateNetworkCellularGatewayUplink
	return r
}

func (r UplinkApiUpdateNetworkCellularGatewayUplinkRequest) Execute() (*InlineResponse20076, *http.Response, error) {
	return r.ApiService.UpdateNetworkCellularGatewayUplinkExecute(r)
}

/*
UpdateNetworkCellularGatewayUplink Updates the uplink settings for your MG network.

Updates the uplink settings for your MG network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return UplinkApiUpdateNetworkCellularGatewayUplinkRequest
*/
func (a *UplinkApiService) UpdateNetworkCellularGatewayUplink(ctx context.Context, networkId string) UplinkApiUpdateNetworkCellularGatewayUplinkRequest {
	return UplinkApiUpdateNetworkCellularGatewayUplinkRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20076
func (a *UplinkApiService) UpdateNetworkCellularGatewayUplinkExecute(r UplinkApiUpdateNetworkCellularGatewayUplinkRequest) (*InlineResponse20076, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20076
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinkApiService.UpdateNetworkCellularGatewayUplink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/uplink"
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
	localVarPostBody = r.updateNetworkCellularGatewayUplink
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
