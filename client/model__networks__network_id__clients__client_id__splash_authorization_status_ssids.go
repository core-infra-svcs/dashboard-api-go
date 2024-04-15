/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids The target SSIDs. Each SSID must be enabled and must have Click-through splash enabled. For each SSID where isAuthorized is true, the expiration time will automatically be set according to the SSID's splash frequency. Not all networks support configuring all SSIDs
type NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids struct {
	Var0 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0 `json:"0,omitempty"`
	Var1 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1 `json:"1,omitempty"`
	Var2 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2 `json:"2,omitempty"`
	Var3 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3 `json:"3,omitempty"`
	Var4 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4 `json:"4,omitempty"`
	Var5 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5 `json:"5,omitempty"`
	Var6 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6 `json:"6,omitempty"`
	Var7 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 `json:"7,omitempty"`
	Var8 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8 `json:"8,omitempty"`
	Var9 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9 `json:"9,omitempty"`
	Var10 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10 `json:"10,omitempty"`
	Var11 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11 `json:"11,omitempty"`
	Var12 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12 `json:"12,omitempty"`
	Var13 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 `json:"13,omitempty"`
	Var14 *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 `json:"14,omitempty"`
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids{}
	return &this
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsidsWithDefaults instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsidsWithDefaults() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar0() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0 {
	if o == nil || isNil(o.Var0) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar0Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0, bool) {
	if o == nil || isNil(o.Var0) {
    return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar0() bool {
	if o != nil && !isNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0 and assigns it to the Var0 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar0(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar1() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1 {
	if o == nil || isNil(o.Var1) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar1Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1, bool) {
	if o == nil || isNil(o.Var1) {
    return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar1() bool {
	if o != nil && !isNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1 and assigns it to the Var1 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar1(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar2() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2 {
	if o == nil || isNil(o.Var2) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar2Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2, bool) {
	if o == nil || isNil(o.Var2) {
    return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar2() bool {
	if o != nil && !isNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2 and assigns it to the Var2 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar2(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar3() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3 {
	if o == nil || isNil(o.Var3) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar3Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3, bool) {
	if o == nil || isNil(o.Var3) {
    return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar3() bool {
	if o != nil && !isNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3 and assigns it to the Var3 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar3(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar4() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4 {
	if o == nil || isNil(o.Var4) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar4Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4, bool) {
	if o == nil || isNil(o.Var4) {
    return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar4() bool {
	if o != nil && !isNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4 and assigns it to the Var4 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar4(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids4) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar5() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5 {
	if o == nil || isNil(o.Var5) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar5Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5, bool) {
	if o == nil || isNil(o.Var5) {
    return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar5() bool {
	if o != nil && !isNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5 and assigns it to the Var5 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar5(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids5) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar6() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6 {
	if o == nil || isNil(o.Var6) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar6Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6, bool) {
	if o == nil || isNil(o.Var6) {
    return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar6() bool {
	if o != nil && !isNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6 and assigns it to the Var6 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar6(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids6) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar7() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 {
	if o == nil || isNil(o.Var7) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar7Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7, bool) {
	if o == nil || isNil(o.Var7) {
    return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar7() bool {
	if o != nil && !isNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7 and assigns it to the Var7 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar7(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids7) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar8() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8 {
	if o == nil || isNil(o.Var8) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar8Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8, bool) {
	if o == nil || isNil(o.Var8) {
    return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar8() bool {
	if o != nil && !isNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8 and assigns it to the Var8 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar8(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids8) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar9() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9 {
	if o == nil || isNil(o.Var9) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar9Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9, bool) {
	if o == nil || isNil(o.Var9) {
    return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar9() bool {
	if o != nil && !isNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9 and assigns it to the Var9 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar9(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids9) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar10() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10 {
	if o == nil || isNil(o.Var10) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar10Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10, bool) {
	if o == nil || isNil(o.Var10) {
    return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar10() bool {
	if o != nil && !isNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10 and assigns it to the Var10 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar10(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids10) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar11() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11 {
	if o == nil || isNil(o.Var11) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar11Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11, bool) {
	if o == nil || isNil(o.Var11) {
    return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar11() bool {
	if o != nil && !isNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11 and assigns it to the Var11 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar11(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids11) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar12() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12 {
	if o == nil || isNil(o.Var12) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar12Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12, bool) {
	if o == nil || isNil(o.Var12) {
    return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar12() bool {
	if o != nil && !isNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12 and assigns it to the Var12 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar12(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids12) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar13() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 {
	if o == nil || isNil(o.Var13) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar13Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13, bool) {
	if o == nil || isNil(o.Var13) {
    return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar13() bool {
	if o != nil && !isNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 and assigns it to the Var13 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar13(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar14() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 {
	if o == nil || isNil(o.Var14) {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) GetVar14Ok() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14, bool) {
	if o == nil || isNil(o.Var14) {
    return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) HasVar14() bool {
	if o != nil && !isNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14 and assigns it to the Var14 field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) SetVar14(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids14) {
	o.Var14 = &v
}

func (o NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids struct {
	value *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
	isSet bool
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) Get() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) Set(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	return &NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


