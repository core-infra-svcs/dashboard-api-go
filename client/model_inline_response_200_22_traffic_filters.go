/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20022TrafficFilters struct for InlineResponse20022TrafficFilters
type InlineResponse20022TrafficFilters struct {
	// Traffic filter type. Must be \"custom\"
	Type string `json:"type"`
	Value InlineResponse20022Value `json:"value"`
}

// NewInlineResponse20022TrafficFilters instantiates a new InlineResponse20022TrafficFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022TrafficFilters(type_ string, value InlineResponse20022Value) *InlineResponse20022TrafficFilters {
	this := InlineResponse20022TrafficFilters{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse20022TrafficFiltersWithDefaults instantiates a new InlineResponse20022TrafficFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022TrafficFiltersWithDefaults() *InlineResponse20022TrafficFilters {
	this := InlineResponse20022TrafficFilters{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20022TrafficFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20022TrafficFilters) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20022TrafficFilters) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse20022TrafficFilters) GetValue() InlineResponse20022Value {
	if o == nil {
		var ret InlineResponse20022Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20022TrafficFilters) GetValueOk() (*InlineResponse20022Value, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse20022TrafficFilters) SetValue(v InlineResponse20022Value) {
	o.Value = v
}

func (o InlineResponse20022TrafficFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022TrafficFilters struct {
	value *InlineResponse20022TrafficFilters
	isSet bool
}

func (v NullableInlineResponse20022TrafficFilters) Get() *InlineResponse20022TrafficFilters {
	return v.value
}

func (v *NullableInlineResponse20022TrafficFilters) Set(val *InlineResponse20022TrafficFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022TrafficFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022TrafficFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022TrafficFilters(val *InlineResponse20022TrafficFilters) *NullableInlineResponse20022TrafficFilters {
	return &NullableInlineResponse20022TrafficFilters{value: val, isSet: true}
}

func (v NullableInlineResponse20022TrafficFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022TrafficFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


