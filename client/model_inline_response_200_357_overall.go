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

// InlineResponse200357Overall The overall usage of all queried interfaces of the wireless LAN controller
type InlineResponse200357Overall struct {
	// The total usage of all queried interfaces during the interval, unit is bit/sec
	Total *int32 `json:"total,omitempty"`
	// The received usage of all queried interfaces during the interval, unit is bit/sec
	Recv *int32 `json:"recv,omitempty"`
	// The sent usage of all queried interfaces during the interval, unit is bit/sec
	Send *int32 `json:"send,omitempty"`
}

// NewInlineResponse200357Overall instantiates a new InlineResponse200357Overall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200357Overall() *InlineResponse200357Overall {
	this := InlineResponse200357Overall{}
	return &this
}

// NewInlineResponse200357OverallWithDefaults instantiates a new InlineResponse200357Overall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200357OverallWithDefaults() *InlineResponse200357Overall {
	this := InlineResponse200357Overall{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200357Overall) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200357Overall) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200357Overall) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200357Overall) SetTotal(v int32) {
	o.Total = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse200357Overall) GetRecv() int32 {
	if o == nil || isNil(o.Recv) {
		var ret int32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200357Overall) GetRecvOk() (*int32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse200357Overall) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given int32 and assigns it to the Recv field.
func (o *InlineResponse200357Overall) SetRecv(v int32) {
	o.Recv = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *InlineResponse200357Overall) GetSend() int32 {
	if o == nil || isNil(o.Send) {
		var ret int32
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200357Overall) GetSendOk() (*int32, bool) {
	if o == nil || isNil(o.Send) {
    return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *InlineResponse200357Overall) HasSend() bool {
	if o != nil && !isNil(o.Send) {
		return true
	}

	return false
}

// SetSend gets a reference to the given int32 and assigns it to the Send field.
func (o *InlineResponse200357Overall) SetSend(v int32) {
	o.Send = &v
}

func (o InlineResponse200357Overall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !isNil(o.Send) {
		toSerialize["send"] = o.Send
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200357Overall struct {
	value *InlineResponse200357Overall
	isSet bool
}

func (v NullableInlineResponse200357Overall) Get() *InlineResponse200357Overall {
	return v.value
}

func (v *NullableInlineResponse200357Overall) Set(val *InlineResponse200357Overall) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200357Overall) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200357Overall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200357Overall(val *InlineResponse200357Overall) *NullableInlineResponse200357Overall {
	return &NullableInlineResponse200357Overall{value: val, isSet: true}
}

func (v NullableInlineResponse200357Overall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200357Overall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


