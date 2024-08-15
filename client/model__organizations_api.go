/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsApi API related settings
type OrganizationsApi struct {
	// Enable API access
	Enabled *bool `json:"enabled,omitempty"`
}

// NewOrganizationsApi instantiates a new OrganizationsApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsApi() *OrganizationsApi {
	this := OrganizationsApi{}
	return &this
}

// NewOrganizationsApiWithDefaults instantiates a new OrganizationsApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsApiWithDefaults() *OrganizationsApi {
	this := OrganizationsApi{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrganizationsApi) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsApi) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrganizationsApi) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrganizationsApi) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrganizationsApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsApi struct {
	value *OrganizationsApi
	isSet bool
}

func (v NullableOrganizationsApi) Get() *OrganizationsApi {
	return v.value
}

func (v *NullableOrganizationsApi) Set(val *OrganizationsApi) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsApi) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsApi(val *OrganizationsApi) *NullableOrganizationsApi {
	return &NullableOrganizationsApi{value: val, isSet: true}
}

func (v NullableOrganizationsApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


