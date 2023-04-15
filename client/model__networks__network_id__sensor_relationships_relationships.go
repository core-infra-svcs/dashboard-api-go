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

// NetworksNetworkIdSensorRelationshipsRelationships An object describing the relationships defined between the device and other devices
type NetworksNetworkIdSensorRelationshipsRelationships struct {
	Livestream *DevicesSerialSensorRelationshipsLivestream `json:"livestream,omitempty"`
}

// NewNetworksNetworkIdSensorRelationshipsRelationships instantiates a new NetworksNetworkIdSensorRelationshipsRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorRelationshipsRelationships() *NetworksNetworkIdSensorRelationshipsRelationships {
	this := NetworksNetworkIdSensorRelationshipsRelationships{}
	return &this
}

// NewNetworksNetworkIdSensorRelationshipsRelationshipsWithDefaults instantiates a new NetworksNetworkIdSensorRelationshipsRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorRelationshipsRelationshipsWithDefaults() *NetworksNetworkIdSensorRelationshipsRelationships {
	this := NetworksNetworkIdSensorRelationshipsRelationships{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorRelationshipsRelationships) GetLivestream() DevicesSerialSensorRelationshipsLivestream {
	if o == nil || isNil(o.Livestream) {
		var ret DevicesSerialSensorRelationshipsLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorRelationshipsRelationships) GetLivestreamOk() (*DevicesSerialSensorRelationshipsLivestream, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorRelationshipsRelationships) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given DevicesSerialSensorRelationshipsLivestream and assigns it to the Livestream field.
func (o *NetworksNetworkIdSensorRelationshipsRelationships) SetLivestream(v DevicesSerialSensorRelationshipsLivestream) {
	o.Livestream = &v
}

func (o NetworksNetworkIdSensorRelationshipsRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorRelationshipsRelationships struct {
	value *NetworksNetworkIdSensorRelationshipsRelationships
	isSet bool
}

func (v NullableNetworksNetworkIdSensorRelationshipsRelationships) Get() *NetworksNetworkIdSensorRelationshipsRelationships {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorRelationshipsRelationships) Set(val *NetworksNetworkIdSensorRelationshipsRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorRelationshipsRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorRelationshipsRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorRelationshipsRelationships(val *NetworksNetworkIdSensorRelationshipsRelationships) *NullableNetworksNetworkIdSensorRelationshipsRelationships {
	return &NullableNetworksNetworkIdSensorRelationshipsRelationships{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorRelationshipsRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorRelationshipsRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


