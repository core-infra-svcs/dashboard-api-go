/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 The origin of the prefix
type NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 struct {
	// Type of the origin
	Type *string `json:"type,omitempty"`
	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces,omitempty"`
}

// NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 instantiates a new NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	this := NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1{}
	return &this
}

// NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1WithDefaults instantiates a new NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1WithDefaults() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	this := NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) SetType(v string) {
	o.Type = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) GetInterfaces() []string {
	if o == nil || isNil(o.Interfaces) {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) GetInterfacesOk() ([]string, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) SetInterfaces(v []string) {
	o.Interfaces = v
}

func (o NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 struct {
	value *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1
	isSet bool
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) Get() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	return v.value
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) Set(val *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1(val *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	return &NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


