/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments An object mapping MAC addresses to IP addresses and client names
type NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments struct {
	// Assigned IP address
	Ip *string `json:"ip,omitempty"`
	// Client name
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments instantiates a new NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments() *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments {
	this := NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments{}
	return &this
}

// NewNetworksNetworkIdApplianceStaticRoutesFixedIpAssignmentsWithDefaults instantiates a new NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceStaticRoutesFixedIpAssignmentsWithDefaults() *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments {
	this := NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments struct {
	value *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) Get() *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) Set(val *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments(val *NetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) *NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments {
	return &NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceStaticRoutesFixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


