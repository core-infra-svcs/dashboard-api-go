/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200340Interfaces struct for InlineResponse200340Interfaces
type InlineResponse200340Interfaces struct {
	// The name of the wireless LAN controller interface
	Name *string `json:"name,omitempty"`
	// The status of packets counter on the interfaces of the wireless LAN controller
	Readings []InlineResponse200340Readings `json:"readings,omitempty"`
}

// NewInlineResponse200340Interfaces instantiates a new InlineResponse200340Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200340Interfaces() *InlineResponse200340Interfaces {
	this := InlineResponse200340Interfaces{}
	return &this
}

// NewInlineResponse200340InterfacesWithDefaults instantiates a new InlineResponse200340Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200340InterfacesWithDefaults() *InlineResponse200340Interfaces {
	this := InlineResponse200340Interfaces{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200340Interfaces) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200340Interfaces) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200340Interfaces) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200340Interfaces) SetName(v string) {
	o.Name = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *InlineResponse200340Interfaces) GetReadings() []InlineResponse200340Readings {
	if o == nil || isNil(o.Readings) {
		var ret []InlineResponse200340Readings
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200340Interfaces) GetReadingsOk() ([]InlineResponse200340Readings, bool) {
	if o == nil || isNil(o.Readings) {
    return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *InlineResponse200340Interfaces) HasReadings() bool {
	if o != nil && !isNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []InlineResponse200340Readings and assigns it to the Readings field.
func (o *InlineResponse200340Interfaces) SetReadings(v []InlineResponse200340Readings) {
	o.Readings = v
}

func (o InlineResponse200340Interfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200340Interfaces struct {
	value *InlineResponse200340Interfaces
	isSet bool
}

func (v NullableInlineResponse200340Interfaces) Get() *InlineResponse200340Interfaces {
	return v.value
}

func (v *NullableInlineResponse200340Interfaces) Set(val *InlineResponse200340Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200340Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200340Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200340Interfaces(val *InlineResponse200340Interfaces) *NullableInlineResponse200340Interfaces {
	return &NullableInlineResponse200340Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse200340Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200340Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


