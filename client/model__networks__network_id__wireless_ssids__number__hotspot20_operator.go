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

// NetworksNetworkIdWirelessSsidsNumberHotspot20Operator Operator settings for this SSID
type NetworksNetworkIdWirelessSsidsNumberHotspot20Operator struct {
	// Operator name
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20Operator instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Operator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20Operator() *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Operator{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20OperatorWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Operator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20OperatorWithDefaults() *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Operator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator struct {
	value *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) Get() *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) Set(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Operator) *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator {
	return &NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Operator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


