/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject7 struct for InlineObject7
type InlineObject7 struct {
	// Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera
	SenseEnabled *bool `json:"senseEnabled,omitempty"`
	// The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera
	MqttBrokerId *string `json:"mqttBrokerId,omitempty"`
	AudioDetection *DevicesSerialCameraSenseAudioDetection `json:"audioDetection,omitempty"`
	// The ID of the object detection model
	DetectionModelId *string `json:"detectionModelId,omitempty"`
}

// NewInlineObject7 instantiates a new InlineObject7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject7() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// NewInlineObject7WithDefaults instantiates a new InlineObject7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject7WithDefaults() *InlineObject7 {
	this := InlineObject7{}
	return &this
}

// GetSenseEnabled returns the SenseEnabled field value if set, zero value otherwise.
func (o *InlineObject7) GetSenseEnabled() bool {
	if o == nil || isNil(o.SenseEnabled) {
		var ret bool
		return ret
	}
	return *o.SenseEnabled
}

// GetSenseEnabledOk returns a tuple with the SenseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetSenseEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SenseEnabled) {
    return nil, false
	}
	return o.SenseEnabled, true
}

// HasSenseEnabled returns a boolean if a field has been set.
func (o *InlineObject7) HasSenseEnabled() bool {
	if o != nil && !isNil(o.SenseEnabled) {
		return true
	}

	return false
}

// SetSenseEnabled gets a reference to the given bool and assigns it to the SenseEnabled field.
func (o *InlineObject7) SetSenseEnabled(v bool) {
	o.SenseEnabled = &v
}

// GetMqttBrokerId returns the MqttBrokerId field value if set, zero value otherwise.
func (o *InlineObject7) GetMqttBrokerId() string {
	if o == nil || isNil(o.MqttBrokerId) {
		var ret string
		return ret
	}
	return *o.MqttBrokerId
}

// GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetMqttBrokerIdOk() (*string, bool) {
	if o == nil || isNil(o.MqttBrokerId) {
    return nil, false
	}
	return o.MqttBrokerId, true
}

// HasMqttBrokerId returns a boolean if a field has been set.
func (o *InlineObject7) HasMqttBrokerId() bool {
	if o != nil && !isNil(o.MqttBrokerId) {
		return true
	}

	return false
}

// SetMqttBrokerId gets a reference to the given string and assigns it to the MqttBrokerId field.
func (o *InlineObject7) SetMqttBrokerId(v string) {
	o.MqttBrokerId = &v
}

// GetAudioDetection returns the AudioDetection field value if set, zero value otherwise.
func (o *InlineObject7) GetAudioDetection() DevicesSerialCameraSenseAudioDetection {
	if o == nil || isNil(o.AudioDetection) {
		var ret DevicesSerialCameraSenseAudioDetection
		return ret
	}
	return *o.AudioDetection
}

// GetAudioDetectionOk returns a tuple with the AudioDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetAudioDetectionOk() (*DevicesSerialCameraSenseAudioDetection, bool) {
	if o == nil || isNil(o.AudioDetection) {
    return nil, false
	}
	return o.AudioDetection, true
}

// HasAudioDetection returns a boolean if a field has been set.
func (o *InlineObject7) HasAudioDetection() bool {
	if o != nil && !isNil(o.AudioDetection) {
		return true
	}

	return false
}

// SetAudioDetection gets a reference to the given DevicesSerialCameraSenseAudioDetection and assigns it to the AudioDetection field.
func (o *InlineObject7) SetAudioDetection(v DevicesSerialCameraSenseAudioDetection) {
	o.AudioDetection = &v
}

// GetDetectionModelId returns the DetectionModelId field value if set, zero value otherwise.
func (o *InlineObject7) GetDetectionModelId() string {
	if o == nil || isNil(o.DetectionModelId) {
		var ret string
		return ret
	}
	return *o.DetectionModelId
}

// GetDetectionModelIdOk returns a tuple with the DetectionModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject7) GetDetectionModelIdOk() (*string, bool) {
	if o == nil || isNil(o.DetectionModelId) {
    return nil, false
	}
	return o.DetectionModelId, true
}

// HasDetectionModelId returns a boolean if a field has been set.
func (o *InlineObject7) HasDetectionModelId() bool {
	if o != nil && !isNil(o.DetectionModelId) {
		return true
	}

	return false
}

// SetDetectionModelId gets a reference to the given string and assigns it to the DetectionModelId field.
func (o *InlineObject7) SetDetectionModelId(v string) {
	o.DetectionModelId = &v
}

func (o InlineObject7) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SenseEnabled) {
		toSerialize["senseEnabled"] = o.SenseEnabled
	}
	if !isNil(o.MqttBrokerId) {
		toSerialize["mqttBrokerId"] = o.MqttBrokerId
	}
	if !isNil(o.AudioDetection) {
		toSerialize["audioDetection"] = o.AudioDetection
	}
	if !isNil(o.DetectionModelId) {
		toSerialize["detectionModelId"] = o.DetectionModelId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject7 struct {
	value *InlineObject7
	isSet bool
}

func (v NullableInlineObject7) Get() *InlineObject7 {
	return v.value
}

func (v *NullableInlineObject7) Set(val *InlineObject7) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject7) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject7(val *InlineObject7) *NullableInlineObject7 {
	return &NullableInlineObject7{value: val, isSet: true}
}

func (v NullableInlineObject7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


