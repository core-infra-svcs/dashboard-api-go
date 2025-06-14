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

// InlineResponse20061 struct for InlineResponse20061
type InlineResponse20061 struct {
	// policies with respective traffic filters for an MX network
	WanTrafficUplinkPreferences []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
}

// NewInlineResponse20061 instantiates a new InlineResponse20061 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20061() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// NewInlineResponse20061WithDefaults instantiates a new InlineResponse20061 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20061WithDefaults() *InlineResponse20061 {
	this := InlineResponse20061{}
	return &this
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineResponse20061) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20061) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineResponse20061) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineResponse20061) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

func (o InlineResponse20061) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20061 struct {
	value *InlineResponse20061
	isSet bool
}

func (v NullableInlineResponse20061) Get() *InlineResponse20061 {
	return v.value
}

func (v *NullableInlineResponse20061) Set(val *InlineResponse20061) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20061) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20061) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20061(val *InlineResponse20061) *NullableInlineResponse20061 {
	return &NullableInlineResponse20061{value: val, isSet: true}
}

func (v NullableInlineResponse20061) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20061) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


