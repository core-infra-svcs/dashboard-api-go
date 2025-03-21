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

// InlineResponse20061 struct for InlineResponse20061
type InlineResponse20061 struct {
	// Intrusion detection mode
	Mode *string `json:"mode,omitempty"`
	// Intrusion detection ruleset
	IdsRulesets *string `json:"idsRulesets,omitempty"`
	ProtectedNetworks *InlineResponse20061ProtectedNetworks `json:"protectedNetworks,omitempty"`
}

// NewInlineResponse20061 instantiates a new InlineResponse20061 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20061() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// NewInlineResponse20061WithDefaults instantiates a new InlineResponse20061 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20061WithDefaults() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse20061) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse20061) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse20061) SetMode(v string) {
	o.Mode = &v
}

// GetIdsRulesets returns the IdsRulesets field value if set, zero value otherwise.
func (o *InlineResponse20061) GetIdsRulesets() string {
	if o == nil || isNil(o.IdsRulesets) {
		var ret string
		return ret
	}
	return *o.IdsRulesets
}

// GetIdsRulesetsOk returns a tuple with the IdsRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetIdsRulesetsOk() (*string, bool) {
	if o == nil || isNil(o.IdsRulesets) {
    return nil, false
	}
	return o.IdsRulesets, true
}

// HasIdsRulesets returns a boolean if a field has been set.
func (o *InlineResponse20061) HasIdsRulesets() bool {
	if o != nil && !isNil(o.IdsRulesets) {
		return true
	}

	return false
}

// SetIdsRulesets gets a reference to the given string and assigns it to the IdsRulesets field.
func (o *InlineResponse20061) SetIdsRulesets(v string) {
	o.IdsRulesets = &v
}

// GetProtectedNetworks returns the ProtectedNetworks field value if set, zero value otherwise.
func (o *InlineResponse20061) GetProtectedNetworks() InlineResponse20061ProtectedNetworks {
	if o == nil || isNil(o.ProtectedNetworks) {
		var ret InlineResponse20061ProtectedNetworks
		return ret
	}
	return *o.ProtectedNetworks
}

// GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetProtectedNetworksOk() (*InlineResponse20061ProtectedNetworks, bool) {
	if o == nil || isNil(o.ProtectedNetworks) {
    return nil, false
	}
	return o.ProtectedNetworks, true
}

// HasProtectedNetworks returns a boolean if a field has been set.
func (o *InlineResponse20061) HasProtectedNetworks() bool {
	if o != nil && !isNil(o.ProtectedNetworks) {
		return true
	}

	return false
}

// SetProtectedNetworks gets a reference to the given InlineResponse20061ProtectedNetworks and assigns it to the ProtectedNetworks field.
func (o *InlineResponse20061) SetProtectedNetworks(v InlineResponse20061ProtectedNetworks) {
	o.ProtectedNetworks = &v
}

func (o InlineResponse20061) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20061 struct {
	value *InlineResponse20061
	isSet bool
}

func (v NullableInlineResponse20061) Get() *InlineResponse20061 {
	return v.value
}

func (v *NullableInlineResponse20061) Set(val *InlineResponse20061) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20061) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20061) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20061(val *InlineResponse20061) *NullableInlineResponse20061 {
	return &NullableInlineResponse20061{value: val, isSet: true}
}

func (v NullableInlineResponse20061) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20061) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


