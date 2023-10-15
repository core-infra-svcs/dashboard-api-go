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

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	// Boolean indicating if external rtsp stream is exposed
	ExternalRtspEnabled *bool `json:"externalRtspEnabled,omitempty"`
}

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetExternalRtspEnabled returns the ExternalRtspEnabled field value if set, zero value otherwise.
func (o *InlineObject8) GetExternalRtspEnabled() bool {
	if o == nil || isNil(o.ExternalRtspEnabled) {
		var ret bool
		return ret
	}
	return *o.ExternalRtspEnabled
}

// GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetExternalRtspEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ExternalRtspEnabled) {
    return nil, false
	}
	return o.ExternalRtspEnabled, true
}

// HasExternalRtspEnabled returns a boolean if a field has been set.
func (o *InlineObject8) HasExternalRtspEnabled() bool {
	if o != nil && !isNil(o.ExternalRtspEnabled) {
		return true
	}

	return false
}

// SetExternalRtspEnabled gets a reference to the given bool and assigns it to the ExternalRtspEnabled field.
func (o *InlineObject8) SetExternalRtspEnabled(v bool) {
	o.ExternalRtspEnabled = &v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalRtspEnabled) {
		toSerialize["externalRtspEnabled"] = o.ExternalRtspEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


