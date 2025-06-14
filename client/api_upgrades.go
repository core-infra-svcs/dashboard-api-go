/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
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


// UpgradesApiService UpgradesApi service
type UpgradesApiService service

type UpgradesApiGetOrganizationFirmwareUpgradesRequest struct {
	ctx context.Context
	ApiService *UpgradesApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	status *[]string
	productTypes *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) PerPage(perPage int32) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) StartingAfter(startingAfter string) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) EndingBefore(endingBefore string) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter the upgrade by status.
func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) Status(status []string) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	r.status = &status
	return r
}

// Optional parameter to filter the upgrade by product type.
func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) ProductTypes(productTypes []string) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	r.productTypes = &productTypes
	return r
}

func (r UpgradesApiGetOrganizationFirmwareUpgradesRequest) Execute() ([]InlineResponse200286, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesExecute(r)
}

/*
GetOrganizationFirmwareUpgrades Get firmware upgrade information for an organization

Get firmware upgrade information for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return UpgradesApiGetOrganizationFirmwareUpgradesRequest
*/
func (a *UpgradesApiService) GetOrganizationFirmwareUpgrades(ctx context.Context, organizationId string) UpgradesApiGetOrganizationFirmwareUpgradesRequest {
	return UpgradesApiGetOrganizationFirmwareUpgradesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200286
func (a *UpgradesApiService) GetOrganizationFirmwareUpgradesExecute(r UpgradesApiGetOrganizationFirmwareUpgradesRequest) ([]InlineResponse200286, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200286
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradesApiService.GetOrganizationFirmwareUpgrades")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades"
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
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, "csv"))
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

type UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest struct {
	ctx context.Context
	ApiService *UpgradesApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	macs *[]string
	firmwareUpgradeBatchIds *[]string
	upgradeStatuses *[]string
	currentUpgradesOnly *bool
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) PerPage(perPage int32) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) StartingAfter(startingAfter string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) EndingBefore(endingBefore string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter by network
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) NetworkIds(networkIds []string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) Serials(serials []string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) Macs(macs []string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter by firmware upgrade batch ids.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeBatchIds = &firmwareUpgradeBatchIds
	return r
}

// Optional parameter to filter by firmware upgrade statuses.
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) UpgradeStatuses(upgradeStatuses []string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.upgradeStatuses = &upgradeStatuses
	return r
}

// Optional parameter to filter to only current or pending upgrade statuses
func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) CurrentUpgradesOnly(currentUpgradesOnly bool) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.currentUpgradesOnly = &currentUpgradesOnly
	return r
}

func (r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) Execute() ([]InlineResponse200287, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesByDeviceExecute(r)
}

/*
GetOrganizationFirmwareUpgradesByDevice Get firmware upgrade status for the filtered devices

Get firmware upgrade status for the filtered devices. This endpoint currently only supports Meraki switches.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest
*/
func (a *UpgradesApiService) GetOrganizationFirmwareUpgradesByDevice(ctx context.Context, organizationId string) UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	return UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200287
func (a *UpgradesApiService) GetOrganizationFirmwareUpgradesByDeviceExecute(r UpgradesApiGetOrganizationFirmwareUpgradesByDeviceRequest) ([]InlineResponse200287, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200287
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpgradesApiService.GetOrganizationFirmwareUpgradesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades/byDevice"
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
	if r.macs != nil {
		localVarQueryParams.Add("macs", parameterToString(*r.macs, "csv"))
	}
	if r.firmwareUpgradeBatchIds != nil {
		localVarQueryParams.Add("firmwareUpgradeBatchIds", parameterToString(*r.firmwareUpgradeBatchIds, "csv"))
	}
	if r.upgradeStatuses != nil {
		localVarQueryParams.Add("upgradeStatuses", parameterToString(*r.upgradeStatuses, "csv"))
	}
	if r.currentUpgradesOnly != nil {
		localVarQueryParams.Add("currentUpgradesOnly", parameterToString(*r.currentUpgradesOnly, ""))
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
