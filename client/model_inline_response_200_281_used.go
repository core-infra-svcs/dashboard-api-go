/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200281Used Information regarding memory usage on the device over the entire timespan
type InlineResponse200281Used struct {
	// Median memory in kB used on the device over the entire timespan rounded up to nearest integer
	Median *int32 `json:"median,omitempty"`
}

// NewInlineResponse200281Used instantiates a new InlineResponse200281Used object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200281Used() *InlineResponse200281Used {
	this := InlineResponse200281Used{}
	return &this
}

// NewInlineResponse200281UsedWithDefaults instantiates a new InlineResponse200281Used object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200281UsedWithDefaults() *InlineResponse200281Used {
	this := InlineResponse200281Used{}
	return &this
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *InlineResponse200281Used) GetMedian() int32 {
	if o == nil || isNil(o.Median) {
		var ret int32
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200281Used) GetMedianOk() (*int32, bool) {
	if o == nil || isNil(o.Median) {
    return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *InlineResponse200281Used) HasMedian() bool {
	if o != nil && !isNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given int32 and assigns it to the Median field.
func (o *InlineResponse200281Used) SetMedian(v int32) {
	o.Median = &v
}

func (o InlineResponse200281Used) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200281Used struct {
	value *InlineResponse200281Used
	isSet bool
}

func (v NullableInlineResponse200281Used) Get() *InlineResponse200281Used {
	return v.value
}

func (v *NullableInlineResponse200281Used) Set(val *InlineResponse200281Used) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200281Used) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200281Used) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200281Used(val *InlineResponse200281Used) *NullableInlineResponse200281Used {
	return &NullableInlineResponse200281Used{value: val, isSet: true}
}

func (v NullableInlineResponse200281Used) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200281Used) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


