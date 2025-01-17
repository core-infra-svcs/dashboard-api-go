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

// InlineResponse20072 struct for InlineResponse20072
type InlineResponse20072 struct {
	// Is the warm spare enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Serial number of the primary appliance
	PrimarySerial *string `json:"primarySerial,omitempty"`
	// Serial number of the warm spare appliance
	SpareSerial *string `json:"spareSerial,omitempty"`
	// Uplink mode, either virtual or public
	UplinkMode *string `json:"uplinkMode,omitempty"`
	Wan1 *InlineResponse20072Wan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20072Wan2 `json:"wan2,omitempty"`
}

// NewInlineResponse20072 instantiates a new InlineResponse20072 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20072() *InlineResponse20072 {
	this := InlineResponse20072{}
	return &this
}

// NewInlineResponse20072WithDefaults instantiates a new InlineResponse20072 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20072WithDefaults() *InlineResponse20072 {
	this := InlineResponse20072{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20072) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20072) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20072) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrimarySerial returns the PrimarySerial field value if set, zero value otherwise.
func (o *InlineResponse20072) GetPrimarySerial() string {
	if o == nil || isNil(o.PrimarySerial) {
		var ret string
		return ret
	}
	return *o.PrimarySerial
}

// GetPrimarySerialOk returns a tuple with the PrimarySerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetPrimarySerialOk() (*string, bool) {
	if o == nil || isNil(o.PrimarySerial) {
    return nil, false
	}
	return o.PrimarySerial, true
}

// HasPrimarySerial returns a boolean if a field has been set.
func (o *InlineResponse20072) HasPrimarySerial() bool {
	if o != nil && !isNil(o.PrimarySerial) {
		return true
	}

	return false
}

// SetPrimarySerial gets a reference to the given string and assigns it to the PrimarySerial field.
func (o *InlineResponse20072) SetPrimarySerial(v string) {
	o.PrimarySerial = &v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineResponse20072) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineResponse20072) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineResponse20072) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

// GetUplinkMode returns the UplinkMode field value if set, zero value otherwise.
func (o *InlineResponse20072) GetUplinkMode() string {
	if o == nil || isNil(o.UplinkMode) {
		var ret string
		return ret
	}
	return *o.UplinkMode
}

// GetUplinkModeOk returns a tuple with the UplinkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetUplinkModeOk() (*string, bool) {
	if o == nil || isNil(o.UplinkMode) {
    return nil, false
	}
	return o.UplinkMode, true
}

// HasUplinkMode returns a boolean if a field has been set.
func (o *InlineResponse20072) HasUplinkMode() bool {
	if o != nil && !isNil(o.UplinkMode) {
		return true
	}

	return false
}

// SetUplinkMode gets a reference to the given string and assigns it to the UplinkMode field.
func (o *InlineResponse20072) SetUplinkMode(v string) {
	o.UplinkMode = &v
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20072) GetWan1() InlineResponse20072Wan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20072Wan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetWan1Ok() (*InlineResponse20072Wan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20072) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20072Wan1 and assigns it to the Wan1 field.
func (o *InlineResponse20072) SetWan1(v InlineResponse20072Wan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20072) GetWan2() InlineResponse20072Wan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20072Wan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20072) GetWan2Ok() (*InlineResponse20072Wan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20072) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20072Wan2 and assigns it to the Wan2 field.
func (o *InlineResponse20072) SetWan2(v InlineResponse20072Wan2) {
	o.Wan2 = &v
}

func (o InlineResponse20072) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrimarySerial) {
		toSerialize["primarySerial"] = o.PrimarySerial
	}
	if !isNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	if !isNil(o.UplinkMode) {
		toSerialize["uplinkMode"] = o.UplinkMode
	}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20072 struct {
	value *InlineResponse20072
	isSet bool
}

func (v NullableInlineResponse20072) Get() *InlineResponse20072 {
	return v.value
}

func (v *NullableInlineResponse20072) Set(val *InlineResponse20072) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20072) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20072) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20072(val *InlineResponse20072) *NullableInlineResponse20072 {
	return &NullableInlineResponse20072{value: val, isSet: true}
}

func (v NullableInlineResponse20072) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20072) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


