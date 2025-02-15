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

// NetworksNetworkIdWirelessSsidsNumberLocalRadius The current setting for Local Authentication, a built-in RADIUS server on the access point. Only valid if authMode is '8021x-localradius'.
type NetworksNetworkIdWirelessSsidsNumberLocalRadius struct {
	// The duration (in seconds) for which LDAP and OCSP lookups are cached.
	CacheTimeout *int32 `json:"cacheTimeout,omitempty"`
	PasswordAuthentication *NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication `json:"passwordAuthentication,omitempty"`
	CertificateAuthentication *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication `json:"certificateAuthentication,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadius instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadius() *NetworksNetworkIdWirelessSsidsNumberLocalRadius {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadius{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalRadius {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadius{}
	return &this
}

// GetCacheTimeout returns the CacheTimeout field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetCacheTimeout() int32 {
	if o == nil || isNil(o.CacheTimeout) {
		var ret int32
		return ret
	}
	return *o.CacheTimeout
}

// GetCacheTimeoutOk returns a tuple with the CacheTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetCacheTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.CacheTimeout) {
    return nil, false
	}
	return o.CacheTimeout, true
}

// HasCacheTimeout returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) HasCacheTimeout() bool {
	if o != nil && !isNil(o.CacheTimeout) {
		return true
	}

	return false
}

// SetCacheTimeout gets a reference to the given int32 and assigns it to the CacheTimeout field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) SetCacheTimeout(v int32) {
	o.CacheTimeout = &v
}

// GetPasswordAuthentication returns the PasswordAuthentication field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetPasswordAuthentication() NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication {
	if o == nil || isNil(o.PasswordAuthentication) {
		var ret NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication
		return ret
	}
	return *o.PasswordAuthentication
}

// GetPasswordAuthenticationOk returns a tuple with the PasswordAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetPasswordAuthenticationOk() (*NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication, bool) {
	if o == nil || isNil(o.PasswordAuthentication) {
    return nil, false
	}
	return o.PasswordAuthentication, true
}

// HasPasswordAuthentication returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) HasPasswordAuthentication() bool {
	if o != nil && !isNil(o.PasswordAuthentication) {
		return true
	}

	return false
}

// SetPasswordAuthentication gets a reference to the given NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication and assigns it to the PasswordAuthentication field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) SetPasswordAuthentication(v NetworksNetworkIdWirelessSsidsNumberLocalRadiusPasswordAuthentication) {
	o.PasswordAuthentication = &v
}

// GetCertificateAuthentication returns the CertificateAuthentication field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetCertificateAuthentication() NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication {
	if o == nil || isNil(o.CertificateAuthentication) {
		var ret NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication
		return ret
	}
	return *o.CertificateAuthentication
}

// GetCertificateAuthenticationOk returns a tuple with the CertificateAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) GetCertificateAuthenticationOk() (*NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication, bool) {
	if o == nil || isNil(o.CertificateAuthentication) {
    return nil, false
	}
	return o.CertificateAuthentication, true
}

// HasCertificateAuthentication returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) HasCertificateAuthentication() bool {
	if o != nil && !isNil(o.CertificateAuthentication) {
		return true
	}

	return false
}

// SetCertificateAuthentication gets a reference to the given NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication and assigns it to the CertificateAuthentication field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadius) SetCertificateAuthentication(v NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) {
	o.CertificateAuthentication = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLocalRadius) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CacheTimeout) {
		toSerialize["cacheTimeout"] = o.CacheTimeout
	}
	if !isNil(o.PasswordAuthentication) {
		toSerialize["passwordAuthentication"] = o.PasswordAuthentication
	}
	if !isNil(o.CertificateAuthentication) {
		toSerialize["certificateAuthentication"] = o.CertificateAuthentication
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius struct {
	value *NetworksNetworkIdWirelessSsidsNumberLocalRadius
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) Get() *NetworksNetworkIdWirelessSsidsNumberLocalRadius {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) Set(val *NetworksNetworkIdWirelessSsidsNumberLocalRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLocalRadius(val *NetworksNetworkIdWirelessSsidsNumberLocalRadius) *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


