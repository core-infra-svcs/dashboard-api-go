/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200290 struct for InlineResponse200290
type InlineResponse200290 struct {
	ResultingNetwork *InlineResponse200290ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse200290 instantiates a new InlineResponse200290 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200290() *InlineResponse200290 {
	this := InlineResponse200290{}
	return &this
}

// NewInlineResponse200290WithDefaults instantiates a new InlineResponse200290 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200290WithDefaults() *InlineResponse200290 {
	this := InlineResponse200290{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse200290) GetResultingNetwork() InlineResponse200290ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse200290ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290) GetResultingNetworkOk() (*InlineResponse200290ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse200290) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse200290ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse200290) SetResultingNetwork(v InlineResponse200290ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse200290) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200290 struct {
	value *InlineResponse200290
	isSet bool
}

func (v NullableInlineResponse200290) Get() *InlineResponse200290 {
	return v.value
}

func (v *NullableInlineResponse200290) Set(val *InlineResponse200290) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200290) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200290) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200290(val *InlineResponse200290) *NullableInlineResponse200290 {
	return &NullableInlineResponse200290{value: val, isSet: true}
}

func (v NullableInlineResponse200290) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200290) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


