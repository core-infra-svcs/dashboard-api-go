/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200218Items struct for InlineResponse200218Items
type InlineResponse200218Items struct {
	// Alert Type
	Type string `json:"type"`
	// Total count of the given alert type
	Count int32 `json:"count"`
}

// NewInlineResponse200218Items instantiates a new InlineResponse200218Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200218Items(type_ string, count int32) *InlineResponse200218Items {
	this := InlineResponse200218Items{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200218ItemsWithDefaults instantiates a new InlineResponse200218Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200218ItemsWithDefaults() *InlineResponse200218Items {
	this := InlineResponse200218Items{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200218Items) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Items) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200218Items) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200218Items) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200218Items) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200218Items) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200218Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200218Items struct {
	value *InlineResponse200218Items
	isSet bool
}

func (v NullableInlineResponse200218Items) Get() *InlineResponse200218Items {
	return v.value
}

func (v *NullableInlineResponse200218Items) Set(val *InlineResponse200218Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200218Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200218Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200218Items(val *InlineResponse200218Items) *NullableInlineResponse200218Items {
	return &NullableInlineResponse200218Items{value: val, isSet: true}
}

func (v NullableInlineResponse200218Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200218Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


