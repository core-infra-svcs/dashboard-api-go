/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200233 struct for InlineResponse200233
type InlineResponse200233 struct {
	// The list of VPN peers
	Peers []InlineResponse200233Peers `json:"peers,omitempty"`
}

// NewInlineResponse200233 instantiates a new InlineResponse200233 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200233() *InlineResponse200233 {
	this := InlineResponse200233{}
	return &this
}

// NewInlineResponse200233WithDefaults instantiates a new InlineResponse200233 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200233WithDefaults() *InlineResponse200233 {
	this := InlineResponse200233{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse200233) GetPeers() []InlineResponse200233Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse200233Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233) GetPeersOk() ([]InlineResponse200233Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse200233) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse200233Peers and assigns it to the Peers field.
func (o *InlineResponse200233) SetPeers(v []InlineResponse200233Peers) {
	o.Peers = v
}

func (o InlineResponse200233) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200233 struct {
	value *InlineResponse200233
	isSet bool
}

func (v NullableInlineResponse200233) Get() *InlineResponse200233 {
	return v.value
}

func (v *NullableInlineResponse200233) Set(val *InlineResponse200233) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200233) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200233) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200233(val *InlineResponse200233) *NullableInlineResponse200233 {
	return &NullableInlineResponse200233{value: val, isSet: true}
}

func (v NullableInlineResponse200233) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200233) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


