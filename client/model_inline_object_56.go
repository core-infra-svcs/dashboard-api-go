/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject56 struct for InlineObject56
type InlineObject56 struct {
	// policies with respective traffic filters for an MX network
	WanTrafficUplinkPreferences []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
}

// NewInlineObject56 instantiates a new InlineObject56 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject56() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// NewInlineObject56WithDefaults instantiates a new InlineObject56 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject56WithDefaults() *InlineObject56 {
	this := InlineObject56{}
	return &this
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineObject56) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject56) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineObject56) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineObject56) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

func (o InlineObject56) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject56 struct {
	value *InlineObject56
	isSet bool
}

func (v NullableInlineObject56) Get() *InlineObject56 {
	return v.value
}

func (v *NullableInlineObject56) Set(val *InlineObject56) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject56) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject56) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject56(val *InlineObject56) *NullableInlineObject56 {
	return &NullableInlineObject56{value: val, isSet: true}
}

func (v NullableInlineObject56) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject56) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


