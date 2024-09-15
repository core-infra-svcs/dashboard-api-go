/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraWirelessProfilesIdentity The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.
type NetworksNetworkIdCameraWirelessProfilesIdentity struct {
	// The username of the identity.
	Username *string `json:"username,omitempty"`
	// The password of the identity.
	Password *string `json:"password,omitempty"`
}

// NewNetworksNetworkIdCameraWirelessProfilesIdentity instantiates a new NetworksNetworkIdCameraWirelessProfilesIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraWirelessProfilesIdentity() *NetworksNetworkIdCameraWirelessProfilesIdentity {
	this := NetworksNetworkIdCameraWirelessProfilesIdentity{}
	return &this
}

// NewNetworksNetworkIdCameraWirelessProfilesIdentityWithDefaults instantiates a new NetworksNetworkIdCameraWirelessProfilesIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraWirelessProfilesIdentityWithDefaults() *NetworksNetworkIdCameraWirelessProfilesIdentity {
	this := NetworksNetworkIdCameraWirelessProfilesIdentity{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NetworksNetworkIdCameraWirelessProfilesIdentity) SetPassword(v string) {
	o.Password = &v
}

func (o NetworksNetworkIdCameraWirelessProfilesIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraWirelessProfilesIdentity struct {
	value *NetworksNetworkIdCameraWirelessProfilesIdentity
	isSet bool
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesIdentity) Get() *NetworksNetworkIdCameraWirelessProfilesIdentity {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesIdentity) Set(val *NetworksNetworkIdCameraWirelessProfilesIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraWirelessProfilesIdentity(val *NetworksNetworkIdCameraWirelessProfilesIdentity) *NullableNetworksNetworkIdCameraWirelessProfilesIdentity {
	return &NullableNetworksNetworkIdCameraWirelessProfilesIdentity{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


