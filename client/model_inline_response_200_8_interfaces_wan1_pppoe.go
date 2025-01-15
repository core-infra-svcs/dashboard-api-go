/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2008InterfacesWan1Pppoe Configuration options for PPPoE.
type InlineResponse2008InterfacesWan1Pppoe struct {
	// Whether PPPoE is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Authentication *InlineResponse2008InterfacesWan1PppoeAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse2008InterfacesWan1Pppoe instantiates a new InlineResponse2008InterfacesWan1Pppoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008InterfacesWan1Pppoe() *InlineResponse2008InterfacesWan1Pppoe {
	this := InlineResponse2008InterfacesWan1Pppoe{}
	return &this
}

// NewInlineResponse2008InterfacesWan1PppoeWithDefaults instantiates a new InlineResponse2008InterfacesWan1Pppoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008InterfacesWan1PppoeWithDefaults() *InlineResponse2008InterfacesWan1Pppoe {
	this := InlineResponse2008InterfacesWan1Pppoe{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1Pppoe) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1Pppoe) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1Pppoe) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2008InterfacesWan1Pppoe) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse2008InterfacesWan1Pppoe) GetAuthentication() InlineResponse2008InterfacesWan1PppoeAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret InlineResponse2008InterfacesWan1PppoeAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008InterfacesWan1Pppoe) GetAuthenticationOk() (*InlineResponse2008InterfacesWan1PppoeAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse2008InterfacesWan1Pppoe) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given InlineResponse2008InterfacesWan1PppoeAuthentication and assigns it to the Authentication field.
func (o *InlineResponse2008InterfacesWan1Pppoe) SetAuthentication(v InlineResponse2008InterfacesWan1PppoeAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse2008InterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008InterfacesWan1Pppoe struct {
	value *InlineResponse2008InterfacesWan1Pppoe
	isSet bool
}

func (v NullableInlineResponse2008InterfacesWan1Pppoe) Get() *InlineResponse2008InterfacesWan1Pppoe {
	return v.value
}

func (v *NullableInlineResponse2008InterfacesWan1Pppoe) Set(val *InlineResponse2008InterfacesWan1Pppoe) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008InterfacesWan1Pppoe) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008InterfacesWan1Pppoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008InterfacesWan1Pppoe(val *InlineResponse2008InterfacesWan1Pppoe) *NullableInlineResponse2008InterfacesWan1Pppoe {
	return &NullableInlineResponse2008InterfacesWan1Pppoe{value: val, isSet: true}
}

func (v NullableInlineResponse2008InterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008InterfacesWan1Pppoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


