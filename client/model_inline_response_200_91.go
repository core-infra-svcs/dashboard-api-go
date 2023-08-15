/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20091 struct for InlineResponse20091
type InlineResponse20091 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Number of connected clients
	ClientCount *int32 `json:"clientCount,omitempty"`
}

// NewInlineResponse20091 instantiates a new InlineResponse20091 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20091() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// NewInlineResponse20091WithDefaults instantiates a new InlineResponse20091 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20091WithDefaults() *InlineResponse20091 {
	this := InlineResponse20091{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20091) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20091) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20091) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20091) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20091) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20091) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetClientCount returns the ClientCount field value if set, zero value otherwise.
func (o *InlineResponse20091) GetClientCount() int32 {
	if o == nil || isNil(o.ClientCount) {
		var ret int32
		return ret
	}
	return *o.ClientCount
}

// GetClientCountOk returns a tuple with the ClientCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20091) GetClientCountOk() (*int32, bool) {
	if o == nil || isNil(o.ClientCount) {
    return nil, false
	}
	return o.ClientCount, true
}

// HasClientCount returns a boolean if a field has been set.
func (o *InlineResponse20091) HasClientCount() bool {
	if o != nil && !isNil(o.ClientCount) {
		return true
	}

	return false
}

// SetClientCount gets a reference to the given int32 and assigns it to the ClientCount field.
func (o *InlineResponse20091) SetClientCount(v int32) {
	o.ClientCount = &v
}

func (o InlineResponse20091) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20091 struct {
	value *InlineResponse20091
	isSet bool
}

func (v NullableInlineResponse20091) Get() *InlineResponse20091 {
	return v.value
}

func (v *NullableInlineResponse20091) Set(val *InlineResponse20091) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20091) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20091) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20091(val *InlineResponse20091) *NullableInlineResponse20091 {
	return &NullableInlineResponse20091{value: val, isSet: true}
}

func (v NullableInlineResponse20091) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20091) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


