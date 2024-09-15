/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
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


// AssignmentsApiService AssignmentsApi service
type AssignmentsApiService service

type AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest struct {
	ctx context.Context
	ApiService *AssignmentsApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	serials *[]string
	productTypes *[]string
	stackIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) PerPage(perPage int32) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) StartingAfter(startingAfter string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) EndingBefore(endingBefore string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) Serials(serials []string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter devices by product types.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) ProductTypes(productTypes []string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter devices by Switch Stack ids.
func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) StackIds(stackIds []string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	r.stackIds = &stackIds
	return r
}

func (r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) Execute() ([]InlineResponse200164, *http.Response, error) {
	return r.ApiService.GetNetworkVlanProfilesAssignmentsByDeviceExecute(r)
}

/*
GetNetworkVlanProfilesAssignmentsByDevice Get the assigned VLAN Profiles for devices in a network

Get the assigned VLAN Profiles for devices in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest
*/
func (a *AssignmentsApiService) GetNetworkVlanProfilesAssignmentsByDevice(ctx context.Context, networkId string) AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest {
	return AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse200164
func (a *AssignmentsApiService) GetNetworkVlanProfilesAssignmentsByDeviceExecute(r AssignmentsApiGetNetworkVlanProfilesAssignmentsByDeviceRequest) ([]InlineResponse200164, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200164
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsApiService.GetNetworkVlanProfilesAssignmentsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/vlanProfiles/assignments/byDevice"
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
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.stackIds != nil {
		localVarQueryParams.Add("stackIds", parameterToString(*r.stackIds, "csv"))
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

type AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest struct {
	ctx context.Context
	ApiService *AssignmentsApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) PerPage(perPage int32) AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) StartingAfter(startingAfter string) AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) EndingBefore(endingBefore string) AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter Sentry Policies by Network Id
func (r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) NetworkIds(networkIds []string) AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	r.networkIds = &networkIds
	return r
}

func (r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) Execute() ([]InlineResponse200280, *http.Response, error) {
	return r.ApiService.GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r)
}

/*
GetOrganizationSmSentryPoliciesAssignmentsByNetwork List the Sentry Policies for an organization ordered in ascending order of priority

List the Sentry Policies for an organization ordered in ascending order of priority

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest
*/
func (a *AssignmentsApiService) GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx context.Context, organizationId string) AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest {
	return AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200280
func (a *AssignmentsApiService) GetOrganizationSmSentryPoliciesAssignmentsByNetworkExecute(r AssignmentsApiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest) ([]InlineResponse200280, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200280
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsApiService.GetOrganizationSmSentryPoliciesAssignmentsByNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork"
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

type AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest struct {
	ctx context.Context
	ApiService *AssignmentsApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	name *string
	mac *string
	serial *string
	model *string
	macs *[]string
	serials *[]string
	models *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) PerPage(perPage int32) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) StartingAfter(startingAfter string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) EndingBefore(endingBefore string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter devices by network.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) NetworkIds(networkIds []string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) ProductTypes(productTypes []string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Name(name string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.name = &name
	return r
}

// Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Mac(mac string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.mac = &mac
	return r
}

// Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Serial(serial string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.serial = &serial
	return r
}

// Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Model(model string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.model = &model
	return r
}

// Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Macs(macs []string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Serials(serials []string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match.
func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Models(models []string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	r.models = &models
	return r
}

func (r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) Execute() ([]InlineResponse200313, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessRfProfilesAssignmentsByDeviceExecute(r)
}

/*
GetOrganizationWirelessRfProfilesAssignmentsByDevice List the RF profiles of an organization by device

List the RF profiles of an organization by device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest
*/
func (a *AssignmentsApiService) GetOrganizationWirelessRfProfilesAssignmentsByDevice(ctx context.Context, organizationId string) AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest {
	return AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200313
func (a *AssignmentsApiService) GetOrganizationWirelessRfProfilesAssignmentsByDeviceExecute(r AssignmentsApiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest) ([]InlineResponse200313, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200313
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsApiService.GetOrganizationWirelessRfProfilesAssignmentsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/rfProfiles/assignments/byDevice"
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
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.mac != nil {
		localVarQueryParams.Add("mac", parameterToString(*r.mac, ""))
	}
	if r.serial != nil {
		localVarQueryParams.Add("serial", parameterToString(*r.serial, ""))
	}
	if r.model != nil {
		localVarQueryParams.Add("model", parameterToString(*r.model, ""))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.models != nil {
		localVarQueryParams.Add("models", parameterToString(*r.models, "csv"))
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

type AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest struct {
	ctx context.Context
	ApiService *AssignmentsApiService
	networkId string
	reassignNetworkVlanProfilesAssignments *InlineObject163
}

func (r AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest) ReassignNetworkVlanProfilesAssignments(reassignNetworkVlanProfilesAssignments InlineObject163) AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest {
	r.reassignNetworkVlanProfilesAssignments = &reassignNetworkVlanProfilesAssignments
	return r
}

func (r AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest) Execute() (*InlineResponse200165, *http.Response, error) {
	return r.ApiService.ReassignNetworkVlanProfilesAssignmentsExecute(r)
}

/*
ReassignNetworkVlanProfilesAssignments Update the assigned VLAN Profile for devices in a network

Update the assigned VLAN Profile for devices in a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest
*/
func (a *AssignmentsApiService) ReassignNetworkVlanProfilesAssignments(ctx context.Context, networkId string) AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest {
	return AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse200165
func (a *AssignmentsApiService) ReassignNetworkVlanProfilesAssignmentsExecute(r AssignmentsApiReassignNetworkVlanProfilesAssignmentsRequest) (*InlineResponse200165, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200165
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsApiService.ReassignNetworkVlanProfilesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/vlanProfiles/assignments/reassign"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reassignNetworkVlanProfilesAssignments == nil {
		return localVarReturnValue, nil, reportError("reassignNetworkVlanProfilesAssignments is required and must be specified")
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
	localVarPostBody = r.reassignNetworkVlanProfilesAssignments
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

type AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest struct {
	ctx context.Context
	ApiService *AssignmentsApiService
	organizationId string
	updateOrganizationSmSentryPoliciesAssignments *InlineObject264
}

func (r AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments InlineObject264) AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	r.updateOrganizationSmSentryPoliciesAssignments = &updateOrganizationSmSentryPoliciesAssignments
	return r
}

func (r AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) Execute() (*InlineResponse200279, *http.Response, error) {
	return r.ApiService.UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r)
}

/*
UpdateOrganizationSmSentryPoliciesAssignments Update an Organizations Sentry Policies using the provided list

Update an Organizations Sentry Policies using the provided list. Sentry Policies are ordered in descending order of priority (i.e. highest priority at the bottom, this is opposite the Dashboard UI). Policies not present in the request will be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest
*/
func (a *AssignmentsApiService) UpdateOrganizationSmSentryPoliciesAssignments(ctx context.Context, organizationId string) AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest {
	return AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200279
func (a *AssignmentsApiService) UpdateOrganizationSmSentryPoliciesAssignmentsExecute(r AssignmentsApiUpdateOrganizationSmSentryPoliciesAssignmentsRequest) (*InlineResponse200279, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200279
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssignmentsApiService.UpdateOrganizationSmSentryPoliciesAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/sm/sentry/policies/assignments"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateOrganizationSmSentryPoliciesAssignments == nil {
		return localVarReturnValue, nil, reportError("updateOrganizationSmSentryPoliciesAssignments is required and must be specified")
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
	localVarPostBody = r.updateOrganizationSmSentryPoliciesAssignments
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
