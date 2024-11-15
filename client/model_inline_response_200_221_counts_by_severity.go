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

// InlineResponse200221CountsBySeverity struct for InlineResponse200221CountsBySeverity
type InlineResponse200221CountsBySeverity struct {
	// Severity Type
	Type string `json:"type"`
	// Total count of the given severity type
	Count int32 `json:"count"`
}

// NewInlineResponse200221CountsBySeverity instantiates a new InlineResponse200221CountsBySeverity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200221CountsBySeverity(type_ string, count int32) *InlineResponse200221CountsBySeverity {
	this := InlineResponse200221CountsBySeverity{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200221CountsBySeverityWithDefaults instantiates a new InlineResponse200221CountsBySeverity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200221CountsBySeverityWithDefaults() *InlineResponse200221CountsBySeverity {
	this := InlineResponse200221CountsBySeverity{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200221CountsBySeverity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221CountsBySeverity) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200221CountsBySeverity) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200221CountsBySeverity) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221CountsBySeverity) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200221CountsBySeverity) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200221CountsBySeverity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200221CountsBySeverity struct {
	value *InlineResponse200221CountsBySeverity
	isSet bool
}

func (v NullableInlineResponse200221CountsBySeverity) Get() *InlineResponse200221CountsBySeverity {
	return v.value
}

func (v *NullableInlineResponse200221CountsBySeverity) Set(val *InlineResponse200221CountsBySeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200221CountsBySeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200221CountsBySeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200221CountsBySeverity(val *InlineResponse200221CountsBySeverity) *NullableInlineResponse200221CountsBySeverity {
	return &NullableInlineResponse200221CountsBySeverity{value: val, isSet: true}
}

func (v NullableInlineResponse200221CountsBySeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200221CountsBySeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


