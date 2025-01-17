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

// InlineResponse20097RangingCompleted Progress information for the ranging process
type InlineResponse20097RangingCompleted struct {
	// Completion percentage of the ranging process
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewInlineResponse20097RangingCompleted instantiates a new InlineResponse20097RangingCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20097RangingCompleted() *InlineResponse20097RangingCompleted {
	this := InlineResponse20097RangingCompleted{}
	return &this
}

// NewInlineResponse20097RangingCompletedWithDefaults instantiates a new InlineResponse20097RangingCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20097RangingCompletedWithDefaults() *InlineResponse20097RangingCompleted {
	this := InlineResponse20097RangingCompleted{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse20097RangingCompleted) GetPercentage() int32 {
	if o == nil || isNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20097RangingCompleted) GetPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse20097RangingCompleted) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *InlineResponse20097RangingCompleted) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o InlineResponse20097RangingCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20097RangingCompleted struct {
	value *InlineResponse20097RangingCompleted
	isSet bool
}

func (v NullableInlineResponse20097RangingCompleted) Get() *InlineResponse20097RangingCompleted {
	return v.value
}

func (v *NullableInlineResponse20097RangingCompleted) Set(val *InlineResponse20097RangingCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20097RangingCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20097RangingCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20097RangingCompleted(val *InlineResponse20097RangingCompleted) *NullableInlineResponse20097RangingCompleted {
	return &NullableInlineResponse20097RangingCompleted{value: val, isSet: true}
}

func (v NullableInlineResponse20097RangingCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20097RangingCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


