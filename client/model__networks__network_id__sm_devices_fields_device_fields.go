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

// NetworksNetworkIdSmDevicesFieldsDeviceFields The new fields of the device. Each field of this object is optional.
type NetworksNetworkIdSmDevicesFieldsDeviceFields struct {
	// New name for the device
	Name *string `json:"name,omitempty"`
	// New notes for the device
	Notes *string `json:"notes,omitempty"`
}

// NewNetworksNetworkIdSmDevicesFieldsDeviceFields instantiates a new NetworksNetworkIdSmDevicesFieldsDeviceFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSmDevicesFieldsDeviceFields() *NetworksNetworkIdSmDevicesFieldsDeviceFields {
	this := NetworksNetworkIdSmDevicesFieldsDeviceFields{}
	return &this
}

// NewNetworksNetworkIdSmDevicesFieldsDeviceFieldsWithDefaults instantiates a new NetworksNetworkIdSmDevicesFieldsDeviceFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSmDevicesFieldsDeviceFieldsWithDefaults() *NetworksNetworkIdSmDevicesFieldsDeviceFields {
	this := NetworksNetworkIdSmDevicesFieldsDeviceFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *NetworksNetworkIdSmDevicesFieldsDeviceFields) SetNotes(v string) {
	o.Notes = &v
}

func (o NetworksNetworkIdSmDevicesFieldsDeviceFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSmDevicesFieldsDeviceFields struct {
	value *NetworksNetworkIdSmDevicesFieldsDeviceFields
	isSet bool
}

func (v NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) Get() *NetworksNetworkIdSmDevicesFieldsDeviceFields {
	return v.value
}

func (v *NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) Set(val *NetworksNetworkIdSmDevicesFieldsDeviceFields) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSmDevicesFieldsDeviceFields(val *NetworksNetworkIdSmDevicesFieldsDeviceFields) *NullableNetworksNetworkIdSmDevicesFieldsDeviceFields {
	return &NullableNetworksNetworkIdSmDevicesFieldsDeviceFields{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSmDevicesFieldsDeviceFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


