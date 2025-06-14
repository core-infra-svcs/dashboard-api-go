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

// NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations struct for NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations
type NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	// The IP address to test connectivity with
	Ip string `json:"ip"`
	// Description of the testing destination. Optional, defaults to an empty string
	Description *string `json:"description,omitempty"`
	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`
}

// NewNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations instantiates a new NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations(ip string) *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	this := NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations{}
	this.Ip = ip
	var description string = ""
	this.Description = &description
	var default_ bool = false
	this.Default = &default_
	return &this
}

// NewNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinationsWithDefaults instantiates a new NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinationsWithDefaults() *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	this := NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations{}
	var description string = ""
	this.Description = &description
	var default_ bool = false
	this.Default = &default_
	return &this
}

// GetIp returns the Ip field value
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) SetIp(v string) {
	o.Ip = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) SetDefault(v bool) {
	o.Default = &v
}

func (o NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	value *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations
	isSet bool
}

func (v NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) Get() *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	return v.value
}

func (v *NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) Set(val *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations(val *NetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) *NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations {
	return &NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCellularGatewayConnectivityMonitoringDestinationsDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


