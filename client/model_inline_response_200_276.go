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

// InlineResponse200276 struct for InlineResponse200276
type InlineResponse200276 struct {
	ResultingNetwork *InlineResponse200276ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse200276 instantiates a new InlineResponse200276 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200276() *InlineResponse200276 {
	this := InlineResponse200276{}
	return &this
}

// NewInlineResponse200276WithDefaults instantiates a new InlineResponse200276 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200276WithDefaults() *InlineResponse200276 {
	this := InlineResponse200276{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse200276) GetResultingNetwork() InlineResponse200276ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse200276ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200276) GetResultingNetworkOk() (*InlineResponse200276ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse200276) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse200276ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse200276) SetResultingNetwork(v InlineResponse200276ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse200276) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200276 struct {
	value *InlineResponse200276
	isSet bool
}

func (v NullableInlineResponse200276) Get() *InlineResponse200276 {
	return v.value
}

func (v *NullableInlineResponse200276) Set(val *InlineResponse200276) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200276) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200276) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200276(val *InlineResponse200276) *NullableInlineResponse200276 {
	return &NullableInlineResponse200276{value: val, isSet: true}
}

func (v NullableInlineResponse200276) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200276) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


