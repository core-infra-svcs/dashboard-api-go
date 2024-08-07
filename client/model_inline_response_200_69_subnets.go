/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069Subnets struct for InlineResponse20069Subnets
type InlineResponse20069Subnets struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet *string `json:"localSubnet,omitempty"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
}

// NewInlineResponse20069Subnets instantiates a new InlineResponse20069Subnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069Subnets() *InlineResponse20069Subnets {
	this := InlineResponse20069Subnets{}
	return &this
}

// NewInlineResponse20069SubnetsWithDefaults instantiates a new InlineResponse20069Subnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069SubnetsWithDefaults() *InlineResponse20069Subnets {
	this := InlineResponse20069Subnets{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value if set, zero value otherwise.
func (o *InlineResponse20069Subnets) GetLocalSubnet() string {
	if o == nil || isNil(o.LocalSubnet) {
		var ret string
		return ret
	}
	return *o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Subnets) GetLocalSubnetOk() (*string, bool) {
	if o == nil || isNil(o.LocalSubnet) {
    return nil, false
	}
	return o.LocalSubnet, true
}

// HasLocalSubnet returns a boolean if a field has been set.
func (o *InlineResponse20069Subnets) HasLocalSubnet() bool {
	if o != nil && !isNil(o.LocalSubnet) {
		return true
	}

	return false
}

// SetLocalSubnet gets a reference to the given string and assigns it to the LocalSubnet field.
func (o *InlineResponse20069Subnets) SetLocalSubnet(v string) {
	o.LocalSubnet = &v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *InlineResponse20069Subnets) GetUseVpn() bool {
	if o == nil || isNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Subnets) GetUseVpnOk() (*bool, bool) {
	if o == nil || isNil(o.UseVpn) {
    return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *InlineResponse20069Subnets) HasUseVpn() bool {
	if o != nil && !isNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *InlineResponse20069Subnets) SetUseVpn(v bool) {
	o.UseVpn = &v
}

func (o InlineResponse20069Subnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalSubnet) {
		toSerialize["localSubnet"] = o.LocalSubnet
	}
	if !isNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069Subnets struct {
	value *InlineResponse20069Subnets
	isSet bool
}

func (v NullableInlineResponse20069Subnets) Get() *InlineResponse20069Subnets {
	return v.value
}

func (v *NullableInlineResponse20069Subnets) Set(val *InlineResponse20069Subnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069Subnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069Subnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069Subnets(val *InlineResponse20069Subnets) *NullableInlineResponse20069Subnets {
	return &NullableInlineResponse20069Subnets{value: val, isSet: true}
}

func (v NullableInlineResponse20069Subnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069Subnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


