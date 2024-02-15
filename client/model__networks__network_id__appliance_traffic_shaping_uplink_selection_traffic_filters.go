/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters struct for NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters
type NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters struct {
	// Type of this traffic filter. Must be one of: 'custom'
	Type string `json:"type"`
	Value NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue `json:"value"`
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters(type_ string, value NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFiltersWithDefaults instantiates a new NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFiltersWithDefaults() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters {
	this := NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters{}
	return &this
}

// GetType returns the Type field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) GetValue() NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue {
	if o == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) GetValueOk() (*NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) SetValue(v NetworksNetworkIdApplianceTrafficShapingUplinkSelectionValue) {
	o.Value = v
}

func (o NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters struct {
	value *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) Get() *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) Set(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters(val *NetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters {
	return &NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceTrafficShapingUplinkSelectionTrafficFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


