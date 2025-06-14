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

// InlineObject183 struct for InlineObject183
type InlineObject183 struct {
	// The currency code of this node group's billing plans
	Currency *string `json:"currency,omitempty"`
	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []NetworksNetworkIdWirelessBillingPlans `json:"plans,omitempty"`
}

// NewInlineObject183 instantiates a new InlineObject183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject183() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// NewInlineObject183WithDefaults instantiates a new InlineObject183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject183WithDefaults() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InlineObject183) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InlineObject183) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InlineObject183) SetCurrency(v string) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *InlineObject183) GetPlans() []NetworksNetworkIdWirelessBillingPlans {
	if o == nil || isNil(o.Plans) {
		var ret []NetworksNetworkIdWirelessBillingPlans
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetPlansOk() ([]NetworksNetworkIdWirelessBillingPlans, bool) {
	if o == nil || isNil(o.Plans) {
    return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *InlineObject183) HasPlans() bool {
	if o != nil && !isNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []NetworksNetworkIdWirelessBillingPlans and assigns it to the Plans field.
func (o *InlineObject183) SetPlans(v []NetworksNetworkIdWirelessBillingPlans) {
	o.Plans = v
}

func (o InlineObject183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !isNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject183 struct {
	value *InlineObject183
	isSet bool
}

func (v NullableInlineObject183) Get() *InlineObject183 {
	return v.value
}

func (v *NullableInlineObject183) Set(val *InlineObject183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject183(val *InlineObject183) *NullableInlineObject183 {
	return &NullableInlineObject183{value: val, isSet: true}
}

func (v NullableInlineObject183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


