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

// NetworksNetworkIdCameraWirelessProfilesSsid1 The details of the SSID config.
type NetworksNetworkIdCameraWirelessProfilesSsid1 struct {
	// The name of the SSID.
	Name *string `json:"name,omitempty"`
	// The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	AuthMode *string `json:"authMode,omitempty"`
	// The encryption mode of the SSID. It can be set to ('wpa', 'wpa-eap'). With 'wpa' mode, the authMode should be 'psk' and with 'wpa-eap' the authMode should be '8021x-radius'
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// The pre-shared key of the SSID.
	Psk *string `json:"psk,omitempty"`
}

// NewNetworksNetworkIdCameraWirelessProfilesSsid1 instantiates a new NetworksNetworkIdCameraWirelessProfilesSsid1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraWirelessProfilesSsid1() *NetworksNetworkIdCameraWirelessProfilesSsid1 {
	this := NetworksNetworkIdCameraWirelessProfilesSsid1{}
	return &this
}

// NewNetworksNetworkIdCameraWirelessProfilesSsid1WithDefaults instantiates a new NetworksNetworkIdCameraWirelessProfilesSsid1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraWirelessProfilesSsid1WithDefaults() *NetworksNetworkIdCameraWirelessProfilesSsid1 {
	this := NetworksNetworkIdCameraWirelessProfilesSsid1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) SetName(v string) {
	o.Name = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetAuthMode() string {
	if o == nil || isNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetAuthModeOk() (*string, bool) {
	if o == nil || isNil(o.AuthMode) {
    return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) HasAuthMode() bool {
	if o != nil && !isNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetEncryptionMode() string {
	if o == nil || isNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionMode) {
    return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) HasEncryptionMode() bool {
	if o != nil && !isNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetPsk returns the Psk field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetPsk() string {
	if o == nil || isNil(o.Psk) {
		var ret string
		return ret
	}
	return *o.Psk
}

// GetPskOk returns a tuple with the Psk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) GetPskOk() (*string, bool) {
	if o == nil || isNil(o.Psk) {
    return nil, false
	}
	return o.Psk, true
}

// HasPsk returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) HasPsk() bool {
	if o != nil && !isNil(o.Psk) {
		return true
	}

	return false
}

// SetPsk gets a reference to the given string and assigns it to the Psk field.
func (o *NetworksNetworkIdCameraWirelessProfilesSsid1) SetPsk(v string) {
	o.Psk = &v
}

func (o NetworksNetworkIdCameraWirelessProfilesSsid1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AuthMode) {
		toSerialize["authMode"] = o.AuthMode
	}
	if !isNil(o.EncryptionMode) {
		toSerialize["encryptionMode"] = o.EncryptionMode
	}
	if !isNil(o.Psk) {
		toSerialize["psk"] = o.Psk
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraWirelessProfilesSsid1 struct {
	value *NetworksNetworkIdCameraWirelessProfilesSsid1
	isSet bool
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesSsid1) Get() *NetworksNetworkIdCameraWirelessProfilesSsid1 {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesSsid1) Set(val *NetworksNetworkIdCameraWirelessProfilesSsid1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesSsid1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesSsid1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraWirelessProfilesSsid1(val *NetworksNetworkIdCameraWirelessProfilesSsid1) *NullableNetworksNetworkIdCameraWirelessProfilesSsid1 {
	return &NullableNetworksNetworkIdCameraWirelessProfilesSsid1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraWirelessProfilesSsid1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraWirelessProfilesSsid1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


