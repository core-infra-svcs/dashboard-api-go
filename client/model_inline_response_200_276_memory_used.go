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

// InlineResponse200276MemoryUsed Information regarding memory usage on the device over the interval
type InlineResponse200276MemoryUsed struct {
	// Minimum memory in kB used on the device over the interval
	Minimum *int32 `json:"minimum,omitempty"`
	// Maximum memory in kB used on the device over the interval
	Maximum *int32 `json:"maximum,omitempty"`
	// Median memory in kB used on the device over the interval rounded up to nearest integer
	Median *int32 `json:"median,omitempty"`
	Percentages *InlineResponse200276MemoryUsedPercentages `json:"percentages,omitempty"`
}

// NewInlineResponse200276MemoryUsed instantiates a new InlineResponse200276MemoryUsed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276MemoryUsed() *InlineResponse200276MemoryUsed {
	this := InlineResponse200276MemoryUsed{}
	return &this
}

// NewInlineResponse200276MemoryUsedWithDefaults instantiates a new InlineResponse200276MemoryUsed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276MemoryUsedWithDefaults() *InlineResponse200276MemoryUsed {
	this := InlineResponse200276MemoryUsed{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *InlineResponse200276MemoryUsed) GetMinimum() int32 {
	if o == nil || isNil(o.Minimum) {
		var ret int32
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276MemoryUsed) GetMinimumOk() (*int32, bool) {
	if o == nil || isNil(o.Minimum) {
    return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *InlineResponse200276MemoryUsed) HasMinimum() bool {
	if o != nil && !isNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given int32 and assigns it to the Minimum field.
func (o *InlineResponse200276MemoryUsed) SetMinimum(v int32) {
	o.Minimum = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *InlineResponse200276MemoryUsed) GetMaximum() int32 {
	if o == nil || isNil(o.Maximum) {
		var ret int32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276MemoryUsed) GetMaximumOk() (*int32, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *InlineResponse200276MemoryUsed) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given int32 and assigns it to the Maximum field.
func (o *InlineResponse200276MemoryUsed) SetMaximum(v int32) {
	o.Maximum = &v
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *InlineResponse200276MemoryUsed) GetMedian() int32 {
	if o == nil || isNil(o.Median) {
		var ret int32
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276MemoryUsed) GetMedianOk() (*int32, bool) {
	if o == nil || isNil(o.Median) {
    return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *InlineResponse200276MemoryUsed) HasMedian() bool {
	if o != nil && !isNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given int32 and assigns it to the Median field.
func (o *InlineResponse200276MemoryUsed) SetMedian(v int32) {
	o.Median = &v
}

// GetPercentages returns the Percentages field value if set, zero value otherwise.
func (o *InlineResponse200276MemoryUsed) GetPercentages() InlineResponse200276MemoryUsedPercentages {
	if o == nil || isNil(o.Percentages) {
		var ret InlineResponse200276MemoryUsedPercentages
		return ret
	}
	return *o.Percentages
}

// GetPercentagesOk returns a tuple with the Percentages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276MemoryUsed) GetPercentagesOk() (*InlineResponse200276MemoryUsedPercentages, bool) {
	if o == nil || isNil(o.Percentages) {
    return nil, false
	}
	return o.Percentages, true
}

// HasPercentages returns a boolean if a field has been set.
func (o *InlineResponse200276MemoryUsed) HasPercentages() bool {
	if o != nil && !isNil(o.Percentages) {
		return true
	}

	return false
}

// SetPercentages gets a reference to the given InlineResponse200276MemoryUsedPercentages and assigns it to the Percentages field.
func (o *InlineResponse200276MemoryUsed) SetPercentages(v InlineResponse200276MemoryUsedPercentages) {
	o.Percentages = &v
}

func (o InlineResponse200276MemoryUsed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	if !isNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	if !isNil(o.Percentages) {
		toSerialize["percentages"] = o.Percentages
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276MemoryUsed struct {
	value *InlineResponse200276MemoryUsed
	isSet bool
}

func (v NullableInlineResponse200276MemoryUsed) Get() *InlineResponse200276MemoryUsed {
	return v.value
}

func (v *NullableInlineResponse200276MemoryUsed) Set(val *InlineResponse200276MemoryUsed) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276MemoryUsed) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276MemoryUsed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276MemoryUsed(val *InlineResponse200276MemoryUsed) *NullableInlineResponse200276MemoryUsed {
	return &NullableInlineResponse200276MemoryUsed{value: val, isSet: true}
}

func (v NullableInlineResponse200276MemoryUsed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276MemoryUsed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


