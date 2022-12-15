/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20098 struct for InlineResponse20098
type InlineResponse20098 struct {
	ResultingNetwork *InlineResponse20098ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse20098 instantiates a new InlineResponse20098 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// NewInlineResponse20098WithDefaults instantiates a new InlineResponse20098 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098WithDefaults() *InlineResponse20098 {
	this := InlineResponse20098{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse20098) GetResultingNetwork() InlineResponse20098ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse20098ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098) GetResultingNetworkOk() (*InlineResponse20098ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse20098) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse20098ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse20098) SetResultingNetwork(v InlineResponse20098ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse20098) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098 struct {
	value *InlineResponse20098
	isSet bool
}

func (v NullableInlineResponse20098) Get() *InlineResponse20098 {
	return v.value
}

func (v *NullableInlineResponse20098) Set(val *InlineResponse20098) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098(val *InlineResponse20098) *NullableInlineResponse20098 {
	return &NullableInlineResponse20098{value: val, isSet: true}
}

func (v NullableInlineResponse20098) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


