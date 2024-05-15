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

// DevicesSerialCameraWirelessProfilesIds The ids of the wireless profile to assign to the given camera
type DevicesSerialCameraWirelessProfilesIds struct {
	// The id of the primary wireless profile
	Primary *string `json:"primary,omitempty"`
	// The id of the secondary wireless profile
	Secondary *string `json:"secondary,omitempty"`
	// The id of the backup wireless profile
	Backup *string `json:"backup,omitempty"`
}

// NewDevicesSerialCameraWirelessProfilesIds instantiates a new DevicesSerialCameraWirelessProfilesIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCameraWirelessProfilesIds() *DevicesSerialCameraWirelessProfilesIds {
	this := DevicesSerialCameraWirelessProfilesIds{}
	return &this
}

// NewDevicesSerialCameraWirelessProfilesIdsWithDefaults instantiates a new DevicesSerialCameraWirelessProfilesIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCameraWirelessProfilesIdsWithDefaults() *DevicesSerialCameraWirelessProfilesIds {
	this := DevicesSerialCameraWirelessProfilesIds{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *DevicesSerialCameraWirelessProfilesIds) GetPrimary() string {
	if o == nil || isNil(o.Primary) {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) GetPrimaryOk() (*string, bool) {
	if o == nil || isNil(o.Primary) {
    return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) HasPrimary() bool {
	if o != nil && !isNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *DevicesSerialCameraWirelessProfilesIds) SetPrimary(v string) {
	o.Primary = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *DevicesSerialCameraWirelessProfilesIds) GetSecondary() string {
	if o == nil || isNil(o.Secondary) {
		var ret string
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) GetSecondaryOk() (*string, bool) {
	if o == nil || isNil(o.Secondary) {
    return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) HasSecondary() bool {
	if o != nil && !isNil(o.Secondary) {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given string and assigns it to the Secondary field.
func (o *DevicesSerialCameraWirelessProfilesIds) SetSecondary(v string) {
	o.Secondary = &v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *DevicesSerialCameraWirelessProfilesIds) GetBackup() string {
	if o == nil || isNil(o.Backup) {
		var ret string
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) GetBackupOk() (*string, bool) {
	if o == nil || isNil(o.Backup) {
    return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *DevicesSerialCameraWirelessProfilesIds) HasBackup() bool {
	if o != nil && !isNil(o.Backup) {
		return true
	}

	return false
}

// SetBackup gets a reference to the given string and assigns it to the Backup field.
func (o *DevicesSerialCameraWirelessProfilesIds) SetBackup(v string) {
	o.Backup = &v
}

func (o DevicesSerialCameraWirelessProfilesIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !isNil(o.Secondary) {
		toSerialize["secondary"] = o.Secondary
	}
	if !isNil(o.Backup) {
		toSerialize["backup"] = o.Backup
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCameraWirelessProfilesIds struct {
	value *DevicesSerialCameraWirelessProfilesIds
	isSet bool
}

func (v NullableDevicesSerialCameraWirelessProfilesIds) Get() *DevicesSerialCameraWirelessProfilesIds {
	return v.value
}

func (v *NullableDevicesSerialCameraWirelessProfilesIds) Set(val *DevicesSerialCameraWirelessProfilesIds) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCameraWirelessProfilesIds) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCameraWirelessProfilesIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCameraWirelessProfilesIds(val *DevicesSerialCameraWirelessProfilesIds) *NullableDevicesSerialCameraWirelessProfilesIds {
	return &NullableDevicesSerialCameraWirelessProfilesIds{value: val, isSet: true}
}

func (v NullableDevicesSerialCameraWirelessProfilesIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCameraWirelessProfilesIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


