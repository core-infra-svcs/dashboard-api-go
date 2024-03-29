/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	UpgradeWindow *InlineResponse20051UpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *InlineResponse20051Products `json:"products,omitempty"`
}

// NewInlineResponse20051 instantiates a new InlineResponse20051 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051WithDefaults() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *InlineResponse20051) GetUpgradeWindow() InlineResponse20051UpgradeWindow {
	if o == nil || isNil(o.UpgradeWindow) {
		var ret InlineResponse20051UpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetUpgradeWindowOk() (*InlineResponse20051UpgradeWindow, bool) {
	if o == nil || isNil(o.UpgradeWindow) {
    return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *InlineResponse20051) HasUpgradeWindow() bool {
	if o != nil && !isNil(o.UpgradeWindow) {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given InlineResponse20051UpgradeWindow and assigns it to the UpgradeWindow field.
func (o *InlineResponse20051) SetUpgradeWindow(v InlineResponse20051UpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineResponse20051) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineResponse20051) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineResponse20051) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineResponse20051) GetProducts() InlineResponse20051Products {
	if o == nil || isNil(o.Products) {
		var ret InlineResponse20051Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetProductsOk() (*InlineResponse20051Products, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineResponse20051) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given InlineResponse20051Products and assigns it to the Products field.
func (o *InlineResponse20051) SetProducts(v InlineResponse20051Products) {
	o.Products = &v
}

func (o InlineResponse20051) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpgradeWindow) {
		toSerialize["upgradeWindow"] = o.UpgradeWindow
	}
	if !isNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051 struct {
	value *InlineResponse20051
	isSet bool
}

func (v NullableInlineResponse20051) Get() *InlineResponse20051 {
	return v.value
}

func (v *NullableInlineResponse20051) Set(val *InlineResponse20051) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051(val *InlineResponse20051) *NullableInlineResponse20051 {
	return &NullableInlineResponse20051{value: val, isSet: true}
}

func (v NullableInlineResponse20051) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


