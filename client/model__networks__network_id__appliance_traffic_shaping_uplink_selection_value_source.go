/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource Source of this custom type traffic filter
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address, or \"any\". E.g.: \"192.168.10.0/24\",  \"192.168.10.1\" (same as \"192.168.10.1/32\"), \"0.0.0.0/0\" (same as \"any\")
	Cidr *string `json:"cidr,omitempty"`
	// VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
	Vlan *int32 `json:"vlan,omitempty"`
	// Host ID in the VLAN, should be used along with 'vlan', and not exceed the vlan subnet capacity. Currently only available under a template network.
	Host *int32 `json:"host,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSourceWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSourceWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetCidr(v string) {
	o.Cidr = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetVlan(v int32) {
	o.Vlan = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetHost() int32 {
	if o == nil || isNil(o.Host) {
		var ret int32
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) GetHostOk() (*int32, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given int32 and assigns it to the Host field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) SetHost(v int32) {
	o.Host = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValueSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


