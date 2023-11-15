/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFirmwareUpgradesRollbacksReasons struct for NetworksNetworkIdFirmwareUpgradesRollbacksReasons
type NetworksNetworkIdFirmwareUpgradesRollbacksReasons struct {
	// Reason for the rollback
	Category string `json:"category"`
	// Additional comment about the rollback
	Comment string `json:"comment"`
}

// NewNetworksNetworkIdFirmwareUpgradesRollbacksReasons instantiates a new NetworksNetworkIdFirmwareUpgradesRollbacksReasons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesRollbacksReasons(category string, comment string) *NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	this := NetworksNetworkIdFirmwareUpgradesRollbacksReasons{}
	this.Category = category
	this.Comment = comment
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesRollbacksReasonsWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesRollbacksReasons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesRollbacksReasonsWithDefaults() *NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	this := NetworksNetworkIdFirmwareUpgradesRollbacksReasons{}
	return &this
}

// GetCategory returns the Category field value
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) SetCategory(v string) {
	o.Category = v
}

// GetComment returns the Comment field value
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) GetCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) SetComment(v string) {
	o.Comment = v
}

func (o NetworksNetworkIdFirmwareUpgradesRollbacksReasons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons struct {
	value *NetworksNetworkIdFirmwareUpgradesRollbacksReasons
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) Get() *NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) Set(val *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons(val *NetworksNetworkIdFirmwareUpgradesRollbacksReasons) *NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	return &NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesRollbacksReasons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


