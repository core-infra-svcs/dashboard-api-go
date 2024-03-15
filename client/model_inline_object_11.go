/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject11 struct for InlineObject11
type InlineObject11 struct {
	// Boolean indicating if external rtsp stream is exposed
	ExternalRtspEnabled *bool `json:"externalRtspEnabled,omitempty"`
}

// NewInlineObject11 instantiates a new InlineObject11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// NewInlineObject11WithDefaults instantiates a new InlineObject11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11WithDefaults() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// GetExternalRtspEnabled returns the ExternalRtspEnabled field value if set, zero value otherwise.
func (o *InlineObject11) GetExternalRtspEnabled() bool {
	if o == nil || isNil(o.ExternalRtspEnabled) {
		var ret bool
		return ret
	}
	return *o.ExternalRtspEnabled
}

// GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetExternalRtspEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ExternalRtspEnabled) {
    return nil, false
	}
	return o.ExternalRtspEnabled, true
}

// HasExternalRtspEnabled returns a boolean if a field has been set.
func (o *InlineObject11) HasExternalRtspEnabled() bool {
	if o != nil && !isNil(o.ExternalRtspEnabled) {
		return true
	}

	return false
}

// SetExternalRtspEnabled gets a reference to the given bool and assigns it to the ExternalRtspEnabled field.
func (o *InlineObject11) SetExternalRtspEnabled(v bool) {
	o.ExternalRtspEnabled = &v
}

func (o InlineObject11) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalRtspEnabled) {
		toSerialize["externalRtspEnabled"] = o.ExternalRtspEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject11 struct {
	value *InlineObject11
	isSet bool
}

func (v NullableInlineObject11) Get() *InlineObject11 {
	return v.value
}

func (v *NullableInlineObject11) Set(val *InlineObject11) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11(val *InlineObject11) *NullableInlineObject11 {
	return &NullableInlineObject11{value: val, isSet: true}
}

func (v NullableInlineObject11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


