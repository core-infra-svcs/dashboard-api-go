/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200116SecurePort A hash of SecureConnect options applied to the Network.
type InlineResponse200116SecurePort struct {
	// Enables / disables SecureConnect on the network. Optional.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200116SecurePort instantiates a new InlineResponse200116SecurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200116SecurePort() *InlineResponse200116SecurePort {
	this := InlineResponse200116SecurePort{}
	return &this
}

// NewInlineResponse200116SecurePortWithDefaults instantiates a new InlineResponse200116SecurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200116SecurePortWithDefaults() *InlineResponse200116SecurePort {
	this := InlineResponse200116SecurePort{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200116SecurePort) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116SecurePort) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200116SecurePort) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200116SecurePort) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200116SecurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200116SecurePort struct {
	value *InlineResponse200116SecurePort
	isSet bool
}

func (v NullableInlineResponse200116SecurePort) Get() *InlineResponse200116SecurePort {
	return v.value
}

func (v *NullableInlineResponse200116SecurePort) Set(val *InlineResponse200116SecurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200116SecurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200116SecurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200116SecurePort(val *InlineResponse200116SecurePort) *NullableInlineResponse200116SecurePort {
	return &NullableInlineResponse200116SecurePort{value: val, isSet: true}
}

func (v NullableInlineResponse200116SecurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200116SecurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

