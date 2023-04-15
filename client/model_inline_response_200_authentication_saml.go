/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200AuthenticationSaml SAML authentication
type InlineResponse200AuthenticationSaml struct {
	// If SAML authentication is enabled for this user
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200AuthenticationSaml instantiates a new InlineResponse200AuthenticationSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200AuthenticationSaml() *InlineResponse200AuthenticationSaml {
	this := InlineResponse200AuthenticationSaml{}
	return &this
}

// NewInlineResponse200AuthenticationSamlWithDefaults instantiates a new InlineResponse200AuthenticationSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200AuthenticationSamlWithDefaults() *InlineResponse200AuthenticationSaml {
	this := InlineResponse200AuthenticationSaml{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200AuthenticationSaml) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200AuthenticationSaml) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200AuthenticationSaml) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200AuthenticationSaml) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200AuthenticationSaml) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200AuthenticationSaml struct {
	value *InlineResponse200AuthenticationSaml
	isSet bool
}

func (v NullableInlineResponse200AuthenticationSaml) Get() *InlineResponse200AuthenticationSaml {
	return v.value
}

func (v *NullableInlineResponse200AuthenticationSaml) Set(val *InlineResponse200AuthenticationSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200AuthenticationSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200AuthenticationSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200AuthenticationSaml(val *InlineResponse200AuthenticationSaml) *NullableInlineResponse200AuthenticationSaml {
	return &NullableInlineResponse200AuthenticationSaml{value: val, isSet: true}
}

func (v NullableInlineResponse200AuthenticationSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200AuthenticationSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


