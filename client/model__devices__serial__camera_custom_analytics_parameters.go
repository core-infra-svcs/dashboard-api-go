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

// DevicesSerialCameraCustomAnalyticsParameters struct for DevicesSerialCameraCustomAnalyticsParameters
type DevicesSerialCameraCustomAnalyticsParameters struct {
	// Name of the parameter
	Name string `json:"name"`
	// Value of the parameter
	Value string `json:"value"`
}

// NewDevicesSerialCameraCustomAnalyticsParameters instantiates a new DevicesSerialCameraCustomAnalyticsParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCameraCustomAnalyticsParameters(name string, value string) *DevicesSerialCameraCustomAnalyticsParameters {
	this := DevicesSerialCameraCustomAnalyticsParameters{}
	this.Name = name
	this.Value = value
	return &this
}

// NewDevicesSerialCameraCustomAnalyticsParametersWithDefaults instantiates a new DevicesSerialCameraCustomAnalyticsParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCameraCustomAnalyticsParametersWithDefaults() *DevicesSerialCameraCustomAnalyticsParameters {
	this := DevicesSerialCameraCustomAnalyticsParameters{}
	return &this
}

// GetName returns the Name field value
func (o *DevicesSerialCameraCustomAnalyticsParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraCustomAnalyticsParameters) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DevicesSerialCameraCustomAnalyticsParameters) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *DevicesSerialCameraCustomAnalyticsParameters) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCameraCustomAnalyticsParameters) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DevicesSerialCameraCustomAnalyticsParameters) SetValue(v string) {
	o.Value = v
}

func (o DevicesSerialCameraCustomAnalyticsParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCameraCustomAnalyticsParameters struct {
	value *DevicesSerialCameraCustomAnalyticsParameters
	isSet bool
}

func (v NullableDevicesSerialCameraCustomAnalyticsParameters) Get() *DevicesSerialCameraCustomAnalyticsParameters {
	return v.value
}

func (v *NullableDevicesSerialCameraCustomAnalyticsParameters) Set(val *DevicesSerialCameraCustomAnalyticsParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCameraCustomAnalyticsParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCameraCustomAnalyticsParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCameraCustomAnalyticsParameters(val *DevicesSerialCameraCustomAnalyticsParameters) *NullableDevicesSerialCameraCustomAnalyticsParameters {
	return &NullableDevicesSerialCameraCustomAnalyticsParameters{value: val, isSet: true}
}

func (v NullableDevicesSerialCameraCustomAnalyticsParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCameraCustomAnalyticsParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


