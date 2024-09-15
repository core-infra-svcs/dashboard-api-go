/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20029 struct for InlineResponse20029
type InlineResponse20029 struct {
	// ID to check the status of the command request
	CommandId *string `json:"commandId,omitempty"`
	// Time when the command was triggered
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Time when the command was completed
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	CreatedBy *DevicesSerialSensorCommandsCreatedBy `json:"createdBy,omitempty"`
	// Operation run on the sensor
	Operation *string `json:"operation,omitempty"`
	// Status of the command request
	Status *string `json:"status,omitempty"`
	// Array of errors if failed
	Errors []string `json:"errors,omitempty"`
}

// NewInlineResponse20029 instantiates a new InlineResponse20029 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029WithDefaults() *InlineResponse20029 {
	this := InlineResponse20029{}
	return &this
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *InlineResponse20029) GetCommandId() string {
	if o == nil || isNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetCommandIdOk() (*string, bool) {
	if o == nil || isNil(o.CommandId) {
    return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *InlineResponse20029) HasCommandId() bool {
	if o != nil && !isNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *InlineResponse20029) SetCommandId(v string) {
	o.CommandId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse20029) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse20029) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse20029) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse20029) GetCompletedAt() time.Time {
	if o == nil || isNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse20029) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *InlineResponse20029) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *InlineResponse20029) GetCreatedBy() DevicesSerialSensorCommandsCreatedBy {
	if o == nil || isNil(o.CreatedBy) {
		var ret DevicesSerialSensorCommandsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetCreatedByOk() (*DevicesSerialSensorCommandsCreatedBy, bool) {
	if o == nil || isNil(o.CreatedBy) {
    return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *InlineResponse20029) HasCreatedBy() bool {
	if o != nil && !isNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given DevicesSerialSensorCommandsCreatedBy and assigns it to the CreatedBy field.
func (o *InlineResponse20029) SetCreatedBy(v DevicesSerialSensorCommandsCreatedBy) {
	o.CreatedBy = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *InlineResponse20029) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
    return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *InlineResponse20029) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *InlineResponse20029) SetOperation(v string) {
	o.Operation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20029) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20029) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20029) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse20029) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse20029) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse20029) SetErrors(v []string) {
	o.Errors = v
}

func (o InlineResponse20029) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CommandId) {
		toSerialize["commandId"] = o.CommandId
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029 struct {
	value *InlineResponse20029
	isSet bool
}

func (v NullableInlineResponse20029) Get() *InlineResponse20029 {
	return v.value
}

func (v *NullableInlineResponse20029) Set(val *InlineResponse20029) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029(val *InlineResponse20029) *NullableInlineResponse20029 {
	return &NullableInlineResponse20029{value: val, isSet: true}
}

func (v NullableInlineResponse20029) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


