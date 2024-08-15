/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2009ZonesZoneId The zone state, dynamic
type InlineResponse2009ZonesZoneId struct {
	// The count per type, dynamic
	Person *int32 `json:"person,omitempty"`
}

// NewInlineResponse2009ZonesZoneId instantiates a new InlineResponse2009ZonesZoneId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009ZonesZoneId() *InlineResponse2009ZonesZoneId {
	this := InlineResponse2009ZonesZoneId{}
	return &this
}

// NewInlineResponse2009ZonesZoneIdWithDefaults instantiates a new InlineResponse2009ZonesZoneId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009ZonesZoneIdWithDefaults() *InlineResponse2009ZonesZoneId {
	this := InlineResponse2009ZonesZoneId{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *InlineResponse2009ZonesZoneId) GetPerson() int32 {
	if o == nil || isNil(o.Person) {
		var ret int32
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009ZonesZoneId) GetPersonOk() (*int32, bool) {
	if o == nil || isNil(o.Person) {
    return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *InlineResponse2009ZonesZoneId) HasPerson() bool {
	if o != nil && !isNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given int32 and assigns it to the Person field.
func (o *InlineResponse2009ZonesZoneId) SetPerson(v int32) {
	o.Person = &v
}

func (o InlineResponse2009ZonesZoneId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009ZonesZoneId struct {
	value *InlineResponse2009ZonesZoneId
	isSet bool
}

func (v NullableInlineResponse2009ZonesZoneId) Get() *InlineResponse2009ZonesZoneId {
	return v.value
}

func (v *NullableInlineResponse2009ZonesZoneId) Set(val *InlineResponse2009ZonesZoneId) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009ZonesZoneId) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009ZonesZoneId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009ZonesZoneId(val *InlineResponse2009ZonesZoneId) *NullableInlineResponse2009ZonesZoneId {
	return &NullableInlineResponse2009ZonesZoneId{value: val, isSet: true}
}

func (v NullableInlineResponse2009ZonesZoneId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009ZonesZoneId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


