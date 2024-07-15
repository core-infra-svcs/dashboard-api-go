/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200110SecurePort A hash of SecureConnect options applied to the Network.
type InlineResponse200110SecurePort struct {
	// Enables / disables SecureConnect on the network. Optional.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200110SecurePort instantiates a new InlineResponse200110SecurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110SecurePort() *InlineResponse200110SecurePort {
	this := InlineResponse200110SecurePort{}
	return &this
}

// NewInlineResponse200110SecurePortWithDefaults instantiates a new InlineResponse200110SecurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110SecurePortWithDefaults() *InlineResponse200110SecurePort {
	this := InlineResponse200110SecurePort{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200110SecurePort) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110SecurePort) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200110SecurePort) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200110SecurePort) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200110SecurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110SecurePort struct {
	value *InlineResponse200110SecurePort
	isSet bool
}

func (v NullableInlineResponse200110SecurePort) Get() *InlineResponse200110SecurePort {
	return v.value
}

func (v *NullableInlineResponse200110SecurePort) Set(val *InlineResponse200110SecurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110SecurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110SecurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110SecurePort(val *InlineResponse200110SecurePort) *NullableInlineResponse200110SecurePort {
	return &NullableInlineResponse200110SecurePort{value: val, isSet: true}
}

func (v NullableInlineResponse200110SecurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110SecurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


