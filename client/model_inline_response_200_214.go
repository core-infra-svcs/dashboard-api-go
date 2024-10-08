/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200214 struct for InlineResponse200214
type InlineResponse200214 struct {
	// The list of VPN peers
	Peers []InlineResponse200214Peers `json:"peers,omitempty"`
}

// NewInlineResponse200214 instantiates a new InlineResponse200214 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200214() *InlineResponse200214 {
	this := InlineResponse200214{}
	return &this
}

// NewInlineResponse200214WithDefaults instantiates a new InlineResponse200214 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200214WithDefaults() *InlineResponse200214 {
	this := InlineResponse200214{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse200214) GetPeers() []InlineResponse200214Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse200214Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200214) GetPeersOk() ([]InlineResponse200214Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse200214) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse200214Peers and assigns it to the Peers field.
func (o *InlineResponse200214) SetPeers(v []InlineResponse200214Peers) {
	o.Peers = v
}

func (o InlineResponse200214) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200214 struct {
	value *InlineResponse200214
	isSet bool
}

func (v NullableInlineResponse200214) Get() *InlineResponse200214 {
	return v.value
}

func (v *NullableInlineResponse200214) Set(val *InlineResponse200214) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200214) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200214) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200214(val *InlineResponse200214) *NullableInlineResponse200214 {
	return &NullableInlineResponse200214{value: val, isSet: true}
}

func (v NullableInlineResponse200214) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200214) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


