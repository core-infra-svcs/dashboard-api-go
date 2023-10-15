/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDevice info related to the device that caused the alert
type NetworksNetworkIdAlertsHistoryDevice struct {
	// device serial
	Serial *string `json:"serial,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDevice instantiates a new NetworksNetworkIdAlertsHistoryDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDevice() *NetworksNetworkIdAlertsHistoryDevice {
	this := NetworksNetworkIdAlertsHistoryDevice{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDeviceWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDeviceWithDefaults() *NetworksNetworkIdAlertsHistoryDevice {
	this := NetworksNetworkIdAlertsHistoryDevice{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDevice) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDevice) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDevice) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdAlertsHistoryDevice) SetSerial(v string) {
	o.Serial = &v
}

func (o NetworksNetworkIdAlertsHistoryDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDevice struct {
	value *NetworksNetworkIdAlertsHistoryDevice
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDevice) Get() *NetworksNetworkIdAlertsHistoryDevice {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDevice) Set(val *NetworksNetworkIdAlertsHistoryDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDevice(val *NetworksNetworkIdAlertsHistoryDevice) *NullableNetworksNetworkIdAlertsHistoryDevice {
	return &NullableNetworksNetworkIdAlertsHistoryDevice{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


