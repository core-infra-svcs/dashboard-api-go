/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
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


// InternetPoliciesApiService InternetPoliciesApi service
type InternetPoliciesApiService service

type InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest struct {
	ctx context.Context
	ApiService *InternetPoliciesApiService
	networkId string
	updateNetworkApplianceSdwanInternetPolicies *InlineObject58
}

func (r InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest) UpdateNetworkApplianceSdwanInternetPolicies(updateNetworkApplianceSdwanInternetPolicies InlineObject58) InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest {
	r.updateNetworkApplianceSdwanInternetPolicies = &updateNetworkApplianceSdwanInternetPolicies
	return r
}

func (r InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest) Execute() (*InlineResponse20060, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceSdwanInternetPoliciesExecute(r)
}

/*
UpdateNetworkApplianceSdwanInternetPolicies Update SDWAN internet traffic preferences for an MX network

Update SDWAN internet traffic preferences for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest
*/
func (a *InternetPoliciesApiService) UpdateNetworkApplianceSdwanInternetPolicies(ctx context.Context, networkId string) InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest {
	return InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return InlineResponse20060
func (a *InternetPoliciesApiService) UpdateNetworkApplianceSdwanInternetPoliciesExecute(r InternetPoliciesApiUpdateNetworkApplianceSdwanInternetPoliciesRequest) (*InlineResponse20060, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse20060
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InternetPoliciesApiService.UpdateNetworkApplianceSdwanInternetPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/sdwan/internetPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	localVarPostBody = r.updateNetworkApplianceSdwanInternetPolicies
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
