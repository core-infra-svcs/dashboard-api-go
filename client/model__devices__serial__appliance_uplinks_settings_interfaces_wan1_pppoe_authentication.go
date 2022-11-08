/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication Settings for PPPoE Authentication.
type DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	// Whether PPPoE authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Username for PPPoE authentication.
	Username *string `json:"username,omitempty"`
	// Password for PPPoE authentication. This parameter is not returned.
	Password *string `json:"password,omitempty"`
}

// NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication() *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
	this := DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{}
	return &this
}

// NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationWithDefaults instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationWithDefaults() *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
	this := DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	value *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication
	isSet bool
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) Get() *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
	return v.value
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) Set(val *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication(val *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
	return &NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{value: val, isSet: true}
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


