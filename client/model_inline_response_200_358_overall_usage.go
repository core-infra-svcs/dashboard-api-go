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

// InlineResponse200358OverallUsage The CPU usage of the wireless LAN controller
type InlineResponse200358OverallUsage struct {
	Average *InlineResponse200358OverallUsageAverage `json:"average,omitempty"`
}

// NewInlineResponse200358OverallUsage instantiates a new InlineResponse200358OverallUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200358OverallUsage() *InlineResponse200358OverallUsage {
	this := InlineResponse200358OverallUsage{}
	return &this
}

// NewInlineResponse200358OverallUsageWithDefaults instantiates a new InlineResponse200358OverallUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200358OverallUsageWithDefaults() *InlineResponse200358OverallUsage {
	this := InlineResponse200358OverallUsage{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *InlineResponse200358OverallUsage) GetAverage() InlineResponse200358OverallUsageAverage {
	if o == nil || isNil(o.Average) {
		var ret InlineResponse200358OverallUsageAverage
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200358OverallUsage) GetAverageOk() (*InlineResponse200358OverallUsageAverage, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *InlineResponse200358OverallUsage) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given InlineResponse200358OverallUsageAverage and assigns it to the Average field.
func (o *InlineResponse200358OverallUsage) SetAverage(v InlineResponse200358OverallUsageAverage) {
	o.Average = &v
}

func (o InlineResponse200358OverallUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200358OverallUsage struct {
	value *InlineResponse200358OverallUsage
	isSet bool
}

func (v NullableInlineResponse200358OverallUsage) Get() *InlineResponse200358OverallUsage {
	return v.value
}

func (v *NullableInlineResponse200358OverallUsage) Set(val *InlineResponse200358OverallUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200358OverallUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200358OverallUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200358OverallUsage(val *InlineResponse200358OverallUsage) *NullableInlineResponse200358OverallUsage {
	return &NullableInlineResponse200358OverallUsage{value: val, isSet: true}
}

func (v NullableInlineResponse200358OverallUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200358OverallUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


