/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20021 struct for InlineResponse20021
type InlineResponse20021 struct {
	//  The L7 firewall application categories and their associated applications for an MX network
	ApplicationCategories []InlineResponse20021ApplicationCategories `json:"applicationCategories,omitempty"`
}

// NewInlineResponse20021 instantiates a new InlineResponse20021 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20021() *InlineResponse20021 {
	this := InlineResponse20021{}
	return &this
}

// NewInlineResponse20021WithDefaults instantiates a new InlineResponse20021 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20021WithDefaults() *InlineResponse20021 {
	this := InlineResponse20021{}
	return &this
}

// GetApplicationCategories returns the ApplicationCategories field value if set, zero value otherwise.
func (o *InlineResponse20021) GetApplicationCategories() []InlineResponse20021ApplicationCategories {
	if o == nil || isNil(o.ApplicationCategories) {
		var ret []InlineResponse20021ApplicationCategories
		return ret
	}
	return o.ApplicationCategories
}

// GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetApplicationCategoriesOk() ([]InlineResponse20021ApplicationCategories, bool) {
	if o == nil || isNil(o.ApplicationCategories) {
    return nil, false
	}
	return o.ApplicationCategories, true
}

// HasApplicationCategories returns a boolean if a field has been set.
func (o *InlineResponse20021) HasApplicationCategories() bool {
	if o != nil && !isNil(o.ApplicationCategories) {
		return true
	}

	return false
}

// SetApplicationCategories gets a reference to the given []InlineResponse20021ApplicationCategories and assigns it to the ApplicationCategories field.
func (o *InlineResponse20021) SetApplicationCategories(v []InlineResponse20021ApplicationCategories) {
	o.ApplicationCategories = v
}

func (o InlineResponse20021) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApplicationCategories) {
		toSerialize["applicationCategories"] = o.ApplicationCategories
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20021 struct {
	value *InlineResponse20021
	isSet bool
}

func (v NullableInlineResponse20021) Get() *InlineResponse20021 {
	return v.value
}

func (v *NullableInlineResponse20021) Set(val *InlineResponse20021) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20021) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20021) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20021(val *InlineResponse20021) *NullableInlineResponse20021 {
	return &NullableInlineResponse20021{value: val, isSet: true}
}

func (v NullableInlineResponse20021) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20021) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


