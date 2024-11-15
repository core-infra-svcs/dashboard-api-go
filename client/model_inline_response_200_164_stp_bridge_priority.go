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

// InlineResponse200164StpBridgePriority struct for InlineResponse200164StpBridgePriority
type InlineResponse200164StpBridgePriority struct {
	// List of switch serial numbers
	Switches []string `json:"switches,omitempty"`
	// List of stack IDs
	Stacks []string `json:"stacks,omitempty"`
	// List of switch template IDs
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// STP priority for switch, stacks, or switch templates
	StpPriority *int32 `json:"stpPriority,omitempty"`
}

// NewInlineResponse200164StpBridgePriority instantiates a new InlineResponse200164StpBridgePriority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200164StpBridgePriority() *InlineResponse200164StpBridgePriority {
	this := InlineResponse200164StpBridgePriority{}
	return &this
}

// NewInlineResponse200164StpBridgePriorityWithDefaults instantiates a new InlineResponse200164StpBridgePriority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200164StpBridgePriorityWithDefaults() *InlineResponse200164StpBridgePriority {
	this := InlineResponse200164StpBridgePriority{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineResponse200164StpBridgePriority) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200164StpBridgePriority) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineResponse200164StpBridgePriority) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *InlineResponse200164StpBridgePriority) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *InlineResponse200164StpBridgePriority) GetStacks() []string {
	if o == nil || isNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200164StpBridgePriority) GetStacksOk() ([]string, bool) {
	if o == nil || isNil(o.Stacks) {
    return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *InlineResponse200164StpBridgePriority) HasStacks() bool {
	if o != nil && !isNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *InlineResponse200164StpBridgePriority) SetStacks(v []string) {
	o.Stacks = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *InlineResponse200164StpBridgePriority) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200164StpBridgePriority) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *InlineResponse200164StpBridgePriority) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *InlineResponse200164StpBridgePriority) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetStpPriority returns the StpPriority field value if set, zero value otherwise.
func (o *InlineResponse200164StpBridgePriority) GetStpPriority() int32 {
	if o == nil || isNil(o.StpPriority) {
		var ret int32
		return ret
	}
	return *o.StpPriority
}

// GetStpPriorityOk returns a tuple with the StpPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200164StpBridgePriority) GetStpPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.StpPriority) {
    return nil, false
	}
	return o.StpPriority, true
}

// HasStpPriority returns a boolean if a field has been set.
func (o *InlineResponse200164StpBridgePriority) HasStpPriority() bool {
	if o != nil && !isNil(o.StpPriority) {
		return true
	}

	return false
}

// SetStpPriority gets a reference to the given int32 and assigns it to the StpPriority field.
func (o *InlineResponse200164StpBridgePriority) SetStpPriority(v int32) {
	o.StpPriority = &v
}

func (o InlineResponse200164StpBridgePriority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !isNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	if !isNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if !isNil(o.StpPriority) {
		toSerialize["stpPriority"] = o.StpPriority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200164StpBridgePriority struct {
	value *InlineResponse200164StpBridgePriority
	isSet bool
}

func (v NullableInlineResponse200164StpBridgePriority) Get() *InlineResponse200164StpBridgePriority {
	return v.value
}

func (v *NullableInlineResponse200164StpBridgePriority) Set(val *InlineResponse200164StpBridgePriority) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200164StpBridgePriority) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200164StpBridgePriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200164StpBridgePriority(val *InlineResponse200164StpBridgePriority) *NullableInlineResponse200164StpBridgePriority {
	return &NullableInlineResponse200164StpBridgePriority{value: val, isSet: true}
}

func (v NullableInlineResponse200164StpBridgePriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200164StpBridgePriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

