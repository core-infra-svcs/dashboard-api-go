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


// CategoriesApiService CategoriesApi service
type CategoriesApiService service

type CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest struct {
	ctx context.Context
	ApiService *CategoriesApiService
	networkId string
}

func (r CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceContentFilteringCategoriesExecute(r)
}

/*
GetNetworkApplianceContentFilteringCategories List all available content filtering categories for an MX network

List all available content filtering categories for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest
*/
func (a *CategoriesApiService) GetNetworkApplianceContentFilteringCategories(ctx context.Context, networkId string) CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest {
	return CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CategoriesApiService) GetNetworkApplianceContentFilteringCategoriesExecute(r CategoriesApiGetNetworkApplianceContentFilteringCategoriesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesApiService.GetNetworkApplianceContentFilteringCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/contentFiltering/categories"
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

type CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest struct {
	ctx context.Context
	ApiService *CategoriesApiService
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
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkTag(networkTag string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkTag = &networkTag
	return r
}

// Match result to an exact device tag
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) DeviceTag(deviceTag string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.deviceTag = &deviceTag
	return r
}

// Match result to an exact network id
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) NetworkId(networkId string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.networkId = &networkId
	return r
}

// Set number of desired results to return. Default is 10.
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Quantity(quantity int32) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.quantity = &quantity
	return r
}

// Filter results by ssid name
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) SsidName(ssidName string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.ssidName = &ssidName
	return r
}

// Filter results by usage uplink
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) UsageUplink(usageUplink string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.usageUplink = &usageUplink
	return r
}

// The beginning of the timespan for the data.
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T0(t0 string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 186 days after t0.
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) T1(t1 string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day.
func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Timespan(timespan float32) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) Execute() ([]InlineResponse200281, *http.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r)
}

/*
GetOrganizationSummaryTopApplicationsCategoriesByUsage Return the top application categories sorted by data usage over given time range

Return the top application categories sorted by data usage over given time range. Default unit is megabytes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest
*/
func (a *CategoriesApiService) GetOrganizationSummaryTopApplicationsCategoriesByUsage(ctx context.Context, organizationId string) CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest {
	return CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200281
func (a *CategoriesApiService) GetOrganizationSummaryTopApplicationsCategoriesByUsageExecute(r CategoriesApiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest) ([]InlineResponse200281, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200281
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoriesApiService.GetOrganizationSummaryTopApplicationsCategoriesByUsage")
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
