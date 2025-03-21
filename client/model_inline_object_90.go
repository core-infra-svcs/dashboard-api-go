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

// InlineObject90 struct for InlineObject90
type InlineObject90 struct {
	Ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids `json:"ssids"`
}

// NewInlineObject90 instantiates a new InlineObject90 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject90(ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *InlineObject90 {
	this := InlineObject90{}
	this.Ssids = ssids
	return &this
}

// NewInlineObject90WithDefaults instantiates a new InlineObject90 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject90WithDefaults() *InlineObject90 {
	this := InlineObject90{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *InlineObject90) GetSsids() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	if o == nil {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *InlineObject90) GetSsidsOk() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *InlineObject90) SetSsids(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	o.Ssids = v
}

func (o InlineObject90) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssids"] = o.Ssids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject90 struct {
	value *InlineObject90
	isSet bool
}

func (v NullableInlineObject90) Get() *InlineObject90 {
	return v.value
}

func (v *NullableInlineObject90) Set(val *InlineObject90) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject90) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject90) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject90(val *InlineObject90) *NullableInlineObject90 {
	return &NullableInlineObject90{value: val, isSet: true}
}

func (v NullableInlineObject90) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject90) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


