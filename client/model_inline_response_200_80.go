/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20080 struct for InlineResponse20080
type InlineResponse20080 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []InlineResponse20080Mappings `json:"mappings,omitempty"`
}

// NewInlineResponse20080 instantiates a new InlineResponse20080 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20080() *InlineResponse20080 {
	this := InlineResponse20080{}
	return &this
}

// NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20080WithDefaults() *InlineResponse20080 {
	this := InlineResponse20080{}
	return &this
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *InlineResponse20080) GetMappings() []InlineResponse20080Mappings {
	if o == nil || isNil(o.Mappings) {
		var ret []InlineResponse20080Mappings
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20080) GetMappingsOk() ([]InlineResponse20080Mappings, bool) {
	if o == nil || isNil(o.Mappings) {
    return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *InlineResponse20080) HasMappings() bool {
	if o != nil && !isNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []InlineResponse20080Mappings and assigns it to the Mappings field.
func (o *InlineResponse20080) SetMappings(v []InlineResponse20080Mappings) {
	o.Mappings = v
}

func (o InlineResponse20080) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20080 struct {
	value *InlineResponse20080
	isSet bool
}

func (v NullableInlineResponse20080) Get() *InlineResponse20080 {
	return v.value
}

func (v *NullableInlineResponse20080) Set(val *InlineResponse20080) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20080) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20080) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20080(val *InlineResponse20080) *NullableInlineResponse20080 {
	return &NullableInlineResponse20080{value: val, isSet: true}
}

func (v NullableInlineResponse20080) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20080) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


