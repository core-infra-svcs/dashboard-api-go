/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	// Custom VPN exclusion rules. Pass an empty array to clear existing rules.
	Custom []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom `json:"custom,omitempty"`
	// Major Application based VPN exclusion rules. Pass an empty array to clear existing rules.
	MajorApplications []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications `json:"majorApplications,omitempty"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *InlineObject63) GetCustom() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom {
	if o == nil || isNil(o.Custom) {
		var ret []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetCustomOk() ([]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom, bool) {
	if o == nil || isNil(o.Custom) {
    return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *InlineObject63) HasCustom() bool {
	if o != nil && !isNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom and assigns it to the Custom field.
func (o *InlineObject63) SetCustom(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsCustom) {
	o.Custom = v
}

// GetMajorApplications returns the MajorApplications field value if set, zero value otherwise.
func (o *InlineObject63) GetMajorApplications() []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications {
	if o == nil || isNil(o.MajorApplications) {
		var ret []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications
		return ret
	}
	return o.MajorApplications
}

// GetMajorApplicationsOk returns a tuple with the MajorApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetMajorApplicationsOk() ([]NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications, bool) {
	if o == nil || isNil(o.MajorApplications) {
    return nil, false
	}
	return o.MajorApplications, true
}

// HasMajorApplications returns a boolean if a field has been set.
func (o *InlineObject63) HasMajorApplications() bool {
	if o != nil && !isNil(o.MajorApplications) {
		return true
	}

	return false
}

// SetMajorApplications gets a reference to the given []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications and assigns it to the MajorApplications field.
func (o *InlineObject63) SetMajorApplications(v []NetworksNetworkIdApplianceTrafficShapingVpnExclusionsMajorApplications) {
	o.MajorApplications = v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !isNil(o.MajorApplications) {
		toSerialize["majorApplications"] = o.MajorApplications
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


