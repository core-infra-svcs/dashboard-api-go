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

// InlineResponse2007InterfacesWan1PppoeAuthentication Settings for PPPoE Authentication.
type InlineResponse2007InterfacesWan1PppoeAuthentication struct {
	// Whether PPPoE authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Username for PPPoE authentication.
	Username *string `json:"username,omitempty"`
}

// NewInlineResponse2007InterfacesWan1PppoeAuthentication instantiates a new InlineResponse2007InterfacesWan1PppoeAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007InterfacesWan1PppoeAuthentication() *InlineResponse2007InterfacesWan1PppoeAuthentication {
	this := InlineResponse2007InterfacesWan1PppoeAuthentication{}
	return &this
}

// NewInlineResponse2007InterfacesWan1PppoeAuthenticationWithDefaults instantiates a new InlineResponse2007InterfacesWan1PppoeAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007InterfacesWan1PppoeAuthenticationWithDefaults() *InlineResponse2007InterfacesWan1PppoeAuthentication {
	this := InlineResponse2007InterfacesWan1PppoeAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse2007InterfacesWan1PppoeAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o InlineResponse2007InterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007InterfacesWan1PppoeAuthentication struct {
	value *InlineResponse2007InterfacesWan1PppoeAuthentication
	isSet bool
}

func (v NullableInlineResponse2007InterfacesWan1PppoeAuthentication) Get() *InlineResponse2007InterfacesWan1PppoeAuthentication {
	return v.value
}

func (v *NullableInlineResponse2007InterfacesWan1PppoeAuthentication) Set(val *InlineResponse2007InterfacesWan1PppoeAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007InterfacesWan1PppoeAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007InterfacesWan1PppoeAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007InterfacesWan1PppoeAuthentication(val *InlineResponse2007InterfacesWan1PppoeAuthentication) *NullableInlineResponse2007InterfacesWan1PppoeAuthentication {
	return &NullableInlineResponse2007InterfacesWan1PppoeAuthentication{value: val, isSet: true}
}

func (v NullableInlineResponse2007InterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007InterfacesWan1PppoeAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


