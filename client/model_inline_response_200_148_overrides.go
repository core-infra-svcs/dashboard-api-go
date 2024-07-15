/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200148Overrides struct for InlineResponse200148Overrides
type InlineResponse200148Overrides struct {
	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches,omitempty"`
	// List of switch template IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// MTU size for the switches or switch templates.
	MtuSize int32 `json:"mtuSize"`
}

// NewInlineResponse200148Overrides instantiates a new InlineResponse200148Overrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200148Overrides(mtuSize int32) *InlineResponse200148Overrides {
	this := InlineResponse200148Overrides{}
	this.MtuSize = mtuSize
	return &this
}

// NewInlineResponse200148OverridesWithDefaults instantiates a new InlineResponse200148Overrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200148OverridesWithDefaults() *InlineResponse200148Overrides {
	this := InlineResponse200148Overrides{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *InlineResponse200148Overrides) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200148Overrides) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *InlineResponse200148Overrides) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *InlineResponse200148Overrides) SetSwitches(v []string) {
	o.Switches = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *InlineResponse200148Overrides) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200148Overrides) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *InlineResponse200148Overrides) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *InlineResponse200148Overrides) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetMtuSize returns the MtuSize field value
func (o *InlineResponse200148Overrides) GetMtuSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MtuSize
}

// GetMtuSizeOk returns a tuple with the MtuSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200148Overrides) GetMtuSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MtuSize, true
}

// SetMtuSize sets field value
func (o *InlineResponse200148Overrides) SetMtuSize(v int32) {
	o.MtuSize = v
}

func (o InlineResponse200148Overrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !isNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if true {
		toSerialize["mtuSize"] = o.MtuSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200148Overrides struct {
	value *InlineResponse200148Overrides
	isSet bool
}

func (v NullableInlineResponse200148Overrides) Get() *InlineResponse200148Overrides {
	return v.value
}

func (v *NullableInlineResponse200148Overrides) Set(val *InlineResponse200148Overrides) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200148Overrides) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200148Overrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200148Overrides(val *InlineResponse200148Overrides) *NullableInlineResponse200148Overrides {
	return &NullableInlineResponse200148Overrides{value: val, isSet: true}
}

func (v NullableInlineResponse200148Overrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200148Overrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


