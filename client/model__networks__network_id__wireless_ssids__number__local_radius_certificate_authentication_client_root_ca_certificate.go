/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate The Client CA Certificate used to sign the client certificate.
type NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate struct {
	// The contents of the Client CA Certificate. Must be in PEM or DER format.
	Contents *string `json:"contents,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificateWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificateWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) GetContents() string {
	if o == nil || isNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) GetContentsOk() (*string, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) SetContents(v string) {
	o.Contents = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate struct {
	value *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) Get() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) Set(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


