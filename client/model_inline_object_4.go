/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// Enable custom analytics
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the custom analytics artifact
	ArtifactId *string `json:"artifactId,omitempty"`
	// Parameters for the custom analytics workload
	Parameters []DevicesSerialCameraCustomAnalyticsParameters `json:"parameters,omitempty"`
}

// NewInlineObject4 instantiates a new InlineObject4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject4() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// NewInlineObject4WithDefaults instantiates a new InlineObject4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject4WithDefaults() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject4) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject4) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject4) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *InlineObject4) GetArtifactId() string {
	if o == nil || isNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetArtifactIdOk() (*string, bool) {
	if o == nil || isNil(o.ArtifactId) {
    return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *InlineObject4) HasArtifactId() bool {
	if o != nil && !isNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *InlineObject4) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *InlineObject4) GetParameters() []DevicesSerialCameraCustomAnalyticsParameters {
	if o == nil || isNil(o.Parameters) {
		var ret []DevicesSerialCameraCustomAnalyticsParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetParametersOk() ([]DevicesSerialCameraCustomAnalyticsParameters, bool) {
	if o == nil || isNil(o.Parameters) {
    return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *InlineObject4) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []DevicesSerialCameraCustomAnalyticsParameters and assigns it to the Parameters field.
func (o *InlineObject4) SetParameters(v []DevicesSerialCameraCustomAnalyticsParameters) {
	o.Parameters = v
}

func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ArtifactId) {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !isNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject4 struct {
	value *InlineObject4
	isSet bool
}

func (v NullableInlineObject4) Get() *InlineObject4 {
	return v.value
}

func (v *NullableInlineObject4) Set(val *InlineObject4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject4(val *InlineObject4) *NullableInlineObject4 {
	return &NullableInlineObject4{value: val, isSet: true}
}

func (v NullableInlineObject4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


