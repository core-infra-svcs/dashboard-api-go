/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20033Stages struct for InlineResponse20033Stages
type InlineResponse20033Stages struct {
	Group *InlineResponse20033Group `json:"group,omitempty"`
	Milestones *InlineResponse20033Milestones `json:"milestones,omitempty"`
	// Current upgrade status of the group
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse20033Stages instantiates a new InlineResponse20033Stages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20033Stages() *InlineResponse20033Stages {
	this := InlineResponse20033Stages{}
	return &this
}

// NewInlineResponse20033StagesWithDefaults instantiates a new InlineResponse20033Stages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20033StagesWithDefaults() *InlineResponse20033Stages {
	this := InlineResponse20033Stages{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20033Stages) GetGroup() InlineResponse20033Group {
	if o == nil || isNil(o.Group) {
		var ret InlineResponse20033Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033Stages) GetGroupOk() (*InlineResponse20033Group, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20033Stages) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given InlineResponse20033Group and assigns it to the Group field.
func (o *InlineResponse20033Stages) SetGroup(v InlineResponse20033Group) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *InlineResponse20033Stages) GetMilestones() InlineResponse20033Milestones {
	if o == nil || isNil(o.Milestones) {
		var ret InlineResponse20033Milestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033Stages) GetMilestonesOk() (*InlineResponse20033Milestones, bool) {
	if o == nil || isNil(o.Milestones) {
    return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *InlineResponse20033Stages) HasMilestones() bool {
	if o != nil && !isNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given InlineResponse20033Milestones and assigns it to the Milestones field.
func (o *InlineResponse20033Stages) SetMilestones(v InlineResponse20033Milestones) {
	o.Milestones = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20033Stages) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20033Stages) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20033Stages) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20033Stages) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse20033Stages) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20033Stages struct {
	value *InlineResponse20033Stages
	isSet bool
}

func (v NullableInlineResponse20033Stages) Get() *InlineResponse20033Stages {
	return v.value
}

func (v *NullableInlineResponse20033Stages) Set(val *InlineResponse20033Stages) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20033Stages) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20033Stages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20033Stages(val *InlineResponse20033Stages) *NullableInlineResponse20033Stages {
	return &NullableInlineResponse20033Stages{value: val, isSet: true}
}

func (v NullableInlineResponse20033Stages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20033Stages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


