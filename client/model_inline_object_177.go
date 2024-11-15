/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject177 struct for InlineObject177
type InlineObject177 struct {
	// Boolean value to enable or disable alternate management interface
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management interface VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp', 'syslog' or 'ldap'
	Protocols []string `json:"protocols,omitempty"`
	// Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put 'accessPoints' in the body when updating template networks. Also, an empty 'accessPoints' array will remove all previous static IP assignments
	AccessPoints []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints `json:"accessPoints,omitempty"`
}

// NewInlineObject177 instantiates a new InlineObject177 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject177() *InlineObject177 {
	this := InlineObject177{}
	return &this
}

// NewInlineObject177WithDefaults instantiates a new InlineObject177 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject177WithDefaults() *InlineObject177 {
	this := InlineObject177{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject177) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject177) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject177) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineObject177) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject177) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineObject177) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *InlineObject177) GetProtocols() []string {
	if o == nil || isNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetProtocolsOk() ([]string, bool) {
	if o == nil || isNil(o.Protocols) {
    return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *InlineObject177) HasProtocols() bool {
	if o != nil && !isNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *InlineObject177) SetProtocols(v []string) {
	o.Protocols = v
}

// GetAccessPoints returns the AccessPoints field value if set, zero value otherwise.
func (o *InlineObject177) GetAccessPoints() []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints {
	if o == nil || isNil(o.AccessPoints) {
		var ret []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints
		return ret
	}
	return o.AccessPoints
}

// GetAccessPointsOk returns a tuple with the AccessPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject177) GetAccessPointsOk() ([]NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints, bool) {
	if o == nil || isNil(o.AccessPoints) {
    return nil, false
	}
	return o.AccessPoints, true
}

// HasAccessPoints returns a boolean if a field has been set.
func (o *InlineObject177) HasAccessPoints() bool {
	if o != nil && !isNil(o.AccessPoints) {
		return true
	}

	return false
}

// SetAccessPoints gets a reference to the given []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints and assigns it to the AccessPoints field.
func (o *InlineObject177) SetAccessPoints(v []NetworksNetworkIdWirelessAlternateManagementInterfaceAccessPoints) {
	o.AccessPoints = v
}

func (o InlineObject177) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.Protocols) {
		toSerialize["protocols"] = o.Protocols
	}
	if !isNil(o.AccessPoints) {
		toSerialize["accessPoints"] = o.AccessPoints
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject177 struct {
	value *InlineObject177
	isSet bool
}

func (v NullableInlineObject177) Get() *InlineObject177 {
	return v.value
}

func (v *NullableInlineObject177) Set(val *InlineObject177) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject177) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject177) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject177(val *InlineObject177) *NullableInlineObject177 {
	return &NullableInlineObject177{value: val, isSet: true}
}

func (v NullableInlineObject177) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject177) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


