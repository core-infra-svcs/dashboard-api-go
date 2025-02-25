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

// InlineResponse200147 struct for InlineResponse200147
type InlineResponse200147 struct {
	// Networks after the split
	ResultingNetworks []InlineResponse20047 `json:"resultingNetworks,omitempty"`
}

// NewInlineResponse200147 instantiates a new InlineResponse200147 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200147() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// NewInlineResponse200147WithDefaults instantiates a new InlineResponse200147 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200147WithDefaults() *InlineResponse200147 {
	this := InlineResponse200147{}
	return &this
}

// GetResultingNetworks returns the ResultingNetworks field value if set, zero value otherwise.
func (o *InlineResponse200147) GetResultingNetworks() []InlineResponse20047 {
	if o == nil || isNil(o.ResultingNetworks) {
		var ret []InlineResponse20047
		return ret
	}
	return o.ResultingNetworks
}

// GetResultingNetworksOk returns a tuple with the ResultingNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200147) GetResultingNetworksOk() ([]InlineResponse20047, bool) {
	if o == nil || isNil(o.ResultingNetworks) {
    return nil, false
	}
	return o.ResultingNetworks, true
}

// HasResultingNetworks returns a boolean if a field has been set.
func (o *InlineResponse200147) HasResultingNetworks() bool {
	if o != nil && !isNil(o.ResultingNetworks) {
		return true
	}

	return false
}

// SetResultingNetworks gets a reference to the given []InlineResponse20047 and assigns it to the ResultingNetworks field.
func (o *InlineResponse200147) SetResultingNetworks(v []InlineResponse20047) {
	o.ResultingNetworks = v
}

func (o InlineResponse200147) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultingNetworks) {
		toSerialize["resultingNetworks"] = o.ResultingNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200147 struct {
	value *InlineResponse200147
	isSet bool
}

func (v NullableInlineResponse200147) Get() *InlineResponse200147 {
	return v.value
}

func (v *NullableInlineResponse200147) Set(val *InlineResponse200147) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200147) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200147) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200147(val *InlineResponse200147) *NullableInlineResponse200147 {
	return &NullableInlineResponse200147{value: val, isSet: true}
}

func (v NullableInlineResponse200147) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200147) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


