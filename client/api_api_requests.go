/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
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


// ApiRequestsApiService ApiRequestsApi service
type ApiRequestsApiService service

type ApiRequestsApiGetOrganizationApiRequestsRequest struct {
	ctx context.Context
	ApiService *ApiRequestsApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	perPage *int32
	startingAfter *string
	endingBefore *string
	adminId *string
	path *string
	method *string
	responseCode *int32
	sourceIp *string
	userAgent *string
	version *int32
	operationIds *[]string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) T0(t0 string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) T1(t1 string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) Timespan(timespan float32) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.timespan = &timespan
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) PerPage(perPage int32) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) StartingAfter(startingAfter string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) EndingBefore(endingBefore string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.endingBefore = &endingBefore
	return r
}

// Filter the results by the ID of the admin who made the API requests
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) AdminId(adminId string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.adminId = &adminId
	return r
}

// Filter the results by the path of the API requests
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) Path(path string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.path = &path
	return r
}

// Filter the results by the method of the API requests (must be &#39;GET&#39;, &#39;PUT&#39;, &#39;POST&#39; or &#39;DELETE&#39;)
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) Method(method string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.method = &method
	return r
}

// Filter the results by the response code of the API requests
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) ResponseCode(responseCode int32) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.responseCode = &responseCode
	return r
}

// Filter the results by the IP address of the originating API request
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) SourceIp(sourceIp string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.sourceIp = &sourceIp
	return r
}

// Filter the results by the user agent string of the API request
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) UserAgent(userAgent string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.userAgent = &userAgent
	return r
}

// Filter the results by the API version of the API request
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) Version(version int32) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.version = &version
	return r
}

// Filter the results by one or more operation IDs for the API request
func (r ApiRequestsApiGetOrganizationApiRequestsRequest) OperationIds(operationIds []string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	r.operationIds = &operationIds
	return r
}

func (r ApiRequestsApiGetOrganizationApiRequestsRequest) Execute() ([]InlineResponse200124, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsExecute(r)
}

/*
GetOrganizationApiRequests List the API requests made by an organization

List the API requests made by an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsApiGetOrganizationApiRequestsRequest
*/
func (a *ApiRequestsApiService) GetOrganizationApiRequests(ctx context.Context, organizationId string) ApiRequestsApiGetOrganizationApiRequestsRequest {
	return ApiRequestsApiGetOrganizationApiRequestsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200124
func (a *ApiRequestsApiService) GetOrganizationApiRequestsExecute(r ApiRequestsApiGetOrganizationApiRequestsRequest) ([]InlineResponse200124, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200124
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsApiService.GetOrganizationApiRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.adminId != nil {
		localVarQueryParams.Add("adminId", parameterToString(*r.adminId, ""))
	}
	if r.path != nil {
		localVarQueryParams.Add("path", parameterToString(*r.path, ""))
	}
	if r.method != nil {
		localVarQueryParams.Add("method", parameterToString(*r.method, ""))
	}
	if r.responseCode != nil {
		localVarQueryParams.Add("responseCode", parameterToString(*r.responseCode, ""))
	}
	if r.sourceIp != nil {
		localVarQueryParams.Add("sourceIp", parameterToString(*r.sourceIp, ""))
	}
	if r.userAgent != nil {
		localVarQueryParams.Add("userAgent", parameterToString(*r.userAgent, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.operationIds != nil {
		localVarQueryParams.Add("operationIds", parameterToString(*r.operationIds, "csv"))
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

type ApiRequestsApiGetOrganizationApiRequestsOverviewRequest struct {
	ctx context.Context
	ApiService *ApiRequestsApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewRequest) T0(t0 string) ApiRequestsApiGetOrganizationApiRequestsOverviewRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewRequest) T1(t1 string) ApiRequestsApiGetOrganizationApiRequestsOverviewRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewRequest) Timespan(timespan float32) ApiRequestsApiGetOrganizationApiRequestsOverviewRequest {
	r.timespan = &timespan
	return r
}

func (r ApiRequestsApiGetOrganizationApiRequestsOverviewRequest) Execute() (*InlineResponse200125, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewExecute(r)
}

/*
GetOrganizationApiRequestsOverview Return an aggregated overview of API requests data

Return an aggregated overview of API requests data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsApiGetOrganizationApiRequestsOverviewRequest
*/
func (a *ApiRequestsApiService) GetOrganizationApiRequestsOverview(ctx context.Context, organizationId string) ApiRequestsApiGetOrganizationApiRequestsOverviewRequest {
	return ApiRequestsApiGetOrganizationApiRequestsOverviewRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200125
func (a *ApiRequestsApiService) GetOrganizationApiRequestsOverviewExecute(r ApiRequestsApiGetOrganizationApiRequestsOverviewRequest) (*InlineResponse200125, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200125
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsApiService.GetOrganizationApiRequestsOverview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct {
	ctx context.Context
	ApiService *ApiRequestsApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
	version *int32
	operationIds *[]string
	sourceIps *[]string
	adminIds *[]string
	userAgent *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T0(t0 string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T1(t1 string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Timespan(timespan float32) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Interval(interval int32) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.interval = &interval
	return r
}

// Filter by API version of the endpoint. Allowable values are: [0, 1]
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Version(version int32) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.version = &version
	return r
}

// Filter by operation ID of the endpoint
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) OperationIds(operationIds []string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.operationIds = &operationIds
	return r
}

// Filter by source IP that made the API request
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) SourceIps(sourceIps []string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.sourceIps = &sourceIps
	return r
}

// Filter by admin ID of user that made the API request
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) AdminIds(adminIds []string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.adminIds = &adminIds
	return r
}

// Filter by user agent string for API request. This will filter by a complete or partial match.
func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) UserAgent(userAgent string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.userAgent = &userAgent
	return r
}

func (r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Execute() ([]InlineResponse200126, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r)
}

/*
GetOrganizationApiRequestsOverviewResponseCodesByInterval Tracks organizations' API requests by response code across a given time period

Tracks organizations' API requests by response code across a given time period

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest
*/
func (a *ApiRequestsApiService) GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx context.Context, organizationId string) ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	return ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200126
func (a *ApiRequestsApiService) GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r ApiRequestsApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) ([]InlineResponse200126, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200126
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiRequestsApiService.GetOrganizationApiRequestsOverviewResponseCodesByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.operationIds != nil {
		localVarQueryParams.Add("operationIds", parameterToString(*r.operationIds, "csv"))
	}
	if r.sourceIps != nil {
		localVarQueryParams.Add("sourceIps", parameterToString(*r.sourceIps, "csv"))
	}
	if r.adminIds != nil {
		localVarQueryParams.Add("adminIds", parameterToString(*r.adminIds, "csv"))
	}
	if r.userAgent != nil {
		localVarQueryParams.Add("userAgent", parameterToString(*r.userAgent, ""))
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
