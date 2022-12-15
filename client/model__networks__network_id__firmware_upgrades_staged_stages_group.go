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

// NetworksNetworkIdFirmwareUpgradesStagedStagesGroup The Staged Upgrade Group
type NetworksNetworkIdFirmwareUpgradesStagedStagesGroup struct {
	// Id of the Staged Upgrade Group
	Id *string `json:"id,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroup() *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesGroup{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroupWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesStagedStagesGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesStagedStagesGroupWithDefaults() *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	this := NetworksNetworkIdFirmwareUpgradesStagedStagesGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) SetDescription(v string) {
	o.Description = &v
}

func (o NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup struct {
	value *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) Get() *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) Set(val *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup(val *NetworksNetworkIdFirmwareUpgradesStagedStagesGroup) *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup {
	return &NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesStagedStagesGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


