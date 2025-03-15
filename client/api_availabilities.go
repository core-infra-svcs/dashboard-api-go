/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
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


// AvailabilitiesApiService AvailabilitiesApi service
type AvailabilitiesApiService service

type AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest struct {
	ctx context.Context
	ApiService *AvailabilitiesApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	productTypes *[]string
	serials *[]string
	tags *[]string
	tagsFilterType *string
	statuses *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) PerPage(perPage int32) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) StartingAfter(startingAfter string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) EndingBefore(endingBefore string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) NetworkIds(networkIds []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. Valid types are wireless, appliance, switch, camera, cellularGateway, sensor, wirelessController, and campusGateway
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) ProductTypes(productTypes []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) Serials(serials []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.serials = &serials
	return r
}

// An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) Tags(tags []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.tags = &tags
	return r
}

// An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) TagsFilterType(tagsFilterType string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.tagsFilterType = &tagsFilterType
	return r
}

// Optional parameter to filter device availabilities by device status. This filter uses multiple exact matches.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) Statuses(statuses []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	r.statuses = &statuses
	return r
}

func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) Execute() ([]InlineResponse200261, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesAvailabilitiesExecute(r)
}

/*
GetOrganizationDevicesAvailabilities List the availability information for devices in an organization

List the availability information for devices in an organization. The data returned by this endpoint is updated every 5 minutes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest
*/
func (a *AvailabilitiesApiService) GetOrganizationDevicesAvailabilities(ctx context.Context, organizationId string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest {
	return AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200261
func (a *AvailabilitiesApiService) GetOrganizationDevicesAvailabilitiesExecute(r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesRequest) ([]InlineResponse200261, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200261
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilitiesApiService.GetOrganizationDevicesAvailabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/availabilities"
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
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.tags != nil {
		localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	}
	if r.tagsFilterType != nil {
		localVarQueryParams.Add("tagsFilterType", parameterToString(*r.tagsFilterType, ""))
	}
	if r.statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.statuses, "csv"))
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

type AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest struct {
	ctx context.Context
	ApiService *AvailabilitiesApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	serials *[]string
	productTypes *[]string
	networkIds *[]string
	statuses *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) PerPage(perPage int32) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) StartingAfter(startingAfter string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) EndingBefore(endingBefore string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) T0(t0 string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) T1(t1 string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Timespan(timespan float32) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.timespan = &timespan
	return r
}

// Optional parameter to filter device availabilities history by device serial numbers
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Serials(serials []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter device availabilities history by device product types
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) ProductTypes(productTypes []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.productTypes = &productTypes
	return r
}

// Optional parameter to filter device availabilities history by network IDs
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) NetworkIds(networkIds []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device availabilities history by device statuses
func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Statuses(statuses []string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	r.statuses = &statuses
	return r
}

func (r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) Execute() ([]InlineResponse200262, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesAvailabilitiesChangeHistoryExecute(r)
}

/*
GetOrganizationDevicesAvailabilitiesChangeHistory List the availability history information for devices in an organization.

List the availability history information for devices in an organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest
*/
func (a *AvailabilitiesApiService) GetOrganizationDevicesAvailabilitiesChangeHistory(ctx context.Context, organizationId string) AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest {
	return AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200262
func (a *AvailabilitiesApiService) GetOrganizationDevicesAvailabilitiesChangeHistoryExecute(r AvailabilitiesApiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest) ([]InlineResponse200262, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200262
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilitiesApiService.GetOrganizationDevicesAvailabilitiesChangeHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/availabilities/changeHistory"
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
	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.statuses, "csv"))
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

type AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest struct {
	ctx context.Context
	ApiService *AvailabilitiesApiService
	organizationId string
	serials *[]string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) Serials(serials []string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.serials = &serials
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) T0(t0 string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) T1(t1 string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) Timespan(timespan float32) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) PerPage(perPage int32) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) StartingAfter(startingAfter string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) EndingBefore(endingBefore string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) Execute() (*InlineResponse200347, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessControllerAvailabilitiesChangeHistoryExecute(r)
}

/*
GetOrganizationWirelessControllerAvailabilitiesChangeHistory List connectivity data of wireless LAN controllers in an organization

List connectivity data of wireless LAN controllers in an organization. If it is HA setup, then only returns active WLC data start from switchover

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest
*/
func (a *AvailabilitiesApiService) GetOrganizationWirelessControllerAvailabilitiesChangeHistory(ctx context.Context, organizationId string) AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest {
	return AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200347
func (a *AvailabilitiesApiService) GetOrganizationWirelessControllerAvailabilitiesChangeHistoryExecute(r AvailabilitiesApiGetOrganizationWirelessControllerAvailabilitiesChangeHistoryRequest) (*InlineResponse200347, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200347
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AvailabilitiesApiService.GetOrganizationWirelessControllerAvailabilitiesChangeHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wirelessController/availabilities/changeHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serials != nil {
		localVarQueryParams.Add("serials", parameterToString(*r.serials, "csv"))
	}
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
