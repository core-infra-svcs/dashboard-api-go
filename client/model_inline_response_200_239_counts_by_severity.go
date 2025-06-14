/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200239CountsBySeverity struct for InlineResponse200239CountsBySeverity
type InlineResponse200239CountsBySeverity struct {
	// Severity Type
	Type string `json:"type"`
	// Total count of the given severity type
	Count int32 `json:"count"`
}

// NewInlineResponse200239CountsBySeverity instantiates a new InlineResponse200239CountsBySeverity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200239CountsBySeverity(type_ string, count int32) *InlineResponse200239CountsBySeverity {
	this := InlineResponse200239CountsBySeverity{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewInlineResponse200239CountsBySeverityWithDefaults instantiates a new InlineResponse200239CountsBySeverity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200239CountsBySeverityWithDefaults() *InlineResponse200239CountsBySeverity {
	this := InlineResponse200239CountsBySeverity{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse200239CountsBySeverity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200239CountsBySeverity) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200239CountsBySeverity) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *InlineResponse200239CountsBySeverity) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200239CountsBySeverity) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InlineResponse200239CountsBySeverity) SetCount(v int32) {
	o.Count = v
}

func (o InlineResponse200239CountsBySeverity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200239CountsBySeverity struct {
	value *InlineResponse200239CountsBySeverity
	isSet bool
}

func (v NullableInlineResponse200239CountsBySeverity) Get() *InlineResponse200239CountsBySeverity {
	return v.value
}

func (v *NullableInlineResponse200239CountsBySeverity) Set(val *InlineResponse200239CountsBySeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200239CountsBySeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200239CountsBySeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200239CountsBySeverity(val *InlineResponse200239CountsBySeverity) *NullableInlineResponse200239CountsBySeverity {
	return &NullableInlineResponse200239CountsBySeverity{value: val, isSet: true}
}

func (v NullableInlineResponse200239CountsBySeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200239CountsBySeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


