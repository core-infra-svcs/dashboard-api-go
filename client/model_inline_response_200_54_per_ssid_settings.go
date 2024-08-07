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

// InlineResponse20054PerSsidSettings Per-SSID radio settings by number.
type InlineResponse20054PerSsidSettings struct {
	Var1 *InlineResponse20054PerSsidSettings1 `json:"1,omitempty"`
	Var2 *InlineResponse20054PerSsidSettings2 `json:"2,omitempty"`
	Var3 *InlineResponse20054PerSsidSettings3 `json:"3,omitempty"`
	Var4 *InlineResponse20054PerSsidSettings4 `json:"4,omitempty"`
}

// NewInlineResponse20054PerSsidSettings instantiates a new InlineResponse20054PerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054PerSsidSettings() *InlineResponse20054PerSsidSettings {
	this := InlineResponse20054PerSsidSettings{}
	return &this
}

// NewInlineResponse20054PerSsidSettingsWithDefaults instantiates a new InlineResponse20054PerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054PerSsidSettingsWithDefaults() *InlineResponse20054PerSsidSettings {
	this := InlineResponse20054PerSsidSettings{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *InlineResponse20054PerSsidSettings) GetVar1() InlineResponse20054PerSsidSettings1 {
	if o == nil || isNil(o.Var1) {
		var ret InlineResponse20054PerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054PerSsidSettings) GetVar1Ok() (*InlineResponse20054PerSsidSettings1, bool) {
	if o == nil || isNil(o.Var1) {
    return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *InlineResponse20054PerSsidSettings) HasVar1() bool {
	if o != nil && !isNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given InlineResponse20054PerSsidSettings1 and assigns it to the Var1 field.
func (o *InlineResponse20054PerSsidSettings) SetVar1(v InlineResponse20054PerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *InlineResponse20054PerSsidSettings) GetVar2() InlineResponse20054PerSsidSettings2 {
	if o == nil || isNil(o.Var2) {
		var ret InlineResponse20054PerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054PerSsidSettings) GetVar2Ok() (*InlineResponse20054PerSsidSettings2, bool) {
	if o == nil || isNil(o.Var2) {
    return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *InlineResponse20054PerSsidSettings) HasVar2() bool {
	if o != nil && !isNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given InlineResponse20054PerSsidSettings2 and assigns it to the Var2 field.
func (o *InlineResponse20054PerSsidSettings) SetVar2(v InlineResponse20054PerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *InlineResponse20054PerSsidSettings) GetVar3() InlineResponse20054PerSsidSettings3 {
	if o == nil || isNil(o.Var3) {
		var ret InlineResponse20054PerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054PerSsidSettings) GetVar3Ok() (*InlineResponse20054PerSsidSettings3, bool) {
	if o == nil || isNil(o.Var3) {
    return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *InlineResponse20054PerSsidSettings) HasVar3() bool {
	if o != nil && !isNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given InlineResponse20054PerSsidSettings3 and assigns it to the Var3 field.
func (o *InlineResponse20054PerSsidSettings) SetVar3(v InlineResponse20054PerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *InlineResponse20054PerSsidSettings) GetVar4() InlineResponse20054PerSsidSettings4 {
	if o == nil || isNil(o.Var4) {
		var ret InlineResponse20054PerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054PerSsidSettings) GetVar4Ok() (*InlineResponse20054PerSsidSettings4, bool) {
	if o == nil || isNil(o.Var4) {
    return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *InlineResponse20054PerSsidSettings) HasVar4() bool {
	if o != nil && !isNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given InlineResponse20054PerSsidSettings4 and assigns it to the Var4 field.
func (o *InlineResponse20054PerSsidSettings) SetVar4(v InlineResponse20054PerSsidSettings4) {
	o.Var4 = &v
}

func (o InlineResponse20054PerSsidSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !isNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !isNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !isNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054PerSsidSettings struct {
	value *InlineResponse20054PerSsidSettings
	isSet bool
}

func (v NullableInlineResponse20054PerSsidSettings) Get() *InlineResponse20054PerSsidSettings {
	return v.value
}

func (v *NullableInlineResponse20054PerSsidSettings) Set(val *InlineResponse20054PerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054PerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054PerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054PerSsidSettings(val *InlineResponse20054PerSsidSettings) *NullableInlineResponse20054PerSsidSettings {
	return &NullableInlineResponse20054PerSsidSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20054PerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054PerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


