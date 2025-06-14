/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20076Subnets struct for InlineResponse20076Subnets
type InlineResponse20076Subnets struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet *string `json:"localSubnet,omitempty"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
	Nat *InlineResponse20076Nat `json:"nat,omitempty"`
}

// NewInlineResponse20076Subnets instantiates a new InlineResponse20076Subnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20076Subnets() *InlineResponse20076Subnets {
	this := InlineResponse20076Subnets{}
	return &this
}

// NewInlineResponse20076SubnetsWithDefaults instantiates a new InlineResponse20076Subnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20076SubnetsWithDefaults() *InlineResponse20076Subnets {
	this := InlineResponse20076Subnets{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value if set, zero value otherwise.
func (o *InlineResponse20076Subnets) GetLocalSubnet() string {
	if o == nil || isNil(o.LocalSubnet) {
		var ret string
		return ret
	}
	return *o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076Subnets) GetLocalSubnetOk() (*string, bool) {
	if o == nil || isNil(o.LocalSubnet) {
    return nil, false
	}
	return o.LocalSubnet, true
}

// HasLocalSubnet returns a boolean if a field has been set.
func (o *InlineResponse20076Subnets) HasLocalSubnet() bool {
	if o != nil && !isNil(o.LocalSubnet) {
		return true
	}

	return false
}

// SetLocalSubnet gets a reference to the given string and assigns it to the LocalSubnet field.
func (o *InlineResponse20076Subnets) SetLocalSubnet(v string) {
	o.LocalSubnet = &v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *InlineResponse20076Subnets) GetUseVpn() bool {
	if o == nil || isNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076Subnets) GetUseVpnOk() (*bool, bool) {
	if o == nil || isNil(o.UseVpn) {
    return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *InlineResponse20076Subnets) HasUseVpn() bool {
	if o != nil && !isNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *InlineResponse20076Subnets) SetUseVpn(v bool) {
	o.UseVpn = &v
}

// GetNat returns the Nat field value if set, zero value otherwise.
func (o *InlineResponse20076Subnets) GetNat() InlineResponse20076Nat {
	if o == nil || isNil(o.Nat) {
		var ret InlineResponse20076Nat
		return ret
	}
	return *o.Nat
}

// GetNatOk returns a tuple with the Nat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20076Subnets) GetNatOk() (*InlineResponse20076Nat, bool) {
	if o == nil || isNil(o.Nat) {
    return nil, false
	}
	return o.Nat, true
}

// HasNat returns a boolean if a field has been set.
func (o *InlineResponse20076Subnets) HasNat() bool {
	if o != nil && !isNil(o.Nat) {
		return true
	}

	return false
}

// SetNat gets a reference to the given InlineResponse20076Nat and assigns it to the Nat field.
func (o *InlineResponse20076Subnets) SetNat(v InlineResponse20076Nat) {
	o.Nat = &v
}

func (o InlineResponse20076Subnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalSubnet) {
		toSerialize["localSubnet"] = o.LocalSubnet
	}
	if !isNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	if !isNil(o.Nat) {
		toSerialize["nat"] = o.Nat
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20076Subnets struct {
	value *InlineResponse20076Subnets
	isSet bool
}

func (v NullableInlineResponse20076Subnets) Get() *InlineResponse20076Subnets {
	return v.value
}

func (v *NullableInlineResponse20076Subnets) Set(val *InlineResponse20076Subnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20076Subnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20076Subnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20076Subnets(val *InlineResponse20076Subnets) *NullableInlineResponse20076Subnets {
	return &NullableInlineResponse20076Subnets{value: val, isSet: true}
}

func (v NullableInlineResponse20076Subnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20076Subnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


