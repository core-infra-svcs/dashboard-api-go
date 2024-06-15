/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200143AlertsEmail Alert settings for DHCP servers
type InlineResponse200143AlertsEmail struct {
	// When enabled, send an email if a new DHCP server is seen. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200143AlertsEmail instantiates a new InlineResponse200143AlertsEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200143AlertsEmail() *InlineResponse200143AlertsEmail {
	this := InlineResponse200143AlertsEmail{}
	return &this
}

// NewInlineResponse200143AlertsEmailWithDefaults instantiates a new InlineResponse200143AlertsEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200143AlertsEmailWithDefaults() *InlineResponse200143AlertsEmail {
	this := InlineResponse200143AlertsEmail{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200143AlertsEmail) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143AlertsEmail) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200143AlertsEmail) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200143AlertsEmail) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200143AlertsEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200143AlertsEmail struct {
	value *InlineResponse200143AlertsEmail
	isSet bool
}

func (v NullableInlineResponse200143AlertsEmail) Get() *InlineResponse200143AlertsEmail {
	return v.value
}

func (v *NullableInlineResponse200143AlertsEmail) Set(val *InlineResponse200143AlertsEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200143AlertsEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200143AlertsEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200143AlertsEmail(val *InlineResponse200143AlertsEmail) *NullableInlineResponse200143AlertsEmail {
	return &NullableInlineResponse200143AlertsEmail{value: val, isSet: true}
}

func (v NullableInlineResponse200143AlertsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200143AlertsEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


