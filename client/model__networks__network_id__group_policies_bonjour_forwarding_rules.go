/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesBonjourForwardingRules struct for NetworksNetworkIdGroupPoliciesBonjourForwardingRules
type NetworksNetworkIdGroupPoliciesBonjourForwardingRules struct {
	// A description for your Bonjour forwarding rule. Optional.
	Description *string `json:"description,omitempty"`
	// The ID of the service VLAN. Required.
	VlanId string `json:"vlanId"`
	// A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AFP', 'AirPlay', 'Apple screen share', 'BitTorrent', 'Chromecast', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners', 'Spotify' and 'SSH'
	Services []string `json:"services"`
}

// NewNetworksNetworkIdGroupPoliciesBonjourForwardingRules instantiates a new NetworksNetworkIdGroupPoliciesBonjourForwardingRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesBonjourForwardingRules(vlanId string, services []string) *NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	this := NetworksNetworkIdGroupPoliciesBonjourForwardingRules{}
	this.VlanId = vlanId
	this.Services = services
	return &this
}

// NewNetworksNetworkIdGroupPoliciesBonjourForwardingRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesBonjourForwardingRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesBonjourForwardingRulesWithDefaults() *NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	this := NetworksNetworkIdGroupPoliciesBonjourForwardingRules{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) SetDescription(v string) {
	o.Description = &v
}

// GetVlanId returns the VlanId field value
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetVlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetVlanIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) SetVlanId(v string) {
	o.VlanId = v
}

// GetServices returns the Services field value
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) GetServicesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) SetServices(v []string) {
	o.Services = v
}

func (o NetworksNetworkIdGroupPoliciesBonjourForwardingRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if true {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules struct {
	value *NetworksNetworkIdGroupPoliciesBonjourForwardingRules
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) Get() *NetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) Set(val *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules(val *NetworksNetworkIdGroupPoliciesBonjourForwardingRules) *NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules {
	return &NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesBonjourForwardingRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


