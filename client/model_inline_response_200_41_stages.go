/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20041Stages struct for InlineResponse20041Stages
type InlineResponse20041Stages struct {
	Group *InlineResponse20041Group `json:"group,omitempty"`
	Milestones *InlineResponse20041Milestones `json:"milestones,omitempty"`
	// Current upgrade status of the group
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse20041Stages instantiates a new InlineResponse20041Stages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20041Stages() *InlineResponse20041Stages {
	this := InlineResponse20041Stages{}
	return &this
}

// NewInlineResponse20041StagesWithDefaults instantiates a new InlineResponse20041Stages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20041StagesWithDefaults() *InlineResponse20041Stages {
	this := InlineResponse20041Stages{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20041Stages) GetGroup() InlineResponse20041Group {
	if o == nil || isNil(o.Group) {
		var ret InlineResponse20041Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041Stages) GetGroupOk() (*InlineResponse20041Group, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20041Stages) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given InlineResponse20041Group and assigns it to the Group field.
func (o *InlineResponse20041Stages) SetGroup(v InlineResponse20041Group) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *InlineResponse20041Stages) GetMilestones() InlineResponse20041Milestones {
	if o == nil || isNil(o.Milestones) {
		var ret InlineResponse20041Milestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041Stages) GetMilestonesOk() (*InlineResponse20041Milestones, bool) {
	if o == nil || isNil(o.Milestones) {
    return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *InlineResponse20041Stages) HasMilestones() bool {
	if o != nil && !isNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given InlineResponse20041Milestones and assigns it to the Milestones field.
func (o *InlineResponse20041Stages) SetMilestones(v InlineResponse20041Milestones) {
	o.Milestones = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20041Stages) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041Stages) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20041Stages) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20041Stages) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse20041Stages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Milestones) {
		toSerialize["milestones"] = o.Milestones
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20041Stages struct {
	value *InlineResponse20041Stages
	isSet bool
}

func (v NullableInlineResponse20041Stages) Get() *InlineResponse20041Stages {
	return v.value
}

func (v *NullableInlineResponse20041Stages) Set(val *InlineResponse20041Stages) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20041Stages) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20041Stages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20041Stages(val *InlineResponse20041Stages) *NullableInlineResponse20041Stages {
	return &NullableInlineResponse20041Stages{value: val, isSet: true}
}

func (v NullableInlineResponse20041Stages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20041Stages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

