/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVpnBgpAuthentication Authentication settings between BGP peers.
type NetworksNetworkIdApplianceVpnBgpAuthentication struct {
	// Password to configure MD5 authentication between BGP peers.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdApplianceVpnBgpAuthentication instantiates a new NetworksNetworkIdApplianceVpnBgpAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVpnBgpAuthentication() *NetworksNetworkIdApplianceVpnBgpAuthentication {
	this := NetworksNetworkIdApplianceVpnBgpAuthentication{}
	return &this
}

// NewNetworksNetworkIdApplianceVpnBgpAuthenticationWithDefaults instantiates a new NetworksNetworkIdApplianceVpnBgpAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVpnBgpAuthenticationWithDefaults() *NetworksNetworkIdApplianceVpnBgpAuthentication {
	this := NetworksNetworkIdApplianceVpnBgpAuthentication{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVpnBgpAuthentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVpnBgpAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVpnBgpAuthentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdApplianceVpnBgpAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdApplianceVpnBgpAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVpnBgpAuthentication struct {
	value *NetworksNetworkIdApplianceVpnBgpAuthentication
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVpnBgpAuthentication) Get() *NetworksNetworkIdApplianceVpnBgpAuthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpAuthentication) Set(val *NetworksNetworkIdApplianceVpnBgpAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVpnBgpAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVpnBgpAuthentication(val *NetworksNetworkIdApplianceVpnBgpAuthentication) *NullableNetworksNetworkIdApplianceVpnBgpAuthentication {
	return &NullableNetworksNetworkIdApplianceVpnBgpAuthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVpnBgpAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVpnBgpAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


