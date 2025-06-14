/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesContentFiltering The content filtering settings for your group policy
type NetworksNetworkIdGroupPoliciesContentFiltering struct {
	AllowedUrlPatterns *NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns `json:"allowedUrlPatterns,omitempty"`
	BlockedUrlPatterns *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns `json:"blockedUrlPatterns,omitempty"`
	BlockedUrlCategories *NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories `json:"blockedUrlCategories,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesContentFiltering instantiates a new NetworksNetworkIdGroupPoliciesContentFiltering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesContentFiltering() *NetworksNetworkIdGroupPoliciesContentFiltering {
	this := NetworksNetworkIdGroupPoliciesContentFiltering{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesContentFilteringWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesContentFiltering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesContentFilteringWithDefaults() *NetworksNetworkIdGroupPoliciesContentFiltering {
	this := NetworksNetworkIdGroupPoliciesContentFiltering{}
	return &this
}

// GetAllowedUrlPatterns returns the AllowedUrlPatterns field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetAllowedUrlPatterns() NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns {
	if o == nil || isNil(o.AllowedUrlPatterns) {
		var ret NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns
		return ret
	}
	return *o.AllowedUrlPatterns
}

// GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetAllowedUrlPatternsOk() (*NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns, bool) {
	if o == nil || isNil(o.AllowedUrlPatterns) {
    return nil, false
	}
	return o.AllowedUrlPatterns, true
}

// HasAllowedUrlPatterns returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) HasAllowedUrlPatterns() bool {
	if o != nil && !isNil(o.AllowedUrlPatterns) {
		return true
	}

	return false
}

// SetAllowedUrlPatterns gets a reference to the given NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns and assigns it to the AllowedUrlPatterns field.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) SetAllowedUrlPatterns(v NetworksNetworkIdGroupPoliciesContentFilteringAllowedUrlPatterns) {
	o.AllowedUrlPatterns = &v
}

// GetBlockedUrlPatterns returns the BlockedUrlPatterns field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetBlockedUrlPatterns() NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns {
	if o == nil || isNil(o.BlockedUrlPatterns) {
		var ret NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns
		return ret
	}
	return *o.BlockedUrlPatterns
}

// GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetBlockedUrlPatternsOk() (*NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns, bool) {
	if o == nil || isNil(o.BlockedUrlPatterns) {
    return nil, false
	}
	return o.BlockedUrlPatterns, true
}

// HasBlockedUrlPatterns returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) HasBlockedUrlPatterns() bool {
	if o != nil && !isNil(o.BlockedUrlPatterns) {
		return true
	}

	return false
}

// SetBlockedUrlPatterns gets a reference to the given NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns and assigns it to the BlockedUrlPatterns field.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) SetBlockedUrlPatterns(v NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlPatterns) {
	o.BlockedUrlPatterns = &v
}

// GetBlockedUrlCategories returns the BlockedUrlCategories field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetBlockedUrlCategories() NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories {
	if o == nil || isNil(o.BlockedUrlCategories) {
		var ret NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories
		return ret
	}
	return *o.BlockedUrlCategories
}

// GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) GetBlockedUrlCategoriesOk() (*NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories, bool) {
	if o == nil || isNil(o.BlockedUrlCategories) {
    return nil, false
	}
	return o.BlockedUrlCategories, true
}

// HasBlockedUrlCategories returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) HasBlockedUrlCategories() bool {
	if o != nil && !isNil(o.BlockedUrlCategories) {
		return true
	}

	return false
}

// SetBlockedUrlCategories gets a reference to the given NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories and assigns it to the BlockedUrlCategories field.
func (o *NetworksNetworkIdGroupPoliciesContentFiltering) SetBlockedUrlCategories(v NetworksNetworkIdGroupPoliciesContentFilteringBlockedUrlCategories) {
	o.BlockedUrlCategories = &v
}

func (o NetworksNetworkIdGroupPoliciesContentFiltering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedUrlPatterns) {
		toSerialize["allowedUrlPatterns"] = o.AllowedUrlPatterns
	}
	if !isNil(o.BlockedUrlPatterns) {
		toSerialize["blockedUrlPatterns"] = o.BlockedUrlPatterns
	}
	if !isNil(o.BlockedUrlCategories) {
		toSerialize["blockedUrlCategories"] = o.BlockedUrlCategories
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesContentFiltering struct {
	value *NetworksNetworkIdGroupPoliciesContentFiltering
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFiltering) Get() *NetworksNetworkIdGroupPoliciesContentFiltering {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFiltering) Set(val *NetworksNetworkIdGroupPoliciesContentFiltering) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFiltering) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFiltering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesContentFiltering(val *NetworksNetworkIdGroupPoliciesContentFiltering) *NullableNetworksNetworkIdGroupPoliciesContentFiltering {
	return &NullableNetworksNetworkIdGroupPoliciesContentFiltering{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesContentFiltering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesContentFiltering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


