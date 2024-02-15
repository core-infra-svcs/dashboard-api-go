/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20081 struct for InlineResponse20081
type InlineResponse20081 struct {
	// Networks after the split
	ResultingNetworks []InlineResponse20021 `json:"resultingNetworks,omitempty"`
}

// NewInlineResponse20081 instantiates a new InlineResponse20081 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20081() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// NewInlineResponse20081WithDefaults instantiates a new InlineResponse20081 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20081WithDefaults() *InlineResponse20081 {
	this := InlineResponse20081{}
	return &this
}

// GetResultingNetworks returns the ResultingNetworks field value if set, zero value otherwise.
func (o *InlineResponse20081) GetResultingNetworks() []InlineResponse20021 {
	if o == nil || isNil(o.ResultingNetworks) {
		var ret []InlineResponse20021
		return ret
	}
	return o.ResultingNetworks
}

// GetResultingNetworksOk returns a tuple with the ResultingNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20081) GetResultingNetworksOk() ([]InlineResponse20021, bool) {
	if o == nil || isNil(o.ResultingNetworks) {
    return nil, false
	}
	return o.ResultingNetworks, true
}

// HasResultingNetworks returns a boolean if a field has been set.
func (o *InlineResponse20081) HasResultingNetworks() bool {
	if o != nil && !isNil(o.ResultingNetworks) {
		return true
	}

	return false
}

// SetResultingNetworks gets a reference to the given []InlineResponse20021 and assigns it to the ResultingNetworks field.
func (o *InlineResponse20081) SetResultingNetworks(v []InlineResponse20021) {
	o.ResultingNetworks = v
}

func (o InlineResponse20081) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetworks) {
		toSerialize["resultingNetworks"] = o.ResultingNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20081 struct {
	value *InlineResponse20081
	isSet bool
}

func (v NullableInlineResponse20081) Get() *InlineResponse20081 {
	return v.value
}

func (v *NullableInlineResponse20081) Set(val *InlineResponse20081) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20081) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20081) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20081(val *InlineResponse20081) *NullableInlineResponse20081 {
	return &NullableInlineResponse20081{value: val, isSet: true}
}

func (v NullableInlineResponse20081) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20081) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


