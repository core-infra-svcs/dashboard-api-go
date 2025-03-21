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

// InlineResponse200348CountsByStatus Client counts by its status
type InlineResponse200348CountsByStatus struct {
	// Number of connected clients
	Online *int32 `json:"online,omitempty"`
}

// NewInlineResponse200348CountsByStatus instantiates a new InlineResponse200348CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200348CountsByStatus() *InlineResponse200348CountsByStatus {
	this := InlineResponse200348CountsByStatus{}
	return &this
}

// NewInlineResponse200348CountsByStatusWithDefaults instantiates a new InlineResponse200348CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200348CountsByStatusWithDefaults() *InlineResponse200348CountsByStatus {
	this := InlineResponse200348CountsByStatus{}
	return &this
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *InlineResponse200348CountsByStatus) GetOnline() int32 {
	if o == nil || isNil(o.Online) {
		var ret int32
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200348CountsByStatus) GetOnlineOk() (*int32, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *InlineResponse200348CountsByStatus) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given int32 and assigns it to the Online field.
func (o *InlineResponse200348CountsByStatus) SetOnline(v int32) {
	o.Online = &v
}

func (o InlineResponse200348CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200348CountsByStatus struct {
	value *InlineResponse200348CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200348CountsByStatus) Get() *InlineResponse200348CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200348CountsByStatus) Set(val *InlineResponse200348CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200348CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200348CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200348CountsByStatus(val *InlineResponse200348CountsByStatus) *NullableInlineResponse200348CountsByStatus {
	return &NullableInlineResponse200348CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200348CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200348CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


