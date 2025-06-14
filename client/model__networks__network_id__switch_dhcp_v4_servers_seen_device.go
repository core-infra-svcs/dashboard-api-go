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

// NetworksNetworkIdSwitchDhcpV4ServersSeenDevice Attributes of the server when it's a device.
type NetworksNetworkIdSwitchDhcpV4ServersSeenDevice struct {
	// Device serial.
	Serial *string `json:"serial,omitempty"`
	// Device name.
	Name *string `json:"name,omitempty"`
	// Url link to device.
	Url *string `json:"url,omitempty"`
	Interface *NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface `json:"interface,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenDevice instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenDevice() *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenDevice{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenDeviceWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenDevice{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) SetUrl(v string) {
	o.Url = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetInterface() NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface {
	if o == nil || isNil(o.Interface) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) GetInterfaceOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface and assigns it to the Interface field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) SetInterface(v NetworksNetworkIdSwitchDhcpV4ServersSeenDeviceInterface) {
	o.Interface = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice(val *NetworksNetworkIdSwitchDhcpV4ServersSeenDevice) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


