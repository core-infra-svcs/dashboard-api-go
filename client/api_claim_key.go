/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// ClaimKeyApiService ClaimKeyApi service
type ClaimKeyApiService service

type ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct {
	ctx context.Context
	ApiService *ClaimKeyApiService
	validateAdministeredLicensingSubscriptionSubscriptionsClaimKey *InlineObject1
}

func (r ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(validateAdministeredLicensingSubscriptionSubscriptionsClaimKey InlineObject1) ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKey = &validateAdministeredLicensingSubscriptionSubscriptionsClaimKey
	return r
}

func (r ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) Execute() (*InlineResponse2003, *http.Response, error) {
	return r.ApiService.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyExecute(r)
}

/*
ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey Find a subscription by claim key

Find a subscription by claim key. Returns 400 if the key has already been claimed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
*/
func (a *ClaimKeyApiService) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx context.Context) ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	return ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InlineResponse2003
func (a *ClaimKeyApiService) ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyExecute(r ClaimKeyApiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) (*InlineResponse2003, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClaimKeyApiService.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/administered/licensing/subscription/subscriptions/claimKey/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKey == nil {
		return localVarReturnValue, nil, reportError("validateAdministeredLicensingSubscriptionSubscriptionsClaimKey is required and must be specified")
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
	localVarPostBody = r.validateAdministeredLicensingSubscriptionSubscriptionsClaimKey
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
