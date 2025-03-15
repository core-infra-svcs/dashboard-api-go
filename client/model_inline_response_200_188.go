/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200188 struct for InlineResponse200188
type InlineResponse200188 struct {
	// Desired ESL hostname of the network
	Hostname *string `json:"hostname,omitempty"`
	// Turn ESL features on and off for this network
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200188 instantiates a new InlineResponse200188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200188() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// NewInlineResponse200188WithDefaults instantiates a new InlineResponse200188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200188WithDefaults() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *InlineResponse200188) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *InlineResponse200188) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *InlineResponse200188) SetHostname(v string) {
	o.Hostname = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200188) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200188) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200188) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200188 struct {
	value *InlineResponse200188
	isSet bool
}

func (v NullableInlineResponse200188) Get() *InlineResponse200188 {
	return v.value
}

func (v *NullableInlineResponse200188) Set(val *InlineResponse200188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200188(val *InlineResponse200188) *NullableInlineResponse200188 {
	return &NullableInlineResponse200188{value: val, isSet: true}
}

func (v NullableInlineResponse200188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


