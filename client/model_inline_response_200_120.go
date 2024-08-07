/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200120 struct for InlineResponse200120
type InlineResponse200120 struct {
	// When the device was first seen as connected to the internet in each connection.
	FirstSeenAt *string `json:"firstSeenAt,omitempty"`
	// When the device was last seen as connected to the internet in each connection.
	LastSeenAt *string `json:"lastSeenAt,omitempty"`
}

// NewInlineResponse200120 instantiates a new InlineResponse200120 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200120() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// NewInlineResponse200120WithDefaults instantiates a new InlineResponse200120 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200120WithDefaults() *InlineResponse200120 {
	this := InlineResponse200120{}
	return &this
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *InlineResponse200120) GetFirstSeenAt() string {
	if o == nil || isNil(o.FirstSeenAt) {
		var ret string
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetFirstSeenAtOk() (*string, bool) {
	if o == nil || isNil(o.FirstSeenAt) {
    return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *InlineResponse200120) HasFirstSeenAt() bool {
	if o != nil && !isNil(o.FirstSeenAt) {
		return true
	}

	return false
}

// SetFirstSeenAt gets a reference to the given string and assigns it to the FirstSeenAt field.
func (o *InlineResponse200120) SetFirstSeenAt(v string) {
	o.FirstSeenAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *InlineResponse200120) GetLastSeenAt() string {
	if o == nil || isNil(o.LastSeenAt) {
		var ret string
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200120) GetLastSeenAtOk() (*string, bool) {
	if o == nil || isNil(o.LastSeenAt) {
    return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *InlineResponse200120) HasLastSeenAt() bool {
	if o != nil && !isNil(o.LastSeenAt) {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given string and assigns it to the LastSeenAt field.
func (o *InlineResponse200120) SetLastSeenAt(v string) {
	o.LastSeenAt = &v
}

func (o InlineResponse200120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FirstSeenAt) {
		toSerialize["firstSeenAt"] = o.FirstSeenAt
	}
	if !isNil(o.LastSeenAt) {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200120 struct {
	value *InlineResponse200120
	isSet bool
}

func (v NullableInlineResponse200120) Get() *InlineResponse200120 {
	return v.value
}

func (v *NullableInlineResponse200120) Set(val *InlineResponse200120) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200120) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200120) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200120(val *InlineResponse200120) *NullableInlineResponse200120 {
	return &NullableInlineResponse200120{value: val, isSet: true}
}

func (v NullableInlineResponse200120) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200120) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


