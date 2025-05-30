/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject91 struct for InlineObject91
type InlineObject91 struct {
	Ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids `json:"ssids"`
}

// NewInlineObject91 instantiates a new InlineObject91 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject91(ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *InlineObject91 {
	this := InlineObject91{}
	this.Ssids = ssids
	return &this
}

// NewInlineObject91WithDefaults instantiates a new InlineObject91 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject91WithDefaults() *InlineObject91 {
	this := InlineObject91{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *InlineObject91) GetSsids() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	if o == nil {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetSsidsOk() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *InlineObject91) SetSsids(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	o.Ssids = v
}

func (o InlineObject91) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssids"] = o.Ssids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject91 struct {
	value *InlineObject91
	isSet bool
}

func (v NullableInlineObject91) Get() *InlineObject91 {
	return v.value
}

func (v *NullableInlineObject91) Set(val *InlineObject91) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject91) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject91) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject91(val *InlineObject91) *NullableInlineObject91 {
	return &NullableInlineObject91{value: val, isSet: true}
}

func (v NullableInlineObject91) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject91) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


