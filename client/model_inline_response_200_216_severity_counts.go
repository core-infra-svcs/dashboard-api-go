/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200216SeverityCounts struct for InlineResponse200216SeverityCounts
type InlineResponse200216SeverityCounts struct {
	// Type
	Type string `json:"type"`
	// Count
	Count int32 `json:"count"`
}

// NewInlineResponse200216SeverityCounts instantiates a new InlineResponse200216SeverityCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200216SeverityCounts(type_ string, count int32) *InlineResponse200216SeverityCounts {
	this := InlineResponse200216SeverityCounts{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200216SeverityCountsWithDefaults instantiates a new InlineResponse200216SeverityCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200216SeverityCountsWithDefaults() *InlineResponse200216SeverityCounts {
	this := InlineResponse200216SeverityCounts{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200216SeverityCounts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216SeverityCounts) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200216SeverityCounts) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200216SeverityCounts) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200216SeverityCounts) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200216SeverityCounts) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200216SeverityCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200216SeverityCounts struct {
	value *InlineResponse200216SeverityCounts
	isSet bool
}

func (v NullableInlineResponse200216SeverityCounts) Get() *InlineResponse200216SeverityCounts {
	return v.value
}

func (v *NullableInlineResponse200216SeverityCounts) Set(val *InlineResponse200216SeverityCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200216SeverityCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200216SeverityCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200216SeverityCounts(val *InlineResponse200216SeverityCounts) *NullableInlineResponse200216SeverityCounts {
	return &NullableInlineResponse200216SeverityCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200216SeverityCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200216SeverityCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


