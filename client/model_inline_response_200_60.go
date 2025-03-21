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

// InlineResponse20060 struct for InlineResponse20060
type InlineResponse20060 struct {
	// policies with respective traffic filters for an MX network
	WanTrafficUplinkPreferences []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences `json:"wanTrafficUplinkPreferences,omitempty"`
}

// NewInlineResponse20060 instantiates a new InlineResponse20060 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060() *InlineResponse20060 {
	this := InlineResponse20060{}
	return &this
}

// NewInlineResponse20060WithDefaults instantiates a new InlineResponse20060 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060WithDefaults() *InlineResponse20060 {
	this := InlineResponse20060{}
	return &this
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *InlineResponse20060) GetWanTrafficUplinkPreferences() []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
		var ret []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20060) GetWanTrafficUplinkPreferencesOk() ([]NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences, bool) {
	if o == nil || isNil(o.WanTrafficUplinkPreferences) {
    return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *InlineResponse20060) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !isNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences and assigns it to the WanTrafficUplinkPreferences field.
func (o *InlineResponse20060) SetWanTrafficUplinkPreferences(v []NetworksNetworkIdApplianceSdwanInternetPoliciesWanTrafficUplinkPreferences) {
	o.WanTrafficUplinkPreferences = v
}

func (o InlineResponse20060) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060 struct {
	value *InlineResponse20060
	isSet bool
}

func (v NullableInlineResponse20060) Get() *InlineResponse20060 {
	return v.value
}

func (v *NullableInlineResponse20060) Set(val *InlineResponse20060) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060(val *InlineResponse20060) *NullableInlineResponse20060 {
	return &NullableInlineResponse20060{value: val, isSet: true}
}

func (v NullableInlineResponse20060) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


