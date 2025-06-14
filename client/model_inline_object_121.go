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

// InlineObject121 struct for InlineObject121
type InlineObject121 struct {
	// Set to true to enable MQTT broker for sensor network
	Enabled bool `json:"enabled"`
}

// NewInlineObject121 instantiates a new InlineObject121 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject121(enabled bool) *InlineObject121 {
	this := InlineObject121{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject121WithDefaults instantiates a new InlineObject121 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject121WithDefaults() *InlineObject121 {
	this := InlineObject121{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject121) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject121) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject121) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineObject121) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject121 struct {
	value *InlineObject121
	isSet bool
}

func (v NullableInlineObject121) Get() *InlineObject121 {
	return v.value
}

func (v *NullableInlineObject121) Set(val *InlineObject121) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject121) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject121) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject121(val *InlineObject121) *NullableInlineObject121 {
	return &NullableInlineObject121{value: val, isSet: true}
}

func (v NullableInlineObject121) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject121) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


