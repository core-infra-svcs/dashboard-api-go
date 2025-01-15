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

// InlineResponse2017Request The parameters of the Wake-on-LAN request
type InlineResponse2017Request struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// The target's VLAN (1 to 4094)
	VlanId *int32 `json:"vlanId,omitempty"`
	// The target's MAC address
	Mac *string `json:"mac,omitempty"`
}

// NewInlineResponse2017Request instantiates a new InlineResponse2017Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2017Request() *InlineResponse2017Request {
	this := InlineResponse2017Request{}
	return &this
}

// NewInlineResponse2017RequestWithDefaults instantiates a new InlineResponse2017Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2017RequestWithDefaults() *InlineResponse2017Request {
	this := InlineResponse2017Request{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse2017Request) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017Request) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse2017Request) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse2017Request) SetSerial(v string) {
	o.Serial = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse2017Request) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017Request) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse2017Request) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineResponse2017Request) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse2017Request) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017Request) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse2017Request) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse2017Request) SetMac(v string) {
	o.Mac = &v
}

func (o InlineResponse2017Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2017Request struct {
	value *InlineResponse2017Request
	isSet bool
}

func (v NullableInlineResponse2017Request) Get() *InlineResponse2017Request {
	return v.value
}

func (v *NullableInlineResponse2017Request) Set(val *InlineResponse2017Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2017Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2017Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2017Request(val *InlineResponse2017Request) *NullableInlineResponse2017Request {
	return &NullableInlineResponse2017Request{value: val, isSet: true}
}

func (v NullableInlineResponse2017Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2017Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

