/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdClientsProvisionClients struct for NetworksNetworkIdClientsProvisionClients
type NetworksNetworkIdClientsProvisionClients struct {
	// The MAC address of the client. Required.
	Mac string `json:"mac"`
	// The display name for the client. Optional. Limited to 255 bytes.
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdClientsProvisionClients instantiates a new NetworksNetworkIdClientsProvisionClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsProvisionClients(mac string) *NetworksNetworkIdClientsProvisionClients {
	this := NetworksNetworkIdClientsProvisionClients{}
	this.Mac = mac
	return &this
}

// NewNetworksNetworkIdClientsProvisionClientsWithDefaults instantiates a new NetworksNetworkIdClientsProvisionClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsProvisionClientsWithDefaults() *NetworksNetworkIdClientsProvisionClients {
	this := NetworksNetworkIdClientsProvisionClients{}
	return &this
}

// GetMac returns the Mac field value
func (o *NetworksNetworkIdClientsProvisionClients) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionClients) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *NetworksNetworkIdClientsProvisionClients) SetMac(v string) {
	o.Mac = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionClients) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionClients) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionClients) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdClientsProvisionClients) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdClientsProvisionClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsProvisionClients struct {
	value *NetworksNetworkIdClientsProvisionClients
	isSet bool
}

func (v NullableNetworksNetworkIdClientsProvisionClients) Get() *NetworksNetworkIdClientsProvisionClients {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsProvisionClients) Set(val *NetworksNetworkIdClientsProvisionClients) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsProvisionClients) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsProvisionClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsProvisionClients(val *NetworksNetworkIdClientsProvisionClients) *NullableNetworksNetworkIdClientsProvisionClients {
	return &NullableNetworksNetworkIdClientsProvisionClients{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsProvisionClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsProvisionClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


