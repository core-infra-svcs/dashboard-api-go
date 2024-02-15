/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject136 struct for InlineObject136
type InlineObject136 struct {
	// The name for your port schedule.
	Name *string `json:"name,omitempty"`
	PortSchedule *NetworksNetworkIdSwitchPortSchedulesPortSchedule1 `json:"portSchedule,omitempty"`
}

// NewInlineObject136 instantiates a new InlineObject136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject136() *InlineObject136 {
	this := InlineObject136{}
	return &this
}

// NewInlineObject136WithDefaults instantiates a new InlineObject136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject136WithDefaults() *InlineObject136 {
	this := InlineObject136{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject136) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject136) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject136) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject136) SetName(v string) {
	o.Name = &v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *InlineObject136) GetPortSchedule() NetworksNetworkIdSwitchPortSchedulesPortSchedule1 {
	if o == nil || isNil(o.PortSchedule) {
		var ret NetworksNetworkIdSwitchPortSchedulesPortSchedule1
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject136) GetPortScheduleOk() (*NetworksNetworkIdSwitchPortSchedulesPortSchedule1, bool) {
	if o == nil || isNil(o.PortSchedule) {
    return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *InlineObject136) HasPortSchedule() bool {
	if o != nil && !isNil(o.PortSchedule) {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given NetworksNetworkIdSwitchPortSchedulesPortSchedule1 and assigns it to the PortSchedule field.
func (o *InlineObject136) SetPortSchedule(v NetworksNetworkIdSwitchPortSchedulesPortSchedule1) {
	o.PortSchedule = &v
}

func (o InlineObject136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PortSchedule) {
		toSerialize["portSchedule"] = o.PortSchedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject136 struct {
	value *InlineObject136
	isSet bool
}

func (v NullableInlineObject136) Get() *InlineObject136 {
	return v.value
}

func (v *NullableInlineObject136) Set(val *InlineObject136) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject136) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject136(val *InlineObject136) *NullableInlineObject136 {
	return &NullableInlineObject136{value: val, isSet: true}
}

func (v NullableInlineObject136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


