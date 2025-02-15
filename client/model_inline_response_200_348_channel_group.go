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

// InlineResponse200348ChannelGroup The channel group of this wireless LAN controller interface
type InlineResponse200348ChannelGroup struct {
	// The interface channel group number
	Number *int32 `json:"number,omitempty"`
}

// NewInlineResponse200348ChannelGroup instantiates a new InlineResponse200348ChannelGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200348ChannelGroup() *InlineResponse200348ChannelGroup {
	this := InlineResponse200348ChannelGroup{}
	return &this
}

// NewInlineResponse200348ChannelGroupWithDefaults instantiates a new InlineResponse200348ChannelGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200348ChannelGroupWithDefaults() *InlineResponse200348ChannelGroup {
	this := InlineResponse200348ChannelGroup{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InlineResponse200348ChannelGroup) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200348ChannelGroup) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineResponse200348ChannelGroup) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *InlineResponse200348ChannelGroup) SetNumber(v int32) {
	o.Number = &v
}

func (o InlineResponse200348ChannelGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200348ChannelGroup struct {
	value *InlineResponse200348ChannelGroup
	isSet bool
}

func (v NullableInlineResponse200348ChannelGroup) Get() *InlineResponse200348ChannelGroup {
	return v.value
}

func (v *NullableInlineResponse200348ChannelGroup) Set(val *InlineResponse200348ChannelGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200348ChannelGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200348ChannelGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200348ChannelGroup(val *InlineResponse200348ChannelGroup) *NullableInlineResponse200348ChannelGroup {
	return &NullableInlineResponse200348ChannelGroup{value: val, isSet: true}
}

func (v NullableInlineResponse200348ChannelGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200348ChannelGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


