/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSensorRelationshipsLivestreamRelatedDevices struct for DevicesSerialSensorRelationshipsLivestreamRelatedDevices
type DevicesSerialSensorRelationshipsLivestreamRelatedDevices struct {
	// The serial of the related device
	Serial *string `json:"serial,omitempty"`
	// The product type of the related device
	ProductType *string `json:"productType,omitempty"`
}

// NewDevicesSerialSensorRelationshipsLivestreamRelatedDevices instantiates a new DevicesSerialSensorRelationshipsLivestreamRelatedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSensorRelationshipsLivestreamRelatedDevices() *DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	this := DevicesSerialSensorRelationshipsLivestreamRelatedDevices{}
	return &this
}

// NewDevicesSerialSensorRelationshipsLivestreamRelatedDevicesWithDefaults instantiates a new DevicesSerialSensorRelationshipsLivestreamRelatedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSensorRelationshipsLivestreamRelatedDevicesWithDefaults() *DevicesSerialSensorRelationshipsLivestreamRelatedDevices {
	this := DevicesSerialSensorRelationshipsLivestreamRelatedDevices{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) SetSerial(v string) {
	o.Serial = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *DevicesSerialSensorRelationshipsLivestreamRelatedDevices) SetProductType(v string) {
	o.ProductType = &v
}

func (o DevicesSerialSensorRelationshipsLivestreamRelatedDevices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
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


