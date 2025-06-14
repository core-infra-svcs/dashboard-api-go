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

// NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy struct for NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy
type NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy struct {
	// Device serial.
	Serial *string `json:"serial,omitempty"`
	// Device name.
	Name *string `json:"name,omitempty"`
	// Url link to device.
	Url *string `json:"url,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy() *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenSeenByWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenSeenByWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) SetUrl(v string) {
	o.Url = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy(val *NetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenSeenBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


