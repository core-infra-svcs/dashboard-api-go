/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFloorPlansTopLeftCorner The longitude and latitude of the top left corner of your floor plan.
type NetworksNetworkIdFloorPlansTopLeftCorner struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewNetworksNetworkIdFloorPlansTopLeftCorner instantiates a new NetworksNetworkIdFloorPlansTopLeftCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansTopLeftCorner() *NetworksNetworkIdFloorPlansTopLeftCorner {
	this := NetworksNetworkIdFloorPlansTopLeftCorner{}
	return &this
}

// NewNetworksNetworkIdFloorPlansTopLeftCornerWithDefaults instantiates a new NetworksNetworkIdFloorPlansTopLeftCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansTopLeftCornerWithDefaults() *NetworksNetworkIdFloorPlansTopLeftCorner {
	this := NetworksNetworkIdFloorPlansTopLeftCorner{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *NetworksNetworkIdFloorPlansTopLeftCorner) SetLng(v float32) {
	o.Lng = &v
}

func (o NetworksNetworkIdFloorPlansTopLeftCorner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansTopLeftCorner struct {
	value *NetworksNetworkIdFloorPlansTopLeftCorner
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansTopLeftCorner) Get() *NetworksNetworkIdFloorPlansTopLeftCorner {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansTopLeftCorner) Set(val *NetworksNetworkIdFloorPlansTopLeftCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansTopLeftCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansTopLeftCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansTopLeftCorner(val *NetworksNetworkIdFloorPlansTopLeftCorner) *NullableNetworksNetworkIdFloorPlansTopLeftCorner {
	return &NullableNetworksNetworkIdFloorPlansTopLeftCorner{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansTopLeftCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansTopLeftCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


