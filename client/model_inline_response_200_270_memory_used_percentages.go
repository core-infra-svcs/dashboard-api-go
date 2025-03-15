/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200270MemoryUsedPercentages Memory utilization percentages on the device over the interval
type InlineResponse200270MemoryUsedPercentages struct {
	// Maximum memory utilization percentage on the device over the interval
	Maximum *int32 `json:"maximum,omitempty"`
}

// NewInlineResponse200270MemoryUsedPercentages instantiates a new InlineResponse200270MemoryUsedPercentages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200270MemoryUsedPercentages() *InlineResponse200270MemoryUsedPercentages {
	this := InlineResponse200270MemoryUsedPercentages{}
	return &this
}

// NewInlineResponse200270MemoryUsedPercentagesWithDefaults instantiates a new InlineResponse200270MemoryUsedPercentages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200270MemoryUsedPercentagesWithDefaults() *InlineResponse200270MemoryUsedPercentages {
	this := InlineResponse200270MemoryUsedPercentages{}
	return &this
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *InlineResponse200270MemoryUsedPercentages) GetMaximum() int32 {
	if o == nil || isNil(o.Maximum) {
		var ret int32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270MemoryUsedPercentages) GetMaximumOk() (*int32, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *InlineResponse200270MemoryUsedPercentages) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given int32 and assigns it to the Maximum field.
func (o *InlineResponse200270MemoryUsedPercentages) SetMaximum(v int32) {
	o.Maximum = &v
}

func (o InlineResponse200270MemoryUsedPercentages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200270MemoryUsedPercentages struct {
	value *InlineResponse200270MemoryUsedPercentages
	isSet bool
}

func (v NullableInlineResponse200270MemoryUsedPercentages) Get() *InlineResponse200270MemoryUsedPercentages {
	return v.value
}

func (v *NullableInlineResponse200270MemoryUsedPercentages) Set(val *InlineResponse200270MemoryUsedPercentages) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200270MemoryUsedPercentages) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200270MemoryUsedPercentages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200270MemoryUsedPercentages(val *InlineResponse200270MemoryUsedPercentages) *NullableInlineResponse200270MemoryUsedPercentages {
	return &NullableInlineResponse200270MemoryUsedPercentages{value: val, isSet: true}
}

func (v NullableInlineResponse200270MemoryUsedPercentages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200270MemoryUsedPercentages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


