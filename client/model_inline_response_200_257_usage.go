/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200257Usage Usage information of all clients across organization
type InlineResponse200257Usage struct {
	Overall *InlineResponse200257UsageOverall `json:"overall,omitempty"`
	// Average data usage (in kb) of each client in organization
	Average *float32 `json:"average,omitempty"`
}

// NewInlineResponse200257Usage instantiates a new InlineResponse200257Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200257Usage() *InlineResponse200257Usage {
	this := InlineResponse200257Usage{}
	return &this
}

// NewInlineResponse200257UsageWithDefaults instantiates a new InlineResponse200257Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200257UsageWithDefaults() *InlineResponse200257Usage {
	this := InlineResponse200257Usage{}
	return &this
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *InlineResponse200257Usage) GetOverall() InlineResponse200257UsageOverall {
	if o == nil || isNil(o.Overall) {
		var ret InlineResponse200257UsageOverall
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200257Usage) GetOverallOk() (*InlineResponse200257UsageOverall, bool) {
	if o == nil || isNil(o.Overall) {
    return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *InlineResponse200257Usage) HasOverall() bool {
	if o != nil && !isNil(o.Overall) {
		return true
	}

	return false
}

// SetOverall gets a reference to the given InlineResponse200257UsageOverall and assigns it to the Overall field.
func (o *InlineResponse200257Usage) SetOverall(v InlineResponse200257UsageOverall) {
	o.Overall = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *InlineResponse200257Usage) GetAverage() float32 {
	if o == nil || isNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200257Usage) GetAverageOk() (*float32, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *InlineResponse200257Usage) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *InlineResponse200257Usage) SetAverage(v float32) {
	o.Average = &v
}

func (o InlineResponse200257Usage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Overall) {
		toSerialize["overall"] = o.Overall
	}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200257Usage struct {
	value *InlineResponse200257Usage
	isSet bool
}

func (v NullableInlineResponse200257Usage) Get() *InlineResponse200257Usage {
	return v.value
}

func (v *NullableInlineResponse200257Usage) Set(val *InlineResponse200257Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200257Usage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200257Usage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200257Usage(val *InlineResponse200257Usage) *NullableInlineResponse200257Usage {
	return &NullableInlineResponse200257Usage{value: val, isSet: true}
}

func (v NullableInlineResponse200257Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200257Usage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


