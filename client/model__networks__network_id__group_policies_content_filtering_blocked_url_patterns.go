/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns Settings for blocked URL patterns
type NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns struct {
	// How URL patterns are applied. Can be 'network default', 'append' or 'override'.
	Settings *string `json:"settings,omitempty"`
	// A list of URL patterns that are blocked
	Patterns []string `json:"patterns,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns {
	this := NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatternsWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatternsWithDefaults() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns {
	this := NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) GetSettings() string {
	if o == nil || o.Settings == nil {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) GetSettingsOk() (*string, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) SetSettings(v string) {
	o.Settings = &v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) GetPatterns() []string {
	if o == nil || o.Patterns == nil {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) GetPatternsOk() ([]string, bool) {
	if o == nil || o.Patterns == nil {
		return nil, false
	}
	return o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) HasPatterns() bool {
	if o != nil && o.Patterns != nil {
		return true
	}

	return false
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) SetPatterns(v []string) {
	o.Patterns = v
}

func (o NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Patterns != nil {
		toSerialize["patterns"] = o.Patterns
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns struct {
	value *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) Get() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) Set(val *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns(val *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns {
	return &NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


