/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	// The name of the new profile. Must be unique.
	Name *string `json:"name,omitempty"`
	// Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	MotionBasedRetentionEnabled *bool `json:"motionBasedRetentionEnabled,omitempty"`
	// Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.
	RestrictedBandwidthModeEnabled *bool `json:"restrictedBandwidthModeEnabled,omitempty"`
	// Whether or not to record audio. Can be either true or false. Defaults to false.
	AudioRecordingEnabled *bool `json:"audioRecordingEnabled,omitempty"`
	// Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.
	CloudArchiveEnabled *bool `json:"cloudArchiveEnabled,omitempty"`
	// The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	MotionDetectorVersion *int32 `json:"motionDetectorVersion,omitempty"`
	// Schedule for which this camera will record video, or 'null' to always record.
	ScheduleId *string `json:"scheduleId,omitempty"`
	// The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days.
	MaxRetentionDays *int32 `json:"maxRetentionDays,omitempty"`
	VideoSettings *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings `json:"videoSettings,omitempty"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject63) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject63) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject63) SetName(v string) {
	o.Name = &v
}

// GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field value if set, zero value otherwise.
func (o *InlineObject63) GetMotionBasedRetentionEnabled() bool {
	if o == nil || isNil(o.MotionBasedRetentionEnabled) {
		var ret bool
		return ret
	}
	return *o.MotionBasedRetentionEnabled
}

// GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMotionBasedRetentionEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MotionBasedRetentionEnabled) {
    return nil, false
	}
	return o.MotionBasedRetentionEnabled, true
}

// HasMotionBasedRetentionEnabled returns a boolean if a field has been set.
func (o *InlineObject63) HasMotionBasedRetentionEnabled() bool {
	if o != nil && !isNil(o.MotionBasedRetentionEnabled) {
		return true
	}

	return false
}

// SetMotionBasedRetentionEnabled gets a reference to the given bool and assigns it to the MotionBasedRetentionEnabled field.
func (o *InlineObject63) SetMotionBasedRetentionEnabled(v bool) {
	o.MotionBasedRetentionEnabled = &v
}

// GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field value if set, zero value otherwise.
func (o *InlineObject63) GetRestrictedBandwidthModeEnabled() bool {
	if o == nil || isNil(o.RestrictedBandwidthModeEnabled) {
		var ret bool
		return ret
	}
	return *o.RestrictedBandwidthModeEnabled
}

// GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetRestrictedBandwidthModeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RestrictedBandwidthModeEnabled) {
    return nil, false
	}
	return o.RestrictedBandwidthModeEnabled, true
}

// HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.
func (o *InlineObject63) HasRestrictedBandwidthModeEnabled() bool {
	if o != nil && !isNil(o.RestrictedBandwidthModeEnabled) {
		return true
	}

	return false
}

// SetRestrictedBandwidthModeEnabled gets a reference to the given bool and assigns it to the RestrictedBandwidthModeEnabled field.
func (o *InlineObject63) SetRestrictedBandwidthModeEnabled(v bool) {
	o.RestrictedBandwidthModeEnabled = &v
}

// GetAudioRecordingEnabled returns the AudioRecordingEnabled field value if set, zero value otherwise.
func (o *InlineObject63) GetAudioRecordingEnabled() bool {
	if o == nil || isNil(o.AudioRecordingEnabled) {
		var ret bool
		return ret
	}
	return *o.AudioRecordingEnabled
}

// GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetAudioRecordingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AudioRecordingEnabled) {
    return nil, false
	}
	return o.AudioRecordingEnabled, true
}

// HasAudioRecordingEnabled returns a boolean if a field has been set.
func (o *InlineObject63) HasAudioRecordingEnabled() bool {
	if o != nil && !isNil(o.AudioRecordingEnabled) {
		return true
	}

	return false
}

// SetAudioRecordingEnabled gets a reference to the given bool and assigns it to the AudioRecordingEnabled field.
func (o *InlineObject63) SetAudioRecordingEnabled(v bool) {
	o.AudioRecordingEnabled = &v
}

// GetCloudArchiveEnabled returns the CloudArchiveEnabled field value if set, zero value otherwise.
func (o *InlineObject63) GetCloudArchiveEnabled() bool {
	if o == nil || isNil(o.CloudArchiveEnabled) {
		var ret bool
		return ret
	}
	return *o.CloudArchiveEnabled
}

// GetCloudArchiveEnabledOk returns a tuple with the CloudArchiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetCloudArchiveEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.CloudArchiveEnabled) {
    return nil, false
	}
	return o.CloudArchiveEnabled, true
}

// HasCloudArchiveEnabled returns a boolean if a field has been set.
func (o *InlineObject63) HasCloudArchiveEnabled() bool {
	if o != nil && !isNil(o.CloudArchiveEnabled) {
		return true
	}

	return false
}

// SetCloudArchiveEnabled gets a reference to the given bool and assigns it to the CloudArchiveEnabled field.
func (o *InlineObject63) SetCloudArchiveEnabled(v bool) {
	o.CloudArchiveEnabled = &v
}

// GetMotionDetectorVersion returns the MotionDetectorVersion field value if set, zero value otherwise.
func (o *InlineObject63) GetMotionDetectorVersion() int32 {
	if o == nil || isNil(o.MotionDetectorVersion) {
		var ret int32
		return ret
	}
	return *o.MotionDetectorVersion
}

// GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMotionDetectorVersionOk() (*int32, bool) {
	if o == nil || isNil(o.MotionDetectorVersion) {
    return nil, false
	}
	return o.MotionDetectorVersion, true
}

// HasMotionDetectorVersion returns a boolean if a field has been set.
func (o *InlineObject63) HasMotionDetectorVersion() bool {
	if o != nil && !isNil(o.MotionDetectorVersion) {
		return true
	}

	return false
}

// SetMotionDetectorVersion gets a reference to the given int32 and assigns it to the MotionDetectorVersion field.
func (o *InlineObject63) SetMotionDetectorVersion(v int32) {
	o.MotionDetectorVersion = &v
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *InlineObject63) GetScheduleId() string {
	if o == nil || isNil(o.ScheduleId) {
		var ret string
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetScheduleIdOk() (*string, bool) {
	if o == nil || isNil(o.ScheduleId) {
    return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *InlineObject63) HasScheduleId() bool {
	if o != nil && !isNil(o.ScheduleId) {
		return true
	}

	return false
}

// SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.
func (o *InlineObject63) SetScheduleId(v string) {
	o.ScheduleId = &v
}

// GetMaxRetentionDays returns the MaxRetentionDays field value if set, zero value otherwise.
func (o *InlineObject63) GetMaxRetentionDays() int32 {
	if o == nil || isNil(o.MaxRetentionDays) {
		var ret int32
		return ret
	}
	return *o.MaxRetentionDays
}

// GetMaxRetentionDaysOk returns a tuple with the MaxRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMaxRetentionDaysOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetentionDays) {
    return nil, false
	}
	return o.MaxRetentionDays, true
}

// HasMaxRetentionDays returns a boolean if a field has been set.
func (o *InlineObject63) HasMaxRetentionDays() bool {
	if o != nil && !isNil(o.MaxRetentionDays) {
		return true
	}

	return false
}

// SetMaxRetentionDays gets a reference to the given int32 and assigns it to the MaxRetentionDays field.
func (o *InlineObject63) SetMaxRetentionDays(v int32) {
	o.MaxRetentionDays = &v
}

// GetVideoSettings returns the VideoSettings field value if set, zero value otherwise.
func (o *InlineObject63) GetVideoSettings() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings {
	if o == nil || isNil(o.VideoSettings) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings
		return ret
	}
	return *o.VideoSettings
}

// GetVideoSettingsOk returns a tuple with the VideoSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetVideoSettingsOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings, bool) {
	if o == nil || isNil(o.VideoSettings) {
    return nil, false
	}
	return o.VideoSettings, true
}

// HasVideoSettings returns a boolean if a field has been set.
func (o *InlineObject63) HasVideoSettings() bool {
	if o != nil && !isNil(o.VideoSettings) {
		return true
	}

	return false
}

// SetVideoSettings gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings and assigns it to the VideoSettings field.
func (o *InlineObject63) SetVideoSettings(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) {
	o.VideoSettings = &v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.MotionBasedRetentionEnabled) {
		toSerialize["motionBasedRetentionEnabled"] = o.MotionBasedRetentionEnabled
	}
	if !isNil(o.RestrictedBandwidthModeEnabled) {
		toSerialize["restrictedBandwidthModeEnabled"] = o.RestrictedBandwidthModeEnabled
	}
	if !isNil(o.AudioRecordingEnabled) {
		toSerialize["audioRecordingEnabled"] = o.AudioRecordingEnabled
	}
	if !isNil(o.CloudArchiveEnabled) {
		toSerialize["cloudArchiveEnabled"] = o.CloudArchiveEnabled
	}
	if !isNil(o.MotionDetectorVersion) {
		toSerialize["motionDetectorVersion"] = o.MotionDetectorVersion
	}
	if !isNil(o.ScheduleId) {
		toSerialize["scheduleId"] = o.ScheduleId
	}
	if !isNil(o.MaxRetentionDays) {
		toSerialize["maxRetentionDays"] = o.MaxRetentionDays
	}
	if !isNil(o.VideoSettings) {
		toSerialize["videoSettings"] = o.VideoSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


