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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// Last 4 characters of the API key
	Suffix *string `json:"suffix,omitempty"`
	// Date that the API key was created
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSuffix() string {
	if o == nil || isNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSuffixOk() (*string, bool) {
	if o == nil || isNil(o.Suffix) {
    return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSuffix() bool {
	if o != nil && !isNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *InlineResponse2001) SetSuffix(v string) {
	o.Suffix = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse2001) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse2001) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse2001) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


