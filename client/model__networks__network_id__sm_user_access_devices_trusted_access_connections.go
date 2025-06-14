/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections struct for NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections
type NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections struct {
	// config id
	TrustedAccessConfigId *string `json:"trustedAccessConfigId,omitempty"`
	// time that config was downloaded
	DownloadedAt *string `json:"downloadedAt,omitempty"`
	// time that SCEP completed
	ScepCompletedAt *time.Time `json:"scepCompletedAt,omitempty"`
	// time of last connection
	LastConnectedAt *time.Time `json:"lastConnectedAt,omitempty"`
}

// NewNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections instantiates a new NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections() *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections {
	this := NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections{}
	return &this
}

// NewNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnectionsWithDefaults instantiates a new NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnectionsWithDefaults() *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections {
	this := NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections{}
	return &this
}

// GetTrustedAccessConfigId returns the TrustedAccessConfigId field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetTrustedAccessConfigId() string {
	if o == nil || isNil(o.TrustedAccessConfigId) {
		var ret string
		return ret
	}
	return *o.TrustedAccessConfigId
}

// GetTrustedAccessConfigIdOk returns a tuple with the TrustedAccessConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetTrustedAccessConfigIdOk() (*string, bool) {
	if o == nil || isNil(o.TrustedAccessConfigId) {
    return nil, false
	}
	return o.TrustedAccessConfigId, true
}

// HasTrustedAccessConfigId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) HasTrustedAccessConfigId() bool {
	if o != nil && !isNil(o.TrustedAccessConfigId) {
		return true
	}

	return false
}

// SetTrustedAccessConfigId gets a reference to the given string and assigns it to the TrustedAccessConfigId field.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) SetTrustedAccessConfigId(v string) {
	o.TrustedAccessConfigId = &v
}

// GetDownloadedAt returns the DownloadedAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetDownloadedAt() string {
	if o == nil || isNil(o.DownloadedAt) {
		var ret string
		return ret
	}
	return *o.DownloadedAt
}

// GetDownloadedAtOk returns a tuple with the DownloadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetDownloadedAtOk() (*string, bool) {
	if o == nil || isNil(o.DownloadedAt) {
    return nil, false
	}
	return o.DownloadedAt, true
}

// HasDownloadedAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) HasDownloadedAt() bool {
	if o != nil && !isNil(o.DownloadedAt) {
		return true
	}

	return false
}

// SetDownloadedAt gets a reference to the given string and assigns it to the DownloadedAt field.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) SetDownloadedAt(v string) {
	o.DownloadedAt = &v
}

// GetScepCompletedAt returns the ScepCompletedAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetScepCompletedAt() time.Time {
	if o == nil || isNil(o.ScepCompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.ScepCompletedAt
}

// GetScepCompletedAtOk returns a tuple with the ScepCompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetScepCompletedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScepCompletedAt) {
    return nil, false
	}
	return o.ScepCompletedAt, true
}

// HasScepCompletedAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) HasScepCompletedAt() bool {
	if o != nil && !isNil(o.ScepCompletedAt) {
		return true
	}

	return false
}

// SetScepCompletedAt gets a reference to the given time.Time and assigns it to the ScepCompletedAt field.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) SetScepCompletedAt(v time.Time) {
	o.ScepCompletedAt = &v
}

// GetLastConnectedAt returns the LastConnectedAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetLastConnectedAt() time.Time {
	if o == nil || isNil(o.LastConnectedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastConnectedAt
}

// GetLastConnectedAtOk returns a tuple with the LastConnectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) GetLastConnectedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastConnectedAt) {
    return nil, false
	}
	return o.LastConnectedAt, true
}

// HasLastConnectedAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) HasLastConnectedAt() bool {
	if o != nil && !isNil(o.LastConnectedAt) {
		return true
	}

	return false
}

// SetLastConnectedAt gets a reference to the given time.Time and assigns it to the LastConnectedAt field.
func (o *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) SetLastConnectedAt(v time.Time) {
	o.LastConnectedAt = &v
}

func (o NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TrustedAccessConfigId) {
		toSerialize["trustedAccessConfigId"] = o.TrustedAccessConfigId
	}
	if !isNil(o.DownloadedAt) {
		toSerialize["downloadedAt"] = o.DownloadedAt
	}
	if !isNil(o.ScepCompletedAt) {
		toSerialize["scepCompletedAt"] = o.ScepCompletedAt
	}
	if !isNil(o.LastConnectedAt) {
		toSerialize["lastConnectedAt"] = o.LastConnectedAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections struct {
	value *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections
	isSet bool
}

func (v NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) Get() *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections {
	return v.value
}

func (v *NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) Set(val *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections(val *NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) *NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections {
	return &NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


