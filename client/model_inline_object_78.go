/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject78 struct for InlineObject78
type InlineObject78 struct {
	Ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids `json:"ssids"`
}

// NewInlineObject78 instantiates a new InlineObject78 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject78(ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *InlineObject78 {
	this := InlineObject78{}
	this.Ssids = ssids
	return &this
}

// NewInlineObject78WithDefaults instantiates a new InlineObject78 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject78WithDefaults() *InlineObject78 {
	this := InlineObject78{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *InlineObject78) GetSsids() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	if o == nil {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *InlineObject78) GetSsidsOk() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *InlineObject78) SetSsids(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	o.Ssids = v
}

func (o InlineObject78) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssids"] = o.Ssids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject78 struct {
	value *InlineObject78
	isSet bool
}

func (v NullableInlineObject78) Get() *InlineObject78 {
	return v.value
}

func (v *NullableInlineObject78) Set(val *InlineObject78) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject78) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject78) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject78(val *InlineObject78) *NullableInlineObject78 {
	return &NullableInlineObject78{value: val, isSet: true}
}

func (v NullableInlineObject78) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject78) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


