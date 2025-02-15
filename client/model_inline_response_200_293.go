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

// InlineResponse200293 struct for InlineResponse200293
type InlineResponse200293 struct {
	// Toggle depicting if SAML SSO settings are enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200293 instantiates a new InlineResponse200293 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200293() *InlineResponse200293 {
	this := InlineResponse200293{}
	return &this
}

// NewInlineResponse200293WithDefaults instantiates a new InlineResponse200293 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200293WithDefaults() *InlineResponse200293 {
	this := InlineResponse200293{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200293) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200293) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200293) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200293) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200293) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200293 struct {
	value *InlineResponse200293
	isSet bool
}

func (v NullableInlineResponse200293) Get() *InlineResponse200293 {
	return v.value
}

func (v *NullableInlineResponse200293) Set(val *InlineResponse200293) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200293) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200293) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200293(val *InlineResponse200293) *NullableInlineResponse200293 {
	return &NullableInlineResponse200293{value: val, isSet: true}
}

func (v NullableInlineResponse200293) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200293) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


