/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200183ApBandSettingsBands Settings related to all bands
type InlineResponse200183ApBandSettingsBands struct {
	// List of enabled bands. Can include [\"2.4\", \"5\", \"6\", \"disabled\"]
	Enabled []string `json:"enabled,omitempty"`
}

// NewInlineResponse200183ApBandSettingsBands instantiates a new InlineResponse200183ApBandSettingsBands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200183ApBandSettingsBands() *InlineResponse200183ApBandSettingsBands {
	this := InlineResponse200183ApBandSettingsBands{}
	return &this
}

// NewInlineResponse200183ApBandSettingsBandsWithDefaults instantiates a new InlineResponse200183ApBandSettingsBands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200183ApBandSettingsBandsWithDefaults() *InlineResponse200183ApBandSettingsBands {
	this := InlineResponse200183ApBandSettingsBands{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200183ApBandSettingsBands) GetEnabled() []string {
	if o == nil || isNil(o.Enabled) {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183ApBandSettingsBands) GetEnabledOk() ([]string, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200183ApBandSettingsBands) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *InlineResponse200183ApBandSettingsBands) SetEnabled(v []string) {
	o.Enabled = v
}

func (o InlineResponse200183ApBandSettingsBands) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200183ApBandSettingsBands struct {
	value *InlineResponse200183ApBandSettingsBands
	isSet bool
}

func (v NullableInlineResponse200183ApBandSettingsBands) Get() *InlineResponse200183ApBandSettingsBands {
	return v.value
}

func (v *NullableInlineResponse200183ApBandSettingsBands) Set(val *InlineResponse200183ApBandSettingsBands) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200183ApBandSettingsBands) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200183ApBandSettingsBands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200183ApBandSettingsBands(val *InlineResponse200183ApBandSettingsBands) *NullableInlineResponse200183ApBandSettingsBands {
	return &NullableInlineResponse200183ApBandSettingsBands{value: val, isSet: true}
}

func (v NullableInlineResponse200183ApBandSettingsBands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200183ApBandSettingsBands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


