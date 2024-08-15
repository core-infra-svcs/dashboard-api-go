/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
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


// CommandsApiService CommandsApi service
type CommandsApiService service

type CommandsApiCreateDeviceSensorCommandRequest struct {
	ctx context.Context
	ApiService *CommandsApiService
	serial string
	createDeviceSensorCommand *InlineObject23
}

func (r CommandsApiCreateDeviceSensorCommandRequest) CreateDeviceSensorCommand(createDeviceSensorCommand InlineObject23) CommandsApiCreateDeviceSensorCommandRequest {
	r.createDeviceSensorCommand = &createDeviceSensorCommand
	return r
}

func (r CommandsApiCreateDeviceSensorCommandRequest) Execute() (*InlineResponse20029, *http.Response, error) {
	return r.ApiService.CreateDeviceSensorCommandExecute(r)
}

/*
CreateDeviceSensorCommand Sends a command to a sensor

Sends a command to a sensor

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return CommandsApiCreateDeviceSensorCommandRequest
*/
func (a *CommandsApiService) CreateDeviceSensorCommand(ctx context.Context, serial string) CommandsApiCreateDeviceSensorCommandRequest {
	return CommandsApiCreateDeviceSensorCommandRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return InlineResponse20029
func (a *CommandsApiService) CreateDeviceSensorCommandExecute(r CommandsApiCreateDeviceSensorCommandRequest) (*InlineResponse20029, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20029
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandsApiService.CreateDeviceSensorCommand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/sensor/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDeviceSensorCommand == nil {
		return localVarReturnValue, nil, reportError("createDeviceSensorCommand is required and must be specified")
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
	localVarPostBody = r.createDeviceSensorCommand
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

type CommandsApiGetDeviceSensorCommandRequest struct {
	ctx context.Context
	ApiService *CommandsApiService
	serial string
	commandId string
}

func (r CommandsApiGetDeviceSensorCommandRequest) Execute() (*InlineResponse20029, *http.Response, error) {
	return r.ApiService.GetDeviceSensorCommandExecute(r)
}

/*
GetDeviceSensorCommand Returns information about the command's execution, including the status

Returns information about the command's execution, including the status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @param commandId Command ID
 @return CommandsApiGetDeviceSensorCommandRequest
*/
func (a *CommandsApiService) GetDeviceSensorCommand(ctx context.Context, serial string, commandId string) CommandsApiGetDeviceSensorCommandRequest {
	return CommandsApiGetDeviceSensorCommandRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
		commandId: commandId,
	}
}

// Execute executes the request
//  @return InlineResponse20029
func (a *CommandsApiService) GetDeviceSensorCommandExecute(r CommandsApiGetDeviceSensorCommandRequest) (*InlineResponse20029, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20029
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandsApiService.GetDeviceSensorCommand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/sensor/commands/{commandId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commandId"+"}", url.PathEscape(parameterToString(r.commandId, "")), -1)

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

type CommandsApiGetDeviceSensorCommandsRequest struct {
	ctx context.Context
	ApiService *CommandsApiService
	serial string
	operations *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	sortOrder *string
	t0 *string
	t1 *string
	timespan *float32
}

// Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData.
func (r CommandsApiGetDeviceSensorCommandsRequest) Operations(operations []string) CommandsApiGetDeviceSensorCommandsRequest {
	r.operations = &operations
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10.
func (r CommandsApiGetDeviceSensorCommandsRequest) PerPage(perPage int32) CommandsApiGetDeviceSensorCommandsRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CommandsApiGetDeviceSensorCommandsRequest) StartingAfter(startingAfter string) CommandsApiGetDeviceSensorCommandsRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r CommandsApiGetDeviceSensorCommandsRequest) EndingBefore(endingBefore string) CommandsApiGetDeviceSensorCommandsRequest {
	r.endingBefore = &endingBefore
	return r
}

// Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;descending&#39;.
func (r CommandsApiGetDeviceSensorCommandsRequest) SortOrder(sortOrder string) CommandsApiGetDeviceSensorCommandsRequest {
	r.sortOrder = &sortOrder
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
func (r CommandsApiGetDeviceSensorCommandsRequest) T0(t0 string) CommandsApiGetDeviceSensorCommandsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 30 days after t0.
func (r CommandsApiGetDeviceSensorCommandsRequest) T1(t1 string) CommandsApiGetDeviceSensorCommandsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days.
func (r CommandsApiGetDeviceSensorCommandsRequest) Timespan(timespan float32) CommandsApiGetDeviceSensorCommandsRequest {
	r.timespan = &timespan
	return r
}

func (r CommandsApiGetDeviceSensorCommandsRequest) Execute() ([]InlineResponse20029, *http.Response, error) {
	return r.ApiService.GetDeviceSensorCommandsExecute(r)
}

/*
GetDeviceSensorCommands Returns a historical log of all commands

Returns a historical log of all commands

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return CommandsApiGetDeviceSensorCommandsRequest
*/
func (a *CommandsApiService) GetDeviceSensorCommands(ctx context.Context, serial string) CommandsApiGetDeviceSensorCommandsRequest {
	return CommandsApiGetDeviceSensorCommandsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return []InlineResponse20029
func (a *CommandsApiService) GetDeviceSensorCommandsExecute(r CommandsApiGetDeviceSensorCommandsRequest) ([]InlineResponse20029, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20029
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommandsApiService.GetDeviceSensorCommands")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/sensor/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.operations != nil {
		localVarQueryParams.Add("operations", parameterToString(*r.operations, "csv"))
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
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
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
