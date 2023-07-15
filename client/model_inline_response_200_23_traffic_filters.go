/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023TrafficFilters struct for InlineResponse20023TrafficFilters
type InlineResponse20023TrafficFilters struct {
	// Traffic filter type. Must be \"custom\"
	Type string `json:"type"`
	Value InlineResponse20023Value `json:"value"`
}

// NewInlineResponse20023TrafficFilters instantiates a new InlineResponse20023TrafficFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023TrafficFilters(type_ string, value InlineResponse20023Value) *InlineResponse20023TrafficFilters {
	this := InlineResponse20023TrafficFilters{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse20023TrafficFiltersWithDefaults instantiates a new InlineResponse20023TrafficFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023TrafficFiltersWithDefaults() *InlineResponse20023TrafficFilters {
	this := InlineResponse20023TrafficFilters{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20023TrafficFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20023TrafficFilters) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20023TrafficFilters) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse20023TrafficFilters) GetValue() InlineResponse20023Value {
	if o == nil {
		var ret InlineResponse20023Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20023TrafficFilters) GetValueOk() (*InlineResponse20023Value, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse20023TrafficFilters) SetValue(v InlineResponse20023Value) {
	o.Value = v
}

func (o InlineResponse20023TrafficFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023TrafficFilters struct {
	value *InlineResponse20023TrafficFilters
	isSet bool
}

func (v NullableInlineResponse20023TrafficFilters) Get() *InlineResponse20023TrafficFilters {
	return v.value
}

func (v *NullableInlineResponse20023TrafficFilters) Set(val *InlineResponse20023TrafficFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023TrafficFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023TrafficFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023TrafficFilters(val *InlineResponse20023TrafficFilters) *NullableInlineResponse20023TrafficFilters {
	return &NullableInlineResponse20023TrafficFilters{value: val, isSet: true}
}

func (v NullableInlineResponse20023TrafficFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023TrafficFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


