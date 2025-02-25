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

// InlineResponse200155 struct for InlineResponse200155
type InlineResponse200155 struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []InlineResponse200155Mappings `json:"mappings,omitempty"`
}

// NewInlineResponse200155 instantiates a new InlineResponse200155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200155() *InlineResponse200155 {
	this := InlineResponse200155{}
	return &this
}

// NewInlineResponse200155WithDefaults instantiates a new InlineResponse200155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200155WithDefaults() *InlineResponse200155 {
	this := InlineResponse200155{}
	return &this
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *InlineResponse200155) GetMappings() []InlineResponse200155Mappings {
	if o == nil || isNil(o.Mappings) {
		var ret []InlineResponse200155Mappings
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200155) GetMappingsOk() ([]InlineResponse200155Mappings, bool) {
	if o == nil || isNil(o.Mappings) {
    return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *InlineResponse200155) HasMappings() bool {
	if o != nil && !isNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []InlineResponse200155Mappings and assigns it to the Mappings field.
func (o *InlineResponse200155) SetMappings(v []InlineResponse200155Mappings) {
	o.Mappings = v
}

func (o InlineResponse200155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200155 struct {
	value *InlineResponse200155
	isSet bool
}

func (v NullableInlineResponse200155) Get() *InlineResponse200155 {
	return v.value
}

func (v *NullableInlineResponse200155) Set(val *InlineResponse200155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200155(val *InlineResponse200155) *NullableInlineResponse200155 {
	return &NullableInlineResponse200155{value: val, isSet: true}
}

func (v NullableInlineResponse200155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


