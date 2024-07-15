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

// InlineObject170 struct for InlineObject170
type InlineObject170 struct {
	// Indicates if this rule will allow, block, or alert.
	Type string `json:"type"`
	Match NetworksNetworkIdWirelessAirMarshalRulesMatch `json:"match"`
}

// NewInlineObject170 instantiates a new InlineObject170 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject170(type_ string, match NetworksNetworkIdWirelessAirMarshalRulesMatch) *InlineObject170 {
	this := InlineObject170{}
	this.Type = type_
	this.Match = match
	return &this
}

// NewInlineObject170WithDefaults instantiates a new InlineObject170 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject170WithDefaults() *InlineObject170 {
	this := InlineObject170{}
	return &this
}

// GetType returns the Type field value
func (o *InlineObject170) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject170) SetType(v string) {
	o.Type = v
}

// GetMatch returns the Match field value
func (o *InlineObject170) GetMatch() NetworksNetworkIdWirelessAirMarshalRulesMatch {
	if o == nil {
		var ret NetworksNetworkIdWirelessAirMarshalRulesMatch
		return ret
	}

	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *InlineObject170) GetMatchOk() (*NetworksNetworkIdWirelessAirMarshalRulesMatch, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value
func (o *InlineObject170) SetMatch(v NetworksNetworkIdWirelessAirMarshalRulesMatch) {
	o.Match = v
}

func (o InlineObject170) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["match"] = o.Match
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject170 struct {
	value *InlineObject170
	isSet bool
}

func (v NullableInlineObject170) Get() *InlineObject170 {
	return v.value
}

func (v *NullableInlineObject170) Set(val *InlineObject170) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject170) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject170) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject170(val *InlineObject170) *NullableInlineObject170 {
	return &NullableInlineObject170{value: val, isSet: true}
}

func (v NullableInlineObject170) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject170) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


