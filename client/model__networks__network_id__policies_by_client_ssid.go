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

// NetworksNetworkIdPoliciesByClientSsid struct for NetworksNetworkIdPoliciesByClientSsid
type NetworksNetworkIdPoliciesByClientSsid struct {
	// number of ssid
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
}

// NewNetworksNetworkIdPoliciesByClientSsid instantiates a new NetworksNetworkIdPoliciesByClientSsid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdPoliciesByClientSsid() *NetworksNetworkIdPoliciesByClientSsid {
	this := NetworksNetworkIdPoliciesByClientSsid{}
	return &this
}

// NewNetworksNetworkIdPoliciesByClientSsidWithDefaults instantiates a new NetworksNetworkIdPoliciesByClientSsid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdPoliciesByClientSsidWithDefaults() *NetworksNetworkIdPoliciesByClientSsid {
	this := NetworksNetworkIdPoliciesByClientSsid{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *NetworksNetworkIdPoliciesByClientSsid) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdPoliciesByClientSsid) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *NetworksNetworkIdPoliciesByClientSsid) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *NetworksNetworkIdPoliciesByClientSsid) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

func (o NetworksNetworkIdPoliciesByClientSsid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdPoliciesByClientSsid struct {
	value *NetworksNetworkIdPoliciesByClientSsid
	isSet bool
}

func (v NullableNetworksNetworkIdPoliciesByClientSsid) Get() *NetworksNetworkIdPoliciesByClientSsid {
	return v.value
}

func (v *NullableNetworksNetworkIdPoliciesByClientSsid) Set(val *NetworksNetworkIdPoliciesByClientSsid) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdPoliciesByClientSsid) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdPoliciesByClientSsid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdPoliciesByClientSsid(val *NetworksNetworkIdPoliciesByClientSsid) *NullableNetworksNetworkIdPoliciesByClientSsid {
	return &NullableNetworksNetworkIdPoliciesByClientSsid{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdPoliciesByClientSsid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdPoliciesByClientSsid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


