/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject76 struct for InlineObject76
type InlineObject76 struct {
	// Enable warm spare
	Enabled bool `json:"enabled"`
	// Serial number of the warm spare appliance
	SpareSerial *string `json:"spareSerial,omitempty"`
	// Uplink mode, either virtual or public
	UplinkMode *string `json:"uplinkMode,omitempty"`
	// The WAN 1 shared IP
	VirtualIp1 *string `json:"virtualIp1,omitempty"`
	// The WAN 2 shared IP
	VirtualIp2 *string `json:"virtualIp2,omitempty"`
}

// NewInlineObject76 instantiates a new InlineObject76 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject76(enabled bool) *InlineObject76 {
	this := InlineObject76{}
	this.Enabled = enabled
	return &this
}

// NewInlineObject76WithDefaults instantiates a new InlineObject76 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject76WithDefaults() *InlineObject76 {
	this := InlineObject76{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineObject76) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineObject76) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSpareSerial returns the SpareSerial field value if set, zero value otherwise.
func (o *InlineObject76) GetSpareSerial() string {
	if o == nil || isNil(o.SpareSerial) {
		var ret string
		return ret
	}
	return *o.SpareSerial
}

// GetSpareSerialOk returns a tuple with the SpareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetSpareSerialOk() (*string, bool) {
	if o == nil || isNil(o.SpareSerial) {
    return nil, false
	}
	return o.SpareSerial, true
}

// HasSpareSerial returns a boolean if a field has been set.
func (o *InlineObject76) HasSpareSerial() bool {
	if o != nil && !isNil(o.SpareSerial) {
		return true
	}

	return false
}

// SetSpareSerial gets a reference to the given string and assigns it to the SpareSerial field.
func (o *InlineObject76) SetSpareSerial(v string) {
	o.SpareSerial = &v
}

// GetUplinkMode returns the UplinkMode field value if set, zero value otherwise.
func (o *InlineObject76) GetUplinkMode() string {
	if o == nil || isNil(o.UplinkMode) {
		var ret string
		return ret
	}
	return *o.UplinkMode
}

// GetUplinkModeOk returns a tuple with the UplinkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetUplinkModeOk() (*string, bool) {
	if o == nil || isNil(o.UplinkMode) {
    return nil, false
	}
	return o.UplinkMode, true
}

// HasUplinkMode returns a boolean if a field has been set.
func (o *InlineObject76) HasUplinkMode() bool {
	if o != nil && !isNil(o.UplinkMode) {
		return true
	}

	return false
}

// SetUplinkMode gets a reference to the given string and assigns it to the UplinkMode field.
func (o *InlineObject76) SetUplinkMode(v string) {
	o.UplinkMode = &v
}

// GetVirtualIp1 returns the VirtualIp1 field value if set, zero value otherwise.
func (o *InlineObject76) GetVirtualIp1() string {
	if o == nil || isNil(o.VirtualIp1) {
		var ret string
		return ret
	}
	return *o.VirtualIp1
}

// GetVirtualIp1Ok returns a tuple with the VirtualIp1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetVirtualIp1Ok() (*string, bool) {
	if o == nil || isNil(o.VirtualIp1) {
    return nil, false
	}
	return o.VirtualIp1, true
}

// HasVirtualIp1 returns a boolean if a field has been set.
func (o *InlineObject76) HasVirtualIp1() bool {
	if o != nil && !isNil(o.VirtualIp1) {
		return true
	}

	return false
}

// SetVirtualIp1 gets a reference to the given string and assigns it to the VirtualIp1 field.
func (o *InlineObject76) SetVirtualIp1(v string) {
	o.VirtualIp1 = &v
}

// GetVirtualIp2 returns the VirtualIp2 field value if set, zero value otherwise.
func (o *InlineObject76) GetVirtualIp2() string {
	if o == nil || isNil(o.VirtualIp2) {
		var ret string
		return ret
	}
	return *o.VirtualIp2
}

// GetVirtualIp2Ok returns a tuple with the VirtualIp2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject76) GetVirtualIp2Ok() (*string, bool) {
	if o == nil || isNil(o.VirtualIp2) {
    return nil, false
	}
	return o.VirtualIp2, true
}

// HasVirtualIp2 returns a boolean if a field has been set.
func (o *InlineObject76) HasVirtualIp2() bool {
	if o != nil && !isNil(o.VirtualIp2) {
		return true
	}

	return false
}

// SetVirtualIp2 gets a reference to the given string and assigns it to the VirtualIp2 field.
func (o *InlineObject76) SetVirtualIp2(v string) {
	o.VirtualIp2 = &v
}

func (o InlineObject76) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SpareSerial) {
		toSerialize["spareSerial"] = o.SpareSerial
	}
	if !isNil(o.UplinkMode) {
		toSerialize["uplinkMode"] = o.UplinkMode
	}
	if !isNil(o.VirtualIp1) {
		toSerialize["virtualIp1"] = o.VirtualIp1
	}
	if !isNil(o.VirtualIp2) {
		toSerialize["virtualIp2"] = o.VirtualIp2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject76 struct {
	value *InlineObject76
	isSet bool
}

func (v NullableInlineObject76) Get() *InlineObject76 {
	return v.value
}

func (v *NullableInlineObject76) Set(val *InlineObject76) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject76) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject76) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject76(val *InlineObject76) *NullableInlineObject76 {
	return &NullableInlineObject76{value: val, isSet: true}
}

func (v NullableInlineObject76) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject76) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


