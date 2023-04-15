/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations struct for NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations
type NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations struct {
	// SSID for which the user is being authorized
	SsidNumber int32 `json:"ssidNumber"`
	// Date for authorization to expire. Default is for authorization to not expire.
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations instantiates a new NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations(ssidNumber int32) *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations {
	this := NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations{}
	this.SsidNumber = ssidNumber
	var expiresAt string = "Never"
	this.ExpiresAt = &expiresAt
	return &this
}

// NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizationsWithDefaults instantiates a new NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizationsWithDefaults() *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations {
	this := NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations{}
	var expiresAt string = "Never"
	this.ExpiresAt = &expiresAt
	return &this
}

// GetSsidNumber returns the SsidNumber field value
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetSsidNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetSsidNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SsidNumber, true
}

// SetSsidNumber sets field value
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) SetSsidNumber(v int32) {
	o.SsidNumber = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetExpiresAt() string {
	if o == nil || isNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetExpiresAtOk() (*string, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations struct {
	value *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations
	isSet bool
}

func (v NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) Get() *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations {
	return v.value
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) Set(val *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations(val *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) *NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations {
	return &NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


