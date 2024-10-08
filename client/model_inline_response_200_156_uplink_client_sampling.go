/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200156UplinkClientSampling Uplink client sampling
type InlineResponse200156UplinkClientSampling struct {
	// Enable client sampling on uplink
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200156UplinkClientSampling instantiates a new InlineResponse200156UplinkClientSampling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200156UplinkClientSampling() *InlineResponse200156UplinkClientSampling {
	this := InlineResponse200156UplinkClientSampling{}
	return &this
}

// NewInlineResponse200156UplinkClientSamplingWithDefaults instantiates a new InlineResponse200156UplinkClientSampling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200156UplinkClientSamplingWithDefaults() *InlineResponse200156UplinkClientSampling {
	this := InlineResponse200156UplinkClientSampling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200156UplinkClientSampling) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200156UplinkClientSampling) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200156UplinkClientSampling) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200156UplinkClientSampling) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200156UplinkClientSampling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200156UplinkClientSampling struct {
	value *InlineResponse200156UplinkClientSampling
	isSet bool
}

func (v NullableInlineResponse200156UplinkClientSampling) Get() *InlineResponse200156UplinkClientSampling {
	return v.value
}

func (v *NullableInlineResponse200156UplinkClientSampling) Set(val *InlineResponse200156UplinkClientSampling) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200156UplinkClientSampling) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200156UplinkClientSampling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200156UplinkClientSampling(val *InlineResponse200156UplinkClientSampling) *NullableInlineResponse200156UplinkClientSampling {
	return &NullableInlineResponse200156UplinkClientSampling{value: val, isSet: true}
}

func (v NullableInlineResponse200156UplinkClientSampling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200156UplinkClientSampling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


