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

// InlineResponse200262 struct for InlineResponse200262
type InlineResponse200262 struct {
	ResultingNetwork *InlineResponse200262ResultingNetwork `json:"resultingNetwork,omitempty"`
}

// NewInlineResponse200262 instantiates a new InlineResponse200262 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200262() *InlineResponse200262 {
	this := InlineResponse200262{}
	return &this
}

// NewInlineResponse200262WithDefaults instantiates a new InlineResponse200262 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200262WithDefaults() *InlineResponse200262 {
	this := InlineResponse200262{}
	return &this
}

// GetResultingNetwork returns the ResultingNetwork field value if set, zero value otherwise.
func (o *InlineResponse200262) GetResultingNetwork() InlineResponse200262ResultingNetwork {
	if o == nil || isNil(o.ResultingNetwork) {
		var ret InlineResponse200262ResultingNetwork
		return ret
	}
	return *o.ResultingNetwork
}

// GetResultingNetworkOk returns a tuple with the ResultingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200262) GetResultingNetworkOk() (*InlineResponse200262ResultingNetwork, bool) {
	if o == nil || isNil(o.ResultingNetwork) {
    return nil, false
	}
	return o.ResultingNetwork, true
}

// HasResultingNetwork returns a boolean if a field has been set.
func (o *InlineResponse200262) HasResultingNetwork() bool {
	if o != nil && !isNil(o.ResultingNetwork) {
		return true
	}

	return false
}

// SetResultingNetwork gets a reference to the given InlineResponse200262ResultingNetwork and assigns it to the ResultingNetwork field.
func (o *InlineResponse200262) SetResultingNetwork(v InlineResponse200262ResultingNetwork) {
	o.ResultingNetwork = &v
}

func (o InlineResponse200262) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetwork) {
		toSerialize["resultingNetwork"] = o.ResultingNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200262 struct {
	value *InlineResponse200262
	isSet bool
}

func (v NullableInlineResponse200262) Get() *InlineResponse200262 {
	return v.value
}

func (v *NullableInlineResponse200262) Set(val *InlineResponse200262) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200262) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200262) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200262(val *InlineResponse200262) *NullableInlineResponse200262 {
	return &NullableInlineResponse200262{value: val, isSet: true}
}

func (v NullableInlineResponse200262) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200262) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


