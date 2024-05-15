/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20013 struct for InlineResponse20013
type InlineResponse20013 struct {
	// Whether custom analytics is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Custom analytics artifact ID
	ArtifactId *string `json:"artifactId,omitempty"`
	// Parameters for the custom analytics workload
	Parameters []InlineResponse20013Parameters `json:"parameters,omitempty"`
}

// NewInlineResponse20013 instantiates a new InlineResponse20013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20013() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// NewInlineResponse20013WithDefaults instantiates a new InlineResponse20013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20013WithDefaults() *InlineResponse20013 {
	this := InlineResponse20013{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20013) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20013) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20013) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *InlineResponse20013) GetArtifactId() string {
	if o == nil || isNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetArtifactIdOk() (*string, bool) {
	if o == nil || isNil(o.ArtifactId) {
    return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *InlineResponse20013) HasArtifactId() bool {
	if o != nil && !isNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *InlineResponse20013) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *InlineResponse20013) GetParameters() []InlineResponse20013Parameters {
	if o == nil || isNil(o.Parameters) {
		var ret []InlineResponse20013Parameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20013) GetParametersOk() ([]InlineResponse20013Parameters, bool) {
	if o == nil || isNil(o.Parameters) {
    return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *InlineResponse20013) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []InlineResponse20013Parameters and assigns it to the Parameters field.
func (o *InlineResponse20013) SetParameters(v []InlineResponse20013Parameters) {
	o.Parameters = v
}

func (o InlineResponse20013) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20013 struct {
	value *InlineResponse20013
	isSet bool
}

func (v NullableInlineResponse20013) Get() *InlineResponse20013 {
	return v.value
}

func (v *NullableInlineResponse20013) Set(val *InlineResponse20013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20013(val *InlineResponse20013) *NullableInlineResponse20013 {
	return &NullableInlineResponse20013{value: val, isSet: true}
}

func (v NullableInlineResponse20013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


