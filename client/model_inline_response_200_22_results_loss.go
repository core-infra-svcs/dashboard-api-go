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

// InlineResponse20022ResultsLoss Lost packets
type InlineResponse20022ResultsLoss struct {
	// Percentage of packets lost
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewInlineResponse20022ResultsLoss instantiates a new InlineResponse20022ResultsLoss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022ResultsLoss() *InlineResponse20022ResultsLoss {
	this := InlineResponse20022ResultsLoss{}
	return &this
}

// NewInlineResponse20022ResultsLossWithDefaults instantiates a new InlineResponse20022ResultsLoss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022ResultsLossWithDefaults() *InlineResponse20022ResultsLoss {
	this := InlineResponse20022ResultsLoss{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse20022ResultsLoss) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022ResultsLoss) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse20022ResultsLoss) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InlineResponse20022ResultsLoss) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o InlineResponse20022ResultsLoss) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022ResultsLoss struct {
	value *InlineResponse20022ResultsLoss
	isSet bool
}

func (v NullableInlineResponse20022ResultsLoss) Get() *InlineResponse20022ResultsLoss {
	return v.value
}

func (v *NullableInlineResponse20022ResultsLoss) Set(val *InlineResponse20022ResultsLoss) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022ResultsLoss) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022ResultsLoss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022ResultsLoss(val *InlineResponse20022ResultsLoss) *NullableInlineResponse20022ResultsLoss {
	return &NullableInlineResponse20022ResultsLoss{value: val, isSet: true}
}

func (v NullableInlineResponse20022ResultsLoss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022ResultsLoss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


