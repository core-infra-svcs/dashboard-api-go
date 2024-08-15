/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject171 struct for InlineObject171
type InlineObject171 struct {
	// Indicates if this rule will allow, block, or alert.
	Type *string `json:"type,omitempty"`
	Match *NetworksNetworkIdWirelessAirMarshalRulesMatch `json:"match,omitempty"`
}

// NewInlineObject171 instantiates a new InlineObject171 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject171() *InlineObject171 {
	this := InlineObject171{}
	return &this
}

// NewInlineObject171WithDefaults instantiates a new InlineObject171 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject171WithDefaults() *InlineObject171 {
	this := InlineObject171{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject171) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject171) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject171) SetType(v string) {
	o.Type = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *InlineObject171) GetMatch() NetworksNetworkIdWirelessAirMarshalRulesMatch {
	if o == nil || isNil(o.Match) {
		var ret NetworksNetworkIdWirelessAirMarshalRulesMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetMatchOk() (*NetworksNetworkIdWirelessAirMarshalRulesMatch, bool) {
	if o == nil || isNil(o.Match) {
    return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *InlineObject171) HasMatch() bool {
	if o != nil && !isNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given NetworksNetworkIdWirelessAirMarshalRulesMatch and assigns it to the Match field.
func (o *InlineObject171) SetMatch(v NetworksNetworkIdWirelessAirMarshalRulesMatch) {
	o.Match = &v
}

func (o InlineObject171) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject171 struct {
	value *InlineObject171
	isSet bool
}

func (v NullableInlineObject171) Get() *InlineObject171 {
	return v.value
}

func (v *NullableInlineObject171) Set(val *InlineObject171) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject171) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject171) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject171(val *InlineObject171) *NullableInlineObject171 {
	return &NullableInlineObject171{value: val, isSet: true}
}

func (v NullableInlineObject171) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject171) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


