/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns Settings for allowed URL patterns
type NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns struct {
	// How URL patterns are applied. Can be 'network default', 'append' or 'override'.
	Settings *string `json:"settings,omitempty"`
	// A list of URL patterns that are allowed
	Patterns []string `json:"patterns,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns() *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns {
	this := NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatternsWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatternsWithDefaults() *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns {
	this := NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) SetSettings(v string) {
	o.Settings = &v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) GetPatterns() []string {
	if o == nil || isNil(o.Patterns) {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) GetPatternsOk() ([]string, bool) {
	if o == nil || isNil(o.Patterns) {
    return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) HasPatterns() bool {
	if o != nil && !isNil(o.Patterns) {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) SetPatterns(v []string) {
	o.Patterns = v
}

func (o NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.Patterns) {
		toSerialize["patterns"] = o.Patterns
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns struct {
	value *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) Get() *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) Set(val *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns(val *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) *NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns {
	return &NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


