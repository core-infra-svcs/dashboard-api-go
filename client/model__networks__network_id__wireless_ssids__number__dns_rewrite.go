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

// NetworksNetworkIdWirelessSsidsNumberDnsRewrite DNS servers rewrite settings
type NetworksNetworkIdWirelessSsidsNumberDnsRewrite struct {
	// Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used
	Enabled *bool `json:"enabled,omitempty"`
	// User specified DNS servers (up to two servers)
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberDnsRewrite instantiates a new NetworksNetworkIdWirelessSsidsNumberDnsRewrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberDnsRewrite() *NetworksNetworkIdWirelessSsidsNumberDnsRewrite {
	this := NetworksNetworkIdWirelessSsidsNumberDnsRewrite{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberDnsRewriteWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberDnsRewrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberDnsRewriteWithDefaults() *NetworksNetworkIdWirelessSsidsNumberDnsRewrite {
	this := NetworksNetworkIdWirelessSsidsNumberDnsRewrite{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) GetDnsCustomNameservers() []string {
	if o == nil || isNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || isNil(o.DnsCustomNameservers) {
    return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) HasDnsCustomNameservers() bool {
	if o != nil && !isNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o NetworksNetworkIdWirelessSsidsNumberDnsRewrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite struct {
	value *NetworksNetworkIdWirelessSsidsNumberDnsRewrite
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) Get() *NetworksNetworkIdWirelessSsidsNumberDnsRewrite {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) Set(val *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite(val *NetworksNetworkIdWirelessSsidsNumberDnsRewrite) *NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite {
	return &NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDnsRewrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


