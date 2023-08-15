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

// InlineObject172 struct for InlineObject172
type InlineObject172 struct {
	// Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false.
	Synchronous *bool `json:"synchronous,omitempty"`
	// A set of changes to make as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Actions []OrganizationsOrganizationIdActionBatchesActions `json:"actions"`
}

// NewInlineObject172 instantiates a new InlineObject172 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject172(actions []OrganizationsOrganizationIdActionBatchesActions) *InlineObject172 {
	this := InlineObject172{}
	this.Actions = actions
	return &this
}

// NewInlineObject172WithDefaults instantiates a new InlineObject172 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject172WithDefaults() *InlineObject172 {
	this := InlineObject172{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineObject172) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineObject172) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineObject172) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineObject172) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineObject172) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineObject172) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetActions returns the Actions field value
func (o *InlineObject172) GetActions() []OrganizationsOrganizationIdActionBatchesActions {
	if o == nil {
		var ret []OrganizationsOrganizationIdActionBatchesActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *InlineObject172) GetActionsOk() ([]OrganizationsOrganizationIdActionBatchesActions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *InlineObject172) SetActions(v []OrganizationsOrganizationIdActionBatchesActions) {
	o.Actions = v
}

func (o InlineObject172) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject172 struct {
	value *InlineObject172
	isSet bool
}

func (v NullableInlineObject172) Get() *InlineObject172 {
	return v.value
}

func (v *NullableInlineObject172) Set(val *InlineObject172) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject172) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject172) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject172(val *InlineObject172) *NullableInlineObject172 {
	return &NullableInlineObject172{value: val, isSet: true}
}

func (v NullableInlineObject172) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject172) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


