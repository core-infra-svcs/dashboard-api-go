/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20098Milestones The Staged Upgrade Milestones for the stage
type InlineResponse20098Milestones struct {
	// Scheduled start time for the group
	ScheduledFor *time.Time `json:"scheduledFor,omitempty"`
	// Start time for the group
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Finish time for the group
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Time that the group was canceled
	CanceledAt *time.Time `json:"canceledAt,omitempty"`
}

// NewInlineResponse20098Milestones instantiates a new InlineResponse20098Milestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20098Milestones() *InlineResponse20098Milestones {
	this := InlineResponse20098Milestones{}
	return &this
}

// NewInlineResponse20098MilestonesWithDefaults instantiates a new InlineResponse20098Milestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20098MilestonesWithDefaults() *InlineResponse20098Milestones {
	this := InlineResponse20098Milestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *InlineResponse20098Milestones) GetScheduledFor() time.Time {
	if o == nil || isNil(o.ScheduledFor) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Milestones) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledFor) {
    return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *InlineResponse20098Milestones) HasScheduledFor() bool {
	if o != nil && !isNil(o.ScheduledFor) {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *InlineResponse20098Milestones) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *InlineResponse20098Milestones) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Milestones) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *InlineResponse20098Milestones) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *InlineResponse20098Milestones) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse20098Milestones) GetCompletedAt() time.Time {
	if o == nil || isNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Milestones) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse20098Milestones) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *InlineResponse20098Milestones) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *InlineResponse20098Milestones) GetCanceledAt() time.Time {
	if o == nil || isNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20098Milestones) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CanceledAt) {
    return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *InlineResponse20098Milestones) HasCanceledAt() bool {
	if o != nil && !isNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *InlineResponse20098Milestones) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

func (o InlineResponse20098Milestones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScheduledFor) {
		toSerialize["scheduledFor"] = o.ScheduledFor
	}
	if !isNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.CanceledAt) {
		toSerialize["canceledAt"] = o.CanceledAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20098Milestones struct {
	value *InlineResponse20098Milestones
	isSet bool
}

func (v NullableInlineResponse20098Milestones) Get() *InlineResponse20098Milestones {
	return v.value
}

func (v *NullableInlineResponse20098Milestones) Set(val *InlineResponse20098Milestones) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20098Milestones) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20098Milestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20098Milestones(val *InlineResponse20098Milestones) *NullableInlineResponse20098Milestones {
	return &NullableInlineResponse20098Milestones{value: val, isSet: true}
}

func (v NullableInlineResponse20098Milestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20098Milestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


