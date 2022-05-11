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

// AnalyticsApiService AnalyticsApi service
type AnalyticsApiService service

type AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest struct {
	ctx _context.Context
	ApiService *AnalyticsApiService
	serial string
}


func (r AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsLiveExecute(r)
}

/*
GetDeviceCameraAnalyticsLive Returns live state from camera of analytics zones

Returns live state from camera of analytics zones

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsLive(ctx _context.Context, serial string) AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest {
	return AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsLiveExecute(r AnalyticsApiApiGetDeviceCameraAnalyticsLiveRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsLive")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/live"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest struct {
	ctx _context.Context
	ApiService *AnalyticsApiService
	serial string
	t0 *string
	t1 *string
	timespan *float32
	objectType *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) T0(t0 string) AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) T1(t1 string) AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) Timespan(timespan float32) AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest {
	r.timespan = &timespan
	return r
}
// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) ObjectType(objectType string) AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsOverviewExecute(r)
}

/*
GetDeviceCameraAnalyticsOverview Returns an overview of aggregate analytics data for a timespan

Returns an overview of aggregate analytics data for a timespan

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsOverview(ctx _context.Context, serial string) AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest {
	return AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsOverviewExecute(r AnalyticsApiApiGetDeviceCameraAnalyticsOverviewRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsOverview")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/overview"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)

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

type AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest struct {
	ctx _context.Context
	ApiService *AnalyticsApiService
	serial string
	objectType *string
}

// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest) ObjectType(objectType string) AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsRecentExecute(r)
}

/*
GetDeviceCameraAnalyticsRecent Returns most recent record for analytics zones

Returns most recent record for analytics zones

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsRecent(ctx _context.Context, serial string) AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest {
	return AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsRecentExecute(r AnalyticsApiApiGetDeviceCameraAnalyticsRecentRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsRecent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/recent"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest struct {
	ctx _context.Context
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
func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) T0(t0 string) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t0 = &t0
	return r
}
// The end of the timespan for the data. t1 can be a maximum of 14 hours after t0.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) T1(t1 string) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.t1 = &t1
	return r
}
// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) Timespan(timespan float32) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.timespan = &timespan
	return r
}
// The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60.
func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) Resolution(resolution int32) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.resolution = &resolution
	return r
}
// [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle].
func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) ObjectType(objectType string) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	r.objectType = &objectType
	return r
}

func (r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsZoneHistoryExecute(r)
}

/*
GetDeviceCameraAnalyticsZoneHistory Return historical records for analytic zones

Return historical records for analytic zones

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @param zoneId
 @return AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZoneHistory(ctx _context.Context, serial string, zoneId string) AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest {
	return AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		zoneId: zoneId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZoneHistoryExecute(r AnalyticsApiApiGetDeviceCameraAnalyticsZoneHistoryRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsZoneHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/zones/{zoneId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneId"+"}", _neturl.PathEscape(parameterToString(r.zoneId, "")), -1)

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

type AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest struct {
	ctx _context.Context
	ApiService *AnalyticsApiService
	serial string
}


func (r AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GetDeviceCameraAnalyticsZonesExecute(r)
}

/*
GetDeviceCameraAnalyticsZones Returns all configured analytic zones for this camera

Returns all configured analytic zones for this camera

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial
 @return AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest
*/
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZones(ctx _context.Context, serial string) AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest {
	return AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *AnalyticsApiService) GetDeviceCameraAnalyticsZonesExecute(r AnalyticsApiApiGetDeviceCameraAnalyticsZonesRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDeviceCameraAnalyticsZones")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/camera/analytics/zones"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", _neturl.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
