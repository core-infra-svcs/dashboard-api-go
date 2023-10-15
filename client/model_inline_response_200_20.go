/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20020 struct for InlineResponse20020
type InlineResponse20020 struct {
	// RF Profiles
	Assigned []InlineResponse20020Assigned `json:"assigned,omitempty"`
}

// NewInlineResponse20020 instantiates a new InlineResponse20020 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// NewInlineResponse20020WithDefaults instantiates a new InlineResponse20020 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020WithDefaults() *InlineResponse20020 {
	this := InlineResponse20020{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *InlineResponse20020) GetAssigned() []InlineResponse20020Assigned {
	if o == nil || isNil(o.Assigned) {
		var ret []InlineResponse20020Assigned
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020) GetAssignedOk() ([]InlineResponse20020Assigned, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *InlineResponse20020) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []InlineResponse20020Assigned and assigns it to the Assigned field.
func (o *InlineResponse20020) SetAssigned(v []InlineResponse20020Assigned) {
	o.Assigned = v
}

func (o InlineResponse20020) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020 struct {
	value *InlineResponse20020
	isSet bool
}

func (v NullableInlineResponse20020) Get() *InlineResponse20020 {
	return v.value
}

func (v *NullableInlineResponse20020) Set(val *InlineResponse20020) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020(val *InlineResponse20020) *NullableInlineResponse20020 {
	return &NullableInlineResponse20020{value: val, isSet: true}
}

func (v NullableInlineResponse20020) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


