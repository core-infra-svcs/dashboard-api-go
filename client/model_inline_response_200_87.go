/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20087 struct for InlineResponse20087
type InlineResponse20087 struct {
	UpgradeWindow *InlineResponse20087UpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *InlineResponse20087Products `json:"products,omitempty"`
}

// NewInlineResponse20087 instantiates a new InlineResponse20087 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087WithDefaults() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *InlineResponse20087) GetUpgradeWindow() InlineResponse20087UpgradeWindow {
	if o == nil || isNil(o.UpgradeWindow) {
		var ret InlineResponse20087UpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetUpgradeWindowOk() (*InlineResponse20087UpgradeWindow, bool) {
	if o == nil || isNil(o.UpgradeWindow) {
    return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *InlineResponse20087) HasUpgradeWindow() bool {
	if o != nil && !isNil(o.UpgradeWindow) {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given InlineResponse20087UpgradeWindow and assigns it to the UpgradeWindow field.
func (o *InlineResponse20087) SetUpgradeWindow(v InlineResponse20087UpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineResponse20087) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineResponse20087) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineResponse20087) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineResponse20087) GetProducts() InlineResponse20087Products {
	if o == nil || isNil(o.Products) {
		var ret InlineResponse20087Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetProductsOk() (*InlineResponse20087Products, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineResponse20087) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given InlineResponse20087Products and assigns it to the Products field.
func (o *InlineResponse20087) SetProducts(v InlineResponse20087Products) {
	o.Products = &v
}

func (o InlineResponse20087) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20087 struct {
	value *InlineResponse20087
	isSet bool
}

func (v NullableInlineResponse20087) Get() *InlineResponse20087 {
	return v.value
}

func (v *NullableInlineResponse20087) Set(val *InlineResponse20087) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087(val *InlineResponse20087) *NullableInlineResponse20087 {
	return &NullableInlineResponse20087{value: val, isSet: true}
}

func (v NullableInlineResponse20087) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


