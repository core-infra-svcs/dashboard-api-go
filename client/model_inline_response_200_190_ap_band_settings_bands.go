/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200190ApBandSettingsBands Settings related to all bands
type InlineResponse200190ApBandSettingsBands struct {
	// List of enabled bands. Can include [\"2.4\", \"5\", \"6\", \"disabled\"]
	Enabled []string `json:"enabled,omitempty"`
}

// NewInlineResponse200190ApBandSettingsBands instantiates a new InlineResponse200190ApBandSettingsBands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200190ApBandSettingsBands() *InlineResponse200190ApBandSettingsBands {
	this := InlineResponse200190ApBandSettingsBands{}
	return &this
}

// NewInlineResponse200190ApBandSettingsBandsWithDefaults instantiates a new InlineResponse200190ApBandSettingsBands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200190ApBandSettingsBandsWithDefaults() *InlineResponse200190ApBandSettingsBands {
	this := InlineResponse200190ApBandSettingsBands{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200190ApBandSettingsBands) GetEnabled() []string {
	if o == nil || isNil(o.Enabled) {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200190ApBandSettingsBands) GetEnabledOk() ([]string, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200190ApBandSettingsBands) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *InlineResponse200190ApBandSettingsBands) SetEnabled(v []string) {
	o.Enabled = v
}

func (o InlineResponse200190ApBandSettingsBands) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200190ApBandSettingsBands struct {
	value *InlineResponse200190ApBandSettingsBands
	isSet bool
}

func (v NullableInlineResponse200190ApBandSettingsBands) Get() *InlineResponse200190ApBandSettingsBands {
	return v.value
}

func (v *NullableInlineResponse200190ApBandSettingsBands) Set(val *InlineResponse200190ApBandSettingsBands) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200190ApBandSettingsBands) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200190ApBandSettingsBands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200190ApBandSettingsBands(val *InlineResponse200190ApBandSettingsBands) *NullableInlineResponse200190ApBandSettingsBands {
	return &NullableInlineResponse200190ApBandSettingsBands{value: val, isSet: true}
}

func (v NullableInlineResponse200190ApBandSettingsBands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200190ApBandSettingsBands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


