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

// InlineResponse200146 struct for InlineResponse200146
type InlineResponse200146 struct {
	// Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set
	Enabled *bool `json:"enabled,omitempty"`
	// Alternate management VLAN, must be between 1 and 4094
	VlanId *int32 `json:"vlanId,omitempty"`
	// Can be one or more of the following values: 'radius', 'snmp' or 'syslog'
	Protocols []string `json:"protocols,omitempty"`
	// Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments
	Switches []InlineResponse200146Switches `json:"switches,omitempty"`
}

// NewInlineResponse200146 instantiates a new InlineResponse200146 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200146() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// NewInlineResponse200146WithDefaults instantiates a new InlineResponse200146 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200146WithDefaults() *InlineResponse200146 {
	this := InlineResponse200146{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200146) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200146) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200146) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse200146) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse200146) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineResponse200146) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *InlineResponse200146) GetProtocols() []string {
	if o == nil || isNil(o.Protocols) {
		var ret []string
		return ret
	}
	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetProtocolsOk() ([]string, bool) {
	if o == nil || isNil(o.Protocols) {
    return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *InlineResponse200146) HasProtocols() bool {
	if o != nil && !isNil(o.Protocols) {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *InlineResponse200146) SetProtocols(v []string) {
	o.Protocols = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineResponse200146) GetSwitches() []InlineResponse200146Switches {
	if o == nil || isNil(o.Switches) {
		var ret []InlineResponse200146Switches
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200146) GetSwitchesOk() ([]InlineResponse200146Switches, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineResponse200146) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []InlineResponse200146Switches and assigns it to the Switches field.
func (o *InlineResponse200146) SetSwitches(v []InlineResponse200146Switches) {
	o.Switches = v
}

func (o InlineResponse200146) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200146 struct {
	value *InlineResponse200146
	isSet bool
}

func (v NullableInlineResponse200146) Get() *InlineResponse200146 {
	return v.value
}

func (v *NullableInlineResponse200146) Set(val *InlineResponse200146) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200146) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200146) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200146(val *InlineResponse200146) *NullableInlineResponse200146 {
	return &NullableInlineResponse200146{value: val, isSet: true}
}

func (v NullableInlineResponse200146) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200146) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


