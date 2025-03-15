/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2012Request Cable test request parameters
type InlineResponse2012Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// A list of ports for which to perform the cable test.
	Ports []string `json:"ports,omitempty"`
}

// NewInlineResponse2012Request instantiates a new InlineResponse2012Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2012Request() *InlineResponse2012Request {
	this := InlineResponse2012Request{}
	return &this
}

// NewInlineResponse2012RequestWithDefaults instantiates a new InlineResponse2012Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2012RequestWithDefaults() *InlineResponse2012Request {
	this := InlineResponse2012Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2012Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2012Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2012Request) SetSerial(v string) {
	o.Serial = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse2012Request) GetPorts() []string {
	if o == nil || isNil(o.Ports) {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012Request) GetPortsOk() ([]string, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse2012Request) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *InlineResponse2012Request) SetPorts(v []string) {
	o.Ports = v
}

func (o InlineResponse2012Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2012Request struct {
	value *InlineResponse2012Request
	isSet bool
}

func (v NullableInlineResponse2012Request) Get() *InlineResponse2012Request {
	return v.value
}

func (v *NullableInlineResponse2012Request) Set(val *InlineResponse2012Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2012Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2012Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2012Request(val *InlineResponse2012Request) *NullableInlineResponse2012Request {
	return &NullableInlineResponse2012Request{value: val, isSet: true}
}

func (v NullableInlineResponse2012Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2012Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


