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

// InlineResponse200217Items struct for InlineResponse200217Items
type InlineResponse200217Items struct {
	// Alert Type
	Type string `json:"type"`
	// Total count of the given alert type
	Count int32 `json:"count"`
}

// NewInlineResponse200217Items instantiates a new InlineResponse200217Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200217Items(type_ string, count int32) *InlineResponse200217Items {
	this := InlineResponse200217Items{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200217ItemsWithDefaults instantiates a new InlineResponse200217Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200217ItemsWithDefaults() *InlineResponse200217Items {
	this := InlineResponse200217Items{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200217Items) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200217Items) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200217Items) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200217Items) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200217Items) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200217Items) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200217Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200217Items struct {
	value *InlineResponse200217Items
	isSet bool
}

func (v NullableInlineResponse200217Items) Get() *InlineResponse200217Items {
	return v.value
}

func (v *NullableInlineResponse200217Items) Set(val *InlineResponse200217Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200217Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200217Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200217Items(val *InlineResponse200217Items) *NullableInlineResponse200217Items {
	return &NullableInlineResponse200217Items{value: val, isSet: true}
}

func (v NullableInlineResponse200217Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200217Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

