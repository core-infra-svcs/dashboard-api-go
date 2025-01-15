/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200204 struct for InlineResponse200204
type InlineResponse200204 struct {
	// ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	Id *string `json:"id,omitempty"`
	// ID of the organization this action batch belongs to
	OrganizationId *string `json:"organizationId,omitempty"`
	// Flag describing whether the action should be previewed before executing or not
	Confirmed *bool `json:"confirmed,omitempty"`
	// Flag describing whether actions should run synchronously or asynchronously
	Synchronous *bool `json:"synchronous,omitempty"`
	Status *OrganizationsOrganizationIdActionBatchesStatus `json:"status,omitempty"`
	// A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Actions []OrganizationsOrganizationIdActionBatchesActions `json:"actions"`
}

// NewInlineResponse200204 instantiates a new InlineResponse200204 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200204(actions []OrganizationsOrganizationIdActionBatchesActions) *InlineResponse200204 {
	this := InlineResponse200204{}
	this.Actions = actions
	return &this
}

// NewInlineResponse200204WithDefaults instantiates a new InlineResponse200204 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200204WithDefaults() *InlineResponse200204 {
	this := InlineResponse200204{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200204) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200204) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200204) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200204) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200204) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse200204) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineResponse200204) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineResponse200204) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineResponse200204) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineResponse200204) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineResponse200204) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineResponse200204) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200204) GetStatus() OrganizationsOrganizationIdActionBatchesStatus {
	if o == nil || isNil(o.Status) {
		var ret OrganizationsOrganizationIdActionBatchesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetStatusOk() (*OrganizationsOrganizationIdActionBatchesStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200204) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationsOrganizationIdActionBatchesStatus and assigns it to the Status field.
func (o *InlineResponse200204) SetStatus(v OrganizationsOrganizationIdActionBatchesStatus) {
	o.Status = &v
}

// GetActions returns the Actions field value
func (o *InlineResponse200204) GetActions() []OrganizationsOrganizationIdActionBatchesActions {
	if o == nil {
		var ret []OrganizationsOrganizationIdActionBatchesActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200204) GetActionsOk() ([]OrganizationsOrganizationIdActionBatchesActions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *InlineResponse200204) SetActions(v []OrganizationsOrganizationIdActionBatchesActions) {
	o.Actions = v
}

func (o InlineResponse200204) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200204 struct {
	value *InlineResponse200204
	isSet bool
}

func (v NullableInlineResponse200204) Get() *InlineResponse200204 {
	return v.value
}

func (v *NullableInlineResponse200204) Set(val *InlineResponse200204) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200204) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200204) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200204(val *InlineResponse200204) *NullableInlineResponse200204 {
	return &NullableInlineResponse200204{value: val, isSet: true}
}

func (v NullableInlineResponse200204) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200204) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


