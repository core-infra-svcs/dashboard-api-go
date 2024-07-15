/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse207Swaps struct for InlineResponse207Swaps
type InlineResponse207Swaps struct {
	// Swap Request ID
	Id string `json:"id"`
	Devices InlineResponse207Devices `json:"devices"`
	// The current status of the swap job.
	Status string `json:"status"`
	// An action to perform on the devices.old object after swap is complete.
	AfterAction string `json:"afterAction"`
	// An iso8601 timestamp for the creation of the swap request.
	CreatedAt string `json:"createdAt"`
	// An iso8601 timestamp for when the swap completed or failed.
	CompletedAt *string `json:"completedAt,omitempty"`
	// A list of error messages for why a swap failed.
	Errors []string `json:"errors,omitempty"`
}

// NewInlineResponse207Swaps instantiates a new InlineResponse207Swaps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207Swaps(id string, devices InlineResponse207Devices, status string, afterAction string, createdAt string) *InlineResponse207Swaps {
	this := InlineResponse207Swaps{}
	this.Id = id
	this.Devices = devices
	this.Status = status
	this.AfterAction = afterAction
	this.CreatedAt = createdAt
	return &this
}

// NewInlineResponse207SwapsWithDefaults instantiates a new InlineResponse207Swaps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207SwapsWithDefaults() *InlineResponse207Swaps {
	this := InlineResponse207Swaps{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse207Swaps) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse207Swaps) SetId(v string) {
	o.Id = v
}

// GetDevices returns the Devices field value
func (o *InlineResponse207Swaps) GetDevices() InlineResponse207Devices {
	if o == nil {
		var ret InlineResponse207Devices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetDevicesOk() (*InlineResponse207Devices, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *InlineResponse207Swaps) SetDevices(v InlineResponse207Devices) {
	o.Devices = v
}

// GetStatus returns the Status field value
func (o *InlineResponse207Swaps) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse207Swaps) SetStatus(v string) {
	o.Status = v
}

// GetAfterAction returns the AfterAction field value
func (o *InlineResponse207Swaps) GetAfterAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfterAction
}

// GetAfterActionOk returns a tuple with the AfterAction field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetAfterActionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfterAction, true
}

// SetAfterAction sets field value
func (o *InlineResponse207Swaps) SetAfterAction(v string) {
	o.AfterAction = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InlineResponse207Swaps) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetCreatedAtOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InlineResponse207Swaps) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse207Swaps) GetCompletedAt() string {
	if o == nil || isNil(o.CompletedAt) {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetCompletedAtOk() (*string, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse207Swaps) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *InlineResponse207Swaps) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse207Swaps) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207Swaps) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse207Swaps) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse207Swaps) SetErrors(v []string) {
	o.Errors = v
}

func (o InlineResponse207Swaps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["devices"] = o.Devices
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["afterAction"] = o.AfterAction
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse207Swaps struct {
	value *InlineResponse207Swaps
	isSet bool
}

func (v NullableInlineResponse207Swaps) Get() *InlineResponse207Swaps {
	return v.value
}

func (v *NullableInlineResponse207Swaps) Set(val *InlineResponse207Swaps) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse207Swaps) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse207Swaps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse207Swaps(val *InlineResponse207Swaps) *NullableInlineResponse207Swaps {
	return &NullableInlineResponse207Swaps{value: val, isSet: true}
}

func (v NullableInlineResponse207Swaps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse207Swaps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


