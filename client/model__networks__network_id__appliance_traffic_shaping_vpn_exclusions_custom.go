/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom struct for NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom
type NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom struct {
	// Protocol.
	Protocol string `json:"protocol"`
	// Destination address; hostname required for DNS, IPv4 otherwise.
	Destination *string `json:"destination,omitempty"`
	// Destination port.
	Port *string `json:"port,omitempty"`
}

// NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom instantiates a new NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom(protocol string) *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom {
	this := NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom{}
	this.Protocol = protocol
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustomWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustomWithDefaults() *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom {
	this := NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetDestination() string {
	if o == nil || isNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetDestinationOk() (*string, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) SetDestination(v string) {
	o.Destination = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) SetPort(v string) {
	o.Port = &v
}

func (o NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom struct {
	value *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) Get() *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) Set(val *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom(val *NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom {
	return &NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


