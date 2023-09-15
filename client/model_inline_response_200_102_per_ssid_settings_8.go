/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200102PerSsidSettings8 Settings for SSID 8
type InlineResponse200102PerSsidSettings8 struct {
	// Name of SSID
	Name *string `json:"name,omitempty"`
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *InlineResponse200102ApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewInlineResponse200102PerSsidSettings8 instantiates a new InlineResponse200102PerSsidSettings8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200102PerSsidSettings8() *InlineResponse200102PerSsidSettings8 {
	this := InlineResponse200102PerSsidSettings8{}
	return &this
}

// NewInlineResponse200102PerSsidSettings8WithDefaults instantiates a new InlineResponse200102PerSsidSettings8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200102PerSsidSettings8WithDefaults() *InlineResponse200102PerSsidSettings8 {
	this := InlineResponse200102PerSsidSettings8{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200102PerSsidSettings8) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102PerSsidSettings8) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200102PerSsidSettings8) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200102PerSsidSettings8) SetName(v string) {
	o.Name = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *InlineResponse200102PerSsidSettings8) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102PerSsidSettings8) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *InlineResponse200102PerSsidSettings8) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *InlineResponse200102PerSsidSettings8) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *InlineResponse200102PerSsidSettings8) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102PerSsidSettings8) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *InlineResponse200102PerSsidSettings8) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *InlineResponse200102PerSsidSettings8) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *InlineResponse200102PerSsidSettings8) GetBands() InlineResponse200102ApBandSettingsBands {
	if o == nil || isNil(o.Bands) {
		var ret InlineResponse200102ApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102PerSsidSettings8) GetBandsOk() (*InlineResponse200102ApBandSettingsBands, bool) {
	if o == nil || isNil(o.Bands) {
    return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *InlineResponse200102PerSsidSettings8) HasBands() bool {
	if o != nil && !isNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given InlineResponse200102ApBandSettingsBands and assigns it to the Bands field.
func (o *InlineResponse200102PerSsidSettings8) SetBands(v InlineResponse200102ApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *InlineResponse200102PerSsidSettings8) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102PerSsidSettings8) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *InlineResponse200102PerSsidSettings8) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *InlineResponse200102PerSsidSettings8) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o InlineResponse200102PerSsidSettings8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
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

type NullableInlineResponse200102PerSsidSettings8 struct {
	value *InlineResponse200102PerSsidSettings8
	isSet bool
}

func (v NullableInlineResponse200102PerSsidSettings8) Get() *InlineResponse200102PerSsidSettings8 {
	return v.value
}

func (v *NullableInlineResponse200102PerSsidSettings8) Set(val *InlineResponse200102PerSsidSettings8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200102PerSsidSettings8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200102PerSsidSettings8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200102PerSsidSettings8(val *InlineResponse200102PerSsidSettings8) *NullableInlineResponse200102PerSsidSettings8 {
	return &NullableInlineResponse200102PerSsidSettings8{value: val, isSet: true}
}

func (v NullableInlineResponse200102PerSsidSettings8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200102PerSsidSettings8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


