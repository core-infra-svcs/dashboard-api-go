/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe Configuration options for PPPoE.
type DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	// Whether PPPoE is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Authentication *DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"`
}

// NewDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe() *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe {
	this := DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe{}
	return &this
}

// NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeWithDefaults instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeWithDefaults() *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe {
	this := DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) GetAuthentication() DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) GetAuthenticationOk() (*DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication and assigns it to the Authentication field.
func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) SetAuthentication(v DevicesSerialApplianceUplinksSettingsInterfacesWan1PppoeAuthentication) {
	o.Authentication = &v
}

func (o DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	value *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe
	isSet bool
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) Get() *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe {
	return v.value
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) Set(val *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe(val *DevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe {
	return &NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe{value: val, isSet: true}
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfacesWan1Pppoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


