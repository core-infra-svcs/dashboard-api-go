/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20018 struct for InlineResponse20018
type InlineResponse20018 struct {
	// RF Profiles
	Assigned []InlineResponse20018Assigned `json:"assigned,omitempty"`
}

// NewInlineResponse20018 instantiates a new InlineResponse20018 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20018() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// NewInlineResponse20018WithDefaults instantiates a new InlineResponse20018 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20018WithDefaults() *InlineResponse20018 {
	this := InlineResponse20018{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *InlineResponse20018) GetAssigned() []InlineResponse20018Assigned {
	if o == nil || isNil(o.Assigned) {
		var ret []InlineResponse20018Assigned
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20018) GetAssignedOk() ([]InlineResponse20018Assigned, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *InlineResponse20018) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []InlineResponse20018Assigned and assigns it to the Assigned field.
func (o *InlineResponse20018) SetAssigned(v []InlineResponse20018Assigned) {
	o.Assigned = v
}

func (o InlineResponse20018) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20018 struct {
	value *InlineResponse20018
	isSet bool
}

func (v NullableInlineResponse20018) Get() *InlineResponse20018 {
	return v.value
}

func (v *NullableInlineResponse20018) Set(val *InlineResponse20018) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20018) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20018) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20018(val *InlineResponse20018) *NullableInlineResponse20018 {
	return &NullableInlineResponse20018{value: val, isSet: true}
}

func (v NullableInlineResponse20018) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20018) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


