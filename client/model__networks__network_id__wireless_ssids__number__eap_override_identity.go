/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity EAP settings for identity requests.
type NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity struct {
	// Maximum number of EAP retries.
	Retries *int32 `json:"retries,omitempty"`
	// EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity instantiates a new NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity() *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity {
	this := NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentityWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentityWithDefaults() *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity {
	this := NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity{}
	return &this
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) GetRetries() int32 {
	if o == nil || isNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) GetRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.Retries) {
    return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) HasRetries() bool {
	if o != nil && !isNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) SetRetries(v int32) {
	o.Retries = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity struct {
	value *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) Get() *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) Set(val *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity(val *NetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) *NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity {
	return &NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberEapOverrideIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


