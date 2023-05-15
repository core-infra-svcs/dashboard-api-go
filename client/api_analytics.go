/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
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


// AnalyticsApiService AnalyticsApi service
type AnalyticsApiService service

type AnalyticsApiGetDeviceCameraAnalyticsLiveRequest struct {
	ctx context.Context
	ApiService *AnalyticsApiService
	serial string
}

func (r AnalyticsApiGetDeviceCameraAnalyticsLiveRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsLiveExecute(r)
}

/*
GetDeviceCameraAnalyticsLive Returns live state from camera of analytics zones

Returns live state from camera of analytics zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return AnalyticsApiGetDeviceCameraAnalyticsLiveRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsLive(ctx context.Context, serial string) AnalyticsApiGetDeviceCameraAnalyticsLiveRequest {
	return AnalyticsApiGetDeviceCameraAnalyticsLiveRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsLiveExecute(r AnalyticsApiGetDeviceCameraAnalyticsLiveRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsLive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/live"
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

type AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest struct {
	ctx context.Context
	ApiService *AnalyticsApiService
	serial string
	t0 *string
	t1 *string
	timespan *float32
	objectType *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) T0(t0 string) AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) T1(t1 string) AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.
func (r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) Timespan(timespan float32) AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest {
	r.timespan = &timespan
	return r
}

// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) ObjectType(objectType string) AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsOverviewExecute(r)
}

/*
GetDeviceCameraAnalyticsOverview Returns an overview of aggregate analytics data for a timespan

Returns an overview of aggregate analytics data for a timespan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsOverview(ctx context.Context, serial string) AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest {
	return AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsOverviewExecute(r AnalyticsApiGetDeviceCameraAnalyticsOverviewRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsOverview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/overview"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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
	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
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

type AnalyticsApiGetDeviceCameraAnalyticsRecentRequest struct {
	ctx context.Context
	ApiService *AnalyticsApiService
	serial string
	objectType *string
}

// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiGetDeviceCameraAnalyticsRecentRequest) ObjectType(objectType string) AnalyticsApiGetDeviceCameraAnalyticsRecentRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiGetDeviceCameraAnalyticsRecentRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsRecentExecute(r)
}

/*
GetDeviceCameraAnalyticsRecent Returns most recent record for analytics zones

Returns most recent record for analytics zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return AnalyticsApiGetDeviceCameraAnalyticsRecentRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsRecent(ctx context.Context, serial string) AnalyticsApiGetDeviceCameraAnalyticsRecentRequest {
	return AnalyticsApiGetDeviceCameraAnalyticsRecentRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsRecentExecute(r AnalyticsApiGetDeviceCameraAnalyticsRecentRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsRecent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/recent"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
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

type AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest struct {
	ctx context.Context
	ApiService *AnalyticsApiService
	serial string
	zoneId string
	t0 *string
	t1 *string
	timespan *float32
	resolution *int32
	objectType *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) T0(t0 string) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.
func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) T1(t1 string) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.
func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) Timespan(timespan float32) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.timespan = &timespan
	return r
}

// The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.
func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) Resolution(resolution int32) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.resolution = &resolution
	return r
}

// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) ObjectType(objectType string) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsZoneHistoryExecute(r)
}

/*
GetDeviceCameraAnalyticsZoneHistory Return historical records for analytic zones

Return historical records for analytic zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param zoneId Zone ID
 @return AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZoneHistory(ctx context.Context, serial string, zoneId string) AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	return AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		zoneId: zoneId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZoneHistoryExecute(r AnalyticsApiGetDeviceCameraAnalyticsZoneHistoryRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsZoneHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/zones/{zoneId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", url.PathEscape(parameterToString(r.zoneId, "")), -1)

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
	if r.objectType != nil {
		localVarQueryParams.Add("objectType", parameterToString(*r.objectType, ""))
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

type AnalyticsApiGetDeviceCameraAnalyticsZonesRequest struct {
	ctx context.Context
	ApiService *AnalyticsApiService
	serial string
}

func (r AnalyticsApiGetDeviceCameraAnalyticsZonesRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsZonesExecute(r)
}

/*
GetDeviceCameraAnalyticsZones Returns all configured analytic zones for this camera

Returns all configured analytic zones for this camera

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return AnalyticsApiGetDeviceCameraAnalyticsZonesRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZones(ctx context.Context, serial string) AnalyticsApiGetDeviceCameraAnalyticsZonesRequest {
	return AnalyticsApiGetDeviceCameraAnalyticsZonesRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZonesExecute(r AnalyticsApiGetDeviceCameraAnalyticsZonesRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsZones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/zones"
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
