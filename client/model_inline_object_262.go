/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject262 struct for InlineObject262
type InlineObject262 struct {
	// device details
	Devices []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices `json:"devices"`
	// Name of the packet capture file
	Name *string `json:"name,omitempty"`
	// Reason for capture
	Notes *string `json:"notes,omitempty"`
	// Duration of the capture in seconds
	Duration *int32 `json:"duration,omitempty"`
	// Filter expression for the capture
	FilterExpression *string `json:"filterExpression,omitempty"`
	// Enable or disable the schedule
	Enabled *bool `json:"enabled,omitempty"`
	Schedule *OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule `json:"schedule,omitempty"`
}

// NewInlineObject262 instantiates a new InlineObject262 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject262(devices []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices) *InlineObject262 {
	this := InlineObject262{}
	this.Devices = devices
	return &this
}

// NewInlineObject262WithDefaults instantiates a new InlineObject262 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject262WithDefaults() *InlineObject262 {
	this := InlineObject262{}
	return &this
}

// GetDevices returns the Devices field value
func (o *InlineObject262) GetDevices() []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices {
	if o == nil {
		var ret []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetDevicesOk() ([]OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices, bool) {
	if o == nil {
    return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *InlineObject262) SetDevices(v []OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices) {
	o.Devices = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject262) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject262) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject262) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject262) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject262) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject262) SetNotes(v string) {
	o.Notes = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineObject262) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineObject262) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineObject262) SetDuration(v int32) {
	o.Duration = &v
}

// GetFilterExpression returns the FilterExpression field value if set, zero value otherwise.
func (o *InlineObject262) GetFilterExpression() string {
	if o == nil || isNil(o.FilterExpression) {
		var ret string
		return ret
	}
	return *o.FilterExpression
}

// GetFilterExpressionOk returns a tuple with the FilterExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetFilterExpressionOk() (*string, bool) {
	if o == nil || isNil(o.FilterExpression) {
    return nil, false
	}
	return o.FilterExpression, true
}

// HasFilterExpression returns a boolean if a field has been set.
func (o *InlineObject262) HasFilterExpression() bool {
	if o != nil && !isNil(o.FilterExpression) {
		return true
	}

	return false
}

// SetFilterExpression gets a reference to the given string and assigns it to the FilterExpression field.
func (o *InlineObject262) SetFilterExpression(v string) {
	o.FilterExpression = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject262) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject262) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject262) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *InlineObject262) GetSchedule() OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule {
	if o == nil || isNil(o.Schedule) {
		var ret OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject262) GetScheduleOk() (*OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *InlineObject262) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule and assigns it to the Schedule field.
func (o *InlineObject262) SetSchedule(v OrganizationsOrganizationIdDevicesPacketCaptureSchedulesSchedule) {
	o.Schedule = &v
}

func (o InlineObject262) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.FilterExpression) {
		toSerialize["filterExpression"] = o.FilterExpression
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject262 struct {
	value *InlineObject262
	isSet bool
}

func (v NullableInlineObject262) Get() *InlineObject262 {
	return v.value
}

func (v *NullableInlineObject262) Set(val *InlineObject262) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject262) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject262) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject262(val *InlineObject262) *NullableInlineObject262 {
	return &NullableInlineObject262{value: val, isSet: true}
}

func (v NullableInlineObject262) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject262) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


