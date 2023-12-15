/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20056 struct for InlineResponse20056
type InlineResponse20056 struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// An array of tags associated with the device.
	Tags []string `json:"tags,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
}

// NewInlineResponse20056 instantiates a new InlineResponse20056 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20056() *InlineResponse20056 {
	this := InlineResponse20056{}
	return &this
}

// NewInlineResponse20056WithDefaults instantiates a new InlineResponse20056 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20056WithDefaults() *InlineResponse20056 {
	this := InlineResponse20056{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20056) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20056) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20056) SetId(v string) {
	o.Id = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20056) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20056) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse20056) SetTags(v []string) {
	o.Tags = v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *InlineResponse20056) GetWifiMac() string {
	if o == nil || isNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056) GetWifiMacOk() (*string, bool) {
	if o == nil || isNil(o.WifiMac) {
    return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *InlineResponse20056) HasWifiMac() bool {
	if o != nil && !isNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *InlineResponse20056) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20056) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20056) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20056) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20056) SetSerial(v string) {
	o.Serial = &v
}

func (o InlineResponse20056) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20056 struct {
	value *InlineResponse20056
	isSet bool
}

func (v NullableInlineResponse20056) Get() *InlineResponse20056 {
	return v.value
}

func (v *NullableInlineResponse20056) Set(val *InlineResponse20056) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20056) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20056) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20056(val *InlineResponse20056) *NullableInlineResponse20056 {
	return &NullableInlineResponse20056{value: val, isSet: true}
}

func (v NullableInlineResponse20056) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20056) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


