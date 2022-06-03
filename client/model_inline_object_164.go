/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject164 struct for InlineObject164
type InlineObject164 struct {
	// The list of VPN peers
	Peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers `json:"peers"`
}

// NewInlineObject164 instantiates a new InlineObject164 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject164(peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) *InlineObject164 {
	this := InlineObject164{}
	this.Peers = peers
	return &this
}

// NewInlineObject164WithDefaults instantiates a new InlineObject164 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject164WithDefaults() *InlineObject164 {
	this := InlineObject164{}
	return &this
}

// GetPeers returns the Peers field value
func (o *InlineObject164) GetPeers() []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *InlineObject164) GetPeersOk() (*[]OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Peers, true
}

// SetPeers sets field value
func (o *InlineObject164) SetPeers(v []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) {
	o.Peers = v
}

func (o InlineObject164) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject164 struct {
	value *InlineObject164
	isSet bool
}

func (v NullableInlineObject164) Get() *InlineObject164 {
	return v.value
}

func (v *NullableInlineObject164) Set(val *InlineObject164) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject164) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject164) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject164(val *InlineObject164) *NullableInlineObject164 {
	return &NullableInlineObject164{value: val, isSet: true}
}

func (v NullableInlineObject164) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject164) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


