/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesRollbacksToVersion Version to downgrade to (if the network has firmware flexibility)
type NetworksNetworkIdFirmwareUpgradesRollbacksToVersion struct {
	// The version ID
	Id *string `json:"id,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesRollbacksToVersion instantiates a new NetworksNetworkIdFirmwareUpgradesRollbacksToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesRollbacksToVersion() *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	this := NetworksNetworkIdFirmwareUpgradesRollbacksToVersion{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesRollbacksToVersionWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesRollbacksToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesRollbacksToVersionWithDefaults() *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	this := NetworksNetworkIdFirmwareUpgradesRollbacksToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) SetId(v string) {
	o.Id = &v
}

func (o NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion struct {
	value *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) Get() *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) Set(val *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion(val *NetworksNetworkIdFirmwareUpgradesRollbacksToVersion) *NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion {
	return &NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


