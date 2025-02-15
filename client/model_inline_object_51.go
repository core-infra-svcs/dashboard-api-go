/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject51 struct for InlineObject51
type InlineObject51 struct {
	SpoofingProtection *NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection `json:"spoofingProtection,omitempty"`
}

// NewInlineObject51 instantiates a new InlineObject51 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject51() *InlineObject51 {
	this := InlineObject51{}
	return &this
}

// NewInlineObject51WithDefaults instantiates a new InlineObject51 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject51WithDefaults() *InlineObject51 {
	this := InlineObject51{}
	return &this
}

// GetSpoofingProtection returns the SpoofingProtection field value if set, zero value otherwise.
func (o *InlineObject51) GetSpoofingProtection() NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection {
	if o == nil || isNil(o.SpoofingProtection) {
		var ret NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection
		return ret
	}
	return *o.SpoofingProtection
}

// GetSpoofingProtectionOk returns a tuple with the SpoofingProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject51) GetSpoofingProtectionOk() (*NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection, bool) {
	if o == nil || isNil(o.SpoofingProtection) {
    return nil, false
	}
	return o.SpoofingProtection, true
}

// HasSpoofingProtection returns a boolean if a field has been set.
func (o *InlineObject51) HasSpoofingProtection() bool {
	if o != nil && !isNil(o.SpoofingProtection) {
		return true
	}

	return false
}

// SetSpoofingProtection gets a reference to the given NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection and assigns it to the SpoofingProtection field.
func (o *InlineObject51) SetSpoofingProtection(v NetworksNetworkIdApplianceFirewallSettingsSpoofingProtection) {
	o.SpoofingProtection = &v
}

func (o InlineObject51) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SpoofingProtection) {
		toSerialize["spoofingProtection"] = o.SpoofingProtection
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject51 struct {
	value *InlineObject51
	isSet bool
}

func (v NullableInlineObject51) Get() *InlineObject51 {
	return v.value
}

func (v *NullableInlineObject51) Set(val *InlineObject51) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject51) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject51) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject51(val *InlineObject51) *NullableInlineObject51 {
	return &NullableInlineObject51{value: val, isSet: true}
}

func (v NullableInlineObject51) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject51) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


