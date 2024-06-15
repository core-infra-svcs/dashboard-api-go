/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200183Transmission Settings related to radio transmission.
type InlineResponse200183Transmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200183Transmission instantiates a new InlineResponse200183Transmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200183Transmission() *InlineResponse200183Transmission {
	this := InlineResponse200183Transmission{}
	return &this
}

// NewInlineResponse200183TransmissionWithDefaults instantiates a new InlineResponse200183Transmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200183TransmissionWithDefaults() *InlineResponse200183Transmission {
	this := InlineResponse200183Transmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200183Transmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183Transmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200183Transmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200183Transmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200183Transmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200183Transmission struct {
	value *InlineResponse200183Transmission
	isSet bool
}

func (v NullableInlineResponse200183Transmission) Get() *InlineResponse200183Transmission {
	return v.value
}

func (v *NullableInlineResponse200183Transmission) Set(val *InlineResponse200183Transmission) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200183Transmission) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200183Transmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200183Transmission(val *InlineResponse200183Transmission) *NullableInlineResponse200183Transmission {
	return &NullableInlineResponse200183Transmission{value: val, isSet: true}
}

func (v NullableInlineResponse200183Transmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200183Transmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


