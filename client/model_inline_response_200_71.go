/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20071 struct for InlineResponse20071
type InlineResponse20071 struct {
	// The list of VPN peers
	Peers []InlineResponse20071Peers `json:"peers,omitempty"`
}

// NewInlineResponse20071 instantiates a new InlineResponse20071 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20071() *InlineResponse20071 {
	this := InlineResponse20071{}
	return &this
}

// NewInlineResponse20071WithDefaults instantiates a new InlineResponse20071 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20071WithDefaults() *InlineResponse20071 {
	this := InlineResponse20071{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse20071) GetPeers() []InlineResponse20071Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse20071Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20071) GetPeersOk() ([]InlineResponse20071Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse20071) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse20071Peers and assigns it to the Peers field.
func (o *InlineResponse20071) SetPeers(v []InlineResponse20071Peers) {
	o.Peers = v
}

func (o InlineResponse20071) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20071 struct {
	value *InlineResponse20071
	isSet bool
}

func (v NullableInlineResponse20071) Get() *InlineResponse20071 {
	return v.value
}

func (v *NullableInlineResponse20071) Set(val *InlineResponse20071) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20071) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20071) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20071(val *InlineResponse20071) *NullableInlineResponse20071 {
	return &NullableInlineResponse20071{value: val, isSet: true}
}

func (v NullableInlineResponse20071) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20071) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


