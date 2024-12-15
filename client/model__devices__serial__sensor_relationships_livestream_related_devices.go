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

// DevicesSerialSensorRelationshipsLivestreamRelatedDevices struct for DevicesSerialSensorRelationshipsLivestreamRelatedDevices
type DevicesSerialSensorRelationshipsLivestreamRelatedDevices struct {
	// The serial of the related device
	Serial string `json:"serial"`
}

// NewDevicesSerialSensorRelationshipsLivestreamRelatedDevices instantiates a new DevicesSerialSensorRelationshipsLivestreamRelatedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSensorRelationshipsLivestreamRelatedDevices(serial string) *DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	this := DevicesSerialSensorRelationshipsLivestreamRelatedDevices{}
	this.Serial = serial
	return &this
}

// NewDevicesSerialSensorRelationshipsLivestreamRelatedDevicesWithDefaults instantiates a new DevicesSerialSensorRelationshipsLivestreamRelatedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSensorRelationshipsLivestreamRelatedDevicesWithDefaults() *DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	this := DevicesSerialSensorRelationshipsLivestreamRelatedDevices{}
	return &this
}

// GetSerial returns the Serial field value
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) SetSerial(v string) {
	o.Serial = v
}

func (o DevicesSerialSensorRelationshipsLivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices struct {
	value *DevicesSerialSensorRelationshipsLivestreamRelatedDevices
	isSet bool
}

func (v NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) Get() *DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	return v.value
}

func (v *NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) Set(val *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices(val *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) *NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	return &NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices{value: val, isSet: true}
}

func (v NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSensorRelationshipsLivestreamRelatedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


