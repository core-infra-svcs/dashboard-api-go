/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20022 struct for InlineResponse20022
type InlineResponse20022 struct {
	// RF Profiles
	Assigned []InlineResponse20022Assigned `json:"assigned,omitempty"`
}

// NewInlineResponse20022 instantiates a new InlineResponse20022 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022WithDefaults() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *InlineResponse20022) GetAssigned() []InlineResponse20022Assigned {
	if o == nil || isNil(o.Assigned) {
		var ret []InlineResponse20022Assigned
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetAssignedOk() ([]InlineResponse20022Assigned, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *InlineResponse20022) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []InlineResponse20022Assigned and assigns it to the Assigned field.
func (o *InlineResponse20022) SetAssigned(v []InlineResponse20022Assigned) {
	o.Assigned = v
}

func (o InlineResponse20022) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022 struct {
	value *InlineResponse20022
	isSet bool
}

func (v NullableInlineResponse20022) Get() *InlineResponse20022 {
	return v.value
}

func (v *NullableInlineResponse20022) Set(val *InlineResponse20022) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022(val *InlineResponse20022) *NullableInlineResponse20022 {
	return &NullableInlineResponse20022{value: val, isSet: true}
}

func (v NullableInlineResponse20022) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


