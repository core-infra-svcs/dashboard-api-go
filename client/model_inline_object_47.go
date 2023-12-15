/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject47 struct for InlineObject47
type InlineObject47 struct {
	// Set mode to 'disabled'/'detection'/'prevention' (optional - omitting will leave current config unchanged)
	Mode *string `json:"mode,omitempty"`
	// Set the detection ruleset 'connectivity'/'balanced'/'security' (optional - omitting will leave current config unchanged). Default value is 'balanced' if none currently saved
	IdsRulesets *string `json:"idsRulesets,omitempty"`
	ProtectedNetworks *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"`
}

// NewInlineObject47 instantiates a new InlineObject47 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject47() *InlineObject47 {
	this := InlineObject47{}
	return &this
}

// NewInlineObject47WithDefaults instantiates a new InlineObject47 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject47WithDefaults() *InlineObject47 {
	this := InlineObject47{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineObject47) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineObject47) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineObject47) SetMode(v string) {
	o.Mode = &v
}

// GetIdsRulesets returns the IdsRulesets field value if set, zero value otherwise.
func (o *InlineObject47) GetIdsRulesets() string {
	if o == nil || isNil(o.IdsRulesets) {
		var ret string
		return ret
	}
	return *o.IdsRulesets
}

// GetIdsRulesetsOk returns a tuple with the IdsRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetIdsRulesetsOk() (*string, bool) {
	if o == nil || isNil(o.IdsRulesets) {
    return nil, false
	}
	return o.IdsRulesets, true
}

// HasIdsRulesets returns a boolean if a field has been set.
func (o *InlineObject47) HasIdsRulesets() bool {
	if o != nil && !isNil(o.IdsRulesets) {
		return true
	}

	return false
}

// SetIdsRulesets gets a reference to the given string and assigns it to the IdsRulesets field.
func (o *InlineObject47) SetIdsRulesets(v string) {
	o.IdsRulesets = &v
}

// GetProtectedNetworks returns the ProtectedNetworks field value if set, zero value otherwise.
func (o *InlineObject47) GetProtectedNetworks() NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	if o == nil || isNil(o.ProtectedNetworks) {
		var ret NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks
		return ret
	}
	return *o.ProtectedNetworks
}

// GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject47) GetProtectedNetworksOk() (*NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks, bool) {
	if o == nil || isNil(o.ProtectedNetworks) {
    return nil, false
	}
	return o.ProtectedNetworks, true
}

// HasProtectedNetworks returns a boolean if a field has been set.
func (o *InlineObject47) HasProtectedNetworks() bool {
	if o != nil && !isNil(o.ProtectedNetworks) {
		return true
	}

	return false
}

// SetProtectedNetworks gets a reference to the given NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks and assigns it to the ProtectedNetworks field.
func (o *InlineObject47) SetProtectedNetworks(v NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) {
	o.ProtectedNetworks = &v
}

func (o InlineObject47) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.IdsRulesets) {
		toSerialize["idsRulesets"] = o.IdsRulesets
	}
	if !isNil(o.ProtectedNetworks) {
		toSerialize["protectedNetworks"] = o.ProtectedNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject47 struct {
	value *InlineObject47
	isSet bool
}

func (v NullableInlineObject47) Get() *InlineObject47 {
	return v.value
}

func (v *NullableInlineObject47) Set(val *InlineObject47) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject47) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject47) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject47(val *InlineObject47) *NullableInlineObject47 {
	return &NullableInlineObject47{value: val, isSet: true}
}

func (v NullableInlineObject47) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject47) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


