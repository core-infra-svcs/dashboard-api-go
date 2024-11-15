/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork Systems Manager network targeted for sentry enrollment.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	// The network ID of the Systems Manager network.
	Id string `json:"id"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork(id string) *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork{}
	this.Id = id
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetworkWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetworkWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) SetId(v string) {
	o.Id = v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSentryEnrollmentSystemsManagerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


