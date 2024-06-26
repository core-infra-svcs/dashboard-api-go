/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200261 struct for InlineResponse200261
type InlineResponse200261 struct {
	ResultingNetwork *InlineResponse200261ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse200261 instantiates a new InlineResponse200261 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200261() *InlineResponse200261 {
	this := InlineResponse200261{}
	return &this
}

// NewInlineResponse200261WithDefaults instantiates a new InlineResponse200261 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200261WithDefaults() *InlineResponse200261 {
	this := InlineResponse200261{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse200261) GetResultingNetwork() InlineResponse200261ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse200261ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200261) GetResultingNetworkOk() (*InlineResponse200261ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse200261) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse200261ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse200261) SetResultingNetwork(v InlineResponse200261ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse200261) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200261 struct {
	value *InlineResponse200261
	isSet bool
}

func (v NullableInlineResponse200261) Get() *InlineResponse200261 {
	return v.value
}

func (v *NullableInlineResponse200261) Set(val *InlineResponse200261) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200261) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200261) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200261(val *InlineResponse200261) *NullableInlineResponse200261 {
	return &NullableInlineResponse200261{value: val, isSet: true}
}

func (v NullableInlineResponse200261) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200261) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


