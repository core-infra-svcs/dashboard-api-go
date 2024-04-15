/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSensorCommandsCreatedBy Information about the admin who triggered the command
type DevicesSerialSensorCommandsCreatedBy struct {
	// ID of the admin
	AdminId *string `json:"adminId,omitempty"`
	// Name of the admin
	Name *string `json:"name,omitempty"`
	// Email of the admin
	Email *string `json:"email,omitempty"`
}

// NewDevicesSerialSensorCommandsCreatedBy instantiates a new DevicesSerialSensorCommandsCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSensorCommandsCreatedBy() *DevicesSerialSensorCommandsCreatedBy {
	this := DevicesSerialSensorCommandsCreatedBy{}
	return &this
}

// NewDevicesSerialSensorCommandsCreatedByWithDefaults instantiates a new DevicesSerialSensorCommandsCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSensorCommandsCreatedByWithDefaults() *DevicesSerialSensorCommandsCreatedBy {
	this := DevicesSerialSensorCommandsCreatedBy{}
	return &this
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *DevicesSerialSensorCommandsCreatedBy) GetAdminId() string {
	if o == nil || isNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) GetAdminIdOk() (*string, bool) {
	if o == nil || isNil(o.AdminId) {
    return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) HasAdminId() bool {
	if o != nil && !isNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *DevicesSerialSensorCommandsCreatedBy) SetAdminId(v string) {
	o.AdminId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevicesSerialSensorCommandsCreatedBy) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevicesSerialSensorCommandsCreatedBy) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DevicesSerialSensorCommandsCreatedBy) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DevicesSerialSensorCommandsCreatedBy) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DevicesSerialSensorCommandsCreatedBy) SetEmail(v string) {
	o.Email = &v
}

func (o DevicesSerialSensorCommandsCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSensorCommandsCreatedBy struct {
	value *DevicesSerialSensorCommandsCreatedBy
	isSet bool
}

func (v NullableDevicesSerialSensorCommandsCreatedBy) Get() *DevicesSerialSensorCommandsCreatedBy {
	return v.value
}

func (v *NullableDevicesSerialSensorCommandsCreatedBy) Set(val *DevicesSerialSensorCommandsCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSensorCommandsCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSensorCommandsCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSensorCommandsCreatedBy(val *DevicesSerialSensorCommandsCreatedBy) *NullableDevicesSerialSensorCommandsCreatedBy {
	return &NullableDevicesSerialSensorCommandsCreatedBy{value: val, isSet: true}
}

func (v NullableDevicesSerialSensorCommandsCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSensorCommandsCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


