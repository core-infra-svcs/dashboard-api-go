/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesFlexRadios Flex radio settings.
type NetworksNetworkIdWirelessRfProfilesFlexRadios struct {
	// Flex radios by model.
	ByModel []NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel `json:"byModel,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesFlexRadios instantiates a new NetworksNetworkIdWirelessRfProfilesFlexRadios object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesFlexRadios() *NetworksNetworkIdWirelessRfProfilesFlexRadios {
	this := NetworksNetworkIdWirelessRfProfilesFlexRadios{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesFlexRadiosWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesFlexRadios object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesFlexRadiosWithDefaults() *NetworksNetworkIdWirelessRfProfilesFlexRadios {
	this := NetworksNetworkIdWirelessRfProfilesFlexRadios{}
	return &this
}

// GetByModel returns the ByModel field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadios) GetByModel() []NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel {
	if o == nil || isNil(o.ByModel) {
		var ret []NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel
		return ret
	}
	return o.ByModel
}

// GetByModelOk returns a tuple with the ByModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadios) GetByModelOk() ([]NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel, bool) {
	if o == nil || isNil(o.ByModel) {
    return nil, false
	}
	return o.ByModel, true
}

// HasByModel returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadios) HasByModel() bool {
	if o != nil && !isNil(o.ByModel) {
		return true
	}

	return false
}

// SetByModel gets a reference to the given []NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel and assigns it to the ByModel field.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadios) SetByModel(v []NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) {
	o.ByModel = v
}

func (o NetworksNetworkIdWirelessRfProfilesFlexRadios) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByModel) {
		toSerialize["byModel"] = o.ByModel
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesFlexRadios struct {
	value *NetworksNetworkIdWirelessRfProfilesFlexRadios
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) Get() *NetworksNetworkIdWirelessRfProfilesFlexRadios {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) Set(val *NetworksNetworkIdWirelessRfProfilesFlexRadios) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesFlexRadios(val *NetworksNetworkIdWirelessRfProfilesFlexRadios) *NullableNetworksNetworkIdWirelessRfProfilesFlexRadios {
	return &NullableNetworksNetworkIdWirelessRfProfilesFlexRadios{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadios) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


