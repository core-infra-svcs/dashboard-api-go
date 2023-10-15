/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSensorRelationshipsLivestream1 A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type DevicesSerialSensorRelationshipsLivestream1 struct {
	// An array of the related devices for the role
	RelatedDevices []DevicesSerialSensorRelationshipsLivestream1RelatedDevices `json:"relatedDevices,omitempty"`
}

// NewDevicesSerialSensorRelationshipsLivestream1 instantiates a new DevicesSerialSensorRelationshipsLivestream1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSensorRelationshipsLivestream1() *DevicesSerialSensorRelationshipsLivestream1 {
	this := DevicesSerialSensorRelationshipsLivestream1{}
	return &this
}

// NewDevicesSerialSensorRelationshipsLivestream1WithDefaults instantiates a new DevicesSerialSensorRelationshipsLivestream1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSensorRelationshipsLivestream1WithDefaults() *DevicesSerialSensorRelationshipsLivestream1 {
	this := DevicesSerialSensorRelationshipsLivestream1{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *DevicesSerialSensorRelationshipsLivestream1) GetRelatedDevices() []DevicesSerialSensorRelationshipsLivestream1RelatedDevices {
	if o == nil || isNil(o.RelatedDevices) {
		var ret []DevicesSerialSensorRelationshipsLivestream1RelatedDevices
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorRelationshipsLivestream1) GetRelatedDevicesOk() ([]DevicesSerialSensorRelationshipsLivestream1RelatedDevices, bool) {
	if o == nil || isNil(o.RelatedDevices) {
    return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *DevicesSerialSensorRelationshipsLivestream1) HasRelatedDevices() bool {
	if o != nil && !isNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []DevicesSerialSensorRelationshipsLivestream1RelatedDevices and assigns it to the RelatedDevices field.
func (o *DevicesSerialSensorRelationshipsLivestream1) SetRelatedDevices(v []DevicesSerialSensorRelationshipsLivestream1RelatedDevices) {
	o.RelatedDevices = v
}

func (o DevicesSerialSensorRelationshipsLivestream1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSensorRelationshipsLivestream1 struct {
	value *DevicesSerialSensorRelationshipsLivestream1
	isSet bool
}

func (v NullableDevicesSerialSensorRelationshipsLivestream1) Get() *DevicesSerialSensorRelationshipsLivestream1 {
	return v.value
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream1) Set(val *DevicesSerialSensorRelationshipsLivestream1) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSensorRelationshipsLivestream1) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSensorRelationshipsLivestream1(val *DevicesSerialSensorRelationshipsLivestream1) *NullableDevicesSerialSensorRelationshipsLivestream1 {
	return &NullableDevicesSerialSensorRelationshipsLivestream1{value: val, isSet: true}
}

func (v NullableDevicesSerialSensorRelationshipsLivestream1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSensorRelationshipsLivestream1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


