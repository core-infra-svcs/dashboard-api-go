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

// InlineResponse200234SeverityCounts struct for InlineResponse200234SeverityCounts
type InlineResponse200234SeverityCounts struct {
	// Type
	Type string `json:"type"`
	// Count
	Count int32 `json:"count"`
}

// NewInlineResponse200234SeverityCounts instantiates a new InlineResponse200234SeverityCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200234SeverityCounts(type_ string, count int32) *InlineResponse200234SeverityCounts {
	this := InlineResponse200234SeverityCounts{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200234SeverityCountsWithDefaults instantiates a new InlineResponse200234SeverityCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200234SeverityCountsWithDefaults() *InlineResponse200234SeverityCounts {
	this := InlineResponse200234SeverityCounts{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200234SeverityCounts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200234SeverityCounts) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200234SeverityCounts) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200234SeverityCounts) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200234SeverityCounts) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200234SeverityCounts) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200234SeverityCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200234SeverityCounts struct {
	value *InlineResponse200234SeverityCounts
	isSet bool
}

func (v NullableInlineResponse200234SeverityCounts) Get() *InlineResponse200234SeverityCounts {
	return v.value
}

func (v *NullableInlineResponse200234SeverityCounts) Set(val *InlineResponse200234SeverityCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200234SeverityCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200234SeverityCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200234SeverityCounts(val *InlineResponse200234SeverityCounts) *NullableInlineResponse200234SeverityCounts {
	return &NullableInlineResponse200234SeverityCounts{value: val, isSet: true}
}

func (v NullableInlineResponse200234SeverityCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200234SeverityCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


