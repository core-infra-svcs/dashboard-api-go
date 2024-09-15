/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems struct for OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
type OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems struct {
	// Service provider account ID
	AccountId *string `json:"accountId,omitempty"`
	// Last updated at
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
	ServiceProvider *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider `json:"serviceProvider,omitempty"`
	// Service provider account name
	Title *string `json:"title,omitempty"`
	// Service provider account username
	Username *string `json:"username,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItemsWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItemsWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
    return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) SetAccountId(v string) {
	o.AccountId = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetLastUpdatedAt() string {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetServiceProvider() OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider {
	if o == nil || isNil(o.ServiceProvider) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetServiceProviderOk() (*OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider, bool) {
	if o == nil || isNil(o.ServiceProvider) {
    return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) HasServiceProvider() bool {
	if o != nil && !isNil(o.ServiceProvider) {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider and assigns it to the ServiceProvider field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) SetServiceProvider(v OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider) {
	o.ServiceProvider = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) SetTitle(v string) {
	o.Title = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) SetUsername(v string) {
	o.Username = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !isNil(o.ServiceProvider) {
		toSerialize["serviceProvider"] = o.ServiceProvider
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) Get() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


