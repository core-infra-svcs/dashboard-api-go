/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories Settings for blocked URL categories
type NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories struct {
	// How URL categories are applied. Can be 'network default', 'append' or 'override'.
	Settings *string `json:"settings,omitempty"`
	// A list of URL categories to block
	Categories []string `json:"categories,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories {
	this := NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategoriesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategoriesWithDefaults() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories {
	this := NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) SetSettings(v string) {
	o.Settings = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) GetCategories() []string {
	if o == nil || isNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) GetCategoriesOk() ([]string, bool) {
	if o == nil || isNil(o.Categories) {
    return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) HasCategories() bool {
	if o != nil && !isNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) SetCategories(v []string) {
	o.Categories = v
}

func (o NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories struct {
	value *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) Get() *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) Set(val *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories(val *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories {
	return &NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


