/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200116 struct for InlineResponse200116
type InlineResponse200116 struct {
	// The numbers of the orders claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses claimed
	Licenses []InlineResponse200116Licenses `json:"licenses,omitempty"`
}

// NewInlineResponse200116 instantiates a new InlineResponse200116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200116() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// NewInlineResponse200116WithDefaults instantiates a new InlineResponse200116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200116WithDefaults() *InlineResponse200116 {
	this := InlineResponse200116{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineResponse200116) GetOrders() []string {
	if o == nil || isNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetOrdersOk() ([]string, bool) {
	if o == nil || isNil(o.Orders) {
    return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineResponse200116) HasOrders() bool {
	if o != nil && !isNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *InlineResponse200116) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse200116) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse200116) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse200116) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *InlineResponse200116) GetLicenses() []InlineResponse200116Licenses {
	if o == nil || isNil(o.Licenses) {
		var ret []InlineResponse200116Licenses
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200116) GetLicensesOk() ([]InlineResponse200116Licenses, bool) {
	if o == nil || isNil(o.Licenses) {
    return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *InlineResponse200116) HasLicenses() bool {
	if o != nil && !isNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []InlineResponse200116Licenses and assigns it to the Licenses field.
func (o *InlineResponse200116) SetLicenses(v []InlineResponse200116Licenses) {
	o.Licenses = v
}

func (o InlineResponse200116) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200116 struct {
	value *InlineResponse200116
	isSet bool
}

func (v NullableInlineResponse200116) Get() *InlineResponse200116 {
	return v.value
}

func (v *NullableInlineResponse200116) Set(val *InlineResponse200116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200116(val *InlineResponse200116) *NullableInlineResponse200116 {
	return &NullableInlineResponse200116{value: val, isSet: true}
}

func (v NullableInlineResponse200116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


