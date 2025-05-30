/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200270Items struct for InlineResponse200270Items
type InlineResponse200270Items struct {
	// Id of scheduled packet capture
	ScheduleId *string `json:"scheduleId,omitempty"`
	// Devices associated to the schedule
	Devices []InlineResponse200270Devices `json:"devices,omitempty"`
	// Name of scheduled packet capture
	Name *string `json:"name,omitempty"`
	Admin *InlineResponse200270Admin `json:"admin,omitempty"`
	// Reason of scheduled packet capture
	Notes *string `json:"notes,omitempty"`
	// Duration of scheduled packet capture
	Duration *int32 `json:"duration,omitempty"`
	// Filter expression for the packet capture
	FilterExpression *string `json:"filterExpression,omitempty"`
	// Time of creation of scheduled packet capture
	CreatedAt *string `json:"createdAt,omitempty"`
	// Time of updation of scheduled packet capture
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The number of pcaps captured/performed
	CaptureCount *int32 `json:"captureCount,omitempty"`
	// Pcap log id of the latest pcap from this schedule
	LastCaptureId *string `json:"lastCaptureId,omitempty"`
	// Whether the packet capture schedule is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Priority of the packet capture
	Priority *int32 `json:"priority,omitempty"`
	Schedule *InlineResponse200270Schedule `json:"schedule,omitempty"`
	// Any warnings pertaining to the schedule and it's nodes
	Warnings []string `json:"warnings,omitempty"`
}

// NewInlineResponse200270Items instantiates a new InlineResponse200270Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200270Items() *InlineResponse200270Items {
	this := InlineResponse200270Items{}
	return &this
}

// NewInlineResponse200270ItemsWithDefaults instantiates a new InlineResponse200270Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200270ItemsWithDefaults() *InlineResponse200270Items {
	this := InlineResponse200270Items{}
	return &this
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetScheduleId() string {
	if o == nil || isNil(o.ScheduleId) {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetScheduleIdOk() (*string, bool) {
	if o == nil || isNil(o.ScheduleId) {
    return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasScheduleId() bool {
	if o != nil && !isNil(o.ScheduleId) {
		return true
	}

	return false
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *InlineResponse200270Items) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetDevices() []InlineResponse200270Devices {
	if o == nil || isNil(o.Devices) {
		var ret []InlineResponse200270Devices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetDevicesOk() ([]InlineResponse200270Devices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []InlineResponse200270Devices and assigns it to the Devices field.
func (o *InlineResponse200270Items) SetDevices(v []InlineResponse200270Devices) {
	o.Devices = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200270Items) SetName(v string) {
	o.Name = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetAdmin() InlineResponse200270Admin {
	if o == nil || isNil(o.Admin) {
		var ret InlineResponse200270Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetAdminOk() (*InlineResponse200270Admin, bool) {
	if o == nil || isNil(o.Admin) {
    return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasAdmin() bool {
	if o != nil && !isNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given InlineResponse200270Admin and assigns it to the Admin field.
func (o *InlineResponse200270Items) SetAdmin(v InlineResponse200270Admin) {
	o.Admin = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse200270Items) SetNotes(v string) {
	o.Notes = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse200270Items) SetDuration(v int32) {
	o.Duration = &v
}

// GetFilterExpression returns the FilterExpression field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetFilterExpression() string {
	if o == nil || isNil(o.FilterExpression) {
		var ret string
		return ret
	}
	return *o.FilterExpression
}

// GetFilterExpressionOk returns a tuple with the FilterExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetFilterExpressionOk() (*string, bool) {
	if o == nil || isNil(o.FilterExpression) {
    return nil, false
	}
	return o.FilterExpression, true
}

// HasFilterExpression returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasFilterExpression() bool {
	if o != nil && !isNil(o.FilterExpression) {
		return true
	}

	return false
}

// SetFilterExpression gets a reference to the given string and assigns it to the FilterExpression field.
func (o *InlineResponse200270Items) SetFilterExpression(v string) {
	o.FilterExpression = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse200270Items) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *InlineResponse200270Items) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCaptureCount returns the CaptureCount field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetCaptureCount() int32 {
	if o == nil || isNil(o.CaptureCount) {
		var ret int32
		return ret
	}
	return *o.CaptureCount
}

// GetCaptureCountOk returns a tuple with the CaptureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetCaptureCountOk() (*int32, bool) {
	if o == nil || isNil(o.CaptureCount) {
    return nil, false
	}
	return o.CaptureCount, true
}

// HasCaptureCount returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasCaptureCount() bool {
	if o != nil && !isNil(o.CaptureCount) {
		return true
	}

	return false
}

// SetCaptureCount gets a reference to the given int32 and assigns it to the CaptureCount field.
func (o *InlineResponse200270Items) SetCaptureCount(v int32) {
	o.CaptureCount = &v
}

// GetLastCaptureId returns the LastCaptureId field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetLastCaptureId() string {
	if o == nil || isNil(o.LastCaptureId) {
		var ret string
		return ret
	}
	return *o.LastCaptureId
}

// GetLastCaptureIdOk returns a tuple with the LastCaptureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetLastCaptureIdOk() (*string, bool) {
	if o == nil || isNil(o.LastCaptureId) {
    return nil, false
	}
	return o.LastCaptureId, true
}

// HasLastCaptureId returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasLastCaptureId() bool {
	if o != nil && !isNil(o.LastCaptureId) {
		return true
	}

	return false
}

// SetLastCaptureId gets a reference to the given string and assigns it to the LastCaptureId field.
func (o *InlineResponse200270Items) SetLastCaptureId(v string) {
	o.LastCaptureId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200270Items) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InlineResponse200270Items) SetPriority(v int32) {
	o.Priority = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetSchedule() InlineResponse200270Schedule {
	if o == nil || isNil(o.Schedule) {
		var ret InlineResponse200270Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetScheduleOk() (*InlineResponse200270Schedule, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given InlineResponse200270Schedule and assigns it to the Schedule field.
func (o *InlineResponse200270Items) SetSchedule(v InlineResponse200270Schedule) {
	o.Schedule = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *InlineResponse200270Items) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200270Items) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
    return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *InlineResponse200270Items) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *InlineResponse200270Items) SetWarnings(v []string) {
	o.Warnings = v
}

func (o InlineResponse200270Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScheduleId) {
		toSerialize["scheduleId"] = o.ScheduleId
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Admin) {
		toSerialize["admin"] = o.Admin
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CaptureCount) {
		toSerialize["captureCount"] = o.CaptureCount
	}
	if !isNil(o.LastCaptureId) {
		toSerialize["lastCaptureId"] = o.LastCaptureId
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200270Items struct {
	value *InlineResponse200270Items
	isSet bool
}

func (v NullableInlineResponse200270Items) Get() *InlineResponse200270Items {
	return v.value
}

func (v *NullableInlineResponse200270Items) Set(val *InlineResponse200270Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200270Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200270Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200270Items(val *InlineResponse200270Items) *NullableInlineResponse200270Items {
	return &NullableInlineResponse200270Items{value: val, isSet: true}
}

func (v NullableInlineResponse200270Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200270Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


