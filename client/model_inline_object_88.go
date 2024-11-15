/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject88 struct for InlineObject88
type InlineObject88 struct {
	Ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids `json:"ssids"`
}

// NewInlineObject88 instantiates a new InlineObject88 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject88(ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *InlineObject88 {
	this := InlineObject88{}
	this.Ssids = ssids
	return &this
}

// NewInlineObject88WithDefaults instantiates a new InlineObject88 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject88WithDefaults() *InlineObject88 {
	this := InlineObject88{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *InlineObject88) GetSsids() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	if o == nil {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *InlineObject88) GetSsidsOk() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *InlineObject88) SetSsids(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	o.Ssids = v
}

func (o InlineObject88) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssids"] = o.Ssids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject88 struct {
	value *InlineObject88
	isSet bool
}

func (v NullableInlineObject88) Get() *InlineObject88 {
	return v.value
}

func (v *NullableInlineObject88) Set(val *InlineObject88) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject88) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject88) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject88(val *InlineObject88) *NullableInlineObject88 {
	return &NullableInlineObject88{value: val, isSet: true}
}

func (v NullableInlineObject88) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject88) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


