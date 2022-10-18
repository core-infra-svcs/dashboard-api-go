/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject198 struct for InlineObject198
type InlineObject198 struct {
	// The numbers of the orders that should be claimed
	Orders *[]string `json:"orders,omitempty"`
	// The serials of the devices that should be claimed
	Serials *[]string `json:"serials,omitempty"`
	// The licenses that should be claimed
	Licenses *[]OrganizationsOrganizationIdInventoryClaimLicenses `json:"licenses,omitempty"`
}

// NewInlineObject198 instantiates a new InlineObject198 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject198() *InlineObject198 {
	this := InlineObject198{}
	return &this
}

// NewInlineObject198WithDefaults instantiates a new InlineObject198 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject198WithDefaults() *InlineObject198 {
	this := InlineObject198{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineObject198) GetOrders() []string {
	if o == nil || o.Orders == nil {
		var ret []string
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject198) GetOrdersOk() (*[]string, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineObject198) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *InlineObject198) SetOrders(v []string) {
	o.Orders = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject198) GetSerials() []string {
	if o == nil || o.Serials == nil {
		var ret []string
		return ret
	}
	return *o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject198) GetSerialsOk() (*[]string, bool) {
	if o == nil || o.Serials == nil {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject198) HasSerials() bool {
	if o != nil && o.Serials != nil {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject198) SetSerials(v []string) {
	o.Serials = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *InlineObject198) GetLicenses() []OrganizationsOrganizationIdInventoryClaimLicenses {
	if o == nil || o.Licenses == nil {
		var ret []OrganizationsOrganizationIdInventoryClaimLicenses
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject198) GetLicensesOk() (*[]OrganizationsOrganizationIdInventoryClaimLicenses, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *InlineObject198) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []OrganizationsOrganizationIdInventoryClaimLicenses and assigns it to the Licenses field.
func (o *InlineObject198) SetLicenses(v []OrganizationsOrganizationIdInventoryClaimLicenses) {
	o.Licenses = &v
}

func (o InlineObject198) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.Serials != nil {
		toSerialize["serials"] = o.Serials
	}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject198 struct {
	value *InlineObject198
	isSet bool
}

func (v NullableInlineObject198) Get() *InlineObject198 {
	return v.value
}

func (v *NullableInlineObject198) Set(val *InlineObject198) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject198) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject198) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject198(val *InlineObject198) *NullableInlineObject198 {
	return &NullableInlineObject198{value: val, isSet: true}
}

func (v NullableInlineObject198) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject198) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


