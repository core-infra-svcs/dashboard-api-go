/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesPerSsidSettings Per-SSID radio settings by number.
type NetworksNetworkIdWirelessRfProfilesPerSsidSettings struct {
	Var0 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings0 `json:"0,omitempty"`
	Var1 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings1 `json:"1,omitempty"`
	Var2 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings2 `json:"2,omitempty"`
	Var3 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings3 `json:"3,omitempty"`
	Var4 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings4 `json:"4,omitempty"`
	Var5 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings5 `json:"5,omitempty"`
	Var6 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 `json:"6,omitempty"`
	Var7 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings7 `json:"7,omitempty"`
	Var8 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings8 `json:"8,omitempty"`
	Var9 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings9 `json:"9,omitempty"`
	Var10 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings10 `json:"10,omitempty"`
	Var11 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 `json:"11,omitempty"`
	Var12 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings12 `json:"12,omitempty"`
	Var13 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings13 `json:"13,omitempty"`
	Var14 *NetworksNetworkIdWirelessRfProfilesPerSsidSettings14 `json:"14,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettingsWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettingsWithDefaults() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar0() NetworksNetworkIdWirelessRfProfilesPerSsidSettings0 {
	if o == nil || isNil(o.Var0) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar0Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings0, bool) {
	if o == nil || isNil(o.Var0) {
    return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar0() bool {
	if o != nil && !isNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings0 and assigns it to the Var0 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar0(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar1() NetworksNetworkIdWirelessRfProfilesPerSsidSettings1 {
	if o == nil || isNil(o.Var1) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar1Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings1, bool) {
	if o == nil || isNil(o.Var1) {
    return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar1() bool {
	if o != nil && !isNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings1 and assigns it to the Var1 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar1(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar2() NetworksNetworkIdWirelessRfProfilesPerSsidSettings2 {
	if o == nil || isNil(o.Var2) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar2Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings2, bool) {
	if o == nil || isNil(o.Var2) {
    return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar2() bool {
	if o != nil && !isNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings2 and assigns it to the Var2 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar2(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar3() NetworksNetworkIdWirelessRfProfilesPerSsidSettings3 {
	if o == nil || isNil(o.Var3) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar3Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings3, bool) {
	if o == nil || isNil(o.Var3) {
    return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar3() bool {
	if o != nil && !isNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings3 and assigns it to the Var3 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar3(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar4() NetworksNetworkIdWirelessRfProfilesPerSsidSettings4 {
	if o == nil || isNil(o.Var4) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar4Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings4, bool) {
	if o == nil || isNil(o.Var4) {
    return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar4() bool {
	if o != nil && !isNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings4 and assigns it to the Var4 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar4(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings4) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar5() NetworksNetworkIdWirelessRfProfilesPerSsidSettings5 {
	if o == nil || isNil(o.Var5) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings5
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar5Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings5, bool) {
	if o == nil || isNil(o.Var5) {
    return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar5() bool {
	if o != nil && !isNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings5 and assigns it to the Var5 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar5(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings5) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar6() NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 {
	if o == nil || isNil(o.Var6) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings6
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar6Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings6, bool) {
	if o == nil || isNil(o.Var6) {
    return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar6() bool {
	if o != nil && !isNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 and assigns it to the Var6 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar6(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar7() NetworksNetworkIdWirelessRfProfilesPerSsidSettings7 {
	if o == nil || isNil(o.Var7) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings7
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar7Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings7, bool) {
	if o == nil || isNil(o.Var7) {
    return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar7() bool {
	if o != nil && !isNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings7 and assigns it to the Var7 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar7(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings7) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar8() NetworksNetworkIdWirelessRfProfilesPerSsidSettings8 {
	if o == nil || isNil(o.Var8) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings8
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar8Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings8, bool) {
	if o == nil || isNil(o.Var8) {
    return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar8() bool {
	if o != nil && !isNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings8 and assigns it to the Var8 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar8(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings8) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar9() NetworksNetworkIdWirelessRfProfilesPerSsidSettings9 {
	if o == nil || isNil(o.Var9) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings9
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar9Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings9, bool) {
	if o == nil || isNil(o.Var9) {
    return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar9() bool {
	if o != nil && !isNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings9 and assigns it to the Var9 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar9(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings9) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar10() NetworksNetworkIdWirelessRfProfilesPerSsidSettings10 {
	if o == nil || isNil(o.Var10) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings10
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar10Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings10, bool) {
	if o == nil || isNil(o.Var10) {
    return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar10() bool {
	if o != nil && !isNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings10 and assigns it to the Var10 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar10(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings10) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar11() NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 {
	if o == nil || isNil(o.Var11) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings11
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar11Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings11, bool) {
	if o == nil || isNil(o.Var11) {
    return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar11() bool {
	if o != nil && !isNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 and assigns it to the Var11 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar11(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar12() NetworksNetworkIdWirelessRfProfilesPerSsidSettings12 {
	if o == nil || isNil(o.Var12) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings12
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar12Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings12, bool) {
	if o == nil || isNil(o.Var12) {
    return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar12() bool {
	if o != nil && !isNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings12 and assigns it to the Var12 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar12(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings12) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar13() NetworksNetworkIdWirelessRfProfilesPerSsidSettings13 {
	if o == nil || isNil(o.Var13) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings13
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar13Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings13, bool) {
	if o == nil || isNil(o.Var13) {
    return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar13() bool {
	if o != nil && !isNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings13 and assigns it to the Var13 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar13(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings13) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar14() NetworksNetworkIdWirelessRfProfilesPerSsidSettings14 {
	if o == nil || isNil(o.Var14) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings14
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) GetVar14Ok() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings14, bool) {
	if o == nil || isNil(o.Var14) {
    return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) HasVar14() bool {
	if o != nil && !isNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings14 and assigns it to the Var14 field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) SetVar14(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings14) {
	o.Var14 = &v
}

func (o NetworksNetworkIdWirelessRfProfilesPerSsidSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
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
	if !isNil(o.Var5) {
		toSerialize["5"] = o.Var5
	}
	if !isNil(o.Var6) {
		toSerialize["6"] = o.Var6
	}
	if !isNil(o.Var7) {
		toSerialize["7"] = o.Var7
	}
	if !isNil(o.Var8) {
		toSerialize["8"] = o.Var8
	}
	if !isNil(o.Var9) {
		toSerialize["9"] = o.Var9
	}
	if !isNil(o.Var10) {
		toSerialize["10"] = o.Var10
	}
	if !isNil(o.Var11) {
		toSerialize["11"] = o.Var11
	}
	if !isNil(o.Var12) {
		toSerialize["12"] = o.Var12
	}
	if !isNil(o.Var13) {
		toSerialize["13"] = o.Var13
	}
	if !isNil(o.Var14) {
		toSerialize["14"] = o.Var14
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings struct {
	value *NetworksNetworkIdWirelessRfProfilesPerSsidSettings
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) Get() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) Set(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings) *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	return &NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


