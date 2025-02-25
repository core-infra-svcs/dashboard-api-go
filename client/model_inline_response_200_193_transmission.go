/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200193Transmission Settings related to radio transmission.
type InlineResponse200193Transmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200193Transmission instantiates a new InlineResponse200193Transmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200193Transmission() *InlineResponse200193Transmission {
	this := InlineResponse200193Transmission{}
	return &this
}

// NewInlineResponse200193TransmissionWithDefaults instantiates a new InlineResponse200193Transmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200193TransmissionWithDefaults() *InlineResponse200193Transmission {
	this := InlineResponse200193Transmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200193Transmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200193Transmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200193Transmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200193Transmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200193Transmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200193Transmission struct {
	value *InlineResponse200193Transmission
	isSet bool
}

func (v NullableInlineResponse200193Transmission) Get() *InlineResponse200193Transmission {
	return v.value
}

func (v *NullableInlineResponse200193Transmission) Set(val *InlineResponse200193Transmission) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200193Transmission) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200193Transmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200193Transmission(val *InlineResponse200193Transmission) *NullableInlineResponse200193Transmission {
	return &NullableInlineResponse200193Transmission{value: val, isSet: true}
}

func (v NullableInlineResponse200193Transmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200193Transmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


