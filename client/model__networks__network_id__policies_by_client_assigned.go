/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdPoliciesByClientAssigned struct for NetworksNetworkIdPoliciesByClientAssigned
type NetworksNetworkIdPoliciesByClientAssigned struct {
	// name of policy
	Name *string `json:"name,omitempty"`
	// type of policy
	Type *string `json:"type,omitempty"`
	// ssid
	Ssid *[]NetworksNetworkIdPoliciesByClientSsid `json:"ssid,omitempty"`
}

// NewNetworksNetworkIdPoliciesByClientAssigned instantiates a new NetworksNetworkIdPoliciesByClientAssigned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdPoliciesByClientAssigned() *NetworksNetworkIdPoliciesByClientAssigned {
	this := NetworksNetworkIdPoliciesByClientAssigned{}
	return &this
}

// NewNetworksNetworkIdPoliciesByClientAssignedWithDefaults instantiates a new NetworksNetworkIdPoliciesByClientAssigned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdPoliciesByClientAssignedWithDefaults() *NetworksNetworkIdPoliciesByClientAssigned {
	this := NetworksNetworkIdPoliciesByClientAssigned{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdPoliciesByClientAssigned) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdPoliciesByClientAssigned) SetType(v string) {
	o.Type = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetSsid() []NetworksNetworkIdPoliciesByClientSsid {
	if o == nil || o.Ssid == nil {
		var ret []NetworksNetworkIdPoliciesByClientSsid
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) GetSsidOk() (*[]NetworksNetworkIdPoliciesByClientSsid, bool) {
	if o == nil || o.Ssid == nil {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NetworksNetworkIdPoliciesByClientAssigned) HasSsid() bool {
	if o != nil && o.Ssid != nil {
		return true
	}

	return false
}

// SetSsid gets a reference to the given []NetworksNetworkIdPoliciesByClientSsid and assigns it to the Ssid field.
func (o *NetworksNetworkIdPoliciesByClientAssigned) SetSsid(v []NetworksNetworkIdPoliciesByClientSsid) {
	o.Ssid = &v
}

func (o NetworksNetworkIdPoliciesByClientAssigned) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Ssid != nil {
		toSerialize["ssid"] = o.Ssid
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdPoliciesByClientAssigned struct {
	value *NetworksNetworkIdPoliciesByClientAssigned
	isSet bool
}

func (v NullableNetworksNetworkIdPoliciesByClientAssigned) Get() *NetworksNetworkIdPoliciesByClientAssigned {
	return v.value
}

func (v *NullableNetworksNetworkIdPoliciesByClientAssigned) Set(val *NetworksNetworkIdPoliciesByClientAssigned) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdPoliciesByClientAssigned) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdPoliciesByClientAssigned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdPoliciesByClientAssigned(val *NetworksNetworkIdPoliciesByClientAssigned) *NullableNetworksNetworkIdPoliciesByClientAssigned {
	return &NullableNetworksNetworkIdPoliciesByClientAssigned{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdPoliciesByClientAssigned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdPoliciesByClientAssigned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


