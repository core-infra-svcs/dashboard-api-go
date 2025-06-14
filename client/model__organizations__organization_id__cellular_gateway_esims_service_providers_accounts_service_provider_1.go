/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 Service Provider information
type OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 struct {
	// Service provider name
	Name *string `json:"name,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1WithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1WithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) SetName(v string) {
	o.Name = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) Get() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1 {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


