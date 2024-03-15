/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200102Overrides struct for InlineResponse200102Overrides
type InlineResponse200102Overrides struct {
	// (optional) List of switch serials for non-template network
	Switches []string `json:"switches,omitempty"`
	// (optional) List of switch stack ids for non-template network
	Stacks []string `json:"stacks,omitempty"`
	// (optional) List of switch templates ids for template network
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// IGMP snooping enabled for switches, switch stacks or switch templates
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic enabled for switches, switch stacks or switch templates
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewInlineResponse200102Overrides instantiates a new InlineResponse200102Overrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200102Overrides() *InlineResponse200102Overrides {
	this := InlineResponse200102Overrides{}
	return &this
}

// NewInlineResponse200102OverridesWithDefaults instantiates a new InlineResponse200102Overrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200102OverridesWithDefaults() *InlineResponse200102Overrides {
	this := InlineResponse200102Overrides{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineResponse200102Overrides) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Overrides) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineResponse200102Overrides) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *InlineResponse200102Overrides) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *InlineResponse200102Overrides) GetStacks() []string {
	if o == nil || isNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Overrides) GetStacksOk() ([]string, bool) {
	if o == nil || isNil(o.Stacks) {
    return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *InlineResponse200102Overrides) HasStacks() bool {
	if o != nil && !isNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *InlineResponse200102Overrides) SetStacks(v []string) {
	o.Stacks = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *InlineResponse200102Overrides) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Overrides) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *InlineResponse200102Overrides) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *InlineResponse200102Overrides) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200102Overrides) GetIgmpSnoopingEnabled() bool {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Overrides) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
    return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200102Overrides) HasIgmpSnoopingEnabled() bool {
	if o != nil && !isNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *InlineResponse200102Overrides) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *InlineResponse200102Overrides) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Overrides) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
    return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *InlineResponse200102Overrides) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *InlineResponse200102Overrides) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o InlineResponse200102Overrides) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.IgmpSnoopingEnabled) {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200102Overrides struct {
	value *InlineResponse200102Overrides
	isSet bool
}

func (v NullableInlineResponse200102Overrides) Get() *InlineResponse200102Overrides {
	return v.value
}

func (v *NullableInlineResponse200102Overrides) Set(val *InlineResponse200102Overrides) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200102Overrides) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200102Overrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200102Overrides(val *InlineResponse200102Overrides) *NullableInlineResponse200102Overrides {
	return &NullableInlineResponse200102Overrides{value: val, isSet: true}
}

func (v NullableInlineResponse200102Overrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200102Overrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

