/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
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


// InventoryApiService InventoryApi service
type InventoryApiService service

type InventoryApiClaimIntoOrganizationInventoryRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	claimIntoOrganizationInventory *InlineObject197
}

func (r InventoryApiClaimIntoOrganizationInventoryRequest) ClaimIntoOrganizationInventory(claimIntoOrganizationInventory InlineObject197) InventoryApiClaimIntoOrganizationInventoryRequest {
	r.claimIntoOrganizationInventory = &claimIntoOrganizationInventory
	return r
}

func (r InventoryApiClaimIntoOrganizationInventoryRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ClaimIntoOrganizationInventoryExecute(r)
}

/*
ClaimIntoOrganizationInventory Claim a list of devices, licenses, and/or orders into an organization inventory

Claim a list of devices, licenses, and/or orders into an organization inventory. When claiming by order, all devices and licenses in the order will be claimed; licenses will be added to the organization and devices will be placed in the organization's inventory. Use /organizations/{organizationId}/inventory/release to release devices from an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiClaimIntoOrganizationInventoryRequest
*/
func (a *InventoryApiService) ClaimIntoOrganizationInventory(ctx context.Context, organizationId string) InventoryApiClaimIntoOrganizationInventoryRequest {
	return InventoryApiClaimIntoOrganizationInventoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InventoryApiService) ClaimIntoOrganizationInventoryExecute(r InventoryApiClaimIntoOrganizationInventoryRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.ClaimIntoOrganizationInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/claim"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.claimIntoOrganizationInventory
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

type InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringImport *InlineObject198
}

func (r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport InlineObject198) InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringImport = &createOrganizationInventoryOnboardingCloudMonitoringImport
	return r
}

func (r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Execute() ([]InlineResponse2016, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
*/
func (a *InventoryApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx context.Context, organizationId string) InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2016
func (a *InventoryApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ([]InlineResponse2016, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2016
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringImport == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringImport is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringImport
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

type InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringPrepare *InlineObject199
}

func (r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare InlineObject199) InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringPrepare = &createOrganizationInventoryOnboardingCloudMonitoringPrepare
	return r
}

func (r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Execute() ([]InlineResponse2017, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringPrepare Initiates or updates an import session

Initiates or updates an import session. An import ID will be generated and used when you are ready to commit the import.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
*/
func (a *InventoryApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx context.Context, organizationId string) InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2017
func (a *InventoryApiService) CreateOrganizationInventoryOnboardingCloudMonitoringPrepareExecute(r InventoryApiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ([]InlineResponse2017, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse2017
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringPrepare == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringPrepare is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringPrepare
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

type InventoryApiGetOrganizationInventoryDeviceRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	serial string
}

func (r InventoryApiGetOrganizationInventoryDeviceRequest) Execute() (*InlineResponse20094, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryDeviceExecute(r)
}

/*
GetOrganizationInventoryDevice Return a single device from the inventory of an organization

Return a single device from the inventory of an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @param serial
 @return InventoryApiGetOrganizationInventoryDeviceRequest
*/
func (a *InventoryApiService) GetOrganizationInventoryDevice(ctx context.Context, organizationId string, serial string) InventoryApiGetOrganizationInventoryDeviceRequest {
	return InventoryApiGetOrganizationInventoryDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		serial: serial,
	}
}

// Execute executes the request
//  @return InlineResponse20094
func (a *InventoryApiService) GetOrganizationInventoryDeviceExecute(r InventoryApiGetOrganizationInventoryDeviceRequest) (*InlineResponse20094, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20094
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.GetOrganizationInventoryDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/devices/{serial}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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

type InventoryApiGetOrganizationInventoryDevicesRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	usedState *string
	search *string
	macs *[]string
	networkIds *[]string
	serials *[]string
	models *[]string
	tags *[]string
	tagsFilterType *string
	productTypes *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) PerPage(perPage int32) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) StartingAfter(startingAfter string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) EndingBefore(endingBefore string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) UsedState(usedState string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.usedState = &usedState
	return r
}

// Search for devices in inventory based on serial number, mac address, or model.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) Search(search string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.search = &search
	return r
}

// Search for devices in inventory based on mac addresses.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) Macs(macs []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.macs = &macs
	return r
}

// Search for devices in inventory based on network ids.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) NetworkIds(networkIds []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.networkIds = &networkIds
	return r
}

// Search for devices in inventory based on serials.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) Serials(serials []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.serials = &serials
	return r
}

// Search for devices in inventory based on model.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) Models(models []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.models = &models
	return r
}

// Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below).
func (r InventoryApiGetOrganizationInventoryDevicesRequest) Tags(tags []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.tags = &tags
	return r
}

// To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) TagsFilterType(tagsFilterType string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

// Filter devices by product type. Accepted values are appliance, camera, cellularGateway, sensor, switch, systemsManager, and wireless.
func (r InventoryApiGetOrganizationInventoryDevicesRequest) ProductTypes(productTypes []string) InventoryApiGetOrganizationInventoryDevicesRequest {
	r.productTypes = &productTypes
	return r
}

func (r InventoryApiGetOrganizationInventoryDevicesRequest) Execute() ([]InlineResponse20094, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryDevicesExecute(r)
}

/*
GetOrganizationInventoryDevices Return the device inventory for an organization

Return the device inventory for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiGetOrganizationInventoryDevicesRequest
*/
func (a *InventoryApiService) GetOrganizationInventoryDevices(ctx context.Context, organizationId string) InventoryApiGetOrganizationInventoryDevicesRequest {
	return InventoryApiGetOrganizationInventoryDevicesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20094
func (a *InventoryApiService) GetOrganizationInventoryDevicesExecute(r InventoryApiGetOrganizationInventoryDevicesRequest) ([]InlineResponse20094, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20094
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.GetOrganizationInventoryDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/devices"
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
	if r.usedState != nil {
		localVarQueryParams.Add("usedState", parameterToString(*r.usedState, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.models != nil {
		localVarQueryParams.Add("models", parameterToString(*r.models, "csv"))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	}
	if r.tagsFilterType != nil {
		localVarQueryParams.Add("tagsFilterType", parameterToString(*r.tagsFilterType, ""))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
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

type InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	importIds *[]string
}

// import ids from an imports
func (r InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ImportIds(importIds []string) InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	r.importIds = &importIds
	return r
}

func (r InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) Execute() ([]InlineResponse20095, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation

Check the status of a committed Import operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest
*/
func (a *InventoryApiService) GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx context.Context, organizationId string) InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	return InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20095
func (a *InventoryApiService) GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r InventoryApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ([]InlineResponse20095, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20095
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.GetOrganizationInventoryOnboardingCloudMonitoringImports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importIds == nil {
		return localVarReturnValue, nil, reportError("importIds is required and must be specified")
	}

	localVarQueryParams.Add("importIds", parameterToString(*r.importIds, "csv"))
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

type InventoryApiReleaseFromOrganizationInventoryRequest struct {
	ctx context.Context
	ApiService *InventoryApiService
	organizationId string
	releaseFromOrganizationInventory *InlineObject200
}

func (r InventoryApiReleaseFromOrganizationInventoryRequest) ReleaseFromOrganizationInventory(releaseFromOrganizationInventory InlineObject200) InventoryApiReleaseFromOrganizationInventoryRequest {
	r.releaseFromOrganizationInventory = &releaseFromOrganizationInventory
	return r
}

func (r InventoryApiReleaseFromOrganizationInventoryRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ReleaseFromOrganizationInventoryExecute(r)
}

/*
ReleaseFromOrganizationInventory Release a list of claimed devices from an organization.

Release a list of claimed devices from an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return InventoryApiReleaseFromOrganizationInventoryRequest
*/
func (a *InventoryApiService) ReleaseFromOrganizationInventory(ctx context.Context, organizationId string) InventoryApiReleaseFromOrganizationInventoryRequest {
	return InventoryApiReleaseFromOrganizationInventoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *InventoryApiService) ReleaseFromOrganizationInventoryExecute(r InventoryApiReleaseFromOrganizationInventoryRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryApiService.ReleaseFromOrganizationInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/release"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.releaseFromOrganizationInventory
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
