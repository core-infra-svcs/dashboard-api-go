/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsProfile Profile attributes
type DevicesSerialSwitchPortsProfile struct {
	// When enabled, override this port's configuration with a port profile.
	Enabled *bool `json:"enabled,omitempty"`
	// When enabled, the ID of the port profile used to override the port's configuration.
	Id *string `json:"id,omitempty"`
	// When enabled, the IName of the profile.
	Iname *string `json:"iname,omitempty"`
}

// NewDevicesSerialSwitchPortsProfile instantiates a new DevicesSerialSwitchPortsProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsProfile() *DevicesSerialSwitchPortsProfile {
	this := DevicesSerialSwitchPortsProfile{}
	return &this
}

// NewDevicesSerialSwitchPortsProfileWithDefaults instantiates a new DevicesSerialSwitchPortsProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsProfileWithDefaults() *DevicesSerialSwitchPortsProfile {
	this := DevicesSerialSwitchPortsProfile{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsProfile) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsProfile) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsProfile) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialSwitchPortsProfile) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsProfile) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsProfile) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsProfile) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DevicesSerialSwitchPortsProfile) SetId(v string) {
	o.Id = &v
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsProfile) GetIname() string {
	if o == nil || isNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsProfile) GetInameOk() (*string, bool) {
	if o == nil || isNil(o.Iname) {
    return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsProfile) HasIname() bool {
	if o != nil && !isNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *DevicesSerialSwitchPortsProfile) SetIname(v string) {
	o.Iname = &v
}

func (o DevicesSerialSwitchPortsProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsProfile struct {
	value *DevicesSerialSwitchPortsProfile
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsProfile) Get() *DevicesSerialSwitchPortsProfile {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsProfile) Set(val *DevicesSerialSwitchPortsProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsProfile(val *DevicesSerialSwitchPortsProfile) *NullableDevicesSerialSwitchPortsProfile {
	return &NullableDevicesSerialSwitchPortsProfile{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


