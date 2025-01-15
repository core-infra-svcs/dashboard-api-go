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

// InlineResponse200200SelfRegistration Self-registration for splash with Meraki authentication.
type InlineResponse200200SelfRegistration struct {
	// Whether or not to allow users to create their own account on the network.
	Enabled *bool `json:"enabled,omitempty"`
	// How created user accounts should be authorized.
	AuthorizationType *string `json:"authorizationType,omitempty"`
}

// NewInlineResponse200200SelfRegistration instantiates a new InlineResponse200200SelfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200SelfRegistration() *InlineResponse200200SelfRegistration {
	this := InlineResponse200200SelfRegistration{}
	return &this
}

// NewInlineResponse200200SelfRegistrationWithDefaults instantiates a new InlineResponse200200SelfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200SelfRegistrationWithDefaults() *InlineResponse200200SelfRegistration {
	this := InlineResponse200200SelfRegistration{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200200SelfRegistration) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200SelfRegistration) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200200SelfRegistration) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200200SelfRegistration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthorizationType returns the AuthorizationType field value if set, zero value otherwise.
func (o *InlineResponse200200SelfRegistration) GetAuthorizationType() string {
	if o == nil || isNil(o.AuthorizationType) {
		var ret string
		return ret
	}
	return *o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200SelfRegistration) GetAuthorizationTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationType) {
    return nil, false
	}
	return o.AuthorizationType, true
}

// HasAuthorizationType returns a boolean if a field has been set.
func (o *InlineResponse200200SelfRegistration) HasAuthorizationType() bool {
	if o != nil && !isNil(o.AuthorizationType) {
		return true
	}

	return false
}

// SetAuthorizationType gets a reference to the given string and assigns it to the AuthorizationType field.
func (o *InlineResponse200200SelfRegistration) SetAuthorizationType(v string) {
	o.AuthorizationType = &v
}

func (o InlineResponse200200SelfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.AuthorizationType) {
		toSerialize["authorizationType"] = o.AuthorizationType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200SelfRegistration struct {
	value *InlineResponse200200SelfRegistration
	isSet bool
}

func (v NullableInlineResponse200200SelfRegistration) Get() *InlineResponse200200SelfRegistration {
	return v.value
}

func (v *NullableInlineResponse200200SelfRegistration) Set(val *InlineResponse200200SelfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200SelfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200SelfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200SelfRegistration(val *InlineResponse200200SelfRegistration) *NullableInlineResponse200200SelfRegistration {
	return &NullableInlineResponse200200SelfRegistration{value: val, isSet: true}
}

func (v NullableInlineResponse200200SelfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200SelfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

