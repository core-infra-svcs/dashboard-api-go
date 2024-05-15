/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject111 struct for InlineObject111
type InlineObject111 struct {
	// Set to true to enable MQTT broker for sensor network
	Enabled bool `json:"enabled"`
}

// NewInlineObject111 instantiates a new InlineObject111 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject111(enabled bool) *InlineObject111 {
	this := InlineObject111{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject111WithDefaults instantiates a new InlineObject111 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject111WithDefaults() *InlineObject111 {
	this := InlineObject111{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject111) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject111) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject111) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineObject111) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject111 struct {
	value *InlineObject111
	isSet bool
}

func (v NullableInlineObject111) Get() *InlineObject111 {
	return v.value
}

func (v *NullableInlineObject111) Set(val *InlineObject111) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject111) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject111) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject111(val *InlineObject111) *NullableInlineObject111 {
	return &NullableInlineObject111{value: val, isSet: true}
}

func (v NullableInlineObject111) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject111) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


