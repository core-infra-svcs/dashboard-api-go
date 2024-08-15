/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks Set the included/excluded networks from the intrusion engine (optional - omitting will leave current config unchanged). This is available only in 'passthrough' mode
type NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks struct {
	// true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved
	UseDefault *bool `json:"useDefault,omitempty"`
	// list of IP addresses or subnets being protected (required if 'useDefault' is false)
	IncludedCidr []string `json:"includedCidr,omitempty"`
	// list of IP addresses or subnets being excluded from protection (required if 'useDefault' is false)
	ExcludedCidr []string `json:"excludedCidr,omitempty"`
}

// NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks instantiates a new NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks() *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	this := NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks{}
	return &this
}

// NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworksWithDefaults instantiates a new NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworksWithDefaults() *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	this := NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks{}
	return &this
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetUseDefault() bool {
	if o == nil || isNil(o.UseDefault) {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetUseDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefault) {
    return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasUseDefault() bool {
	if o != nil && !isNil(o.UseDefault) {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetUseDefault(v bool) {
	o.UseDefault = &v
}

// GetIncludedCidr returns the IncludedCidr field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetIncludedCidr() []string {
	if o == nil || isNil(o.IncludedCidr) {
		var ret []string
		return ret
	}
	return o.IncludedCidr
}

// GetIncludedCidrOk returns a tuple with the IncludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetIncludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedCidr) {
    return nil, false
	}
	return o.IncludedCidr, true
}

// HasIncludedCidr returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasIncludedCidr() bool {
	if o != nil && !isNil(o.IncludedCidr) {
		return true
	}

	return false
}

// SetIncludedCidr gets a reference to the given []string and assigns it to the IncludedCidr field.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetIncludedCidr(v []string) {
	o.IncludedCidr = v
}

// GetExcludedCidr returns the ExcludedCidr field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetExcludedCidr() []string {
	if o == nil || isNil(o.ExcludedCidr) {
		var ret []string
		return ret
	}
	return o.ExcludedCidr
}

// GetExcludedCidrOk returns a tuple with the ExcludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) GetExcludedCidrOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedCidr) {
    return nil, false
	}
	return o.ExcludedCidr, true
}

// HasExcludedCidr returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) HasExcludedCidr() bool {
	if o != nil && !isNil(o.ExcludedCidr) {
		return true
	}

	return false
}

// SetExcludedCidr gets a reference to the given []string and assigns it to the ExcludedCidr field.
func (o *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) SetExcludedCidr(v []string) {
	o.ExcludedCidr = v
}

func (o NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UseDefault) {
		toSerialize["useDefault"] = o.UseDefault
	}
	if !isNil(o.IncludedCidr) {
		toSerialize["includedCidr"] = o.IncludedCidr
	}
	if !isNil(o.ExcludedCidr) {
		toSerialize["excludedCidr"] = o.ExcludedCidr
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks struct {
	value *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) Get() *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) Set(val *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks(val *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) *NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	return &NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


