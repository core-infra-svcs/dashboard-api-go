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

// InlineResponse200338Usage The specific core CPU usage of the wireless LAN controller
type InlineResponse200338Usage struct {
	Average *InlineResponse200338UsageAverage `json:"average,omitempty"`
}

// NewInlineResponse200338Usage instantiates a new InlineResponse200338Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200338Usage() *InlineResponse200338Usage {
	this := InlineResponse200338Usage{}
	return &this
}

// NewInlineResponse200338UsageWithDefaults instantiates a new InlineResponse200338Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200338UsageWithDefaults() *InlineResponse200338Usage {
	this := InlineResponse200338Usage{}
	return &this
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *InlineResponse200338Usage) GetAverage() InlineResponse200338UsageAverage {
	if o == nil || isNil(o.Average) {
		var ret InlineResponse200338UsageAverage
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Usage) GetAverageOk() (*InlineResponse200338UsageAverage, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *InlineResponse200338Usage) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given InlineResponse200338UsageAverage and assigns it to the Average field.
func (o *InlineResponse200338Usage) SetAverage(v InlineResponse200338UsageAverage) {
	o.Average = &v
}

func (o InlineResponse200338Usage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200338Usage struct {
	value *InlineResponse200338Usage
	isSet bool
}

func (v NullableInlineResponse200338Usage) Get() *InlineResponse200338Usage {
	return v.value
}

func (v *NullableInlineResponse200338Usage) Set(val *InlineResponse200338Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200338Usage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200338Usage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200338Usage(val *InlineResponse200338Usage) *NullableInlineResponse200338Usage {
	return &NullableInlineResponse200338Usage{value: val, isSet: true}
}

func (v NullableInlineResponse200338Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200338Usage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


