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

// InlineResponse20059 struct for InlineResponse20059
type InlineResponse20059 struct {
	// RF Profiles
	Assigned []InlineResponse20059Assigned `json:"assigned,omitempty"`
}

// NewInlineResponse20059 instantiates a new InlineResponse20059 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// NewInlineResponse20059WithDefaults instantiates a new InlineResponse20059 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059WithDefaults() *InlineResponse20059 {
	this := InlineResponse20059{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *InlineResponse20059) GetAssigned() []InlineResponse20059Assigned {
	if o == nil || isNil(o.Assigned) {
		var ret []InlineResponse20059Assigned
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059) GetAssignedOk() ([]InlineResponse20059Assigned, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *InlineResponse20059) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []InlineResponse20059Assigned and assigns it to the Assigned field.
func (o *InlineResponse20059) SetAssigned(v []InlineResponse20059Assigned) {
	o.Assigned = v
}

func (o InlineResponse20059) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059 struct {
	value *InlineResponse20059
	isSet bool
}

func (v NullableInlineResponse20059) Get() *InlineResponse20059 {
	return v.value
}

func (v *NullableInlineResponse20059) Set(val *InlineResponse20059) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059(val *InlineResponse20059) *NullableInlineResponse20059 {
	return &NullableInlineResponse20059{value: val, isSet: true}
}

func (v NullableInlineResponse20059) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


