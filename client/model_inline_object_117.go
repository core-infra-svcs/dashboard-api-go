/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject117 struct for InlineObject117
type InlineObject117 struct {
	// Set to true to enable MQTT broker for sensor network
	Enabled bool `json:"enabled"`
}

// NewInlineObject117 instantiates a new InlineObject117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject117(enabled bool) *InlineObject117 {
	this := InlineObject117{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject117WithDefaults instantiates a new InlineObject117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject117WithDefaults() *InlineObject117 {
	this := InlineObject117{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject117) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject117) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject117) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineObject117) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject117 struct {
	value *InlineObject117
	isSet bool
}

func (v NullableInlineObject117) Get() *InlineObject117 {
	return v.value
}

func (v *NullableInlineObject117) Set(val *InlineObject117) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject117) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject117(val *InlineObject117) *NullableInlineObject117 {
	return &NullableInlineObject117{value: val, isSet: true}
}

func (v NullableInlineObject117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


