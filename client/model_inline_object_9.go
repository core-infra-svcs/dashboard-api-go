/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	// An array of port forwarding params
	Rules *[]DevicesSerialCellularGatewayPortForwardingRulesRules `json:"rules,omitempty"`
}

// NewInlineObject9 instantiates a new InlineObject9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// NewInlineObject9WithDefaults instantiates a new InlineObject9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9WithDefaults() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject9) GetRules() []DevicesSerialCellularGatewayPortForwardingRulesRules {
	if o == nil || o.Rules == nil {
		var ret []DevicesSerialCellularGatewayPortForwardingRulesRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetRulesOk() (*[]DevicesSerialCellularGatewayPortForwardingRulesRules, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject9) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []DevicesSerialCellularGatewayPortForwardingRulesRules and assigns it to the Rules field.
func (o *InlineObject9) SetRules(v []DevicesSerialCellularGatewayPortForwardingRulesRules) {
	o.Rules = &v
}

func (o InlineObject9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject9 struct {
	value *InlineObject9
	isSet bool
}

func (v NullableInlineObject9) Get() *InlineObject9 {
	return v.value
}

func (v *NullableInlineObject9) Set(val *InlineObject9) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9(val *InlineObject9) *NullableInlineObject9 {
	return &NullableInlineObject9{value: val, isSet: true}
}

func (v NullableInlineObject9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


