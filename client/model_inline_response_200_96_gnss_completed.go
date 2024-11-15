/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20096GnssCompleted Progress information for the GNSS acquisition process
type InlineResponse20096GnssCompleted struct {
	// Completion percentage of the GNSS acquisition process
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewInlineResponse20096GnssCompleted instantiates a new InlineResponse20096GnssCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20096GnssCompleted() *InlineResponse20096GnssCompleted {
	this := InlineResponse20096GnssCompleted{}
	return &this
}

// NewInlineResponse20096GnssCompletedWithDefaults instantiates a new InlineResponse20096GnssCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20096GnssCompletedWithDefaults() *InlineResponse20096GnssCompleted {
	this := InlineResponse20096GnssCompleted{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse20096GnssCompleted) GetPercentage() int32 {
	if o == nil || isNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20096GnssCompleted) GetPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse20096GnssCompleted) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *InlineResponse20096GnssCompleted) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o InlineResponse20096GnssCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20096GnssCompleted struct {
	value *InlineResponse20096GnssCompleted
	isSet bool
}

func (v NullableInlineResponse20096GnssCompleted) Get() *InlineResponse20096GnssCompleted {
	return v.value
}

func (v *NullableInlineResponse20096GnssCompleted) Set(val *InlineResponse20096GnssCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20096GnssCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20096GnssCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20096GnssCompleted(val *InlineResponse20096GnssCompleted) *NullableInlineResponse20096GnssCompleted {
	return &NullableInlineResponse20096GnssCompleted{value: val, isSet: true}
}

func (v NullableInlineResponse20096GnssCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20096GnssCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


