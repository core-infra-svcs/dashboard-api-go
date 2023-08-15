/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject199 struct for InlineObject199
type InlineObject199 struct {
	// A list of network IDs to apply the opt-in to
	LimitScopeToNetworks []string `json:"limitScopeToNetworks,omitempty"`
}

// NewInlineObject199 instantiates a new InlineObject199 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject199() *InlineObject199 {
	this := InlineObject199{}
	return &this
}

// NewInlineObject199WithDefaults instantiates a new InlineObject199 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject199WithDefaults() *InlineObject199 {
	this := InlineObject199{}
	return &this
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineObject199) GetLimitScopeToNetworks() []string {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []string
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject199) GetLimitScopeToNetworksOk() ([]string, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineObject199) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []string and assigns it to the LimitScopeToNetworks field.
func (o *InlineObject199) SetLimitScopeToNetworks(v []string) {
	o.LimitScopeToNetworks = v
}

func (o InlineObject199) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject199 struct {
	value *InlineObject199
	isSet bool
}

func (v NullableInlineObject199) Get() *InlineObject199 {
	return v.value
}

func (v *NullableInlineObject199) Set(val *InlineObject199) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject199) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject199) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject199(val *InlineObject199) *NullableInlineObject199 {
	return &NullableInlineObject199{value: val, isSet: true}
}

func (v NullableInlineObject199) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject199) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


