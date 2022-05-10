/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject127 struct for InlineObject127
type InlineObject127 struct {
	// Boolean value to enable or disable alternate management interface
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management interface VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp', 'syslog' or 'ldap'
	Protocols []string `json:"protocols,omitempty"`
	// Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put 'accessPoints' in the body when updating template networks. Also, an empty 'accessPoints' array will remove all previous static IP assignments
	AccessPoints []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints `json:"accessPoints,omitempty"`
}

// NewInlineObject127 instantiates a new InlineObject127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject127() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// NewInlineObject127WithDefaults instantiates a new InlineObject127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject127WithDefaults() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject127) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject127) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject127) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineObject127) GetVlanId() int32 {
	if o == nil || o.VlanId == nil {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetVlanIdOk() (*int32, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject127) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineObject127) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *InlineObject127) GetProtocols() []string {
	if o == nil || o.Protocols == nil {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetProtocolsOk() ([]string, bool) {
	if o == nil || o.Protocols == nil {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *InlineObject127) HasProtocols() bool {
	if o != nil && o.Protocols != nil {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *InlineObject127) SetProtocols(v []string) {
	o.Protocols = v
}

// GetAccessPoints returns the AccessPoints field value if set, zero value otherwise.
func (o *InlineObject127) GetAccessPoints() []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	if o == nil || o.AccessPoints == nil {
		var ret []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints
		return ret
	}
	return o.AccessPoints
}

// GetAccessPointsOk returns a tuple with the AccessPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetAccessPointsOk() ([]NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints, bool) {
	if o == nil || o.AccessPoints == nil {
		return nil, false
	}
	return o.AccessPoints, true
}

// HasAccessPoints returns a boolean if a field has been set.
func (o *InlineObject127) HasAccessPoints() bool {
	if o != nil && o.AccessPoints != nil {
		return true
	}

	return false
}

// SetAccessPoints gets a reference to the given []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints and assigns it to the AccessPoints field.
func (o *InlineObject127) SetAccessPoints(v []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) {
	o.AccessPoints = v
}

func (o InlineObject127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.VlanId != nil {
		toSerialize["vlanId"] = o.VlanId
	}
	if o.Protocols != nil {
		toSerialize["protocols"] = o.Protocols
	}
	if o.AccessPoints != nil {
		toSerialize["accessPoints"] = o.AccessPoints
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject127 struct {
	value *InlineObject127
	isSet bool
}

func (v NullableInlineObject127) Get() *InlineObject127 {
	return v.value
}

func (v *NullableInlineObject127) Set(val *InlineObject127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject127(val *InlineObject127) *NullableInlineObject127 {
	return &NullableInlineObject127{value: val, isSet: true}
}

func (v NullableInlineObject127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


