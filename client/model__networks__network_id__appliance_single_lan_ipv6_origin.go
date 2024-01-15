/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSingleLanIpv6Origin The origin of the prefix
type NetworksNetworkIdApplianceSingleLanIpv6Origin struct {
	// Type of the origin
	Type string `json:"type"`
	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces,omitempty"`
}

// NewNetworksNetworkIdApplianceSingleLanIpv6Origin instantiates a new NetworksNetworkIdApplianceSingleLanIpv6Origin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSingleLanIpv6Origin(type_ string) *NetworksNetworkIdApplianceSingleLanIpv6Origin {
	this := NetworksNetworkIdApplianceSingleLanIpv6Origin{}
	this.Type = type_
	return &this
}

// NewNetworksNetworkIdApplianceSingleLanIpv6OriginWithDefaults instantiates a new NetworksNetworkIdApplianceSingleLanIpv6Origin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSingleLanIpv6OriginWithDefaults() *NetworksNetworkIdApplianceSingleLanIpv6Origin {
	this := NetworksNetworkIdApplianceSingleLanIpv6Origin{}
	return &this
}

// GetType returns the Type field value
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) SetType(v string) {
	o.Type = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) GetInterfaces() []string {
	if o == nil || isNil(o.Interfaces) {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) GetInterfacesOk() ([]string, bool) {
	if o == nil || isNil(o.Interfaces) {
    return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) HasInterfaces() bool {
	if o != nil && !isNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6Origin) SetInterfaces(v []string) {
	o.Interfaces = v
}

func (o NetworksNetworkIdApplianceSingleLanIpv6Origin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSingleLanIpv6Origin struct {
	value *NetworksNetworkIdApplianceSingleLanIpv6Origin
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) Get() *NetworksNetworkIdApplianceSingleLanIpv6Origin {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) Set(val *NetworksNetworkIdApplianceSingleLanIpv6Origin) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSingleLanIpv6Origin(val *NetworksNetworkIdApplianceSingleLanIpv6Origin) *NullableNetworksNetworkIdApplianceSingleLanIpv6Origin {
	return &NullableNetworksNetworkIdApplianceSingleLanIpv6Origin{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6Origin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


