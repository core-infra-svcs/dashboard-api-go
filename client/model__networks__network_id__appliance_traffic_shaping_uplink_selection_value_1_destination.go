/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination Destination of this custom type traffic filter
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination struct {
	// E.g.: \"any\", \"0\" (also means \"any\"), \"8080\", \"1-1024\"
	Port *string `json:"port,omitempty"`
	// CIDR format address, or \"any\". E.g.: \"192.168.10.0/24\",  \"192.168.10.1\" (same as \"192.168.10.1/32\"), \"0.0.0.0/0\" (same as \"any\")
	Cidr *string `json:"cidr,omitempty"`
	// Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: \"L_12345678\".
	Network *string `json:"network,omitempty"`
	// VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.
	Vlan *int32 `json:"vlan,omitempty"`
	// Host ID in the VLAN, should be used along with 'vlan', and not exceed the vlan subnet capacity. Currently only available under a template network.
	Host *int32 `json:"host,omitempty"`
	// FQDN format address. Currently only availabe in 'destination' of 'vpnTrafficUplinkPreference' object. E.g.: 'www.google.com'
	Fqdn *string `json:"fqdn,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination{}
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1DestinationWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1DestinationWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetPort(v string) {
	o.Port = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetCidr(v string) {
	o.Cidr = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetNetwork() string {
	if o == nil || isNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetNetworkOk() (*string, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetNetwork(v string) {
	o.Network = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetVlan(v int32) {
	o.Vlan = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetHost() int32 {
	if o == nil || isNil(o.Host) {
		var ret int32
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetHostOk() (*int32, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given int32 and assigns it to the Host field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetHost(v int32) {
	o.Host = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) SetFqdn(v string) {
	o.Fqdn = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1Destination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


