/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
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


// ApplicationsApiService ApplicationsApi service
type ApplicationsApiService service

type ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	networkId string
	applicationId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
}

// The beginning of the timespan for the data. The maximum lookback period is 7 days from today.
func (r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) T0(t0 string) ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) T1(t1 string) ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours.
func (r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) Timespan(timespan float32) ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300.
func (r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) Resolution(resolution int32) ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest {
	r.resolution = &resolution
	return r
}

func (r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) Execute() ([]InlineResponse20097, *http.Response, error) {
	return r.ApiService.GetNetworkInsightApplicationHealthByTimeExecute(r)
}

/*
GetNetworkInsightApplicationHealthByTime Get application health by time

Get application health by time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param applicationId Application ID
 @return ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest
*/
func (a *ApplicationsApiService) GetNetworkInsightApplicationHealthByTime(ctx context.Context, networkId string, applicationId string) ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest {
	return ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []InlineResponse20097
func (a *ApplicationsApiService) GetNetworkInsightApplicationHealthByTimeExecute(r ApplicationsApiGetNetworkInsightApplicationHealthByTimeRequest) ([]InlineResponse20097, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20097
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetNetworkInsightApplicationHealthByTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/insight/applications/{applicationId}/healthByTime"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.resolution != nil {
		localVarQueryParams.Add("resolution", parameterToString(*r.resolution, ""))
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

type ApplicationsApiGetOrganizationInsightApplicationsRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	organizationId string
}

func (r ApplicationsApiGetOrganizationInsightApplicationsRequest) Execute() ([]InlineResponse200249, *http.Response, error) {
	return r.ApiService.GetOrganizationInsightApplicationsExecute(r)
}

/*
GetOrganizationInsightApplications List all Insight tracked applications

List all Insight tracked applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsApiGetOrganizationInsightApplicationsRequest
*/
func (a *ApplicationsApiService) GetOrganizationInsightApplications(ctx context.Context, organizationId string) ApplicationsApiGetOrganizationInsightApplicationsRequest {
	return ApplicationsApiGetOrganizationInsightApplicationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200249
func (a *ApplicationsApiService) GetOrganizationInsightApplicationsExecute(r ApplicationsApiGetOrganizationInsightApplicationsRequest) ([]InlineResponse200249, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200249
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetOrganizationInsightApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/insight/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	organizationId string
	networkTag *string
	device *string
	networkId *string
	quantity *int32
	ssidName *string
	usageUplink *string
	t0 *string
	t1 *string
	timespan *float32
}

// Match result to an exact network tag
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) NetworkTag(networkTag string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) Device(device string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.device = &device
	return r
}

// Match result to an exact network id
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) NetworkId(networkId string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) Quantity(quantity int32) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) SsidName(ssidName string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) UsageUplink(usageUplink string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) T0(t0 string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) T1(t1 string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) Timespan(timespan float32) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) Execute() ([]InlineResponse200280, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopApplicationsByUsageExecute(r)
}

/*
GetOrganizationSummaryTopApplicationsByUsage Return the top applications sorted by data usage over given time range

Return the top applications sorted by data usage over given time range. Default unit is megabytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest
*/
func (a *ApplicationsApiService) GetOrganizationSummaryTopApplicationsByUsage(ctx context.Context, organizationId string) ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest {
	return ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200280
func (a *ApplicationsApiService) GetOrganizationSummaryTopApplicationsByUsageExecute(r ApplicationsApiGetOrganizationSummaryTopApplicationsByUsageRequest) ([]InlineResponse200280, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200280
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetOrganizationSummaryTopApplicationsByUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/applications/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkTag != nil {
		localVarQueryParams.Add("networkTag", parameterToString(*r.networkTag, ""))
	}
	if r.device != nil {
		localVarQueryParams.Add("device", parameterToString(*r.device, ""))
	}
	if r.networkId != nil {
		localVarQueryParams.Add("networkId", parameterToString(*r.networkId, ""))
	}
	if r.quantity != nil {
		localVarQueryParams.Add("quantity", parameterToString(*r.quantity, ""))
	}
	if r.ssidName != nil {
		localVarQueryParams.Add("ssidName", parameterToString(*r.ssidName, ""))
	}
	if r.usageUplink != nil {
		localVarQueryParams.Add("usageUplink", parameterToString(*r.usageUplink, ""))
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

type ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest struct {
	ctx context.Context
	ApiService *ApplicationsApiService
	organizationId string
	networkTag *string
	deviceTag *string
	networkId *string
	quantity *int32
	ssidName *string
	usageUplink *string
	t0 *string
	t1 *string
	timespan *float32
}

// Match result to an exact network tag
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkTag(networkTag string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) DeviceTag(deviceTag string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.deviceTag = &deviceTag
	return r
}

// Match result to an exact network id
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkId(networkId string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Quantity(quantity int32) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) SsidName(ssidName string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) UsageUplink(usageUplink string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T0(t0 string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T1(t1 string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.
func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Timespan(timespan float32) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Execute() ([]InlineResponse200281, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r)
}

/*
GetOrganizationSummaryTopApplicationsCategoriesByUsage Return the top application categories sorted by data usage over given time range

Return the top application categories sorted by data usage over given time range. Default unit is megabytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest
*/
func (a *ApplicationsApiService) GetOrganizationSummaryTopApplicationsCategoriesByUsage(ctx context.Context, organizationId string) ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	return ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200281
func (a *ApplicationsApiService) GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r ApplicationsApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) ([]InlineResponse200281, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200281
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsApiService.GetOrganizationSummaryTopApplicationsCategoriesByUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/applications/categories/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkTag != nil {
		localVarQueryParams.Add("networkTag", parameterToString(*r.networkTag, ""))
	}
	if r.deviceTag != nil {
		localVarQueryParams.Add("deviceTag", parameterToString(*r.deviceTag, ""))
	}
	if r.networkId != nil {
		localVarQueryParams.Add("networkId", parameterToString(*r.networkId, ""))
	}
	if r.quantity != nil {
		localVarQueryParams.Add("quantity", parameterToString(*r.quantity, ""))
	}
	if r.ssidName != nil {
		localVarQueryParams.Add("ssidName", parameterToString(*r.ssidName, ""))
	}
	if r.usageUplink != nil {
		localVarQueryParams.Add("usageUplink", parameterToString(*r.usageUplink, ""))
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
