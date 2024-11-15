/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200183 struct for InlineResponse200183
type InlineResponse200183 struct {
	// Desired ESL hostname of the network
	Hostname *string `json:"hostname,omitempty"`
	// Turn ESL features on and off for this network
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200183 instantiates a new InlineResponse200183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200183() *InlineResponse200183 {
	this := InlineResponse200183{}
	return &this
}

// NewInlineResponse200183WithDefaults instantiates a new InlineResponse200183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200183WithDefaults() *InlineResponse200183 {
	this := InlineResponse200183{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *InlineResponse200183) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *InlineResponse200183) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *InlineResponse200183) SetHostname(v string) {
	o.Hostname = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200183) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200183) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200183) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200183) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200183 struct {
	value *InlineResponse200183
	isSet bool
}

func (v NullableInlineResponse200183) Get() *InlineResponse200183 {
	return v.value
}

func (v *NullableInlineResponse200183) Set(val *InlineResponse200183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200183(val *InlineResponse200183) *NullableInlineResponse200183 {
	return &NullableInlineResponse200183{value: val, isSet: true}
}

func (v NullableInlineResponse200183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


