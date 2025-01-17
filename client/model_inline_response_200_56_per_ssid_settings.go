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

// InlineResponse20056PerSsidSettings Per-SSID radio settings by number.
type InlineResponse20056PerSsidSettings struct {
	Var1 *InlineResponse20056PerSsidSettings1 `json:"1,omitempty"`
	Var2 *InlineResponse20056PerSsidSettings2 `json:"2,omitempty"`
	Var3 *InlineResponse20056PerSsidSettings3 `json:"3,omitempty"`
	Var4 *InlineResponse20056PerSsidSettings4 `json:"4,omitempty"`
}

// NewInlineResponse20056PerSsidSettings instantiates a new InlineResponse20056PerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20056PerSsidSettings() *InlineResponse20056PerSsidSettings {
	this := InlineResponse20056PerSsidSettings{}
	return &this
}

// NewInlineResponse20056PerSsidSettingsWithDefaults instantiates a new InlineResponse20056PerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20056PerSsidSettingsWithDefaults() *InlineResponse20056PerSsidSettings {
	this := InlineResponse20056PerSsidSettings{}
	return &this
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *InlineResponse20056PerSsidSettings) GetVar1() InlineResponse20056PerSsidSettings1 {
	if o == nil || isNil(o.Var1) {
		var ret InlineResponse20056PerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056PerSsidSettings) GetVar1Ok() (*InlineResponse20056PerSsidSettings1, bool) {
	if o == nil || isNil(o.Var1) {
    return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *InlineResponse20056PerSsidSettings) HasVar1() bool {
	if o != nil && !isNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given InlineResponse20056PerSsidSettings1 and assigns it to the Var1 field.
func (o *InlineResponse20056PerSsidSettings) SetVar1(v InlineResponse20056PerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *InlineResponse20056PerSsidSettings) GetVar2() InlineResponse20056PerSsidSettings2 {
	if o == nil || isNil(o.Var2) {
		var ret InlineResponse20056PerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056PerSsidSettings) GetVar2Ok() (*InlineResponse20056PerSsidSettings2, bool) {
	if o == nil || isNil(o.Var2) {
    return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *InlineResponse20056PerSsidSettings) HasVar2() bool {
	if o != nil && !isNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given InlineResponse20056PerSsidSettings2 and assigns it to the Var2 field.
func (o *InlineResponse20056PerSsidSettings) SetVar2(v InlineResponse20056PerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *InlineResponse20056PerSsidSettings) GetVar3() InlineResponse20056PerSsidSettings3 {
	if o == nil || isNil(o.Var3) {
		var ret InlineResponse20056PerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056PerSsidSettings) GetVar3Ok() (*InlineResponse20056PerSsidSettings3, bool) {
	if o == nil || isNil(o.Var3) {
    return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *InlineResponse20056PerSsidSettings) HasVar3() bool {
	if o != nil && !isNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given InlineResponse20056PerSsidSettings3 and assigns it to the Var3 field.
func (o *InlineResponse20056PerSsidSettings) SetVar3(v InlineResponse20056PerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *InlineResponse20056PerSsidSettings) GetVar4() InlineResponse20056PerSsidSettings4 {
	if o == nil || isNil(o.Var4) {
		var ret InlineResponse20056PerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056PerSsidSettings) GetVar4Ok() (*InlineResponse20056PerSsidSettings4, bool) {
	if o == nil || isNil(o.Var4) {
    return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *InlineResponse20056PerSsidSettings) HasVar4() bool {
	if o != nil && !isNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given InlineResponse20056PerSsidSettings4 and assigns it to the Var4 field.
func (o *InlineResponse20056PerSsidSettings) SetVar4(v InlineResponse20056PerSsidSettings4) {
	o.Var4 = &v
}

func (o InlineResponse20056PerSsidSettings) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20056PerSsidSettings struct {
	value *InlineResponse20056PerSsidSettings
	isSet bool
}

func (v NullableInlineResponse20056PerSsidSettings) Get() *InlineResponse20056PerSsidSettings {
	return v.value
}

func (v *NullableInlineResponse20056PerSsidSettings) Set(val *InlineResponse20056PerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20056PerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20056PerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20056PerSsidSettings(val *InlineResponse20056PerSsidSettings) *NullableInlineResponse20056PerSsidSettings {
	return &NullableInlineResponse20056PerSsidSettings{value: val, isSet: true}
}

func (v NullableInlineResponse20056PerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20056PerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


