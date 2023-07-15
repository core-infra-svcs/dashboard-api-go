/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20079Overrides struct for InlineResponse20079Overrides
type InlineResponse20079Overrides struct {
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

// NewInlineResponse20079Overrides instantiates a new InlineResponse20079Overrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079Overrides() *InlineResponse20079Overrides {
	this := InlineResponse20079Overrides{}
	return &this
}

// NewInlineResponse20079OverridesWithDefaults instantiates a new InlineResponse20079Overrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079OverridesWithDefaults() *InlineResponse20079Overrides {
	this := InlineResponse20079Overrides{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineResponse20079Overrides) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Overrides) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineResponse20079Overrides) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *InlineResponse20079Overrides) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *InlineResponse20079Overrides) GetStacks() []string {
	if o == nil || isNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Overrides) GetStacksOk() ([]string, bool) {
	if o == nil || isNil(o.Stacks) {
    return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *InlineResponse20079Overrides) HasStacks() bool {
	if o != nil && !isNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *InlineResponse20079Overrides) SetStacks(v []string) {
	o.Stacks = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *InlineResponse20079Overrides) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Overrides) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *InlineResponse20079Overrides) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *InlineResponse20079Overrides) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *InlineResponse20079Overrides) GetIgmpSnoopingEnabled() bool {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Overrides) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
    return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *InlineResponse20079Overrides) HasIgmpSnoopingEnabled() bool {
	if o != nil && !isNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *InlineResponse20079Overrides) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *InlineResponse20079Overrides) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079Overrides) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
    return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *InlineResponse20079Overrides) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *InlineResponse20079Overrides) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o InlineResponse20079Overrides) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20079Overrides struct {
	value *InlineResponse20079Overrides
	isSet bool
}

func (v NullableInlineResponse20079Overrides) Get() *InlineResponse20079Overrides {
	return v.value
}

func (v *NullableInlineResponse20079Overrides) Set(val *InlineResponse20079Overrides) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079Overrides) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079Overrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079Overrides(val *InlineResponse20079Overrides) *NullableInlineResponse20079Overrides {
	return &NullableInlineResponse20079Overrides{value: val, isSet: true}
}

func (v NullableInlineResponse20079Overrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079Overrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


