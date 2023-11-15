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

// InlineResponse2004ResultsLoss Lost packets
type InlineResponse2004ResultsLoss struct {
	// Percentage of packets lost
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewInlineResponse2004ResultsLoss instantiates a new InlineResponse2004ResultsLoss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004ResultsLoss() *InlineResponse2004ResultsLoss {
	this := InlineResponse2004ResultsLoss{}
	return &this
}

// NewInlineResponse2004ResultsLossWithDefaults instantiates a new InlineResponse2004ResultsLoss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004ResultsLossWithDefaults() *InlineResponse2004ResultsLoss {
	this := InlineResponse2004ResultsLoss{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse2004ResultsLoss) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004ResultsLoss) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse2004ResultsLoss) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InlineResponse2004ResultsLoss) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o InlineResponse2004ResultsLoss) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004ResultsLoss struct {
	value *InlineResponse2004ResultsLoss
	isSet bool
}

func (v NullableInlineResponse2004ResultsLoss) Get() *InlineResponse2004ResultsLoss {
	return v.value
}

func (v *NullableInlineResponse2004ResultsLoss) Set(val *InlineResponse2004ResultsLoss) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004ResultsLoss) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004ResultsLoss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004ResultsLoss(val *InlineResponse2004ResultsLoss) *NullableInlineResponse2004ResultsLoss {
	return &NullableInlineResponse2004ResultsLoss{value: val, isSet: true}
}

func (v NullableInlineResponse2004ResultsLoss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004ResultsLoss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

