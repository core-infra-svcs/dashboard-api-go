/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject59 struct for InlineObject59
type InlineObject59 struct {
	// Set mode to 'disabled'/'detection'/'prevention' (optional - omitting will leave current config unchanged)
	Mode *string `json:"mode,omitempty"`
	// Set the detection ruleset 'connectivity'/'balanced'/'security' (optional - omitting will leave current config unchanged). Default value is 'balanced' if none currently saved
	IdsRulesets *string `json:"idsRulesets,omitempty"`
	ProtectedNetworks *NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks `json:"protectedNetworks,omitempty"`
}

// NewInlineObject59 instantiates a new InlineObject59 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject59() *InlineObject59 {
	this := InlineObject59{}
	return &this
}

// NewInlineObject59WithDefaults instantiates a new InlineObject59 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject59WithDefaults() *InlineObject59 {
	this := InlineObject59{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineObject59) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineObject59) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineObject59) SetMode(v string) {
	o.Mode = &v
}

// GetIdsRulesets returns the IdsRulesets field value if set, zero value otherwise.
func (o *InlineObject59) GetIdsRulesets() string {
	if o == nil || isNil(o.IdsRulesets) {
		var ret string
		return ret
	}
	return *o.IdsRulesets
}

// GetIdsRulesetsOk returns a tuple with the IdsRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetIdsRulesetsOk() (*string, bool) {
	if o == nil || isNil(o.IdsRulesets) {
    return nil, false
	}
	return o.IdsRulesets, true
}

// HasIdsRulesets returns a boolean if a field has been set.
func (o *InlineObject59) HasIdsRulesets() bool {
	if o != nil && !isNil(o.IdsRulesets) {
		return true
	}

	return false
}

// SetIdsRulesets gets a reference to the given string and assigns it to the IdsRulesets field.
func (o *InlineObject59) SetIdsRulesets(v string) {
	o.IdsRulesets = &v
}

// GetProtectedNetworks returns the ProtectedNetworks field value if set, zero value otherwise.
func (o *InlineObject59) GetProtectedNetworks() NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks {
	if o == nil || isNil(o.ProtectedNetworks) {
		var ret NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks
		return ret
	}
	return *o.ProtectedNetworks
}

// GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetProtectedNetworksOk() (*NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks, bool) {
	if o == nil || isNil(o.ProtectedNetworks) {
    return nil, false
	}
	return o.ProtectedNetworks, true
}

// HasProtectedNetworks returns a boolean if a field has been set.
func (o *InlineObject59) HasProtectedNetworks() bool {
	if o != nil && !isNil(o.ProtectedNetworks) {
		return true
	}

	return false
}

// SetProtectedNetworks gets a reference to the given NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks and assigns it to the ProtectedNetworks field.
func (o *InlineObject59) SetProtectedNetworks(v NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks) {
	o.ProtectedNetworks = &v
}

func (o InlineObject59) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject59 struct {
	value *InlineObject59
	isSet bool
}

func (v NullableInlineObject59) Get() *InlineObject59 {
	return v.value
}

func (v *NullableInlineObject59) Set(val *InlineObject59) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject59) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject59) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject59(val *InlineObject59) *NullableInlineObject59 {
	return &NullableInlineObject59{value: val, isSet: true}
}

func (v NullableInlineObject59) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject59) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


