/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsRadiusAccountingServers struct for NetworksNetworkIdWirelessSsidsRadiusAccountingServers
type NetworksNetworkIdWirelessSsidsRadiusAccountingServers struct {
	// IP address (or FQDN) to which the APs will send RADIUS accounting messages
	Host *string `json:"host,omitempty"`
	// Port on the RADIUS server that is listening for accounting messages
	Port *int32 `json:"port,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsRadiusAccountingServers instantiates a new NetworksNetworkIdWirelessSsidsRadiusAccountingServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsRadiusAccountingServers() *NetworksNetworkIdWirelessSsidsRadiusAccountingServers {
	this := NetworksNetworkIdWirelessSsidsRadiusAccountingServers{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsRadiusAccountingServersWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsRadiusAccountingServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsRadiusAccountingServersWithDefaults() *NetworksNetworkIdWirelessSsidsRadiusAccountingServers {
	this := NetworksNetworkIdWirelessSsidsRadiusAccountingServers{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) SetPort(v int32) {
	o.Port = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetOpenRoamingCertificateId() int32 {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || isNil(o.OpenRoamingCertificateId) {
    return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) HasOpenRoamingCertificateId() bool {
	if o != nil && !isNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetCaCertificate() string {
	if o == nil || isNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) GetCaCertificateOk() (*string, bool) {
	if o == nil || isNil(o.CaCertificate) {
    return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) HasCaCertificate() bool {
	if o != nil && !isNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o NetworksNetworkIdWirelessSsidsRadiusAccountingServers) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers struct {
	value *NetworksNetworkIdWirelessSsidsRadiusAccountingServers
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) Get() *NetworksNetworkIdWirelessSsidsRadiusAccountingServers {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) Set(val *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers(val *NetworksNetworkIdWirelessSsidsRadiusAccountingServers) *NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers {
	return &NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsRadiusAccountingServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


