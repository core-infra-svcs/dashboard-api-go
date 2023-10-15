/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20025TrafficFilters1 struct for InlineResponse20025TrafficFilters1
type InlineResponse20025TrafficFilters1 struct {
	// Traffic filter type. Must be one of: 'applicationCategory', 'application' or 'custom'
	Type string `json:"type"`
	Value InlineResponse20025Value1 `json:"value"`
}

// NewInlineResponse20025TrafficFilters1 instantiates a new InlineResponse20025TrafficFilters1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025TrafficFilters1(type_ string, value InlineResponse20025Value1) *InlineResponse20025TrafficFilters1 {
	this := InlineResponse20025TrafficFilters1{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse20025TrafficFilters1WithDefaults instantiates a new InlineResponse20025TrafficFilters1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025TrafficFilters1WithDefaults() *InlineResponse20025TrafficFilters1 {
	this := InlineResponse20025TrafficFilters1{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20025TrafficFilters1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20025TrafficFilters1) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20025TrafficFilters1) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse20025TrafficFilters1) GetValue() InlineResponse20025Value1 {
	if o == nil {
		var ret InlineResponse20025Value1
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20025TrafficFilters1) GetValueOk() (*InlineResponse20025Value1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse20025TrafficFilters1) SetValue(v InlineResponse20025Value1) {
	o.Value = v
}

func (o InlineResponse20025TrafficFilters1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025TrafficFilters1 struct {
	value *InlineResponse20025TrafficFilters1
	isSet bool
}

func (v NullableInlineResponse20025TrafficFilters1) Get() *InlineResponse20025TrafficFilters1 {
	return v.value
}

func (v *NullableInlineResponse20025TrafficFilters1) Set(val *InlineResponse20025TrafficFilters1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025TrafficFilters1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025TrafficFilters1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025TrafficFilters1(val *InlineResponse20025TrafficFilters1) *NullableInlineResponse20025TrafficFilters1 {
	return &NullableInlineResponse20025TrafficFilters1{value: val, isSet: true}
}

func (v NullableInlineResponse20025TrafficFilters1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025TrafficFilters1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


