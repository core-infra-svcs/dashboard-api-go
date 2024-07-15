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

// InlineResponse200138 struct for InlineResponse200138
type InlineResponse200138 struct {
	// Networks after the split
	ResultingNetworks []InlineResponse20043 `json:"resultingNetworks,omitempty"`
}

// NewInlineResponse200138 instantiates a new InlineResponse200138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200138() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// NewInlineResponse200138WithDefaults instantiates a new InlineResponse200138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200138WithDefaults() *InlineResponse200138 {
	this := InlineResponse200138{}
	return &this
}

// GetResultingNetworks returns the ResultingNetworks field value if set, zero value otherwise.
func (o *InlineResponse200138) GetResultingNetworks() []InlineResponse20043 {
	if o == nil || isNil(o.ResultingNetworks) {
		var ret []InlineResponse20043
		return ret
	}
	return o.ResultingNetworks
}

// GetResultingNetworksOk returns a tuple with the ResultingNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200138) GetResultingNetworksOk() ([]InlineResponse20043, bool) {
	if o == nil || isNil(o.ResultingNetworks) {
    return nil, false
	}
	return o.ResultingNetworks, true
}

// HasResultingNetworks returns a boolean if a field has been set.
func (o *InlineResponse200138) HasResultingNetworks() bool {
	if o != nil && !isNil(o.ResultingNetworks) {
		return true
	}

	return false
}

// SetResultingNetworks gets a reference to the given []InlineResponse20043 and assigns it to the ResultingNetworks field.
func (o *InlineResponse200138) SetResultingNetworks(v []InlineResponse20043) {
	o.ResultingNetworks = v
}

func (o InlineResponse200138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetworks) {
		toSerialize["resultingNetworks"] = o.ResultingNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200138 struct {
	value *InlineResponse200138
	isSet bool
}

func (v NullableInlineResponse200138) Get() *InlineResponse200138 {
	return v.value
}

func (v *NullableInlineResponse200138) Set(val *InlineResponse200138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200138(val *InlineResponse200138) *NullableInlineResponse200138 {
	return &NullableInlineResponse200138{value: val, isSet: true}
}

func (v NullableInlineResponse200138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


