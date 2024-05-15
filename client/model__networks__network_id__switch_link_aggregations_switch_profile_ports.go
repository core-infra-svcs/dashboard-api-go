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

// NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts struct for NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
type NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts struct {
	// Profile identifier.
	Profile string `json:"profile"`
	// Port identifier of switch port. For modules, the identifier is \"SlotNumber_ModuleType_PortNumber\" (Ex: \"1_8X10G_1\"), otherwise it is just the port number (Ex: \"8\").
	PortId string `json:"portId"`
}

// NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts(profile string, portId string) *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	this := NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts{}
	this.Profile = profile
	this.PortId = portId
	return &this
}

// NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePortsWithDefaults instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePortsWithDefaults() *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	this := NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts{}
	return &this
}

// GetProfile returns the Profile field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetProfileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) SetProfile(v string) {
	o.Profile = v
}

// GetPortId returns the PortId field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetPortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) GetPortIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PortId, true
}

// SetPortId sets field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) SetPortId(v string) {
	o.PortId = v
}

func (o NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["portId"] = o.PortId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts struct {
	value *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) Get() *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) Set(val *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts(val *NetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts {
	return &NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchProfilePorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


