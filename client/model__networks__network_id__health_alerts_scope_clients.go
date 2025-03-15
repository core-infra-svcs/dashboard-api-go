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

// NetworksNetworkIdHealthAlertsScopeClients struct for NetworksNetworkIdHealthAlertsScopeClients
type NetworksNetworkIdHealthAlertsScopeClients struct {
	// Mac address of the client
	Mac *string `json:"mac,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopeClients instantiates a new NetworksNetworkIdHealthAlertsScopeClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopeClients() *NetworksNetworkIdHealthAlertsScopeClients {
	this := NetworksNetworkIdHealthAlertsScopeClients{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeClientsWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeClientsWithDefaults() *NetworksNetworkIdHealthAlertsScopeClients {
	this := NetworksNetworkIdHealthAlertsScopeClients{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeClients) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeClients) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeClients) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworksNetworkIdHealthAlertsScopeClients) SetMac(v string) {
	o.Mac = &v
}

func (o NetworksNetworkIdHealthAlertsScopeClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopeClients struct {
	value *NetworksNetworkIdHealthAlertsScopeClients
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopeClients) Get() *NetworksNetworkIdHealthAlertsScopeClients {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeClients) Set(val *NetworksNetworkIdHealthAlertsScopeClients) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopeClients) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopeClients(val *NetworksNetworkIdHealthAlertsScopeClients) *NullableNetworksNetworkIdHealthAlertsScopeClients {
	return &NullableNetworksNetworkIdHealthAlertsScopeClients{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopeClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


