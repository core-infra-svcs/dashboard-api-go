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

// InlineResponse20053Stages struct for InlineResponse20053Stages
type InlineResponse20053Stages struct {
	Group *InlineResponse20053Group `json:"group,omitempty"`
	Milestones *InlineResponse20053Milestones `json:"milestones,omitempty"`
	// Current upgrade status of the group
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse20053Stages instantiates a new InlineResponse20053Stages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20053Stages() *InlineResponse20053Stages {
	this := InlineResponse20053Stages{}
	return &this
}

// NewInlineResponse20053StagesWithDefaults instantiates a new InlineResponse20053Stages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20053StagesWithDefaults() *InlineResponse20053Stages {
	this := InlineResponse20053Stages{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20053Stages) GetGroup() InlineResponse20053Group {
	if o == nil || isNil(o.Group) {
		var ret InlineResponse20053Group
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Stages) GetGroupOk() (*InlineResponse20053Group, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20053Stages) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given InlineResponse20053Group and assigns it to the Group field.
func (o *InlineResponse20053Stages) SetGroup(v InlineResponse20053Group) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *InlineResponse20053Stages) GetMilestones() InlineResponse20053Milestones {
	if o == nil || isNil(o.Milestones) {
		var ret InlineResponse20053Milestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Stages) GetMilestonesOk() (*InlineResponse20053Milestones, bool) {
	if o == nil || isNil(o.Milestones) {
    return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *InlineResponse20053Stages) HasMilestones() bool {
	if o != nil && !isNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given InlineResponse20053Milestones and assigns it to the Milestones field.
func (o *InlineResponse20053Stages) SetMilestones(v InlineResponse20053Milestones) {
	o.Milestones = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20053Stages) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20053Stages) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20053Stages) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20053Stages) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse20053Stages) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20053Stages struct {
	value *InlineResponse20053Stages
	isSet bool
}

func (v NullableInlineResponse20053Stages) Get() *InlineResponse20053Stages {
	return v.value
}

func (v *NullableInlineResponse20053Stages) Set(val *InlineResponse20053Stages) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20053Stages) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20053Stages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20053Stages(val *InlineResponse20053Stages) *NullableInlineResponse20053Stages {
	return &NullableInlineResponse20053Stages{value: val, isSet: true}
}

func (v NullableInlineResponse20053Stages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20053Stages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


