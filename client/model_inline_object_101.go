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

// InlineObject101 struct for InlineObject101
type InlineObject101 struct {
	// The name for your port schedule.
	Name *string `json:"name,omitempty"`
	PortSchedule *NetworksNetworkIdSwitchPortSchedulesPortSchedule `json:"portSchedule,omitempty"`
}

// NewInlineObject101 instantiates a new InlineObject101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject101() *InlineObject101 {
	this := InlineObject101{}
	return &this
}

// NewInlineObject101WithDefaults instantiates a new InlineObject101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject101WithDefaults() *InlineObject101 {
	this := InlineObject101{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject101) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject101) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject101) SetName(v string) {
	o.Name = &v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *InlineObject101) GetPortSchedule() NetworksNetworkIdSwitchPortSchedulesPortSchedule {
	if o == nil || o.PortSchedule == nil {
		var ret NetworksNetworkIdSwitchPortSchedulesPortSchedule
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetPortScheduleOk() (*NetworksNetworkIdSwitchPortSchedulesPortSchedule, bool) {
	if o == nil || o.PortSchedule == nil {
		return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *InlineObject101) HasPortSchedule() bool {
	if o != nil && o.PortSchedule != nil {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortSchedule and assigns it to the PortSchedule field.
func (o *InlineObject101) SetPortSchedule(v NetworksNetworkIdSwitchPortSchedulesPortSchedule) {
	o.PortSchedule = &v
}

func (o InlineObject101) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PortSchedule != nil {
		toSerialize["portSchedule"] = o.PortSchedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject101 struct {
	value *InlineObject101
	isSet bool
}

func (v NullableInlineObject101) Get() *InlineObject101 {
	return v.value
}

func (v *NullableInlineObject101) Set(val *InlineObject101) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject101) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject101) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject101(val *InlineObject101) *NullableInlineObject101 {
	return &NullableInlineObject101{value: val, isSet: true}
}

func (v NullableInlineObject101) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject101) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

