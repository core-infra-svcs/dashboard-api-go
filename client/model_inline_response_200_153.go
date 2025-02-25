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

// InlineResponse200153 struct for InlineResponse200153
type InlineResponse200153 struct {
	// ID of the trusted server.
	TrustedServerId *string `json:"trustedServerId,omitempty"`
	// Mac address of the trusted server.
	Mac *string `json:"mac,omitempty"`
	// Vlan ID of the trusted server.
	Vlan *int32 `json:"vlan,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 `json:"ipv4,omitempty"`
}

// NewInlineResponse200153 instantiates a new InlineResponse200153 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200153() *InlineResponse200153 {
	this := InlineResponse200153{}
	return &this
}

// NewInlineResponse200153WithDefaults instantiates a new InlineResponse200153 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200153WithDefaults() *InlineResponse200153 {
	this := InlineResponse200153{}
	return &this
}

// GetTrustedServerId returns the TrustedServerId field value if set, zero value otherwise.
func (o *InlineResponse200153) GetTrustedServerId() string {
	if o == nil || isNil(o.TrustedServerId) {
		var ret string
		return ret
	}
	return *o.TrustedServerId
}

// GetTrustedServerIdOk returns a tuple with the TrustedServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200153) GetTrustedServerIdOk() (*string, bool) {
	if o == nil || isNil(o.TrustedServerId) {
    return nil, false
	}
	return o.TrustedServerId, true
}

// HasTrustedServerId returns a boolean if a field has been set.
func (o *InlineResponse200153) HasTrustedServerId() bool {
	if o != nil && !isNil(o.TrustedServerId) {
		return true
	}

	return false
}

// SetTrustedServerId gets a reference to the given string and assigns it to the TrustedServerId field.
func (o *InlineResponse200153) SetTrustedServerId(v string) {
	o.TrustedServerId = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200153) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200153) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200153) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200153) SetMac(v string) {
	o.Mac = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineResponse200153) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200153) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineResponse200153) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineResponse200153) SetVlan(v int32) {
	o.Vlan = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *InlineResponse200153) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200153) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *InlineResponse200153) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4 and assigns it to the Ipv4 field.
func (o *InlineResponse200153) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4) {
	o.Ipv4 = &v
}

func (o InlineResponse200153) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrustedServerId) {
		toSerialize["trustedServerId"] = o.TrustedServerId
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200153 struct {
	value *InlineResponse200153
	isSet bool
}

func (v NullableInlineResponse200153) Get() *InlineResponse200153 {
	return v.value
}

func (v *NullableInlineResponse200153) Set(val *InlineResponse200153) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200153) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200153) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200153(val *InlineResponse200153) *NullableInlineResponse200153 {
	return &NullableInlineResponse200153{value: val, isSet: true}
}

func (v NullableInlineResponse200153) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200153) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


