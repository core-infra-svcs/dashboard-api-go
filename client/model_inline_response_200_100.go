/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200100 struct for InlineResponse200100
type InlineResponse200100 struct {
	// The currency code of this node group's billing plans
	Currency *string `json:"currency,omitempty"`
	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []InlineResponse200100Plans `json:"plans,omitempty"`
}

// NewInlineResponse200100 instantiates a new InlineResponse200100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200100() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// NewInlineResponse200100WithDefaults instantiates a new InlineResponse200100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200100WithDefaults() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InlineResponse200100) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InlineResponse200100) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InlineResponse200100) SetCurrency(v string) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *InlineResponse200100) GetPlans() []InlineResponse200100Plans {
	if o == nil || isNil(o.Plans) {
		var ret []InlineResponse200100Plans
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetPlansOk() ([]InlineResponse200100Plans, bool) {
	if o == nil || isNil(o.Plans) {
    return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *InlineResponse200100) HasPlans() bool {
	if o != nil && !isNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []InlineResponse200100Plans and assigns it to the Plans field.
func (o *InlineResponse200100) SetPlans(v []InlineResponse200100Plans) {
	o.Plans = v
}

func (o InlineResponse200100) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !isNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200100 struct {
	value *InlineResponse200100
	isSet bool
}

func (v NullableInlineResponse200100) Get() *InlineResponse200100 {
	return v.value
}

func (v *NullableInlineResponse200100) Set(val *InlineResponse200100) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200100) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200100(val *InlineResponse200100) *NullableInlineResponse200100 {
	return &NullableInlineResponse200100{value: val, isSet: true}
}

func (v NullableInlineResponse200100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


