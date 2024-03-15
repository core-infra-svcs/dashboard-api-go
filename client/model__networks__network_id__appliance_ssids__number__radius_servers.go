/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSsidsNumberRadiusServers struct for NetworksNetworkIdApplianceSsidsNumberRadiusServers
type NetworksNetworkIdApplianceSsidsNumberRadiusServers struct {
	// The IP address of your RADIUS server.
	Host *string `json:"host,omitempty"`
	// The UDP port your RADIUS servers listens on for Access-requests.
	Port *int32 `json:"port,omitempty"`
	// The RADIUS client shared secret.
	Secret *string `json:"secret,omitempty"`
}

// NewNetworksNetworkIdApplianceSsidsNumberRadiusServers instantiates a new NetworksNetworkIdApplianceSsidsNumberRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSsidsNumberRadiusServers() *NetworksNetworkIdApplianceSsidsNumberRadiusServers {
	this := NetworksNetworkIdApplianceSsidsNumberRadiusServers{}
	return &this
}

// NewNetworksNetworkIdApplianceSsidsNumberRadiusServersWithDefaults instantiates a new NetworksNetworkIdApplianceSsidsNumberRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSsidsNumberRadiusServersWithDefaults() *NetworksNetworkIdApplianceSsidsNumberRadiusServers {
	this := NetworksNetworkIdApplianceSsidsNumberRadiusServers{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworksNetworkIdApplianceSsidsNumberRadiusServers) SetSecret(v string) {
	o.Secret = &v
}

func (o NetworksNetworkIdApplianceSsidsNumberRadiusServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers struct {
	value *NetworksNetworkIdApplianceSsidsNumberRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) Get() *NetworksNetworkIdApplianceSsidsNumberRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) Set(val *NetworksNetworkIdApplianceSsidsNumberRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSsidsNumberRadiusServers(val *NetworksNetworkIdApplianceSsidsNumberRadiusServers) *NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers {
	return &NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSsidsNumberRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


