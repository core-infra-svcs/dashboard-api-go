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

// InlineResponse200110NamedVlans A hash of Named VLANs options applied to the Network.
type InlineResponse200110NamedVlans struct {
	// Enables / disables Named VLANs on the Network.
	Enabled bool `json:"enabled"`
}

// NewInlineResponse200110NamedVlans instantiates a new InlineResponse200110NamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110NamedVlans(enabled bool) *InlineResponse200110NamedVlans {
	this := InlineResponse200110NamedVlans{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse200110NamedVlansWithDefaults instantiates a new InlineResponse200110NamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110NamedVlansWithDefaults() *InlineResponse200110NamedVlans {
	this := InlineResponse200110NamedVlans{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse200110NamedVlans) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200110NamedVlans) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse200110NamedVlans) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse200110NamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110NamedVlans struct {
	value *InlineResponse200110NamedVlans
	isSet bool
}

func (v NullableInlineResponse200110NamedVlans) Get() *InlineResponse200110NamedVlans {
	return v.value
}

func (v *NullableInlineResponse200110NamedVlans) Set(val *InlineResponse200110NamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110NamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110NamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110NamedVlans(val *InlineResponse200110NamedVlans) *NullableInlineResponse200110NamedVlans {
	return &NullableInlineResponse200110NamedVlans{value: val, isSet: true}
}

func (v NullableInlineResponse200110NamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110NamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

