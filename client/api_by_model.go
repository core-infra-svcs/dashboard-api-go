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


// ByModelApiService ByModelApi service
type ByModelApiService service

type ByModelApiGetOrganizationDevicesOverviewByModelRequest struct {
	ctx context.Context
	ApiService *ByModelApiService
	organizationId string
	models *[]string
	networkIds *[]string
	productTypes *[]string
}

// Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.
func (r ByModelApiGetOrganizationDevicesOverviewByModelRequest) Models(models []string) ByModelApiGetOrganizationDevicesOverviewByModelRequest {
	r.models = &models
	return r
}

// Optional parameter to filter devices by networkId.
func (r ByModelApiGetOrganizationDevicesOverviewByModelRequest) NetworkIds(networkIds []string) ByModelApiGetOrganizationDevicesOverviewByModelRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter device by device product types. This filter uses multiple exact matches.
func (r ByModelApiGetOrganizationDevicesOverviewByModelRequest) ProductTypes(productTypes []string) ByModelApiGetOrganizationDevicesOverviewByModelRequest {
	r.productTypes = &productTypes
	return r
}

func (r ByModelApiGetOrganizationDevicesOverviewByModelRequest) Execute() (*InlineResponse200252, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesOverviewByModelExecute(r)
}

/*
GetOrganizationDevicesOverviewByModel Lists the count for each device model

Lists the count for each device model

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByModelApiGetOrganizationDevicesOverviewByModelRequest
*/
func (a *ByModelApiService) GetOrganizationDevicesOverviewByModel(ctx context.Context, organizationId string) ByModelApiGetOrganizationDevicesOverviewByModelRequest {
	return ByModelApiGetOrganizationDevicesOverviewByModelRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200252
func (a *ByModelApiService) GetOrganizationDevicesOverviewByModelExecute(r ByModelApiGetOrganizationDevicesOverviewByModelRequest) (*InlineResponse200252, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200252
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByModelApiService.GetOrganizationDevicesOverviewByModel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/overview/byModel"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.models != nil {
		localVarQueryParams.Add("models", parameterToString(*r.models, "csv"))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	if r.productTypes != nil {
		localVarQueryParams.Add("productTypes", parameterToString(*r.productTypes, "csv"))
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
