/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesSecurePort The Secure Port status of the port.
type DevicesSerialSwitchPortsStatusesSecurePort struct {
	// Whether Secure Port is turned on for this port.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether Secure Port is currently active for this port.
	Active *bool `json:"active,omitempty"`
	// The current Secure Port status.
	AuthenticationStatus *string `json:"authenticationStatus,omitempty"`
	ConfigOverrides *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides `json:"configOverrides,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesSecurePort instantiates a new DevicesSerialSwitchPortsStatusesSecurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesSecurePort() *DevicesSerialSwitchPortsStatusesSecurePort {
	this := DevicesSerialSwitchPortsStatusesSecurePort{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesSecurePortWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesSecurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesSecurePortWithDefaults() *DevicesSerialSwitchPortsStatusesSecurePort {
	this := DevicesSerialSwitchPortsStatusesSecurePort{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) SetActive(v bool) {
	o.Active = &v
}

// GetAuthenticationStatus returns the AuthenticationStatus field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetAuthenticationStatus() string {
	if o == nil || isNil(o.AuthenticationStatus) {
		var ret string
		return ret
	}
	return *o.AuthenticationStatus
}

// GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetAuthenticationStatusOk() (*string, bool) {
	if o == nil || isNil(o.AuthenticationStatus) {
    return nil, false
	}
	return o.AuthenticationStatus, true
}

// HasAuthenticationStatus returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) HasAuthenticationStatus() bool {
	if o != nil && !isNil(o.AuthenticationStatus) {
		return true
	}

	return false
}

// SetAuthenticationStatus gets a reference to the given string and assigns it to the AuthenticationStatus field.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) SetAuthenticationStatus(v string) {
	o.AuthenticationStatus = &v
}

// GetConfigOverrides returns the ConfigOverrides field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetConfigOverrides() DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides {
	if o == nil || isNil(o.ConfigOverrides) {
		var ret DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides
		return ret
	}
	return *o.ConfigOverrides
}

// GetConfigOverridesOk returns a tuple with the ConfigOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) GetConfigOverridesOk() (*DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides, bool) {
	if o == nil || isNil(o.ConfigOverrides) {
    return nil, false
	}
	return o.ConfigOverrides, true
}

// HasConfigOverrides returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) HasConfigOverrides() bool {
	if o != nil && !isNil(o.ConfigOverrides) {
		return true
	}

	return false
}

// SetConfigOverrides gets a reference to the given DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides and assigns it to the ConfigOverrides field.
func (o *DevicesSerialSwitchPortsStatusesSecurePort) SetConfigOverrides(v DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) {
	o.ConfigOverrides = &v
}

func (o DevicesSerialSwitchPortsStatusesSecurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.AuthenticationStatus) {
		toSerialize["authenticationStatus"] = o.AuthenticationStatus
	}
	if !isNil(o.ConfigOverrides) {
		toSerialize["configOverrides"] = o.ConfigOverrides
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesSecurePort struct {
	value *DevicesSerialSwitchPortsStatusesSecurePort
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePort) Get() *DevicesSerialSwitchPortsStatusesSecurePort {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePort) Set(val *DevicesSerialSwitchPortsStatusesSecurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesSecurePort(val *DevicesSerialSwitchPortsStatusesSecurePort) *NullableDevicesSerialSwitchPortsStatusesSecurePort {
	return &NullableDevicesSerialSwitchPortsStatusesSecurePort{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


