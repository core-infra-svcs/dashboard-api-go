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

// InlineResponse200311Energy How much energy (in watt-hours) has been delivered by this port during the interval.
type InlineResponse200311Energy struct {
	Usage *InlineResponse200311EnergyUsage `json:"usage,omitempty"`
}

// NewInlineResponse200311Energy instantiates a new InlineResponse200311Energy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200311Energy() *InlineResponse200311Energy {
	this := InlineResponse200311Energy{}
	return &this
}

// NewInlineResponse200311EnergyWithDefaults instantiates a new InlineResponse200311Energy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200311EnergyWithDefaults() *InlineResponse200311Energy {
	this := InlineResponse200311Energy{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200311Energy) GetUsage() InlineResponse200311EnergyUsage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200311EnergyUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200311Energy) GetUsageOk() (*InlineResponse200311EnergyUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200311Energy) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200311EnergyUsage and assigns it to the Usage field.
func (o *InlineResponse200311Energy) SetUsage(v InlineResponse200311EnergyUsage) {
	o.Usage = &v
}

func (o InlineResponse200311Energy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200311Energy struct {
	value *InlineResponse200311Energy
	isSet bool
}

func (v NullableInlineResponse200311Energy) Get() *InlineResponse200311Energy {
	return v.value
}

func (v *NullableInlineResponse200311Energy) Set(val *InlineResponse200311Energy) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200311Energy) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200311Energy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200311Energy(val *InlineResponse200311Energy) *NullableInlineResponse200311Energy {
	return &NullableInlineResponse200311Energy{value: val, isSet: true}
}

func (v NullableInlineResponse200311Energy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200311Energy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

