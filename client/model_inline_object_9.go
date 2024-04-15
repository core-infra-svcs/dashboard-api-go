/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	// The ID of a quality and retention profile to assign to the camera. The profile's settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera.
	ProfileId *string `json:"profileId,omitempty"`
	// Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera.
	MotionBasedRetentionEnabled *bool `json:"motionBasedRetentionEnabled,omitempty"`
	// Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera
	AudioRecordingEnabled *bool `json:"audioRecordingEnabled,omitempty"`
	// Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras.
	RestrictedBandwidthModeEnabled *bool `json:"restrictedBandwidthModeEnabled,omitempty"`
	// Quality of the camera. Can be one of 'Standard', 'High' or 'Enhanced'. Not all qualities are supported by every camera model.
	Quality *string `json:"quality,omitempty"`
	// Resolution of the camera. Can be one of '1280x720', '1920x1080', '1080x1080', '2112x2112', '2880x2880', '2688x1512' or '3840x2160'.Not all resolutions are supported by every camera model.
	Resolution *string `json:"resolution,omitempty"`
	// The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.
	MotionDetectorVersion *int32 `json:"motionDetectorVersion,omitempty"`
}

// NewInlineObject9 instantiates a new InlineObject9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// NewInlineObject9WithDefaults instantiates a new InlineObject9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9WithDefaults() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *InlineObject9) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *InlineObject9) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *InlineObject9) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetMotionBasedRetentionEnabled returns the MotionBasedRetentionEnabled field value if set, zero value otherwise.
func (o *InlineObject9) GetMotionBasedRetentionEnabled() bool {
	if o == nil || isNil(o.MotionBasedRetentionEnabled) {
		var ret bool
		return ret
	}
	return *o.MotionBasedRetentionEnabled
}

// GetMotionBasedRetentionEnabledOk returns a tuple with the MotionBasedRetentionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetMotionBasedRetentionEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MotionBasedRetentionEnabled) {
    return nil, false
	}
	return o.MotionBasedRetentionEnabled, true
}

// HasMotionBasedRetentionEnabled returns a boolean if a field has been set.
func (o *InlineObject9) HasMotionBasedRetentionEnabled() bool {
	if o != nil && !isNil(o.MotionBasedRetentionEnabled) {
		return true
	}

	return false
}

// SetMotionBasedRetentionEnabled gets a reference to the given bool and assigns it to the MotionBasedRetentionEnabled field.
func (o *InlineObject9) SetMotionBasedRetentionEnabled(v bool) {
	o.MotionBasedRetentionEnabled = &v
}

// GetAudioRecordingEnabled returns the AudioRecordingEnabled field value if set, zero value otherwise.
func (o *InlineObject9) GetAudioRecordingEnabled() bool {
	if o == nil || isNil(o.AudioRecordingEnabled) {
		var ret bool
		return ret
	}
	return *o.AudioRecordingEnabled
}

// GetAudioRecordingEnabledOk returns a tuple with the AudioRecordingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetAudioRecordingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AudioRecordingEnabled) {
    return nil, false
	}
	return o.AudioRecordingEnabled, true
}

// HasAudioRecordingEnabled returns a boolean if a field has been set.
func (o *InlineObject9) HasAudioRecordingEnabled() bool {
	if o != nil && !isNil(o.AudioRecordingEnabled) {
		return true
	}

	return false
}

// SetAudioRecordingEnabled gets a reference to the given bool and assigns it to the AudioRecordingEnabled field.
func (o *InlineObject9) SetAudioRecordingEnabled(v bool) {
	o.AudioRecordingEnabled = &v
}

// GetRestrictedBandwidthModeEnabled returns the RestrictedBandwidthModeEnabled field value if set, zero value otherwise.
func (o *InlineObject9) GetRestrictedBandwidthModeEnabled() bool {
	if o == nil || isNil(o.RestrictedBandwidthModeEnabled) {
		var ret bool
		return ret
	}
	return *o.RestrictedBandwidthModeEnabled
}

// GetRestrictedBandwidthModeEnabledOk returns a tuple with the RestrictedBandwidthModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetRestrictedBandwidthModeEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RestrictedBandwidthModeEnabled) {
    return nil, false
	}
	return o.RestrictedBandwidthModeEnabled, true
}

// HasRestrictedBandwidthModeEnabled returns a boolean if a field has been set.
func (o *InlineObject9) HasRestrictedBandwidthModeEnabled() bool {
	if o != nil && !isNil(o.RestrictedBandwidthModeEnabled) {
		return true
	}

	return false
}

// SetRestrictedBandwidthModeEnabled gets a reference to the given bool and assigns it to the RestrictedBandwidthModeEnabled field.
func (o *InlineObject9) SetRestrictedBandwidthModeEnabled(v bool) {
	o.RestrictedBandwidthModeEnabled = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *InlineObject9) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *InlineObject9) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *InlineObject9) SetQuality(v string) {
	o.Quality = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *InlineObject9) GetResolution() string {
	if o == nil || isNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetResolutionOk() (*string, bool) {
	if o == nil || isNil(o.Resolution) {
    return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *InlineObject9) HasResolution() bool {
	if o != nil && !isNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *InlineObject9) SetResolution(v string) {
	o.Resolution = &v
}

// GetMotionDetectorVersion returns the MotionDetectorVersion field value if set, zero value otherwise.
func (o *InlineObject9) GetMotionDetectorVersion() int32 {
	if o == nil || isNil(o.MotionDetectorVersion) {
		var ret int32
		return ret
	}
	return *o.MotionDetectorVersion
}

// GetMotionDetectorVersionOk returns a tuple with the MotionDetectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetMotionDetectorVersionOk() (*int32, bool) {
	if o == nil || isNil(o.MotionDetectorVersion) {
    return nil, false
	}
	return o.MotionDetectorVersion, true
}

// HasMotionDetectorVersion returns a boolean if a field has been set.
func (o *InlineObject9) HasMotionDetectorVersion() bool {
	if o != nil && !isNil(o.MotionDetectorVersion) {
		return true
	}

	return false
}

// SetMotionDetectorVersion gets a reference to the given int32 and assigns it to the MotionDetectorVersion field.
func (o *InlineObject9) SetMotionDetectorVersion(v int32) {
	o.MotionDetectorVersion = &v
}

func (o InlineObject9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.MotionBasedRetentionEnabled) {
		toSerialize["motionBasedRetentionEnabled"] = o.MotionBasedRetentionEnabled
	}
	if !isNil(o.AudioRecordingEnabled) {
		toSerialize["audioRecordingEnabled"] = o.AudioRecordingEnabled
	}
	if !isNil(o.RestrictedBandwidthModeEnabled) {
		toSerialize["restrictedBandwidthModeEnabled"] = o.RestrictedBandwidthModeEnabled
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !isNil(o.MotionDetectorVersion) {
		toSerialize["motionDetectorVersion"] = o.MotionDetectorVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject9 struct {
	value *InlineObject9
	isSet bool
}

func (v NullableInlineObject9) Get() *InlineObject9 {
	return v.value
}

func (v *NullableInlineObject9) Set(val *InlineObject9) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9(val *InlineObject9) *NullableInlineObject9 {
	return &NullableInlineObject9{value: val, isSet: true}
}

func (v NullableInlineObject9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


