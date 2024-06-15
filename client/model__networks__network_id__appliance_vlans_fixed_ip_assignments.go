/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceVlansFixedIpAssignments IP assignment information, keyed by MAC address of the device
type NetworksNetworkIdApplianceVlansFixedIpAssignments struct {
	// IP address of the assignment
	Ip *string `json:"ip,omitempty"`
	// Name of the IP assignment
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdApplianceVlansFixedIpAssignments instantiates a new NetworksNetworkIdApplianceVlansFixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVlansFixedIpAssignments() *NetworksNetworkIdApplianceVlansFixedIpAssignments {
	this := NetworksNetworkIdApplianceVlansFixedIpAssignments{}
	return &this
}

// NewNetworksNetworkIdApplianceVlansFixedIpAssignmentsWithDefaults instantiates a new NetworksNetworkIdApplianceVlansFixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVlansFixedIpAssignmentsWithDefaults() *NetworksNetworkIdApplianceVlansFixedIpAssignments {
	this := NetworksNetworkIdApplianceVlansFixedIpAssignments{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceVlansFixedIpAssignments) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdApplianceVlansFixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVlansFixedIpAssignments struct {
	value *NetworksNetworkIdApplianceVlansFixedIpAssignments
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) Get() *NetworksNetworkIdApplianceVlansFixedIpAssignments {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) Set(val *NetworksNetworkIdApplianceVlansFixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVlansFixedIpAssignments(val *NetworksNetworkIdApplianceVlansFixedIpAssignments) *NullableNetworksNetworkIdApplianceVlansFixedIpAssignments {
	return &NullableNetworksNetworkIdApplianceVlansFixedIpAssignments{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVlansFixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


