/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20093 struct for InlineResponse20093
type InlineResponse20093 struct {
	ResultingNetwork *InlineResponse20093ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse20093 instantiates a new InlineResponse20093 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20093() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// NewInlineResponse20093WithDefaults instantiates a new InlineResponse20093 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20093WithDefaults() *InlineResponse20093 {
	this := InlineResponse20093{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse20093) GetResultingNetwork() InlineResponse20093ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse20093ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20093) GetResultingNetworkOk() (*InlineResponse20093ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse20093) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse20093ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse20093) SetResultingNetwork(v InlineResponse20093ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse20093) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20093 struct {
	value *InlineResponse20093
	isSet bool
}

func (v NullableInlineResponse20093) Get() *InlineResponse20093 {
	return v.value
}

func (v *NullableInlineResponse20093) Set(val *InlineResponse20093) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20093) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20093) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20093(val *InlineResponse20093) *NullableInlineResponse20093 {
	return &NullableInlineResponse20093{value: val, isSet: true}
}

func (v NullableInlineResponse20093) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20093) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

