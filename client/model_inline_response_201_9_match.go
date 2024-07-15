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

// InlineResponse2019Match Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)
type InlineResponse2019Match struct {
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)
	String *string `json:"string,omitempty"`
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse2019Match instantiates a new InlineResponse2019Match object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2019Match() *InlineResponse2019Match {
	this := InlineResponse2019Match{}
	return &this
}

// NewInlineResponse2019MatchWithDefaults instantiates a new InlineResponse2019Match object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2019MatchWithDefaults() *InlineResponse2019Match {
	this := InlineResponse2019Match{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *InlineResponse2019Match) GetString() string {
	if o == nil || isNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2019Match) GetStringOk() (*string, bool) {
	if o == nil || isNil(o.String) {
    return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *InlineResponse2019Match) HasString() bool {
	if o != nil && !isNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *InlineResponse2019Match) SetString(v string) {
	o.String = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse2019Match) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2019Match) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse2019Match) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse2019Match) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse2019Match) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2019Match struct {
	value *InlineResponse2019Match
	isSet bool
}

func (v NullableInlineResponse2019Match) Get() *InlineResponse2019Match {
	return v.value
}

func (v *NullableInlineResponse2019Match) Set(val *InlineResponse2019Match) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2019Match) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2019Match) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2019Match(val *InlineResponse2019Match) *NullableInlineResponse2019Match {
	return &NullableInlineResponse2019Match{value: val, isSet: true}
}

func (v NullableInlineResponse2019Match) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2019Match) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


