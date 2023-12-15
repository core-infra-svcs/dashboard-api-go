/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject119 struct for InlineObject119
type InlineObject119 struct {
	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols,omitempty"`
	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches `json:"switches,omitempty"`
}

// NewInlineObject119 instantiates a new InlineObject119 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject119() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// NewInlineObject119WithDefaults instantiates a new InlineObject119 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject119WithDefaults() *InlineObject119 {
	this := InlineObject119{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject119) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject119) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject119) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineObject119) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject119) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineObject119) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *InlineObject119) GetProtocols() []string {
	if o == nil || isNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetProtocolsOk() ([]string, bool) {
	if o == nil || isNil(o.Protocols) {
    return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *InlineObject119) HasProtocols() bool {
	if o != nil && !isNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *InlineObject119) SetProtocols(v []string) {
	o.Protocols = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineObject119) GetSwitches() []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches {
	if o == nil || isNil(o.Switches) {
		var ret []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject119) GetSwitchesOk() ([]NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineObject119) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches and assigns it to the Switches field.
func (o *InlineObject119) SetSwitches(v []NetworksNetworkIdSwitchAlternateManagementInterfaceSwitches) {
	o.Switches = v
}

func (o InlineObject119) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject119 struct {
	value *InlineObject119
	isSet bool
}

func (v NullableInlineObject119) Get() *InlineObject119 {
	return v.value
}

func (v *NullableInlineObject119) Set(val *InlineObject119) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject119) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject119) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject119(val *InlineObject119) *NullableInlineObject119 {
	return &NullableInlineObject119{value: val, isSet: true}
}

func (v NullableInlineObject119) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject119) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


