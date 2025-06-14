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
	"strings"
)


// IpsecApiService IpsecApi service
type IpsecApiService service

type IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct {
	ctx context.Context
	ApiService *IpsecApiService
	organizationId string
}

func (r IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest) Execute() (*InlineResponse200233, *http.Response, error) {
	return r.ApiService.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasExecute(r)
}

/*
GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas Get the list of available IPsec SLA policies for an organization

Get the list of available IPsec SLA policies for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest
*/
func (a *IpsecApiService) GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx context.Context, organizationId string) IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest {
	return IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200233
func (a *IpsecApiService) GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasExecute(r IpsecApiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest) (*InlineResponse200233, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200233
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IpsecApiService.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct {
	ctx context.Context
	ApiService *IpsecApiService
	organizationId string
	updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas *InlineObject233
}

func (r IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest) UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas InlineObject233) IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest {
	r.updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas = &updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas
	return r
}

func (r IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest) Execute() (*InlineResponse200234, *http.Response, error) {
	return r.ApiService.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasExecute(r)
}

/*
UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas Update the IPsec SLA policies for an organization

Update the IPsec SLA policies for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest
*/
func (a *IpsecApiService) UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx context.Context, organizationId string) IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest {
	return IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return InlineResponse200234
func (a *IpsecApiService) UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasExecute(r IpsecApiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest) (*InlineResponse200234, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse200234
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IpsecApiService.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	localVarPostBody = r.updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas
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
