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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// The device serial
	Serial *string `json:"serial,omitempty"`
	// RF Profile ID
	RfProfileId *string `json:"rfProfileId,omitempty"`
	TwoFourGhzSettings *InlineResponse2001TwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *InlineResponse2001FiveGhzSettings `json:"fiveGhzSettings,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2001) SetSerial(v string) {
	o.Serial = &v
}

// GetRfProfileId returns the RfProfileId field value if set, zero value otherwise.
func (o *InlineResponse2001) GetRfProfileId() string {
	if o == nil || isNil(o.RfProfileId) {
		var ret string
		return ret
	}
	return *o.RfProfileId
}

// GetRfProfileIdOk returns a tuple with the RfProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetRfProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.RfProfileId) {
    return nil, false
	}
	return o.RfProfileId, true
}

// HasRfProfileId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasRfProfileId() bool {
	if o != nil && !isNil(o.RfProfileId) {
		return true
	}

	return false
}

// SetRfProfileId gets a reference to the given string and assigns it to the RfProfileId field.
func (o *InlineResponse2001) SetRfProfileId(v string) {
	o.RfProfileId = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse2001) GetTwoFourGhzSettings() InlineResponse2001TwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret InlineResponse2001TwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTwoFourGhzSettingsOk() (*InlineResponse2001TwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse2001) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given InlineResponse2001TwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineResponse2001) SetTwoFourGhzSettings(v InlineResponse2001TwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse2001) GetFiveGhzSettings() InlineResponse2001FiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret InlineResponse2001FiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetFiveGhzSettingsOk() (*InlineResponse2001FiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse2001) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given InlineResponse2001FiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineResponse2001) SetFiveGhzSettings(v InlineResponse2001FiveGhzSettings) {
	o.FiveGhzSettings = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.RfProfileId) {
		toSerialize["rfProfileId"] = o.RfProfileId
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


