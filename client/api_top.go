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

// TopApiService TopApi service
type TopApiService service

type TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) Execute() ([]InlineResponse2004, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopAppliancesByUtilizationExecute(r)
}

/*
GetOrganizationSummaryTopAppliancesByUtilization Return the top 10 appliances sorted by utilization over given time range.

Return the top 10 appliances sorted by utilization over given time range.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopAppliancesByUtilization(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest {
	return TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2004
func (a *TopApiService) GetOrganizationSummaryTopAppliancesByUtilizationExecute(r TopApiApiGetOrganizationSummaryTopAppliancesByUtilizationRequest) ([]InlineResponse2004, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2004
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopAppliancesByUtilization")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/appliances/byUtilization"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopClientsByUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopClientsByUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopClientsByUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopClientsByUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopClientsByUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopClientsByUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopClientsByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopClientsByUsageRequest) Execute() ([]InlineResponse2005, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopClientsByUsageExecute(r)
}

/*
GetOrganizationSummaryTopClientsByUsage Return metrics for organization's top 10 clients by data usage (in mb) over given time range.

Return metrics for organization's top 10 clients by data usage (in mb) over given time range.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopClientsByUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopClientsByUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopClientsByUsageRequest {
	return TopApiApiGetOrganizationSummaryTopClientsByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2005
func (a *TopApiService) GetOrganizationSummaryTopClientsByUsageExecute(r TopApiApiGetOrganizationSummaryTopClientsByUsageRequest) ([]InlineResponse2005, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2005
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopClientsByUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/clients/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest) Execute() ([]InlineResponse2006, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopClientsManufacturersByUsageExecute(r)
}

/*
GetOrganizationSummaryTopClientsManufacturersByUsage Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.

Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopClientsManufacturersByUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest {
	return TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2006
func (a *TopApiService) GetOrganizationSummaryTopClientsManufacturersByUsageExecute(r TopApiApiGetOrganizationSummaryTopClientsManufacturersByUsageRequest) ([]InlineResponse2006, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2006
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopClientsManufacturersByUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/clients/manufacturers/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest) Execute() ([]InlineResponse2007, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopDevicesByUsageExecute(r)
}

/*
GetOrganizationSummaryTopDevicesByUsage Return metrics for organization's top 10 devices sorted by data usage over given time range

Return metrics for organization's top 10 devices sorted by data usage over given time range. Default unit is megabytes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopDevicesByUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest {
	return TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2007
func (a *TopApiService) GetOrganizationSummaryTopDevicesByUsageExecute(r TopApiApiGetOrganizationSummaryTopDevicesByUsageRequest) ([]InlineResponse2007, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2007
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopDevicesByUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/devices/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopDevicesModelsByUsageExecute(r)
}

/*
GetOrganizationSummaryTopDevicesModelsByUsage Return metrics for organization's top 10 device models sorted by data usage over given time range

Return metrics for organization's top 10 device models sorted by data usage over given time range. Default unit is megabytes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopDevicesModelsByUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest {
	return TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TopApiService) GetOrganizationSummaryTopDevicesModelsByUsageExecute(r TopApiApiGetOrganizationSummaryTopDevicesModelsByUsageRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopDevicesModelsByUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/devices/models/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopSsidsByUsageExecute(r)
}

/*
GetOrganizationSummaryTopSsidsByUsage Return metrics for organization's top 10 ssids by data usage over given time range

Return metrics for organization's top 10 ssids by data usage over given time range. Default unit is megabytes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopSsidsByUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest {
	return TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *TopApiService) GetOrganizationSummaryTopSsidsByUsageExecute(r TopApiApiGetOrganizationSummaryTopSsidsByUsageRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopSsidsByUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/ssids/byUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest struct {
	ctx _context.Context
	ApiService *TopApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
}

// The beginning of the timespan for the data.
func (r TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T0(t0 string) TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) T1(t1 string) TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
func (r TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Timespan(timespan float32) TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	r.timespan = &timespan
	return r
}

func (r TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) Execute() ([]InlineResponse2008, *_nethttp.Response, error) {
	return r.ApiService.GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r)
}

/*
GetOrganizationSummaryTopSwitchesByEnergyUsage Return metrics for organization's top 10 switches by energy usage over given time range

Return metrics for organization's top 10 switches by energy usage over given time range. Default unit is joules.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest
*/
func (a *TopApiService) GetOrganizationSummaryTopSwitchesByEnergyUsage(ctx _context.Context, organizationId string) TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest {
	return TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse2008
func (a *TopApiService) GetOrganizationSummaryTopSwitchesByEnergyUsageExecute(r TopApiApiGetOrganizationSummaryTopSwitchesByEnergyUsageRequest) ([]InlineResponse2008, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []InlineResponse2008
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopApiService.GetOrganizationSummaryTopSwitchesByEnergyUsage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/summary/top/switches/byEnergyUsage"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
