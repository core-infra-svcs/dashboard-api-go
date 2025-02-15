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

// NetworksNetworkIdWirelessSsidsRadiusServers struct for NetworksNetworkIdWirelessSsidsRadiusServers
type NetworksNetworkIdWirelessSsidsRadiusServers struct {
	// IP address (or FQDN) of your RADIUS server
	Host *string `json:"host,omitempty"`
	// UDP port the RADIUS server listens on for Access-requests
	Port *int32 `json:"port,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsRadiusServers instantiates a new NetworksNetworkIdWirelessSsidsRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsRadiusServers() *NetworksNetworkIdWirelessSsidsRadiusServers {
	this := NetworksNetworkIdWirelessSsidsRadiusServers{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsRadiusServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsRadiusServersWithDefaults() *NetworksNetworkIdWirelessSsidsRadiusServers {
	this := NetworksNetworkIdWirelessSsidsRadiusServers{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetPort(v int32) {
	o.Port = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetOpenRoamingCertificateId() int32 {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
    return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasOpenRoamingCertificateId() bool {
	if o != nil && !isNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetCaCertificate() string {
	if o == nil || isNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) GetCaCertificateOk() (*string, bool) {
	if o == nil || isNil(o.CaCertificate) {
    return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) HasCaCertificate() bool {
	if o != nil && !isNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *NetworksNetworkIdWirelessSsidsRadiusServers) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o NetworksNetworkIdWirelessSsidsRadiusServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.OpenRoamingCertificateId) {
		toSerialize["openRoamingCertificateId"] = o.OpenRoamingCertificateId
	}
	if !isNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsRadiusServers struct {
	value *NetworksNetworkIdWirelessSsidsRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusServers) Get() *NetworksNetworkIdWirelessSsidsRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusServers) Set(val *NetworksNetworkIdWirelessSsidsRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsRadiusServers(val *NetworksNetworkIdWirelessSsidsRadiusServers) *NullableNetworksNetworkIdWirelessSsidsRadiusServers {
	return &NullableNetworksNetworkIdWirelessSsidsRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


