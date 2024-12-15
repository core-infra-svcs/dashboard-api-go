/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices struct for NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices
type NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices struct {
	// Serial for device to publish position for
	Serial string `json:"serial"`
	// Latitude
	Lat float32 `json:"lat"`
	// Longitude
	Lng float32 `json:"lng"`
	AutoLocate *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate `json:"autoLocate,omitempty"`
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices(serial string, lat float32, lng float32) *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices{}
	this.Serial = serial
	this.Lat = lat
	this.Lng = lng
	return &this
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevicesWithDefaults instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevicesWithDefaults() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) SetSerial(v string) {
	o.Serial = v
}

// GetLat returns the Lat field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetLat() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetLatOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) SetLat(v float32) {
	o.Lat = v
}

// GetLng returns the Lng field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetLng() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Lng
}

// GetLngOk returns a tuple with the Lng field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetLngOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lng, true
}

// SetLng sets field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) SetLng(v float32) {
	o.Lng = v
}

// GetAutoLocate returns the AutoLocate field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetAutoLocate() NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate {
	if o == nil || isNil(o.AutoLocate) {
		var ret NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate
		return ret
	}
	return *o.AutoLocate
}

// GetAutoLocateOk returns a tuple with the AutoLocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) GetAutoLocateOk() (*NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate, bool) {
	if o == nil || isNil(o.AutoLocate) {
    return nil, false
	}
	return o.AutoLocate, true
}

// HasAutoLocate returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) HasAutoLocate() bool {
	if o != nil && !isNil(o.AutoLocate) {
		return true
	}

	return false
}

// SetAutoLocate gets a reference to the given NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate and assigns it to the AutoLocate field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) SetAutoLocate(v NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishAutoLocate) {
	o.AutoLocate = &v
}

func (o NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["lat"] = o.Lat
	}
	if true {
		toSerialize["lng"] = o.Lng
	}
	if !isNil(o.AutoLocate) {
		toSerialize["autoLocate"] = o.AutoLocate
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices struct {
	value *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) Get() *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) Set(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices(val *NetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices {
	return &NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsJobIdPublishDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


