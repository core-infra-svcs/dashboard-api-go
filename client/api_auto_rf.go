/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
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


// AutoRfApiService AutoRfApi service
type AutoRfApiService service

type AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest struct {
	ctx context.Context
	ApiService *AutoRfApiService
	organizationId string
	recalculateOrganizationWirelessRadioAutoRfChannels *InlineObject275
}

func (r AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest) RecalculateOrganizationWirelessRadioAutoRfChannels(recalculateOrganizationWirelessRadioAutoRfChannels InlineObject275) AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest {
	r.recalculateOrganizationWirelessRadioAutoRfChannels = &recalculateOrganizationWirelessRadioAutoRfChannels
	return r
}

func (r AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest) Execute() (*InlineResponse200328, *http.Response, error) {
	return r.ApiService.RecalculateOrganizationWirelessRadioAutoRfChannelsExecute(r)
}

/*
RecalculateOrganizationWirelessRadioAutoRfChannels Recalculates automatically assigned channels for every AP within specified the specified network(s)

Recalculates automatically assigned channels for every AP within specified the specified network(s). Note: This could cause a brief loss in connectivity for wireless clients.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest
*/
func (a *AutoRfApiService) RecalculateOrganizationWirelessRadioAutoRfChannels(ctx context.Context, organizationId string) AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest {
	return AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200328
func (a *AutoRfApiService) RecalculateOrganizationWirelessRadioAutoRfChannelsExecute(r AutoRfApiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest) (*InlineResponse200328, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200328
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutoRfApiService.RecalculateOrganizationWirelessRadioAutoRfChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/radio/autoRf/channels/recalculate"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recalculateOrganizationWirelessRadioAutoRfChannels == nil {
		return localVarReturnValue, nil, reportError("recalculateOrganizationWirelessRadioAutoRfChannels is required and must be specified")
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
	localVarPostBody = r.recalculateOrganizationWirelessRadioAutoRfChannels
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
