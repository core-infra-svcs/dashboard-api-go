/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject116 struct for InlineObject116
type InlineObject116 struct {
	// The wifiMac of the device to be modified.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The id of the device to be modified.
	Id *string `json:"id,omitempty"`
	// The serial of the device to be modified.
	Serial *string `json:"serial,omitempty"`
	DeviceFields NetworksNetworkIdSmDevicesFieldsDeviceFields `json:"deviceFields"`
}

// NewInlineObject116 instantiates a new InlineObject116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject116(deviceFields NetworksNetworkIdSmDevicesFieldsDeviceFields) *InlineObject116 {
	this := InlineObject116{}
	this.DeviceFields = deviceFields
	return &this
}

// NewInlineObject116WithDefaults instantiates a new InlineObject116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject116WithDefaults() *InlineObject116 {
	this := InlineObject116{}
	return &this
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineObject116) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineObject116) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineObject116) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject116) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject116) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject116) SetId(v string) {
	o.Id = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineObject116) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineObject116) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineObject116) SetSerial(v string) {
	o.Serial = &v
}

// GetDeviceFields returns the DeviceFields field value
func (o *InlineObject116) GetDeviceFields() NetworksNetworkIdSmDevicesFieldsDeviceFields {
	if o == nil {
		var ret NetworksNetworkIdSmDevicesFieldsDeviceFields
		return ret
	}

	return o.DeviceFields
}

// GetDeviceFieldsOk returns a tuple with the DeviceFields field value
// and a boolean to check if the value has been set.
func (o *InlineObject116) GetDeviceFieldsOk() (*NetworksNetworkIdSmDevicesFieldsDeviceFields, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DeviceFields, true
}

// SetDeviceFields sets field value
func (o *InlineObject116) SetDeviceFields(v NetworksNetworkIdSmDevicesFieldsDeviceFields) {
	o.DeviceFields = v
}

func (o InlineObject116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["deviceFields"] = o.DeviceFields
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject116 struct {
	value *InlineObject116
	isSet bool
}

func (v NullableInlineObject116) Get() *InlineObject116 {
	return v.value
}

func (v *NullableInlineObject116) Set(val *InlineObject116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject116(val *InlineObject116) *NullableInlineObject116 {
	return &NullableInlineObject116{value: val, isSet: true}
}

func (v NullableInlineObject116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


