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

// InlineResponse200336Readings struct for InlineResponse200336Readings
type InlineResponse200336Readings struct {
	// The name of the wireless LAN controller interface
	Name *string `json:"name,omitempty"`
	// The MAC address of the wireless controller interface
	Mac *string `json:"mac,omitempty"`
	// The volume of data, in bytes/sec, received by wireless controller interface
	Recv *int32 `json:"recv,omitempty"`
	// The volume of data, in bytes/sec, transmitted by wireless controller interface
	Send *int32 `json:"send,omitempty"`
}

// NewInlineResponse200336Readings instantiates a new InlineResponse200336Readings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200336Readings() *InlineResponse200336Readings {
	this := InlineResponse200336Readings{}
	return &this
}

// NewInlineResponse200336ReadingsWithDefaults instantiates a new InlineResponse200336Readings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200336ReadingsWithDefaults() *InlineResponse200336Readings {
	this := InlineResponse200336Readings{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200336Readings) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Readings) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200336Readings) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200336Readings) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200336Readings) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Readings) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200336Readings) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200336Readings) SetMac(v string) {
	o.Mac = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse200336Readings) GetRecv() int32 {
	if o == nil || isNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Readings) GetRecvOk() (*int32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse200336Readings) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *InlineResponse200336Readings) SetRecv(v int32) {
	o.Recv = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *InlineResponse200336Readings) GetSend() int32 {
	if o == nil || isNil(o.Send) {
		var ret int32
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336Readings) GetSendOk() (*int32, bool) {
	if o == nil || isNil(o.Send) {
    return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *InlineResponse200336Readings) HasSend() bool {
	if o != nil && !isNil(o.Send) {
		return true
	}

	return false
}

// SetSend gets a reference to the given int32 and assigns it to the Send field.
func (o *InlineResponse200336Readings) SetSend(v int32) {
	o.Send = &v
}

func (o InlineResponse200336Readings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !isNil(o.Send) {
		toSerialize["send"] = o.Send
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200336Readings struct {
	value *InlineResponse200336Readings
	isSet bool
}

func (v NullableInlineResponse200336Readings) Get() *InlineResponse200336Readings {
	return v.value
}

func (v *NullableInlineResponse200336Readings) Set(val *InlineResponse200336Readings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200336Readings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200336Readings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200336Readings(val *InlineResponse200336Readings) *NullableInlineResponse200336Readings {
	return &NullableInlineResponse200336Readings{value: val, isSet: true}
}

func (v NullableInlineResponse200336Readings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200336Readings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

