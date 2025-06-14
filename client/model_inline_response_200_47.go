/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20047 struct for InlineResponse20047
type InlineResponse20047 struct {
	// SSID status list
	BasicServiceSets []InlineResponse20047BasicServiceSets `json:"basicServiceSets,omitempty"`
}

// NewInlineResponse20047 instantiates a new InlineResponse20047 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047() *InlineResponse20047 {
	this := InlineResponse20047{}
	return &this
}

// NewInlineResponse20047WithDefaults instantiates a new InlineResponse20047 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047WithDefaults() *InlineResponse20047 {
	this := InlineResponse20047{}
	return &this
}

// GetBasicServiceSets returns the BasicServiceSets field value if set, zero value otherwise.
func (o *InlineResponse20047) GetBasicServiceSets() []InlineResponse20047BasicServiceSets {
	if o == nil || isNil(o.BasicServiceSets) {
		var ret []InlineResponse20047BasicServiceSets
		return ret
	}
	return o.BasicServiceSets
}

// GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047) GetBasicServiceSetsOk() ([]InlineResponse20047BasicServiceSets, bool) {
	if o == nil || isNil(o.BasicServiceSets) {
    return nil, false
	}
	return o.BasicServiceSets, true
}

// HasBasicServiceSets returns a boolean if a field has been set.
func (o *InlineResponse20047) HasBasicServiceSets() bool {
	if o != nil && !isNil(o.BasicServiceSets) {
		return true
	}

	return false
}

// SetBasicServiceSets gets a reference to the given []InlineResponse20047BasicServiceSets and assigns it to the BasicServiceSets field.
func (o *InlineResponse20047) SetBasicServiceSets(v []InlineResponse20047BasicServiceSets) {
	o.BasicServiceSets = v
}

func (o InlineResponse20047) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BasicServiceSets) {
		toSerialize["basicServiceSets"] = o.BasicServiceSets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047 struct {
	value *InlineResponse20047
	isSet bool
}

func (v NullableInlineResponse20047) Get() *InlineResponse20047 {
	return v.value
}

func (v *NullableInlineResponse20047) Set(val *InlineResponse20047) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047(val *InlineResponse20047) *NullableInlineResponse20047 {
	return &NullableInlineResponse20047{value: val, isSet: true}
}

func (v NullableInlineResponse20047) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


