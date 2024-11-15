/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200151 struct for InlineResponse200151
type InlineResponse200151 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []InlineResponse200151Mappings `json:"mappings,omitempty"`
}

// NewInlineResponse200151 instantiates a new InlineResponse200151 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200151() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// NewInlineResponse200151WithDefaults instantiates a new InlineResponse200151 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200151WithDefaults() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *InlineResponse200151) GetMappings() []InlineResponse200151Mappings {
	if o == nil || isNil(o.Mappings) {
		var ret []InlineResponse200151Mappings
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetMappingsOk() ([]InlineResponse200151Mappings, bool) {
	if o == nil || isNil(o.Mappings) {
    return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *InlineResponse200151) HasMappings() bool {
	if o != nil && !isNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []InlineResponse200151Mappings and assigns it to the Mappings field.
func (o *InlineResponse200151) SetMappings(v []InlineResponse200151Mappings) {
	o.Mappings = v
}

func (o InlineResponse200151) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200151 struct {
	value *InlineResponse200151
	isSet bool
}

func (v NullableInlineResponse200151) Get() *InlineResponse200151 {
	return v.value
}

func (v *NullableInlineResponse200151) Set(val *InlineResponse200151) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200151) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200151) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200151(val *InlineResponse200151) *NullableInlineResponse200151 {
	return &NullableInlineResponse200151{value: val, isSet: true}
}

func (v NullableInlineResponse200151) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200151) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


