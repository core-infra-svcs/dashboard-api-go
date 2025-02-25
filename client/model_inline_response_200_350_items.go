/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200350Items struct for InlineResponse200350Items
type InlineResponse200350Items struct {
	// The cloud ID of the wireless LAN controller
	Serial *string `json:"serial,omitempty"`
	// The usages of layer 2 interfaces of the wireless LAN controller. Usage is in bytes
	Readings []InlineResponse200350Readings `json:"readings,omitempty"`
}

// NewInlineResponse200350Items instantiates a new InlineResponse200350Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200350Items() *InlineResponse200350Items {
	this := InlineResponse200350Items{}
	return &this
}

// NewInlineResponse200350ItemsWithDefaults instantiates a new InlineResponse200350Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200350ItemsWithDefaults() *InlineResponse200350Items {
	this := InlineResponse200350Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200350Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200350Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200350Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200350Items) SetSerial(v string) {
	o.Serial = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *InlineResponse200350Items) GetReadings() []InlineResponse200350Readings {
	if o == nil || isNil(o.Readings) {
		var ret []InlineResponse200350Readings
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200350Items) GetReadingsOk() ([]InlineResponse200350Readings, bool) {
	if o == nil || isNil(o.Readings) {
    return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *InlineResponse200350Items) HasReadings() bool {
	if o != nil && !isNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []InlineResponse200350Readings and assigns it to the Readings field.
func (o *InlineResponse200350Items) SetReadings(v []InlineResponse200350Readings) {
	o.Readings = v
}

func (o InlineResponse200350Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200350Items struct {
	value *InlineResponse200350Items
	isSet bool
}

func (v NullableInlineResponse200350Items) Get() *InlineResponse200350Items {
	return v.value
}

func (v *NullableInlineResponse200350Items) Set(val *InlineResponse200350Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200350Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200350Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200350Items(val *InlineResponse200350Items) *NullableInlineResponse200350Items {
	return &NullableInlineResponse200350Items{value: val, isSet: true}
}

func (v NullableInlineResponse200350Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200350Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


