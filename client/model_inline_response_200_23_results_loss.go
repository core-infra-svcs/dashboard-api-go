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

// InlineResponse20023ResultsLoss Lost packets
type InlineResponse20023ResultsLoss struct {
	// Percentage of packets lost
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewInlineResponse20023ResultsLoss instantiates a new InlineResponse20023ResultsLoss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023ResultsLoss() *InlineResponse20023ResultsLoss {
	this := InlineResponse20023ResultsLoss{}
	return &this
}

// NewInlineResponse20023ResultsLossWithDefaults instantiates a new InlineResponse20023ResultsLoss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023ResultsLossWithDefaults() *InlineResponse20023ResultsLoss {
	this := InlineResponse20023ResultsLoss{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse20023ResultsLoss) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023ResultsLoss) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse20023ResultsLoss) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InlineResponse20023ResultsLoss) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o InlineResponse20023ResultsLoss) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023ResultsLoss struct {
	value *InlineResponse20023ResultsLoss
	isSet bool
}

func (v NullableInlineResponse20023ResultsLoss) Get() *InlineResponse20023ResultsLoss {
	return v.value
}

func (v *NullableInlineResponse20023ResultsLoss) Set(val *InlineResponse20023ResultsLoss) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023ResultsLoss) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023ResultsLoss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023ResultsLoss(val *InlineResponse20023ResultsLoss) *NullableInlineResponse20023ResultsLoss {
	return &NullableInlineResponse20023ResultsLoss{value: val, isSet: true}
}

func (v NullableInlineResponse20023ResultsLoss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023ResultsLoss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

