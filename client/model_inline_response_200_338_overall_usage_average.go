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

// InlineResponse200338OverallUsageAverage The average CPU usage of the wireless LAN controller
type InlineResponse200338OverallUsageAverage struct {
	// The CPU usage percentage of the wireless LAN controller
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewInlineResponse200338OverallUsageAverage instantiates a new InlineResponse200338OverallUsageAverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200338OverallUsageAverage() *InlineResponse200338OverallUsageAverage {
	this := InlineResponse200338OverallUsageAverage{}
	return &this
}

// NewInlineResponse200338OverallUsageAverageWithDefaults instantiates a new InlineResponse200338OverallUsageAverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200338OverallUsageAverageWithDefaults() *InlineResponse200338OverallUsageAverage {
	this := InlineResponse200338OverallUsageAverage{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *InlineResponse200338OverallUsageAverage) GetPercentage() float32 {
	if o == nil || isNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338OverallUsageAverage) GetPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.Percentage) {
    return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *InlineResponse200338OverallUsageAverage) HasPercentage() bool {
	if o != nil && !isNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *InlineResponse200338OverallUsageAverage) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o InlineResponse200338OverallUsageAverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200338OverallUsageAverage struct {
	value *InlineResponse200338OverallUsageAverage
	isSet bool
}

func (v NullableInlineResponse200338OverallUsageAverage) Get() *InlineResponse200338OverallUsageAverage {
	return v.value
}

func (v *NullableInlineResponse200338OverallUsageAverage) Set(val *InlineResponse200338OverallUsageAverage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200338OverallUsageAverage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200338OverallUsageAverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200338OverallUsageAverage(val *InlineResponse200338OverallUsageAverage) *NullableInlineResponse200338OverallUsageAverage {
	return &NullableInlineResponse200338OverallUsageAverage{value: val, isSet: true}
}

func (v NullableInlineResponse200338OverallUsageAverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200338OverallUsageAverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

