/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
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


// PayloadTemplatesApiService PayloadTemplatesApi service
type PayloadTemplatesApiService service

type PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest struct {
	ctx context.Context
	ApiService *PayloadTemplatesApiService
	networkId string
	createNetworkWebhooksPayloadTemplate *InlineObject143
}

func (r PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest) CreateNetworkWebhooksPayloadTemplate(createNetworkWebhooksPayloadTemplate InlineObject143) PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest {
	r.createNetworkWebhooksPayloadTemplate = &createNetworkWebhooksPayloadTemplate
	return r
}

func (r PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest) Execute() (*InlineResponse20056, *http.Response, error) {
	return r.ApiService.CreateNetworkWebhooksPayloadTemplateExecute(r)
}

/*
CreateNetworkWebhooksPayloadTemplate Create a webhook payload template for a network

Create a webhook payload template for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest
*/
func (a *PayloadTemplatesApiService) CreateNetworkWebhooksPayloadTemplate(ctx context.Context, networkId string) PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest {
	return PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20056
func (a *PayloadTemplatesApiService) CreateNetworkWebhooksPayloadTemplateExecute(r PayloadTemplatesApiCreateNetworkWebhooksPayloadTemplateRequest) (*InlineResponse20056, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20056
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayloadTemplatesApiService.CreateNetworkWebhooksPayloadTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/payloadTemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkWebhooksPayloadTemplate == nil {
		return localVarReturnValue, nil, reportError("createNetworkWebhooksPayloadTemplate is required and must be specified")
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
	localVarPostBody = r.createNetworkWebhooksPayloadTemplate
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

type PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest struct {
	ctx context.Context
	ApiService *PayloadTemplatesApiService
	networkId string
	payloadTemplateId string
}

func (r PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkWebhooksPayloadTemplateExecute(r)
}

/*
DeleteNetworkWebhooksPayloadTemplate Destroy a webhook payload template for a network

Destroy a webhook payload template for a network. Does not work for included templates ('wpt_00001', 'wpt_00002', 'wpt_00003', 'wpt_00004', 'wpt_00005' or 'wpt_00006')

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param payloadTemplateId
 @return PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest
*/
func (a *PayloadTemplatesApiService) DeleteNetworkWebhooksPayloadTemplate(ctx context.Context, networkId string, payloadTemplateId string) PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest {
	return PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		payloadTemplateId: payloadTemplateId,
	}
}

// Execute executes the request
func (a *PayloadTemplatesApiService) DeleteNetworkWebhooksPayloadTemplateExecute(r PayloadTemplatesApiDeleteNetworkWebhooksPayloadTemplateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayloadTemplatesApiService.DeleteNetworkWebhooksPayloadTemplate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payloadTemplateId"+"}", url.PathEscape(parameterToString(r.payloadTemplateId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest struct {
	ctx context.Context
	ApiService *PayloadTemplatesApiService
	networkId string
	payloadTemplateId string
}

func (r PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest) Execute() (*InlineResponse20056, *http.Response, error) {
	return r.ApiService.GetNetworkWebhooksPayloadTemplateExecute(r)
}

/*
GetNetworkWebhooksPayloadTemplate Get the webhook payload template for a network

Get the webhook payload template for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param payloadTemplateId
 @return PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest
*/
func (a *PayloadTemplatesApiService) GetNetworkWebhooksPayloadTemplate(ctx context.Context, networkId string, payloadTemplateId string) PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest {
	return PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		payloadTemplateId: payloadTemplateId,
	}
}

// Execute executes the request
//  @return InlineResponse20056
func (a *PayloadTemplatesApiService) GetNetworkWebhooksPayloadTemplateExecute(r PayloadTemplatesApiGetNetworkWebhooksPayloadTemplateRequest) (*InlineResponse20056, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20056
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayloadTemplatesApiService.GetNetworkWebhooksPayloadTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payloadTemplateId"+"}", url.PathEscape(parameterToString(r.payloadTemplateId, "")), -1)

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

type PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest struct {
	ctx context.Context
	ApiService *PayloadTemplatesApiService
	networkId string
}

func (r PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest) Execute() ([]InlineResponse20056, *http.Response, error) {
	return r.ApiService.GetNetworkWebhooksPayloadTemplatesExecute(r)
}

/*
GetNetworkWebhooksPayloadTemplates List the webhook payload templates for a network

List the webhook payload templates for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest
*/
func (a *PayloadTemplatesApiService) GetNetworkWebhooksPayloadTemplates(ctx context.Context, networkId string) PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest {
	return PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []InlineResponse20056
func (a *PayloadTemplatesApiService) GetNetworkWebhooksPayloadTemplatesExecute(r PayloadTemplatesApiGetNetworkWebhooksPayloadTemplatesRequest) ([]InlineResponse20056, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse20056
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayloadTemplatesApiService.GetNetworkWebhooksPayloadTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/payloadTemplates"
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

type PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest struct {
	ctx context.Context
	ApiService *PayloadTemplatesApiService
	networkId string
	payloadTemplateId string
	updateNetworkWebhooksPayloadTemplate *InlineObject144
}

func (r PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest) UpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate InlineObject144) PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest {
	r.updateNetworkWebhooksPayloadTemplate = &updateNetworkWebhooksPayloadTemplate
	return r
}

func (r PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest) Execute() (*InlineResponse20056, *http.Response, error) {
	return r.ApiService.UpdateNetworkWebhooksPayloadTemplateExecute(r)
}

/*
UpdateNetworkWebhooksPayloadTemplate Update a webhook payload template for a network

Update a webhook payload template for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @param payloadTemplateId
 @return PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest
*/
func (a *PayloadTemplatesApiService) UpdateNetworkWebhooksPayloadTemplate(ctx context.Context, networkId string, payloadTemplateId string) PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest {
	return PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		payloadTemplateId: payloadTemplateId,
	}
}

// Execute executes the request
//  @return InlineResponse20056
func (a *PayloadTemplatesApiService) UpdateNetworkWebhooksPayloadTemplateExecute(r PayloadTemplatesApiUpdateNetworkWebhooksPayloadTemplateRequest) (*InlineResponse20056, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20056
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PayloadTemplatesApiService.UpdateNetworkWebhooksPayloadTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"payloadTemplateId"+"}", url.PathEscape(parameterToString(r.payloadTemplateId, "")), -1)

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
	localVarPostBody = r.updateNetworkWebhooksPayloadTemplate
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
