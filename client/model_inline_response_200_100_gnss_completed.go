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

// InlineResponse200100GnssCompleted Progress information for the GNSS acquisition process
type InlineResponse200100GnssCompleted struct {
	// Completion percentage of the GNSS acquisition process
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewInlineResponse200100GnssCompleted instantiates a new InlineResponse200100GnssCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200100GnssCompleted() *InlineResponse200100GnssCompleted {
	this := InlineResponse200100GnssCompleted{}
	return &this
}

// NewInlineResponse200100GnssCompletedWithDefaults instantiates a new InlineResponse200100GnssCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200100GnssCompletedWithDefaults() *InlineResponse200100GnssCompleted {
	this := InlineResponse200100GnssCompleted{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse200100GnssCompleted) GetPercentage() int32 {
	if o == nil || isNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100GnssCompleted) GetPercentageOk() (*int32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse200100GnssCompleted) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *InlineResponse200100GnssCompleted) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o InlineResponse200100GnssCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200100GnssCompleted struct {
	value *InlineResponse200100GnssCompleted
	isSet bool
}

func (v NullableInlineResponse200100GnssCompleted) Get() *InlineResponse200100GnssCompleted {
	return v.value
}

func (v *NullableInlineResponse200100GnssCompleted) Set(val *InlineResponse200100GnssCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200100GnssCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200100GnssCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200100GnssCompleted(val *InlineResponse200100GnssCompleted) *NullableInlineResponse200100GnssCompleted {
	return &NullableInlineResponse200100GnssCompleted{value: val, isSet: true}
}

func (v NullableInlineResponse200100GnssCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200100GnssCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


