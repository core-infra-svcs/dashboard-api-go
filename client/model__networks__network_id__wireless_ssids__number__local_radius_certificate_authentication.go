/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication The current setting for certificate verification.
type NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication struct {
	// Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether or not to verify the certificate with LDAP.
	UseLdap *bool `json:"useLdap,omitempty"`
	// Whether or not to verify the certificate with OCSP.
	UseOcsp *bool `json:"useOcsp,omitempty"`
	// (Optional) The URL of the OCSP responder to verify client certificate status.
	OcspResponderUrl *string `json:"ocspResponderUrl,omitempty"`
	ClientRootCaCertificate *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate `json:"clientRootCaCertificate,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationWithDefaults() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication {
	this := NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUseLdap returns the UseLdap field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseLdap() bool {
	if o == nil || isNil(o.UseLdap) {
		var ret bool
		return ret
	}
	return *o.UseLdap
}

// GetUseLdapOk returns a tuple with the UseLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseLdapOk() (*bool, bool) {
	if o == nil || isNil(o.UseLdap) {
    return nil, false
	}
	return o.UseLdap, true
}

// HasUseLdap returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasUseLdap() bool {
	if o != nil && !isNil(o.UseLdap) {
		return true
	}

	return false
}

// SetUseLdap gets a reference to the given bool and assigns it to the UseLdap field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetUseLdap(v bool) {
	o.UseLdap = &v
}

// GetUseOcsp returns the UseOcsp field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseOcsp() bool {
	if o == nil || isNil(o.UseOcsp) {
		var ret bool
		return ret
	}
	return *o.UseOcsp
}

// GetUseOcspOk returns a tuple with the UseOcsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetUseOcspOk() (*bool, bool) {
	if o == nil || isNil(o.UseOcsp) {
    return nil, false
	}
	return o.UseOcsp, true
}

// HasUseOcsp returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasUseOcsp() bool {
	if o != nil && !isNil(o.UseOcsp) {
		return true
	}

	return false
}

// SetUseOcsp gets a reference to the given bool and assigns it to the UseOcsp field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetUseOcsp(v bool) {
	o.UseOcsp = &v
}

// GetOcspResponderUrl returns the OcspResponderUrl field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetOcspResponderUrl() string {
	if o == nil || isNil(o.OcspResponderUrl) {
		var ret string
		return ret
	}
	return *o.OcspResponderUrl
}

// GetOcspResponderUrlOk returns a tuple with the OcspResponderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetOcspResponderUrlOk() (*string, bool) {
	if o == nil || isNil(o.OcspResponderUrl) {
    return nil, false
	}
	return o.OcspResponderUrl, true
}

// HasOcspResponderUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasOcspResponderUrl() bool {
	if o != nil && !isNil(o.OcspResponderUrl) {
		return true
	}

	return false
}

// SetOcspResponderUrl gets a reference to the given string and assigns it to the OcspResponderUrl field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetOcspResponderUrl(v string) {
	o.OcspResponderUrl = &v
}

// GetClientRootCaCertificate returns the ClientRootCaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetClientRootCaCertificate() NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	if o == nil || isNil(o.ClientRootCaCertificate) {
		var ret NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate
		return ret
	}
	return *o.ClientRootCaCertificate
}

// GetClientRootCaCertificateOk returns a tuple with the ClientRootCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) GetClientRootCaCertificateOk() (*NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate, bool) {
	if o == nil || isNil(o.ClientRootCaCertificate) {
    return nil, false
	}
	return o.ClientRootCaCertificate, true
}

// HasClientRootCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) HasClientRootCaCertificate() bool {
	if o != nil && !isNil(o.ClientRootCaCertificate) {
		return true
	}

	return false
}

// SetClientRootCaCertificate gets a reference to the given NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate and assigns it to the ClientRootCaCertificate field.
func (o *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) SetClientRootCaCertificate(v NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthenticationClientRootCaCertificate) {
	o.ClientRootCaCertificate = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.UseLdap) {
		toSerialize["useLdap"] = o.UseLdap
	}
	if !isNil(o.UseOcsp) {
		toSerialize["useOcsp"] = o.UseOcsp
	}
	if !isNil(o.OcspResponderUrl) {
		toSerialize["ocspResponderUrl"] = o.OcspResponderUrl
	}
	if !isNil(o.ClientRootCaCertificate) {
		toSerialize["clientRootCaCertificate"] = o.ClientRootCaCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication struct {
	value *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) Get() *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) Set(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication(val *NetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication {
	return &NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberLocalRadiusCertificateAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


