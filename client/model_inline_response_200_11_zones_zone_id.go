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

// InlineResponse20011ZonesZoneId The zone state, dynamic
type InlineResponse20011ZonesZoneId struct {
	// The count per type, dynamic
	Person *int32 `json:"person,omitempty"`
}

// NewInlineResponse20011ZonesZoneId instantiates a new InlineResponse20011ZonesZoneId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011ZonesZoneId() *InlineResponse20011ZonesZoneId {
	this := InlineResponse20011ZonesZoneId{}
	return &this
}

// NewInlineResponse20011ZonesZoneIdWithDefaults instantiates a new InlineResponse20011ZonesZoneId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011ZonesZoneIdWithDefaults() *InlineResponse20011ZonesZoneId {
	this := InlineResponse20011ZonesZoneId{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *InlineResponse20011ZonesZoneId) GetPerson() int32 {
	if o == nil || isNil(o.Person) {
		var ret int32
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011ZonesZoneId) GetPersonOk() (*int32, bool) {
	if o == nil || isNil(o.Person) {
    return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *InlineResponse20011ZonesZoneId) HasPerson() bool {
	if o != nil && !isNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given int32 and assigns it to the Person field.
func (o *InlineResponse20011ZonesZoneId) SetPerson(v int32) {
	o.Person = &v
}

func (o InlineResponse20011ZonesZoneId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011ZonesZoneId struct {
	value *InlineResponse20011ZonesZoneId
	isSet bool
}

func (v NullableInlineResponse20011ZonesZoneId) Get() *InlineResponse20011ZonesZoneId {
	return v.value
}

func (v *NullableInlineResponse20011ZonesZoneId) Set(val *InlineResponse20011ZonesZoneId) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011ZonesZoneId) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011ZonesZoneId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011ZonesZoneId(val *InlineResponse20011ZonesZoneId) *NullableInlineResponse20011ZonesZoneId {
	return &NullableInlineResponse20011ZonesZoneId{value: val, isSet: true}
}

func (v NullableInlineResponse20011ZonesZoneId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011ZonesZoneId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


