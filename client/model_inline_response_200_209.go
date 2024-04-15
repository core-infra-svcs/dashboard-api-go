/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200209 struct for InlineResponse200209
type InlineResponse200209 struct {
	// The list of VPN peers
	Peers []InlineResponse200209Peers `json:"peers,omitempty"`
}

// NewInlineResponse200209 instantiates a new InlineResponse200209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200209() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200209WithDefaults() *InlineResponse200209 {
	this := InlineResponse200209{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse200209) GetPeers() []InlineResponse200209Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse200209Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200209) GetPeersOk() ([]InlineResponse200209Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse200209) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse200209Peers and assigns it to the Peers field.
func (o *InlineResponse200209) SetPeers(v []InlineResponse200209Peers) {
	o.Peers = v
}

func (o InlineResponse200209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200209 struct {
	value *InlineResponse200209
	isSet bool
}

func (v NullableInlineResponse200209) Get() *InlineResponse200209 {
	return v.value
}

func (v *NullableInlineResponse200209) Set(val *InlineResponse200209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200209(val *InlineResponse200209) *NullableInlineResponse200209 {
	return &NullableInlineResponse200209{value: val, isSet: true}
}

func (v NullableInlineResponse200209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


