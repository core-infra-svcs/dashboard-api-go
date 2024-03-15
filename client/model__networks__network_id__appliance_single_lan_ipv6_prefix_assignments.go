/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments struct for NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments
type NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments struct {
	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous *bool `json:"autonomous,omitempty"`
	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix *string `json:"staticPrefix,omitempty"`
	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 *string `json:"staticApplianceIp6,omitempty"`
	Origin *NetworksNetworkIdApplianceSingleLanIpv6Origin `json:"origin,omitempty"`
}

// NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments instantiates a new NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments() *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments {
	this := NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments{}
	return &this
}

// NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignmentsWithDefaults instantiates a new NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignmentsWithDefaults() *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments {
	this := NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments{}
	return &this
}

// GetAutonomous returns the Autonomous field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetAutonomous() bool {
	if o == nil || isNil(o.Autonomous) {
		var ret bool
		return ret
	}
	return *o.Autonomous
}

// GetAutonomousOk returns a tuple with the Autonomous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetAutonomousOk() (*bool, bool) {
	if o == nil || isNil(o.Autonomous) {
    return nil, false
	}
	return o.Autonomous, true
}

// HasAutonomous returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasAutonomous() bool {
	if o != nil && !isNil(o.Autonomous) {
		return true
	}

	return false
}

// SetAutonomous gets a reference to the given bool and assigns it to the Autonomous field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetAutonomous(v bool) {
	o.Autonomous = &v
}

// GetStaticPrefix returns the StaticPrefix field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticPrefix() string {
	if o == nil || isNil(o.StaticPrefix) {
		var ret string
		return ret
	}
	return *o.StaticPrefix
}

// GetStaticPrefixOk returns a tuple with the StaticPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticPrefixOk() (*string, bool) {
	if o == nil || isNil(o.StaticPrefix) {
    return nil, false
	}
	return o.StaticPrefix, true
}

// HasStaticPrefix returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasStaticPrefix() bool {
	if o != nil && !isNil(o.StaticPrefix) {
		return true
	}

	return false
}

// SetStaticPrefix gets a reference to the given string and assigns it to the StaticPrefix field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetStaticPrefix(v string) {
	o.StaticPrefix = &v
}

// GetStaticApplianceIp6 returns the StaticApplianceIp6 field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticApplianceIp6() string {
	if o == nil || isNil(o.StaticApplianceIp6) {
		var ret string
		return ret
	}
	return *o.StaticApplianceIp6
}

// GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetStaticApplianceIp6Ok() (*string, bool) {
	if o == nil || isNil(o.StaticApplianceIp6) {
    return nil, false
	}
	return o.StaticApplianceIp6, true
}

// HasStaticApplianceIp6 returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasStaticApplianceIp6() bool {
	if o != nil && !isNil(o.StaticApplianceIp6) {
		return true
	}

	return false
}

// SetStaticApplianceIp6 gets a reference to the given string and assigns it to the StaticApplianceIp6 field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetStaticApplianceIp6(v string) {
	o.StaticApplianceIp6 = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetOrigin() NetworksNetworkIdApplianceSingleLanIpv6Origin {
	if o == nil || isNil(o.Origin) {
		var ret NetworksNetworkIdApplianceSingleLanIpv6Origin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) GetOriginOk() (*NetworksNetworkIdApplianceSingleLanIpv6Origin, bool) {
	if o == nil || isNil(o.Origin) {
    return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) HasOrigin() bool {
	if o != nil && !isNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NetworksNetworkIdApplianceSingleLanIpv6Origin and assigns it to the Origin field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) SetOrigin(v NetworksNetworkIdApplianceSingleLanIpv6Origin) {
	o.Origin = &v
}

func (o NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Autonomous) {
		toSerialize["autonomous"] = o.Autonomous
	}
	if !isNil(o.StaticPrefix) {
		toSerialize["staticPrefix"] = o.StaticPrefix
	}
	if !isNil(o.StaticApplianceIp6) {
		toSerialize["staticApplianceIp6"] = o.StaticApplianceIp6
	}
	if !isNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments struct {
	value *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) Get() *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) Set(val *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments(val *NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) *NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments {
	return &NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


