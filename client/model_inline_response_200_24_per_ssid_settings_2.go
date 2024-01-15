/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20024PerSsidSettings2 Settings for SSID 2.
type InlineResponse20024PerSsidSettings2 struct {
	// Band mode of this SSID
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewInlineResponse20024PerSsidSettings2 instantiates a new InlineResponse20024PerSsidSettings2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20024PerSsidSettings2() *InlineResponse20024PerSsidSettings2 {
	this := InlineResponse20024PerSsidSettings2{}
	return &this
}

// NewInlineResponse20024PerSsidSettings2WithDefaults instantiates a new InlineResponse20024PerSsidSettings2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20024PerSsidSettings2WithDefaults() *InlineResponse20024PerSsidSettings2 {
	this := InlineResponse20024PerSsidSettings2{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *InlineResponse20024PerSsidSettings2) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024PerSsidSettings2) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *InlineResponse20024PerSsidSettings2) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *InlineResponse20024PerSsidSettings2) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *InlineResponse20024PerSsidSettings2) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024PerSsidSettings2) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *InlineResponse20024PerSsidSettings2) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *InlineResponse20024PerSsidSettings2) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o InlineResponse20024PerSsidSettings2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20024PerSsidSettings2 struct {
	value *InlineResponse20024PerSsidSettings2
	isSet bool
}

func (v NullableInlineResponse20024PerSsidSettings2) Get() *InlineResponse20024PerSsidSettings2 {
	return v.value
}

func (v *NullableInlineResponse20024PerSsidSettings2) Set(val *InlineResponse20024PerSsidSettings2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20024PerSsidSettings2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20024PerSsidSettings2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20024PerSsidSettings2(val *InlineResponse20024PerSsidSettings2) *NullableInlineResponse20024PerSsidSettings2 {
	return &NullableInlineResponse20024PerSsidSettings2{value: val, isSet: true}
}

func (v NullableInlineResponse20024PerSsidSettings2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20024PerSsidSettings2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


