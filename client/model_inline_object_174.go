/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject174 struct for InlineObject174
type InlineObject174 struct {
	// The currency code of this node group's billing plans
	Currency *string `json:"currency,omitempty"`
	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []NetworksNetworkIdWirelessBillingPlans `json:"plans,omitempty"`
}

// NewInlineObject174 instantiates a new InlineObject174 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject174() *InlineObject174 {
	this := InlineObject174{}
	return &this
}

// NewInlineObject174WithDefaults instantiates a new InlineObject174 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject174WithDefaults() *InlineObject174 {
	this := InlineObject174{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InlineObject174) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject174) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InlineObject174) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InlineObject174) SetCurrency(v string) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *InlineObject174) GetPlans() []NetworksNetworkIdWirelessBillingPlans {
	if o == nil || isNil(o.Plans) {
		var ret []NetworksNetworkIdWirelessBillingPlans
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject174) GetPlansOk() ([]NetworksNetworkIdWirelessBillingPlans, bool) {
	if o == nil || isNil(o.Plans) {
    return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *InlineObject174) HasPlans() bool {
	if o != nil && !isNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []NetworksNetworkIdWirelessBillingPlans and assigns it to the Plans field.
func (o *InlineObject174) SetPlans(v []NetworksNetworkIdWirelessBillingPlans) {
	o.Plans = v
}

func (o InlineObject174) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !isNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject174 struct {
	value *InlineObject174
	isSet bool
}

func (v NullableInlineObject174) Get() *InlineObject174 {
	return v.value
}

func (v *NullableInlineObject174) Set(val *InlineObject174) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject174) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject174) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject174(val *InlineObject174) *NullableInlineObject174 {
	return &NullableInlineObject174{value: val, isSet: true}
}

func (v NullableInlineObject174) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject174) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


