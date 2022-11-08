/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20066SentryEnrollment Systems Manager sentry enrollment splash settings.
type InlineResponse20066SentryEnrollment struct {
	SystemsManagerNetwork *InlineResponse20066SentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"`
	// The strength of the enforcement of selected system types.
	Strength *string `json:"strength,omitempty"`
	// The system types that the Sentry enforces.
	EnforcedSystems []string `json:"enforcedSystems,omitempty"`
}

// NewInlineResponse20066SentryEnrollment instantiates a new InlineResponse20066SentryEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20066SentryEnrollment() *InlineResponse20066SentryEnrollment {
	this := InlineResponse20066SentryEnrollment{}
	return &this
}

// NewInlineResponse20066SentryEnrollmentWithDefaults instantiates a new InlineResponse20066SentryEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20066SentryEnrollmentWithDefaults() *InlineResponse20066SentryEnrollment {
	this := InlineResponse20066SentryEnrollment{}
	return &this
}

// GetSystemsManagerNetwork returns the SystemsManagerNetwork field value if set, zero value otherwise.
func (o *InlineResponse20066SentryEnrollment) GetSystemsManagerNetwork() InlineResponse20066SentryEnrollmentSystemsManagerNetwork {
	if o == nil || isNil(o.SystemsManagerNetwork) {
		var ret InlineResponse20066SentryEnrollmentSystemsManagerNetwork
		return ret
	}
	return *o.SystemsManagerNetwork
}

// GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066SentryEnrollment) GetSystemsManagerNetworkOk() (*InlineResponse20066SentryEnrollmentSystemsManagerNetwork, bool) {
	if o == nil || isNil(o.SystemsManagerNetwork) {
    return nil, false
	}
	return o.SystemsManagerNetwork, true
}

// HasSystemsManagerNetwork returns a boolean if a field has been set.
func (o *InlineResponse20066SentryEnrollment) HasSystemsManagerNetwork() bool {
	if o != nil && !isNil(o.SystemsManagerNetwork) {
		return true
	}

	return false
}

// SetSystemsManagerNetwork gets a reference to the given InlineResponse20066SentryEnrollmentSystemsManagerNetwork and assigns it to the SystemsManagerNetwork field.
func (o *InlineResponse20066SentryEnrollment) SetSystemsManagerNetwork(v InlineResponse20066SentryEnrollmentSystemsManagerNetwork) {
	o.SystemsManagerNetwork = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *InlineResponse20066SentryEnrollment) GetStrength() string {
	if o == nil || isNil(o.Strength) {
		var ret string
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066SentryEnrollment) GetStrengthOk() (*string, bool) {
	if o == nil || isNil(o.Strength) {
    return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *InlineResponse20066SentryEnrollment) HasStrength() bool {
	if o != nil && !isNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given string and assigns it to the Strength field.
func (o *InlineResponse20066SentryEnrollment) SetStrength(v string) {
	o.Strength = &v
}

// GetEnforcedSystems returns the EnforcedSystems field value if set, zero value otherwise.
func (o *InlineResponse20066SentryEnrollment) GetEnforcedSystems() []string {
	if o == nil || isNil(o.EnforcedSystems) {
		var ret []string
		return ret
	}
	return o.EnforcedSystems
}

// GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20066SentryEnrollment) GetEnforcedSystemsOk() ([]string, bool) {
	if o == nil || isNil(o.EnforcedSystems) {
    return nil, false
	}
	return o.EnforcedSystems, true
}

// HasEnforcedSystems returns a boolean if a field has been set.
func (o *InlineResponse20066SentryEnrollment) HasEnforcedSystems() bool {
	if o != nil && !isNil(o.EnforcedSystems) {
		return true
	}

	return false
}

// SetEnforcedSystems gets a reference to the given []string and assigns it to the EnforcedSystems field.
func (o *InlineResponse20066SentryEnrollment) SetEnforcedSystems(v []string) {
	o.EnforcedSystems = v
}

func (o InlineResponse20066SentryEnrollment) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20066SentryEnrollment struct {
	value *InlineResponse20066SentryEnrollment
	isSet bool
}

func (v NullableInlineResponse20066SentryEnrollment) Get() *InlineResponse20066SentryEnrollment {
	return v.value
}

func (v *NullableInlineResponse20066SentryEnrollment) Set(val *InlineResponse20066SentryEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20066SentryEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20066SentryEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20066SentryEnrollment(val *InlineResponse20066SentryEnrollment) *NullableInlineResponse20066SentryEnrollment {
	return &NullableInlineResponse20066SentryEnrollment{value: val, isSet: true}
}

func (v NullableInlineResponse20066SentryEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20066SentryEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


