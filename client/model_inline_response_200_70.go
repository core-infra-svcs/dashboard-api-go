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

// InlineResponse20070 struct for InlineResponse20070
type InlineResponse20070 struct {
	// Is the warm spare enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Serial number of the primary appliance
	PrimarySerial *string `json:"primarySerial,omitempty"`
	// Serial number of the warm spare appliance
	SpareSerial *string `json:"spareSerial,omitempty"`
	// Uplink mode, either virtual or public
	UplinkMode *string `json:"uplinkMode,omitempty"`
	Wan1 *InlineResponse20070Wan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse20070Wan2 `json:"wan2,omitempty"`
}

// NewInlineResponse20070 instantiates a new InlineResponse20070 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20070() *InlineResponse20070 {
	this := InlineResponse20070{}
	return &this
}

// NewInlineResponse20070WithDefaults instantiates a new InlineResponse20070 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20070WithDefaults() *InlineResponse20070 {
	this := InlineResponse20070{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20070) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20070) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20070) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrimarySerial returns the PrimarySerial field value if set, zero value otherwise.
func (o *InlineResponse20070) GetPrimarySerial() string {
	if o == nil || isNil(o.PrimarySerial) {
		var ret string
		return ret
	}
	return *o.PrimarySerial
}

// GetPrimarySerialOk returns a tuple with the PrimarySerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetPrimarySerialOk() (*string, bool) {
	if o == nil || isNil(o.PrimarySerial) {
    return nil, false
	}
	return o.PrimarySerial, true
}

// HasPrimarySerial returns a boolean if a field has been set.
func (o *InlineResponse20070) HasPrimarySerial() bool {
	if o != nil && !isNil(o.PrimarySerial) {
		return true
	}

	return false
}

// SetPrimarySerial gets a reference to the given string and assigns it to the PrimarySerial field.
func (o *InlineResponse20070) SetPrimarySerial(v string) {
	o.PrimarySerial = &v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineResponse20070) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineResponse20070) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineResponse20070) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

// GetUplinkMode returns the UplinkMode field value if set, zero value otherwise.
func (o *InlineResponse20070) GetUplinkMode() string {
	if o == nil || isNil(o.UplinkMode) {
		var ret string
		return ret
	}
	return *o.UplinkMode
}

// GetUplinkModeOk returns a tuple with the UplinkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetUplinkModeOk() (*string, bool) {
	if o == nil || isNil(o.UplinkMode) {
    return nil, false
	}
	return o.UplinkMode, true
}

// HasUplinkMode returns a boolean if a field has been set.
func (o *InlineResponse20070) HasUplinkMode() bool {
	if o != nil && !isNil(o.UplinkMode) {
		return true
	}

	return false
}

// SetUplinkMode gets a reference to the given string and assigns it to the UplinkMode field.
func (o *InlineResponse20070) SetUplinkMode(v string) {
	o.UplinkMode = &v
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse20070) GetWan1() InlineResponse20070Wan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse20070Wan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetWan1Ok() (*InlineResponse20070Wan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse20070) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse20070Wan1 and assigns it to the Wan1 field.
func (o *InlineResponse20070) SetWan1(v InlineResponse20070Wan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse20070) GetWan2() InlineResponse20070Wan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse20070Wan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20070) GetWan2Ok() (*InlineResponse20070Wan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse20070) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse20070Wan2 and assigns it to the Wan2 field.
func (o *InlineResponse20070) SetWan2(v InlineResponse20070Wan2) {
	o.Wan2 = &v
}

func (o InlineResponse20070) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20070 struct {
	value *InlineResponse20070
	isSet bool
}

func (v NullableInlineResponse20070) Get() *InlineResponse20070 {
	return v.value
}

func (v *NullableInlineResponse20070) Set(val *InlineResponse20070) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20070) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20070) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20070(val *InlineResponse20070) *NullableInlineResponse20070 {
	return &NullableInlineResponse20070{value: val, isSet: true}
}

func (v NullableInlineResponse20070) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20070) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


