/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides The configuration overrides applied to this port when Secure Port is active.
type DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides struct {
	// The type of the  ('trunk', 'access' or 'stack').
	Type *string `json:"type,omitempty"`
	// The VLAN of the . For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the . Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the . Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
}

// NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides instantiates a new DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides() *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides {
	this := DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides{}
	return &this
}

// NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverridesWithDefaults instantiates a new DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchPortsStatusesSecurePortConfigOverridesWithDefaults() *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides {
	this := DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVoiceVlan() int32 {
	if o == nil || isNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || isNil(o.VoiceVlan) {
    return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasVoiceVlan() bool {
	if o != nil && !isNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetAllowedVlans() string {
	if o == nil || isNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) GetAllowedVlansOk() (*string, bool) {
	if o == nil || isNil(o.AllowedVlans) {
    return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) HasAllowedVlans() bool {
	if o != nil && !isNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

func (o DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.VoiceVlan) {
		toSerialize["voiceVlan"] = o.VoiceVlan
	}
	if !isNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides struct {
	value *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides
	isSet bool
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) Get() *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides {
	return v.value
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) Set(val *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides(val *DevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) *NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides {
	return &NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchPortsStatusesSecurePortConfigOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


