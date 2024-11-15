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

// InlineResponse20043 struct for InlineResponse20043
type InlineResponse20043 struct {
	// SSID status list
	BasicServiceSets []InlineResponse20043BasicServiceSets `json:"basicServiceSets,omitempty"`
}

// NewInlineResponse20043 instantiates a new InlineResponse20043 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20043() *InlineResponse20043 {
	this := InlineResponse20043{}
	return &this
}

// NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20043WithDefaults() *InlineResponse20043 {
	this := InlineResponse20043{}
	return &this
}

// GetBasicServiceSets returns the BasicServiceSets field value if set, zero value otherwise.
func (o *InlineResponse20043) GetBasicServiceSets() []InlineResponse20043BasicServiceSets {
	if o == nil || isNil(o.BasicServiceSets) {
		var ret []InlineResponse20043BasicServiceSets
		return ret
	}
	return o.BasicServiceSets
}

// GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetBasicServiceSetsOk() ([]InlineResponse20043BasicServiceSets, bool) {
	if o == nil || isNil(o.BasicServiceSets) {
    return nil, false
	}
	return o.BasicServiceSets, true
}

// HasBasicServiceSets returns a boolean if a field has been set.
func (o *InlineResponse20043) HasBasicServiceSets() bool {
	if o != nil && !isNil(o.BasicServiceSets) {
		return true
	}

	return false
}

// SetBasicServiceSets gets a reference to the given []InlineResponse20043BasicServiceSets and assigns it to the BasicServiceSets field.
func (o *InlineResponse20043) SetBasicServiceSets(v []InlineResponse20043BasicServiceSets) {
	o.BasicServiceSets = v
}

func (o InlineResponse20043) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BasicServiceSets) {
		toSerialize["basicServiceSets"] = o.BasicServiceSets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20043 struct {
	value *InlineResponse20043
	isSet bool
}

func (v NullableInlineResponse20043) Get() *InlineResponse20043 {
	return v.value
}

func (v *NullableInlineResponse20043) Set(val *InlineResponse20043) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20043) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20043) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20043(val *InlineResponse20043) *NullableInlineResponse20043 {
	return &NullableInlineResponse20043{value: val, isSet: true}
}

func (v NullableInlineResponse20043) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20043) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


