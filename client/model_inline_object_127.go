/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject127 struct for InlineObject127
type InlineObject127 struct {
	// The name for your port schedule. Required
	Name string `json:"name"`
	PortSchedule *NetworksNetworkIdSwitchPortSchedulesPortSchedule1 `json:"portSchedule,omitempty"`
}

// NewInlineObject127 instantiates a new InlineObject127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject127(name string) *InlineObject127 {
	this := InlineObject127{}
	this.Name = name
	return &this
}

// NewInlineObject127WithDefaults instantiates a new InlineObject127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject127WithDefaults() *InlineObject127 {
	this := InlineObject127{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject127) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject127) SetName(v string) {
	o.Name = v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *InlineObject127) GetPortSchedule() NetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	if o == nil || isNil(o.PortSchedule) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortSchedule1
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject127) GetPortScheduleOk() (*NetworksNetworkIdSwitchPortSchedulesPortSchedule1, bool) {
	if o == nil || isNil(o.PortSchedule) {
    return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *InlineObject127) HasPortSchedule() bool {
	if o != nil && !isNil(o.PortSchedule) {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortSchedule1 and assigns it to the PortSchedule field.
func (o *InlineObject127) SetPortSchedule(v NetworksNetworkIdSwitchPortSchedulesPortSchedule1) {
	o.PortSchedule = &v
}

func (o InlineObject127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PortSchedule) {
		toSerialize["portSchedule"] = o.PortSchedule
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


