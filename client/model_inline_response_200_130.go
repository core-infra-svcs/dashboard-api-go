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

// InlineResponse200130 struct for InlineResponse200130
type InlineResponse200130 struct {
	// When the device was first seen as connected to the internet in each connection.
	FirstSeenAt *string `json:"firstSeenAt,omitempty"`
	// When the device was last seen as connected to the internet in each connection.
	LastSeenAt *string `json:"lastSeenAt,omitempty"`
}

// NewInlineResponse200130 instantiates a new InlineResponse200130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200130() *InlineResponse200130 {
	this := InlineResponse200130{}
	return &this
}

// NewInlineResponse200130WithDefaults instantiates a new InlineResponse200130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200130WithDefaults() *InlineResponse200130 {
	this := InlineResponse200130{}
	return &this
}

// GetFirstSeenAt returns the FirstSeenAt field value if set, zero value otherwise.
func (o *InlineResponse200130) GetFirstSeenAt() string {
	if o == nil || isNil(o.FirstSeenAt) {
		var ret string
		return ret
	}
	return *o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200130) GetFirstSeenAtOk() (*string, bool) {
	if o == nil || isNil(o.FirstSeenAt) {
    return nil, false
	}
	return o.FirstSeenAt, true
}

// HasFirstSeenAt returns a boolean if a field has been set.
func (o *InlineResponse200130) HasFirstSeenAt() bool {
	if o != nil && !isNil(o.FirstSeenAt) {
		return true
	}

	return false
}

// SetFirstSeenAt gets a reference to the given string and assigns it to the FirstSeenAt field.
func (o *InlineResponse200130) SetFirstSeenAt(v string) {
	o.FirstSeenAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *InlineResponse200130) GetLastSeenAt() string {
	if o == nil || isNil(o.LastSeenAt) {
		var ret string
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200130) GetLastSeenAtOk() (*string, bool) {
	if o == nil || isNil(o.LastSeenAt) {
    return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *InlineResponse200130) HasLastSeenAt() bool {
	if o != nil && !isNil(o.LastSeenAt) {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given string and assigns it to the LastSeenAt field.
func (o *InlineResponse200130) SetLastSeenAt(v string) {
	o.LastSeenAt = &v
}

func (o InlineResponse200130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FirstSeenAt) {
		toSerialize["firstSeenAt"] = o.FirstSeenAt
	}
	if !isNil(o.LastSeenAt) {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200130 struct {
	value *InlineResponse200130
	isSet bool
}

func (v NullableInlineResponse200130) Get() *InlineResponse200130 {
	return v.value
}

func (v *NullableInlineResponse200130) Set(val *InlineResponse200130) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200130) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200130(val *InlineResponse200130) *NullableInlineResponse200130 {
	return &NullableInlineResponse200130{value: val, isSet: true}
}

func (v NullableInlineResponse200130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


