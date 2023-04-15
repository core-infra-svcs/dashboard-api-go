/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject182 struct for InlineObject182
type InlineObject182 struct {
	// The list of VPN peers
	Peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers `json:"peers"`
}

// NewInlineObject182 instantiates a new InlineObject182 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject182(peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) *InlineObject182 {
	this := InlineObject182{}
	this.Peers = peers
	return &this
}

// NewInlineObject182WithDefaults instantiates a new InlineObject182 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject182WithDefaults() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// GetPeers returns the Peers field value
func (o *InlineObject182) GetPeers() []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetPeersOk() ([]OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *InlineObject182) SetPeers(v []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) {
	o.Peers = v
}

func (o InlineObject182) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject182 struct {
	value *InlineObject182
	isSet bool
}

func (v NullableInlineObject182) Get() *InlineObject182 {
	return v.value
}

func (v *NullableInlineObject182) Set(val *InlineObject182) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject182) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject182) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject182(val *InlineObject182) *NullableInlineObject182 {
	return &NullableInlineObject182{value: val, isSet: true}
}

func (v NullableInlineObject182) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject182) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


