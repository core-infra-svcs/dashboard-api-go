/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20087 struct for InlineResponse20087
type InlineResponse20087 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Number of connected clients
	ClientCount *int32 `json:"clientCount,omitempty"`
}

// NewInlineResponse20087 instantiates a new InlineResponse20087 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087WithDefaults() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20087) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20087) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20087) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20087) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20087) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20087) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetClientCount returns the ClientCount field value if set, zero value otherwise.
func (o *InlineResponse20087) GetClientCount() int32 {
	if o == nil || isNil(o.ClientCount) {
		var ret int32
		return ret
	}
	return *o.ClientCount
}

// GetClientCountOk returns a tuple with the ClientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetClientCountOk() (*int32, bool) {
	if o == nil || isNil(o.ClientCount) {
    return nil, false
	}
	return o.ClientCount, true
}

// HasClientCount returns a boolean if a field has been set.
func (o *InlineResponse20087) HasClientCount() bool {
	if o != nil && !isNil(o.ClientCount) {
		return true
	}

	return false
}

// SetClientCount gets a reference to the given int32 and assigns it to the ClientCount field.
func (o *InlineResponse20087) SetClientCount(v int32) {
	o.ClientCount = &v
}

func (o InlineResponse20087) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.ClientCount) {
		toSerialize["clientCount"] = o.ClientCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20087 struct {
	value *InlineResponse20087
	isSet bool
}

func (v NullableInlineResponse20087) Get() *InlineResponse20087 {
	return v.value
}

func (v *NullableInlineResponse20087) Set(val *InlineResponse20087) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087(val *InlineResponse20087) *NullableInlineResponse20087 {
	return &NullableInlineResponse20087{value: val, isSet: true}
}

func (v NullableInlineResponse20087) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


