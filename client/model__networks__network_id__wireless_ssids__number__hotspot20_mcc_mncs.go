/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs struct for NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs
type NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs struct {
	// MCC value
	Mcc *string `json:"mcc,omitempty"`
	// MNC value
	Mnc *string `json:"mnc,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs() *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) GetMcc() string {
	if o == nil || isNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) GetMccOk() (*string, bool) {
	if o == nil || isNil(o.Mcc) {
    return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) HasMcc() bool {
	if o != nil && !isNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) GetMnc() string {
	if o == nil || isNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) GetMncOk() (*string, bool) {
	if o == nil || isNil(o.Mnc) {
    return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) HasMnc() bool {
	if o != nil && !isNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) SetMnc(v string) {
	o.Mnc = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !isNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs struct {
	value *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) Get() *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) Set(val *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs(val *NetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs {
	return &NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20MccMncs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


