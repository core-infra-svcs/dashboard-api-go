/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject250 struct for InlineObject250
type InlineObject250 struct {
	// Short name of the early access feature
	ShortName string `json:"shortName"`
	// A list of network IDs to apply the opt-in to
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"`
}

// NewInlineObject250 instantiates a new InlineObject250 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject250(shortName string) *InlineObject250 {
	this := InlineObject250{}
	this.ShortName = shortName
	return &this
}

// NewInlineObject250WithDefaults instantiates a new InlineObject250 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject250WithDefaults() *InlineObject250 {
	this := InlineObject250{}
	return &this
}

// GetShortName returns the ShortName field value
func (o *InlineObject250) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *InlineObject250) GetShortNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *InlineObject250) SetShortName(v string) {
	o.ShortName = v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineObject250) GetLimitScopeToNetworks() []string {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []string
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject250) GetLimitScopeToNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineObject250) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []string and assigns it to the LimitScopeToNetworks field.
func (o *InlineObject250) SetLimitScopeToNetworks(v []string) {
	o.LimitScopeToNetworks = v
}

func (o InlineObject250) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject250 struct {
	value *InlineObject250
	isSet bool
}

func (v NullableInlineObject250) Get() *InlineObject250 {
	return v.value
}

func (v *NullableInlineObject250) Set(val *InlineObject250) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject250) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject250) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject250(val *InlineObject250) *NullableInlineObject250 {
	return &NullableInlineObject250{value: val, isSet: true}
}

func (v NullableInlineObject250) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject250) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


