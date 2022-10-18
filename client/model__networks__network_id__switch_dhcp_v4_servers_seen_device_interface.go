/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface Interface attributes of the server. Only for configured servers.
type NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface struct {
	// Interface name.
	Name *string `json:"name,omitempty"`
	// Url link to interface.
	Url *string `json:"url,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface() *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterfaceWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterfaceWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) SetUrl(v string) {
	o.Url = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface(val *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


