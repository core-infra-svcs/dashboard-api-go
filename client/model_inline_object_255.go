/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject255 struct for InlineObject255
type InlineObject255 struct {
	// The numbers of the orders that should be claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices that should be claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses that should be claimed
	Licenses []OrganizationsOrganizationIdInventoryClaimLicenses `json:"licenses,omitempty"`
}

// NewInlineObject255 instantiates a new InlineObject255 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject255() *InlineObject255 {
	this := InlineObject255{}
	return &this
}

// NewInlineObject255WithDefaults instantiates a new InlineObject255 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject255WithDefaults() *InlineObject255 {
	this := InlineObject255{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineObject255) GetOrders() []string {
	if o == nil || isNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject255) GetOrdersOk() ([]string, bool) {
	if o == nil || isNil(o.Orders) {
    return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineObject255) HasOrders() bool {
	if o != nil && !isNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *InlineObject255) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject255) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject255) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject255) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject255) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *InlineObject255) GetLicenses() []OrganizationsOrganizationIdInventoryClaimLicenses {
	if o == nil || isNil(o.Licenses) {
		var ret []OrganizationsOrganizationIdInventoryClaimLicenses
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject255) GetLicensesOk() ([]OrganizationsOrganizationIdInventoryClaimLicenses, bool) {
	if o == nil || isNil(o.Licenses) {
    return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *InlineObject255) HasLicenses() bool {
	if o != nil && !isNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []OrganizationsOrganizationIdInventoryClaimLicenses and assigns it to the Licenses field.
func (o *InlineObject255) SetLicenses(v []OrganizationsOrganizationIdInventoryClaimLicenses) {
	o.Licenses = v
}

func (o InlineObject255) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject255 struct {
	value *InlineObject255
	isSet bool
}

func (v NullableInlineObject255) Get() *InlineObject255 {
	return v.value
}

func (v *NullableInlineObject255) Set(val *InlineObject255) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject255) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject255) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject255(val *InlineObject255) *NullableInlineObject255 {
	return &NullableInlineObject255{value: val, isSet: true}
}

func (v NullableInlineObject255) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject255) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


