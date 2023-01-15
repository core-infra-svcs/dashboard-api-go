/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApi API-specific settings
type OrganizationsOrganizationIdApi struct {
	// If true, enable the access to the Cisco Meraki Dashboard API
	Enabled *bool `json:"enabled,omitempty"`
}

// NewOrganizationsOrganizationIdApi instantiates a new OrganizationsOrganizationIdApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApi() *OrganizationsOrganizationIdApi {
	this := OrganizationsOrganizationIdApi{}
	return &this
}

// NewOrganizationsOrganizationIdApiWithDefaults instantiates a new OrganizationsOrganizationIdApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApiWithDefaults() *OrganizationsOrganizationIdApi {
	this := OrganizationsOrganizationIdApi{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApi) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApi) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApi) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsOrganizationIdApi) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrganizationsOrganizationIdApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApi struct {
	value *OrganizationsOrganizationIdApi
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApi) Get() *OrganizationsOrganizationIdApi {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApi) Set(val *OrganizationsOrganizationIdApi) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApi) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApi(val *OrganizationsOrganizationIdApi) *NullableOrganizationsOrganizationIdApi {
	return &NullableOrganizationsOrganizationIdApi{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


