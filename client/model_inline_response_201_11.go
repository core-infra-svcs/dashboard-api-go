/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20111 struct for InlineResponse20111
type InlineResponse20111 struct {
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
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse20111 instantiates a new InlineResponse20111 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20111(actions []OrganizationsOrganizationIdActionBatchesActions) *InlineResponse20111 {
	this := InlineResponse20111{}
	this.Actions = actions
	return &this
}

// NewInlineResponse20111WithDefaults instantiates a new InlineResponse20111 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20111WithDefaults() *InlineResponse20111 {
	this := InlineResponse20111{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20111) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20111) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20111) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse20111) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse20111) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse20111) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineResponse20111) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineResponse20111) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineResponse20111) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineResponse20111) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineResponse20111) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineResponse20111) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20111) GetStatus() OrganizationsOrganizationIdActionBatchesStatus {
	if o == nil || isNil(o.Status) {
		var ret OrganizationsOrganizationIdActionBatchesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetStatusOk() (*OrganizationsOrganizationIdActionBatchesStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20111) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationsOrganizationIdActionBatchesStatus and assigns it to the Status field.
func (o *InlineResponse20111) SetStatus(v OrganizationsOrganizationIdActionBatchesStatus) {
	o.Status = &v
}

// GetActions returns the Actions field value
func (o *InlineResponse20111) GetActions() []OrganizationsOrganizationIdActionBatchesActions {
	if o == nil {
		var ret []OrganizationsOrganizationIdActionBatchesActions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetActionsOk() ([]OrganizationsOrganizationIdActionBatchesActions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *InlineResponse20111) SetActions(v []OrganizationsOrganizationIdActionBatchesActions) {
	o.Actions = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse20111) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20111) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse20111) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse20111) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse20111) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20111 struct {
	value *InlineResponse20111
	isSet bool
}

func (v NullableInlineResponse20111) Get() *InlineResponse20111 {
	return v.value
}

func (v *NullableInlineResponse20111) Set(val *InlineResponse20111) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20111) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20111) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20111(val *InlineResponse20111) *NullableInlineResponse20111 {
	return &NullableInlineResponse20111{value: val, isSet: true}
}

func (v NullableInlineResponse20111) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20111) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


