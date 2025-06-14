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

// InlineResponse200373Items struct for InlineResponse200373Items
type InlineResponse200373Items struct {
	// The cloud ID of the wireless LAN controller
	Serial *string `json:"serial,omitempty"`
	// The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes
	Readings []InlineResponse200373Readings `json:"readings,omitempty"`
}

// NewInlineResponse200373Items instantiates a new InlineResponse200373Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200373Items() *InlineResponse200373Items {
	this := InlineResponse200373Items{}
	return &this
}

// NewInlineResponse200373ItemsWithDefaults instantiates a new InlineResponse200373Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200373ItemsWithDefaults() *InlineResponse200373Items {
	this := InlineResponse200373Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200373Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200373Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200373Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200373Items) SetSerial(v string) {
	o.Serial = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *InlineResponse200373Items) GetReadings() []InlineResponse200373Readings {
	if o == nil || isNil(o.Readings) {
		var ret []InlineResponse200373Readings
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200373Items) GetReadingsOk() ([]InlineResponse200373Readings, bool) {
	if o == nil || isNil(o.Readings) {
    return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *InlineResponse200373Items) HasReadings() bool {
	if o != nil && !isNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []InlineResponse200373Readings and assigns it to the Readings field.
func (o *InlineResponse200373Items) SetReadings(v []InlineResponse200373Readings) {
	o.Readings = v
}

func (o InlineResponse200373Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200373Items struct {
	value *InlineResponse200373Items
	isSet bool
}

func (v NullableInlineResponse200373Items) Get() *InlineResponse200373Items {
	return v.value
}

func (v *NullableInlineResponse200373Items) Set(val *InlineResponse200373Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200373Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200373Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200373Items(val *InlineResponse200373Items) *NullableInlineResponse200373Items {
	return &NullableInlineResponse200373Items{value: val, isSet: true}
}

func (v NullableInlineResponse200373Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200373Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


