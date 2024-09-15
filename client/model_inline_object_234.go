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

// InlineObject234 struct for InlineObject234
type InlineObject234 struct {
	// Short name of the early access feature
	ShortName string `json:"shortName"`
	// A list of network IDs to apply the opt-in to
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"`
}

// NewInlineObject234 instantiates a new InlineObject234 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject234(shortName string) *InlineObject234 {
	this := InlineObject234{}
	this.ShortName = shortName
	return &this
}

// NewInlineObject234WithDefaults instantiates a new InlineObject234 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject234WithDefaults() *InlineObject234 {
	this := InlineObject234{}
	return &this
}

// GetShortName returns the ShortName field value
func (o *InlineObject234) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *InlineObject234) GetShortNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *InlineObject234) SetShortName(v string) {
	o.ShortName = v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineObject234) GetLimitScopeToNetworks() []string {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []string
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject234) GetLimitScopeToNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineObject234) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []string and assigns it to the LimitScopeToNetworks field.
func (o *InlineObject234) SetLimitScopeToNetworks(v []string) {
	o.LimitScopeToNetworks = v
}

func (o InlineObject234) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject234 struct {
	value *InlineObject234
	isSet bool
}

func (v NullableInlineObject234) Get() *InlineObject234 {
	return v.value
}

func (v *NullableInlineObject234) Set(val *InlineObject234) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject234) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject234) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject234(val *InlineObject234) *NullableInlineObject234 {
	return &NullableInlineObject234{value: val, isSet: true}
}

func (v NullableInlineObject234) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject234) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


