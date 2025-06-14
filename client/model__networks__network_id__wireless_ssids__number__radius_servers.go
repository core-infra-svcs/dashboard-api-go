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

// NetworksNetworkIdWirelessSsidsNumberRadiusServers struct for NetworksNetworkIdWirelessSsidsNumberRadiusServers
type NetworksNetworkIdWirelessSsidsNumberRadiusServers struct {
	// IP address (or FQDN) of your RADIUS server
	Host string `json:"host"`
	// UDP port the RADIUS server listens on for Access-requests
	Port *int32 `json:"port,omitempty"`
	// RADIUS client shared secret
	Secret *string `json:"secret,omitempty"`
	// Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.
	RadsecEnabled *bool `json:"radsecEnabled,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server.
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberRadiusServers instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberRadiusServers(host string) *NetworksNetworkIdWirelessSsidsNumberRadiusServers {
	this := NetworksNetworkIdWirelessSsidsNumberRadiusServers{}
	this.Host = host
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberRadiusServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberRadiusServersWithDefaults() *NetworksNetworkIdWirelessSsidsNumberRadiusServers {
	this := NetworksNetworkIdWirelessSsidsNumberRadiusServers{}
	return &this
}

// GetHost returns the Host field value
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetSecret(v string) {
	o.Secret = &v
}

// GetRadsecEnabled returns the RadsecEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetRadsecEnabled() bool {
	if o == nil || isNil(o.RadsecEnabled) {
		var ret bool
		return ret
	}
	return *o.RadsecEnabled
}

// GetRadsecEnabledOk returns a tuple with the RadsecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetRadsecEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RadsecEnabled) {
    return nil, false
	}
	return o.RadsecEnabled, true
}

// HasRadsecEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasRadsecEnabled() bool {
	if o != nil && !isNil(o.RadsecEnabled) {
		return true
	}

	return false
}

// SetRadsecEnabled gets a reference to the given bool and assigns it to the RadsecEnabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetRadsecEnabled(v bool) {
	o.RadsecEnabled = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetOpenRoamingCertificateId() int32 {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
    return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasOpenRoamingCertificateId() bool {
	if o != nil && !isNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetCaCertificate() string {
	if o == nil || isNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) GetCaCertificateOk() (*string, bool) {
	if o == nil || isNil(o.CaCertificate) {
    return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) HasCaCertificate() bool {
	if o != nil && !isNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *NetworksNetworkIdWirelessSsidsNumberRadiusServers) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberRadiusServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !isNil(o.RadsecEnabled) {
		toSerialize["radsecEnabled"] = o.RadsecEnabled
	}
	if !isNil(o.OpenRoamingCertificateId) {
		toSerialize["openRoamingCertificateId"] = o.OpenRoamingCertificateId
	}
	if !isNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers struct {
	value *NetworksNetworkIdWirelessSsidsNumberRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) Get() *NetworksNetworkIdWirelessSsidsNumberRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) Set(val *NetworksNetworkIdWirelessSsidsNumberRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberRadiusServers(val *NetworksNetworkIdWirelessSsidsNumberRadiusServers) *NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers {
	return &NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


