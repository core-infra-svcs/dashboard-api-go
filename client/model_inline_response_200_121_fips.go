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

// InlineResponse200121Fips A hash of FIPS options applied to the Network
type InlineResponse200121Fips struct {
	// Enables / disables FIPS on the network.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200121Fips instantiates a new InlineResponse200121Fips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200121Fips() *InlineResponse200121Fips {
	this := InlineResponse200121Fips{}
	return &this
}

// NewInlineResponse200121FipsWithDefaults instantiates a new InlineResponse200121Fips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200121FipsWithDefaults() *InlineResponse200121Fips {
	this := InlineResponse200121Fips{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200121Fips) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200121Fips) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200121Fips) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200121Fips) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200121Fips) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200121Fips struct {
	value *InlineResponse200121Fips
	isSet bool
}

func (v NullableInlineResponse200121Fips) Get() *InlineResponse200121Fips {
	return v.value
}

func (v *NullableInlineResponse200121Fips) Set(val *InlineResponse200121Fips) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200121Fips) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200121Fips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200121Fips(val *InlineResponse200121Fips) *NullableInlineResponse200121Fips {
	return &NullableInlineResponse200121Fips{value: val, isSet: true}
}

func (v NullableInlineResponse200121Fips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200121Fips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


