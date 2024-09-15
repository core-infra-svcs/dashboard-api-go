/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20057 struct for InlineResponse20057
type InlineResponse20057 struct {
	// Intrusion detection mode
	Mode *string `json:"mode,omitempty"`
	// Intrusion detection ruleset
	IdsRulesets *string `json:"idsRulesets,omitempty"`
	ProtectedNetworks *InlineResponse20057ProtectedNetworks `json:"protectedNetworks,omitempty"`
}

// NewInlineResponse20057 instantiates a new InlineResponse20057 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// NewInlineResponse20057WithDefaults instantiates a new InlineResponse20057 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057WithDefaults() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse20057) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse20057) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse20057) SetMode(v string) {
	o.Mode = &v
}

// GetIdsRulesets returns the IdsRulesets field value if set, zero value otherwise.
func (o *InlineResponse20057) GetIdsRulesets() string {
	if o == nil || isNil(o.IdsRulesets) {
		var ret string
		return ret
	}
	return *o.IdsRulesets
}

// GetIdsRulesetsOk returns a tuple with the IdsRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetIdsRulesetsOk() (*string, bool) {
	if o == nil || isNil(o.IdsRulesets) {
    return nil, false
	}
	return o.IdsRulesets, true
}

// HasIdsRulesets returns a boolean if a field has been set.
func (o *InlineResponse20057) HasIdsRulesets() bool {
	if o != nil && !isNil(o.IdsRulesets) {
		return true
	}

	return false
}

// SetIdsRulesets gets a reference to the given string and assigns it to the IdsRulesets field.
func (o *InlineResponse20057) SetIdsRulesets(v string) {
	o.IdsRulesets = &v
}

// GetProtectedNetworks returns the ProtectedNetworks field value if set, zero value otherwise.
func (o *InlineResponse20057) GetProtectedNetworks() InlineResponse20057ProtectedNetworks {
	if o == nil || isNil(o.ProtectedNetworks) {
		var ret InlineResponse20057ProtectedNetworks
		return ret
	}
	return *o.ProtectedNetworks
}

// GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetProtectedNetworksOk() (*InlineResponse20057ProtectedNetworks, bool) {
	if o == nil || isNil(o.ProtectedNetworks) {
    return nil, false
	}
	return o.ProtectedNetworks, true
}

// HasProtectedNetworks returns a boolean if a field has been set.
func (o *InlineResponse20057) HasProtectedNetworks() bool {
	if o != nil && !isNil(o.ProtectedNetworks) {
		return true
	}

	return false
}

// SetProtectedNetworks gets a reference to the given InlineResponse20057ProtectedNetworks and assigns it to the ProtectedNetworks field.
func (o *InlineResponse20057) SetProtectedNetworks(v InlineResponse20057ProtectedNetworks) {
	o.ProtectedNetworks = &v
}

func (o InlineResponse20057) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20057 struct {
	value *InlineResponse20057
	isSet bool
}

func (v NullableInlineResponse20057) Get() *InlineResponse20057 {
	return v.value
}

func (v *NullableInlineResponse20057) Set(val *InlineResponse20057) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057(val *InlineResponse20057) *NullableInlineResponse20057 {
	return &NullableInlineResponse20057{value: val, isSet: true}
}

func (v NullableInlineResponse20057) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


