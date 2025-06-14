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

// NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate The auto locate position for this device
type NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate struct {
	// Whether or not this location should be saved as a user-defined anchor
	IsAnchor bool `json:"isAnchor"`
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate(isAnchor bool) *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate{}
	this.IsAnchor = isAnchor
	return &this
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocateWithDefaults instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocateWithDefaults() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate{}
	return &this
}

// GetIsAnchor returns the IsAnchor field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetIsAnchor() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAnchor
}

// GetIsAnchorOk returns a tuple with the IsAnchor field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetIsAnchorOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsAnchor, true
}

// SetIsAnchor sets field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) SetIsAnchor(v bool) {
	o.IsAnchor = v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) SetLng(v float32) {
	o.Lng = &v
}

func (o NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isAnchor"] = o.IsAnchor
	}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate struct {
	value *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) Get() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) Set(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate {
	return &NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdRecalculateAutoLocate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


