/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200110Transmission Settings related to radio transmission.
type InlineResponse200110Transmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200110Transmission instantiates a new InlineResponse200110Transmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110Transmission() *InlineResponse200110Transmission {
	this := InlineResponse200110Transmission{}
	return &this
}

// NewInlineResponse200110TransmissionWithDefaults instantiates a new InlineResponse200110Transmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110TransmissionWithDefaults() *InlineResponse200110Transmission {
	this := InlineResponse200110Transmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200110Transmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110Transmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200110Transmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200110Transmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200110Transmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110Transmission struct {
	value *InlineResponse200110Transmission
	isSet bool
}

func (v NullableInlineResponse200110Transmission) Get() *InlineResponse200110Transmission {
	return v.value
}

func (v *NullableInlineResponse200110Transmission) Set(val *InlineResponse200110Transmission) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110Transmission) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110Transmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110Transmission(val *InlineResponse200110Transmission) *NullableInlineResponse200110Transmission {
	return &NullableInlineResponse200110Transmission{value: val, isSet: true}
}

func (v NullableInlineResponse200110Transmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110Transmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


