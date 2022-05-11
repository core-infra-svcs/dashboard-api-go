/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject83 struct for InlineObject83
type InlineObject83 struct {
	// The wifiMac of the device to be modified.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The id of the device to be modified.
	Id *string `json:"id,omitempty"`
	// The serial of the device to be modified.
	Serial *string `json:"serial,omitempty"`
	DeviceFields NetworksNetworkIdSmDevicesFieldsDeviceFields `json:"deviceFields"`
}

// NewInlineObject83 instantiates a new InlineObject83 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject83(deviceFields NetworksNetworkIdSmDevicesFieldsDeviceFields) *InlineObject83 {
	this := InlineObject83{}
	this.DeviceFields = deviceFields
	return &this
}

// NewInlineObject83WithDefaults instantiates a new InlineObject83 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject83WithDefaults() *InlineObject83 {
	this := InlineObject83{}
	return &this
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineObject83) GetWifiMac() string {
	if o == nil || o.WifiMac == nil {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetWifiMacOk() (*string, bool) {
	if o == nil || o.WifiMac == nil {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineObject83) HasWifiMac() bool {
	if o != nil && o.WifiMac != nil {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineObject83) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject83) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject83) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject83) SetId(v string) {
	o.Id = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineObject83) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineObject83) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineObject83) SetSerial(v string) {
	o.Serial = &v
}

// GetDeviceFields returns the DeviceFields field value
func (o *InlineObject83) GetDeviceFields() NetworksNetworkIdSmDevicesFieldsDeviceFields {
	if o == nil {
		var ret NetworksNetworkIdSmDevicesFieldsDeviceFields
		return ret
	}

	return o.DeviceFields
}

// GetDeviceFieldsOk returns a tuple with the DeviceFields field value
// and a boolean to check if the value has been set.
func (o *InlineObject83) GetDeviceFieldsOk() (*NetworksNetworkIdSmDevicesFieldsDeviceFields, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeviceFields, true
}

// SetDeviceFields sets field value
func (o *InlineObject83) SetDeviceFields(v NetworksNetworkIdSmDevicesFieldsDeviceFields) {
	o.DeviceFields = v
}

func (o InlineObject83) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WifiMac != nil {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["deviceFields"] = o.DeviceFields
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject83 struct {
	value *InlineObject83
	isSet bool
}

func (v NullableInlineObject83) Get() *InlineObject83 {
	return v.value
}

func (v *NullableInlineObject83) Set(val *InlineObject83) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject83) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject83) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject83(val *InlineObject83) *NullableInlineObject83 {
	return &NullableInlineObject83{value: val, isSet: true}
}

func (v NullableInlineObject83) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject83) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


