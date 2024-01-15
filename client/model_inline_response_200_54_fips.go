/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20054Fips A hash of FIPS options applied to the Network
type InlineResponse20054Fips struct {
	// Enables / disables FIPS on the network.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20054Fips instantiates a new InlineResponse20054Fips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054Fips() *InlineResponse20054Fips {
	this := InlineResponse20054Fips{}
	return &this
}

// NewInlineResponse20054FipsWithDefaults instantiates a new InlineResponse20054Fips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054FipsWithDefaults() *InlineResponse20054Fips {
	this := InlineResponse20054Fips{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20054Fips) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054Fips) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20054Fips) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20054Fips) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20054Fips) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054Fips struct {
	value *InlineResponse20054Fips
	isSet bool
}

func (v NullableInlineResponse20054Fips) Get() *InlineResponse20054Fips {
	return v.value
}

func (v *NullableInlineResponse20054Fips) Set(val *InlineResponse20054Fips) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054Fips) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054Fips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054Fips(val *InlineResponse20054Fips) *NullableInlineResponse20054Fips {
	return &NullableInlineResponse20054Fips{value: val, isSet: true}
}

func (v NullableInlineResponse20054Fips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054Fips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

