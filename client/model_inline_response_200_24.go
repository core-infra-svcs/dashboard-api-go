/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20024 struct for InlineResponse20024
type InlineResponse20024 struct {
	// SSID status list
	BasicServiceSets []InlineResponse20024BasicServiceSets `json:"basicServiceSets,omitempty"`
}

// NewInlineResponse20024 instantiates a new InlineResponse20024 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20024() *InlineResponse20024 {
	this := InlineResponse20024{}
	return &this
}

// NewInlineResponse20024WithDefaults instantiates a new InlineResponse20024 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20024WithDefaults() *InlineResponse20024 {
	this := InlineResponse20024{}
	return &this
}

// GetBasicServiceSets returns the BasicServiceSets field value if set, zero value otherwise.
func (o *InlineResponse20024) GetBasicServiceSets() []InlineResponse20024BasicServiceSets {
	if o == nil || isNil(o.BasicServiceSets) {
		var ret []InlineResponse20024BasicServiceSets
		return ret
	}
	return o.BasicServiceSets
}

// GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024) GetBasicServiceSetsOk() ([]InlineResponse20024BasicServiceSets, bool) {
	if o == nil || isNil(o.BasicServiceSets) {
    return nil, false
	}
	return o.BasicServiceSets, true
}

// HasBasicServiceSets returns a boolean if a field has been set.
func (o *InlineResponse20024) HasBasicServiceSets() bool {
	if o != nil && !isNil(o.BasicServiceSets) {
		return true
	}

	return false
}

// SetBasicServiceSets gets a reference to the given []InlineResponse20024BasicServiceSets and assigns it to the BasicServiceSets field.
func (o *InlineResponse20024) SetBasicServiceSets(v []InlineResponse20024BasicServiceSets) {
	o.BasicServiceSets = v
}

func (o InlineResponse20024) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BasicServiceSets) {
		toSerialize["basicServiceSets"] = o.BasicServiceSets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20024 struct {
	value *InlineResponse20024
	isSet bool
}

func (v NullableInlineResponse20024) Get() *InlineResponse20024 {
	return v.value
}

func (v *NullableInlineResponse20024) Set(val *InlineResponse20024) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20024) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20024) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20024(val *InlineResponse20024) *NullableInlineResponse20024 {
	return &NullableInlineResponse20024{value: val, isSet: true}
}

func (v NullableInlineResponse20024) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20024) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


