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

// InlineResponse200190ApBandSettings Settings that will be enabled if selectionType is set to 'ap'.
type InlineResponse200190ApBandSettings struct {
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *InlineResponse200190ApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band. Can be either true or false. Defaults to true.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewInlineResponse200190ApBandSettings instantiates a new InlineResponse200190ApBandSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200190ApBandSettings() *InlineResponse200190ApBandSettings {
	this := InlineResponse200190ApBandSettings{}
	return &this
}

// NewInlineResponse200190ApBandSettingsWithDefaults instantiates a new InlineResponse200190ApBandSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200190ApBandSettingsWithDefaults() *InlineResponse200190ApBandSettings {
	this := InlineResponse200190ApBandSettings{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *InlineResponse200190ApBandSettings) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200190ApBandSettings) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *InlineResponse200190ApBandSettings) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *InlineResponse200190ApBandSettings) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *InlineResponse200190ApBandSettings) GetBands() InlineResponse200190ApBandSettingsBands {
	if o == nil || isNil(o.Bands) {
		var ret InlineResponse200190ApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200190ApBandSettings) GetBandsOk() (*InlineResponse200190ApBandSettingsBands, bool) {
	if o == nil || isNil(o.Bands) {
    return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *InlineResponse200190ApBandSettings) HasBands() bool {
	if o != nil && !isNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given InlineResponse200190ApBandSettingsBands and assigns it to the Bands field.
func (o *InlineResponse200190ApBandSettings) SetBands(v InlineResponse200190ApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *InlineResponse200190ApBandSettings) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200190ApBandSettings) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *InlineResponse200190ApBandSettings) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *InlineResponse200190ApBandSettings) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o InlineResponse200190ApBandSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200190ApBandSettings struct {
	value *InlineResponse200190ApBandSettings
	isSet bool
}

func (v NullableInlineResponse200190ApBandSettings) Get() *InlineResponse200190ApBandSettings {
	return v.value
}

func (v *NullableInlineResponse200190ApBandSettings) Set(val *InlineResponse200190ApBandSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200190ApBandSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200190ApBandSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200190ApBandSettings(val *InlineResponse200190ApBandSettings) *NullableInlineResponse200190ApBandSettings {
	return &NullableInlineResponse200190ApBandSettings{value: val, isSet: true}
}

func (v NullableInlineResponse200190ApBandSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200190ApBandSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

