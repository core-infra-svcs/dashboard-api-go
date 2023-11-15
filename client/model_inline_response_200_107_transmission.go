/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200107Transmission Settings related to radio transmission.
type InlineResponse200107Transmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200107Transmission instantiates a new InlineResponse200107Transmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200107Transmission() *InlineResponse200107Transmission {
	this := InlineResponse200107Transmission{}
	return &this
}

// NewInlineResponse200107TransmissionWithDefaults instantiates a new InlineResponse200107Transmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200107TransmissionWithDefaults() *InlineResponse200107Transmission {
	this := InlineResponse200107Transmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200107Transmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200107Transmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200107Transmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200107Transmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200107Transmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200107Transmission struct {
	value *InlineResponse200107Transmission
	isSet bool
}

func (v NullableInlineResponse200107Transmission) Get() *InlineResponse200107Transmission {
	return v.value
}

func (v *NullableInlineResponse200107Transmission) Set(val *InlineResponse200107Transmission) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200107Transmission) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200107Transmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200107Transmission(val *InlineResponse200107Transmission) *NullableInlineResponse200107Transmission {
	return &NullableInlineResponse200107Transmission{value: val, isSet: true}
}

func (v NullableInlineResponse200107Transmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200107Transmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


