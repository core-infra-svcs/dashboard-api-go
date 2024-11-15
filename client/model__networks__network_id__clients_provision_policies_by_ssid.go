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

// NetworksNetworkIdClientsProvisionPoliciesBySsid An object, describing the policy-connection associations for each active SSID within the network. Keys should be the number of enabled SSIDs, mapping to an object describing the client's policy
type NetworksNetworkIdClientsProvisionPoliciesBySsid struct {
	Var0 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"0,omitempty"`
	Var1 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"1,omitempty"`
	Var2 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"2,omitempty"`
	Var3 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"3,omitempty"`
	Var4 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"4,omitempty"`
	Var5 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"5,omitempty"`
	Var6 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"6,omitempty"`
	Var7 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"7,omitempty"`
	Var8 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"8,omitempty"`
	Var9 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"9,omitempty"`
	Var10 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"10,omitempty"`
	Var11 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"11,omitempty"`
	Var12 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"12,omitempty"`
	Var13 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"13,omitempty"`
	Var14 *NetworksNetworkIdClientsProvisionPoliciesBySsid0 `json:"14,omitempty"`
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySsid instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsProvisionPoliciesBySsid() *NetworksNetworkIdClientsProvisionPoliciesBySsid {
	this := NetworksNetworkIdClientsProvisionPoliciesBySsid{}
	return &this
}

// NewNetworksNetworkIdClientsProvisionPoliciesBySsidWithDefaults instantiates a new NetworksNetworkIdClientsProvisionPoliciesBySsid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsProvisionPoliciesBySsidWithDefaults() *NetworksNetworkIdClientsProvisionPoliciesBySsid {
	this := NetworksNetworkIdClientsProvisionPoliciesBySsid{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar0() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var0) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar0Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var0) {
    return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar0() bool {
	if o != nil && !isNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var0 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar0(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar1() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var1) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar1Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var1) {
    return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar1() bool {
	if o != nil && !isNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var1 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar1(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar2() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var2) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar2Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var2) {
    return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar2() bool {
	if o != nil && !isNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var2 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar2(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar3() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var3) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar3Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var3) {
    return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar3() bool {
	if o != nil && !isNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var3 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar3(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar4() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var4) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar4Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var4) {
    return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar4() bool {
	if o != nil && !isNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var4 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar4(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar5() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var5) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar5Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var5) {
    return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar5() bool {
	if o != nil && !isNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var5 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar5(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar6() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var6) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar6Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var6) {
    return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar6() bool {
	if o != nil && !isNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var6 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar6(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar7() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var7) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar7Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var7) {
    return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar7() bool {
	if o != nil && !isNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var7 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar7(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar8() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var8) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar8Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var8) {
    return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar8() bool {
	if o != nil && !isNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var8 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar8(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar9() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var9) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar9Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var9) {
    return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar9() bool {
	if o != nil && !isNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var9 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar9(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar10() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var10) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar10Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var10) {
    return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar10() bool {
	if o != nil && !isNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var10 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar10(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar11() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var11) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar11Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var11) {
    return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar11() bool {
	if o != nil && !isNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var11 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar11(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar12() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var12) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar12Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var12) {
    return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar12() bool {
	if o != nil && !isNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var12 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar12(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar13() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var13) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar13Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var13) {
    return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar13() bool {
	if o != nil && !isNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var13 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar13(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar14() NetworksNetworkIdClientsProvisionPoliciesBySsid0 {
	if o == nil || isNil(o.Var14) {
		var ret NetworksNetworkIdClientsProvisionPoliciesBySsid0
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) GetVar14Ok() (*NetworksNetworkIdClientsProvisionPoliciesBySsid0, bool) {
	if o == nil || isNil(o.Var14) {
    return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) HasVar14() bool {
	if o != nil && !isNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given NetworksNetworkIdClientsProvisionPoliciesBySsid0 and assigns it to the Var14 field.
func (o *NetworksNetworkIdClientsProvisionPoliciesBySsid) SetVar14(v NetworksNetworkIdClientsProvisionPoliciesBySsid0) {
	o.Var14 = &v
}

func (o NetworksNetworkIdClientsProvisionPoliciesBySsid) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdClientsProvisionPoliciesBySsid struct {
	value *NetworksNetworkIdClientsProvisionPoliciesBySsid
	isSet bool
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) Get() *NetworksNetworkIdClientsProvisionPoliciesBySsid {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) Set(val *NetworksNetworkIdClientsProvisionPoliciesBySsid) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsProvisionPoliciesBySsid(val *NetworksNetworkIdClientsProvisionPoliciesBySsid) *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid {
	return &NullableNetworksNetworkIdClientsProvisionPoliciesBySsid{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsProvisionPoliciesBySsid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


