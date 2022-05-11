/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberOauth The OAuth settings of this SSID. Only valid if splashPage is 'Google OAuth'.
type NetworksNetworkIdWirelessSsidsNumberOauth struct {
	// (Optional) The list of domains allowed access to the network.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberOauth instantiates a new NetworksNetworkIdWirelessSsidsNumberOauth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberOauth() *NetworksNetworkIdWirelessSsidsNumberOauth {
	this := NetworksNetworkIdWirelessSsidsNumberOauth{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberOauthWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberOauth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberOauthWithDefaults() *NetworksNetworkIdWirelessSsidsNumberOauth {
	this := NetworksNetworkIdWirelessSsidsNumberOauth{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberOauth) GetAllowedDomains() []string {
	if o == nil || o.AllowedDomains == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomains
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberOauth) GetAllowedDomainsOk() (*[]string, bool) {
	if o == nil || o.AllowedDomains == nil {
		return nil, false
	}
	return o.AllowedDomains, true
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberOauth) HasAllowedDomains() bool {
	if o != nil && o.AllowedDomains != nil {
		return true
	}

	return false
}

// SetAllowedDomains gets a reference to the given []string and assigns it to the AllowedDomains field.
func (o *NetworksNetworkIdWirelessSsidsNumberOauth) SetAllowedDomains(v []string) {
	o.AllowedDomains = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberOauth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedDomains != nil {
		toSerialize["allowedDomains"] = o.AllowedDomains
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberOauth struct {
	value *NetworksNetworkIdWirelessSsidsNumberOauth
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberOauth) Get() *NetworksNetworkIdWirelessSsidsNumberOauth {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberOauth) Set(val *NetworksNetworkIdWirelessSsidsNumberOauth) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberOauth) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberOauth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberOauth(val *NetworksNetworkIdWirelessSsidsNumberOauth) *NullableNetworksNetworkIdWirelessSsidsNumberOauth {
	return &NullableNetworksNetworkIdWirelessSsidsNumberOauth{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberOauth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberOauth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


