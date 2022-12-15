/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 struct for NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 struct {
	// Type of this traffic filter. Must be one of: 'applicationCategory', 'application' or 'custom'
	Type string `json:"type"`
	Value NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 `json:"value"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1(type_ string, value NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1WithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1WithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1{}
	return &this
}

// GetType returns the Type field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) GetValue() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1 {
	if o == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) GetValueOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) SetValue(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue1) {
	o.Value = v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1 {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


