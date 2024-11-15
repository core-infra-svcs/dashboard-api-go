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

// InlineResponse200219 struct for InlineResponse200219
type InlineResponse200219 struct {
	// The list of VPN peers
	Peers []InlineResponse200219Peers `json:"peers,omitempty"`
}

// NewInlineResponse200219 instantiates a new InlineResponse200219 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200219() *InlineResponse200219 {
	this := InlineResponse200219{}
	return &this
}

// NewInlineResponse200219WithDefaults instantiates a new InlineResponse200219 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200219WithDefaults() *InlineResponse200219 {
	this := InlineResponse200219{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse200219) GetPeers() []InlineResponse200219Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse200219Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200219) GetPeersOk() ([]InlineResponse200219Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse200219) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse200219Peers and assigns it to the Peers field.
func (o *InlineResponse200219) SetPeers(v []InlineResponse200219Peers) {
	o.Peers = v
}

func (o InlineResponse200219) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200219 struct {
	value *InlineResponse200219
	isSet bool
}

func (v NullableInlineResponse200219) Get() *InlineResponse200219 {
	return v.value
}

func (v *NullableInlineResponse200219) Set(val *InlineResponse200219) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200219) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200219) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200219(val *InlineResponse200219) *NullableInlineResponse200219 {
	return &NullableInlineResponse200219{value: val, isSet: true}
}

func (v NullableInlineResponse200219) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200219) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


