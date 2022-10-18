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

// NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin WAN1/WAN2/Independent prefix.
type NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin struct {
	// Origin type
	Type *string `json:"type,omitempty"`
	// Uplink provided or independent
	Interfaces *[]string `json:"interfaces,omitempty"`
}

// NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin instantiates a new NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin {
	this := NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin{}
	return &this
}

// NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOriginWithDefaults instantiates a new NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOriginWithDefaults() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin {
	this := NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) SetType(v string) {
	o.Type = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) GetInterfaces() []string {
	if o == nil || o.Interfaces == nil {
		var ret []string
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) GetInterfacesOk() (*[]string, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) SetInterfaces(v []string) {
	o.Interfaces = &v
}

func (o NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Interfaces != nil {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin struct {
	value *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin
	isSet bool
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) Get() *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin {
	return v.value
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) Set(val *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin(val *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin {
	return &NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


