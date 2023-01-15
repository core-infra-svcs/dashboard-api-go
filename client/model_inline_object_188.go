/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject188 struct for InlineObject188
type InlineObject188 struct {
	// The numbers of the orders that should be claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices that should be claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses that should be claimed
	Licenses []OrganizationsOrganizationIdClaimLicenses `json:"licenses,omitempty"`
}

// NewInlineObject188 instantiates a new InlineObject188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject188() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// NewInlineObject188WithDefaults instantiates a new InlineObject188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject188WithDefaults() *InlineObject188 {
	this := InlineObject188{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineObject188) GetOrders() []string {
	if o == nil || isNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetOrdersOk() ([]string, bool) {
	if o == nil || isNil(o.Orders) {
    return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineObject188) HasOrders() bool {
	if o != nil && !isNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *InlineObject188) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject188) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject188) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject188) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *InlineObject188) GetLicenses() []OrganizationsOrganizationIdClaimLicenses {
	if o == nil || isNil(o.Licenses) {
		var ret []OrganizationsOrganizationIdClaimLicenses
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject188) GetLicensesOk() ([]OrganizationsOrganizationIdClaimLicenses, bool) {
	if o == nil || isNil(o.Licenses) {
    return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *InlineObject188) HasLicenses() bool {
	if o != nil && !isNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []OrganizationsOrganizationIdClaimLicenses and assigns it to the Licenses field.
func (o *InlineObject188) SetLicenses(v []OrganizationsOrganizationIdClaimLicenses) {
	o.Licenses = v
}

func (o InlineObject188) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject188 struct {
	value *InlineObject188
	isSet bool
}

func (v NullableInlineObject188) Get() *InlineObject188 {
	return v.value
}

func (v *NullableInlineObject188) Set(val *InlineObject188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject188(val *InlineObject188) *NullableInlineObject188 {
	return &NullableInlineObject188{value: val, isSet: true}
}

func (v NullableInlineObject188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


