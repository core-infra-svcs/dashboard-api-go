/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSsidsRadiusServers struct for NetworksNetworkIdApplianceSsidsRadiusServers
type NetworksNetworkIdApplianceSsidsRadiusServers struct {
	// The IP address of your RADIUS server.
	Host *string `json:"host,omitempty"`
	// The UDP port your RADIUS servers listens on for Access-requests.
	Port *int32 `json:"port,omitempty"`
}

// NewNetworksNetworkIdApplianceSsidsRadiusServers instantiates a new NetworksNetworkIdApplianceSsidsRadiusServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSsidsRadiusServers() *NetworksNetworkIdApplianceSsidsRadiusServers {
	this := NetworksNetworkIdApplianceSsidsRadiusServers{}
	return &this
}

// NewNetworksNetworkIdApplianceSsidsRadiusServersWithDefaults instantiates a new NetworksNetworkIdApplianceSsidsRadiusServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSsidsRadiusServersWithDefaults() *NetworksNetworkIdApplianceSsidsRadiusServers {
	this := NetworksNetworkIdApplianceSsidsRadiusServers{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworksNetworkIdApplianceSsidsRadiusServers) SetPort(v int32) {
	o.Port = &v
}

func (o NetworksNetworkIdApplianceSsidsRadiusServers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSsidsRadiusServers struct {
	value *NetworksNetworkIdApplianceSsidsRadiusServers
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSsidsRadiusServers) Get() *NetworksNetworkIdApplianceSsidsRadiusServers {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSsidsRadiusServers) Set(val *NetworksNetworkIdApplianceSsidsRadiusServers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSsidsRadiusServers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSsidsRadiusServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSsidsRadiusServers(val *NetworksNetworkIdApplianceSsidsRadiusServers) *NullableNetworksNetworkIdApplianceSsidsRadiusServers {
	return &NullableNetworksNetworkIdApplianceSsidsRadiusServers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSsidsRadiusServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSsidsRadiusServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


