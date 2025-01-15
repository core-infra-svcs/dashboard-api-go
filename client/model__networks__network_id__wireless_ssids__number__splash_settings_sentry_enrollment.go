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

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment Systems Manager sentry enrollment splash settings.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment struct {
	SystemsManagerNetwork *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork `json:"systemsManagerNetwork,omitempty"`
	// The strength of the enforcement of selected system types. Must be one of: 'focused', 'click-through', and 'strict'.
	Strength *string `json:"strength,omitempty"`
	// The system types that the Sentry enforces. Must be included in: 'iOS, 'Android', 'macOS', and 'Windows'.
	EnforcedSystems []string `json:"enforcedSystems,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment{}
	return &this
}

// GetSystemsManagerNetwork returns the SystemsManagerNetwork field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetSystemsManagerNetwork() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork {
	if o == nil || isNil(o.SystemsManagerNetwork) {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork
		return ret
	}
	return *o.SystemsManagerNetwork
}

// GetSystemsManagerNetworkOk returns a tuple with the SystemsManagerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetSystemsManagerNetworkOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork, bool) {
	if o == nil || isNil(o.SystemsManagerNetwork) {
    return nil, false
	}
	return o.SystemsManagerNetwork, true
}

// HasSystemsManagerNetwork returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) HasSystemsManagerNetwork() bool {
	if o != nil && !isNil(o.SystemsManagerNetwork) {
		return true
	}

	return false
}

// SetSystemsManagerNetwork gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork and assigns it to the SystemsManagerNetwork field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) SetSystemsManagerNetwork(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) {
	o.SystemsManagerNetwork = &v
}

// GetStrength returns the Strength field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetStrength() string {
	if o == nil || isNil(o.Strength) {
		var ret string
		return ret
	}
	return *o.Strength
}

// GetStrengthOk returns a tuple with the Strength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetStrengthOk() (*string, bool) {
	if o == nil || isNil(o.Strength) {
    return nil, false
	}
	return o.Strength, true
}

// HasStrength returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) HasStrength() bool {
	if o != nil && !isNil(o.Strength) {
		return true
	}

	return false
}

// SetStrength gets a reference to the given string and assigns it to the Strength field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) SetStrength(v string) {
	o.Strength = &v
}

// GetEnforcedSystems returns the EnforcedSystems field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetEnforcedSystems() []string {
	if o == nil || isNil(o.EnforcedSystems) {
		var ret []string
		return ret
	}
	return o.EnforcedSystems
}

// GetEnforcedSystemsOk returns a tuple with the EnforcedSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) GetEnforcedSystemsOk() ([]string, bool) {
	if o == nil || isNil(o.EnforcedSystems) {
    return nil, false
	}
	return o.EnforcedSystems, true
}

// HasEnforcedSystems returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) HasEnforcedSystems() bool {
	if o != nil && !isNil(o.EnforcedSystems) {
		return true
	}

	return false
}

// SetEnforcedSystems gets a reference to the given []string and assigns it to the EnforcedSystems field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) SetEnforcedSystems(v []string) {
	o.EnforcedSystems = v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


