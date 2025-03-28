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

// InlineResponse200204SentryEnrollment Systems Manager sentry enrollment splash settings.
type InlineResponse200204SentryEnrollment struct {
	SystemsManagerNetwork *InlineResponse200204SentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"`
	// The strength of the enforcement of selected system types.
	Strength *string `json:"strength,omitempty"`
	// The system types that the Sentry enforces.
	EnforcedSystems []string `json:"enforcedSystems,omitempty"`
}

// NewInlineResponse200204SentryEnrollment instantiates a new InlineResponse200204SentryEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200204SentryEnrollment() *InlineResponse200204SentryEnrollment {
	this := InlineResponse200204SentryEnrollment{}
	return &this
}

// NewInlineResponse200204SentryEnrollmentWithDefaults instantiates a new InlineResponse200204SentryEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200204SentryEnrollmentWithDefaults() *InlineResponse200204SentryEnrollment {
	this := InlineResponse200204SentryEnrollment{}
	return &this
}

// GetSystemsManagerNetwork returns the SystemsManagerNetwork field value if set, zero value otherwise.
func (o *InlineResponse200204SentryEnrollment) GetSystemsManagerNetwork() InlineResponse200204SentryEnrollmentSystemsManagerNetwork {
	if o == nil || isNil(o.SystemsManagerNetwork) {
		var ret InlineResponse200204SentryEnrollmentSystemsManagerNetwork
		return ret
	}
	return *o.SystemsManagerNetwork
}

// GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204SentryEnrollment) GetSystemsManagerNetworkOk() (*InlineResponse200204SentryEnrollmentSystemsManagerNetwork, bool) {
	if o == nil || isNil(o.SystemsManagerNetwork) {
    return nil, false
	}
	return o.SystemsManagerNetwork, true
}

// HasSystemsManagerNetwork returns a boolean if a field has been set.
func (o *InlineResponse200204SentryEnrollment) HasSystemsManagerNetwork() bool {
	if o != nil && !isNil(o.SystemsManagerNetwork) {
		return true
	}

	return false
}

// SetSystemsManagerNetwork gets a reference to the given InlineResponse200204SentryEnrollmentSystemsManagerNetwork and assigns it to the SystemsManagerNetwork field.
func (o *InlineResponse200204SentryEnrollment) SetSystemsManagerNetwork(v InlineResponse200204SentryEnrollmentSystemsManagerNetwork) {
	o.SystemsManagerNetwork = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *InlineResponse200204SentryEnrollment) GetStrength() string {
	if o == nil || isNil(o.Strength) {
		var ret string
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204SentryEnrollment) GetStrengthOk() (*string, bool) {
	if o == nil || isNil(o.Strength) {
    return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *InlineResponse200204SentryEnrollment) HasStrength() bool {
	if o != nil && !isNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given string and assigns it to the Strength field.
func (o *InlineResponse200204SentryEnrollment) SetStrength(v string) {
	o.Strength = &v
}

// GetEnforcedSystems returns the EnforcedSystems field value if set, zero value otherwise.
func (o *InlineResponse200204SentryEnrollment) GetEnforcedSystems() []string {
	if o == nil || isNil(o.EnforcedSystems) {
		var ret []string
		return ret
	}
	return o.EnforcedSystems
}

// GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200204SentryEnrollment) GetEnforcedSystemsOk() ([]string, bool) {
	if o == nil || isNil(o.EnforcedSystems) {
    return nil, false
	}
	return o.EnforcedSystems, true
}

// HasEnforcedSystems returns a boolean if a field has been set.
func (o *InlineResponse200204SentryEnrollment) HasEnforcedSystems() bool {
	if o != nil && !isNil(o.EnforcedSystems) {
		return true
	}

	return false
}

// SetEnforcedSystems gets a reference to the given []string and assigns it to the EnforcedSystems field.
func (o *InlineResponse200204SentryEnrollment) SetEnforcedSystems(v []string) {
	o.EnforcedSystems = v
}

func (o InlineResponse200204SentryEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SystemsManagerNetwork) {
		toSerialize["systemsManagerNetwork"] = o.SystemsManagerNetwork
	}
	if !isNil(o.Strength) {
		toSerialize["strength"] = o.Strength
	}
	if !isNil(o.EnforcedSystems) {
		toSerialize["enforcedSystems"] = o.EnforcedSystems
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200204SentryEnrollment struct {
	value *InlineResponse200204SentryEnrollment
	isSet bool
}

func (v NullableInlineResponse200204SentryEnrollment) Get() *InlineResponse200204SentryEnrollment {
	return v.value
}

func (v *NullableInlineResponse200204SentryEnrollment) Set(val *InlineResponse200204SentryEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200204SentryEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200204SentryEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200204SentryEnrollment(val *InlineResponse200204SentryEnrollment) *NullableInlineResponse200204SentryEnrollment {
	return &NullableInlineResponse200204SentryEnrollment{value: val, isSet: true}
}

func (v NullableInlineResponse200204SentryEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200204SentryEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


