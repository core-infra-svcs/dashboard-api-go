/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject240 struct for InlineObject240
type InlineObject240 struct {
	// Service provider account ID
	AccountId string `json:"accountId"`
	// Service provider account API key
	ApiKey string `json:"apiKey"`
	ServiceProvider OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 `json:"serviceProvider"`
	// Service provider account name
	Title string `json:"title"`
	// Service provider account username
	Username string `json:"username"`
}

// NewInlineObject240 instantiates a new InlineObject240 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject240(accountId string, apiKey string, serviceProvider OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1, title string, username string) *InlineObject240 {
	this := InlineObject240{}
	this.AccountId = accountId
	this.ApiKey = apiKey
	this.ServiceProvider = serviceProvider
	this.Title = title
	this.Username = username
	return &this
}

// NewInlineObject240WithDefaults instantiates a new InlineObject240 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject240WithDefaults() *InlineObject240 {
	this := InlineObject240{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *InlineObject240) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetAccountIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *InlineObject240) SetAccountId(v string) {
	o.AccountId = v
}

// GetApiKey returns the ApiKey field value
func (o *InlineObject240) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetApiKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *InlineObject240) SetApiKey(v string) {
	o.ApiKey = v
}

// GetServiceProvider returns the ServiceProvider field value
func (o *InlineObject240) GetServiceProvider() OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 {
	if o == nil {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1
		return ret
	}

	return o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetServiceProviderOk() (*OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceProvider, true
}

// SetServiceProvider sets field value
func (o *InlineObject240) SetServiceProvider(v OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) {
	o.ServiceProvider = v
}

// GetTitle returns the Title field value
func (o *InlineObject240) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineObject240) SetTitle(v string) {
	o.Title = v
}

// GetUsername returns the Username field value
func (o *InlineObject240) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InlineObject240) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InlineObject240) SetUsername(v string) {
	o.Username = v
}

func (o InlineObject240) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["apiKey"] = o.ApiKey
	}
	if true {
		toSerialize["serviceProvider"] = o.ServiceProvider
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject240 struct {
	value *InlineObject240
	isSet bool
}

func (v NullableInlineObject240) Get() *InlineObject240 {
	return v.value
}

func (v *NullableInlineObject240) Set(val *InlineObject240) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject240) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject240) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject240(val *InlineObject240) *NullableInlineObject240 {
	return &NullableInlineObject240{value: val, isSet: true}
}

func (v NullableInlineObject240) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject240) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


