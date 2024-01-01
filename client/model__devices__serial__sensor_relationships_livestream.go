/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSensorRelationshipsLivestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type DevicesSerialSensorRelationshipsLivestream struct {
	// An array of the related devices for the role
	RelatedDevices []DevicesSerialSensorRelationshipsLivestreamRelatedDevices `json:"relatedDevices,omitempty"`
}

// NewDevicesSerialSensorRelationshipsLivestream instantiates a new DevicesSerialSensorRelationshipsLivestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSensorRelationshipsLivestream() *DevicesSerialSensorRelationshipsLivestream {
	this := DevicesSerialSensorRelationshipsLivestream{}
	return &this
}

// NewDevicesSerialSensorRelationshipsLivestreamWithDefaults instantiates a new DevicesSerialSensorRelationshipsLivestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSensorRelationshipsLivestreamWithDefaults() *DevicesSerialSensorRelationshipsLivestream {
	this := DevicesSerialSensorRelationshipsLivestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *DevicesSerialSensorRelationshipsLivestream) GetRelatedDevices() []DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []DevicesSerialSensorRelationshipsLivestreamRelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorRelationshipsLivestream) GetRelatedDevicesOk() ([]DevicesSerialSensorRelationshipsLivestreamRelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *DevicesSerialSensorRelationshipsLivestream) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []DevicesSerialSensorRelationshipsLivestreamRelatedDevices and assigns it to the RelatedDevices field.
func (o *DevicesSerialSensorRelationshipsLivestream) SetRelatedDevices(v []DevicesSerialSensorRelationshipsLivestreamRelatedDevices) {
	o.RelatedDevices = v
}

func (o DevicesSerialSensorRelationshipsLivestream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSensorRelationshipsLivestream struct {
	value *DevicesSerialSensorRelationshipsLivestream
	isSet bool
}

func (v NullableDevicesSerialSensorRelationshipsLivestream) Get() *DevicesSerialSensorRelationshipsLivestream {
	return v.value
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream) Set(val *DevicesSerialSensorRelationshipsLivestream) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSensorRelationshipsLivestream) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSensorRelationshipsLivestream(val *DevicesSerialSensorRelationshipsLivestream) *NullableDevicesSerialSensorRelationshipsLivestream {
	return &NullableDevicesSerialSensorRelationshipsLivestream{value: val, isSet: true}
}

func (v NullableDevicesSerialSensorRelationshipsLivestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


