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

// InlineResponse20095 struct for InlineResponse20095
type InlineResponse20095 struct {
	UpgradeWindow *InlineResponse20095UpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *InlineResponse20095Products `json:"products,omitempty"`
}

// NewInlineResponse20095 instantiates a new InlineResponse20095 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// NewInlineResponse20095WithDefaults instantiates a new InlineResponse20095 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095WithDefaults() *InlineResponse20095 {
	this := InlineResponse20095{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *InlineResponse20095) GetUpgradeWindow() InlineResponse20095UpgradeWindow {
	if o == nil || isNil(o.UpgradeWindow) {
		var ret InlineResponse20095UpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetUpgradeWindowOk() (*InlineResponse20095UpgradeWindow, bool) {
	if o == nil || isNil(o.UpgradeWindow) {
    return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *InlineResponse20095) HasUpgradeWindow() bool {
	if o != nil && !isNil(o.UpgradeWindow) {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given InlineResponse20095UpgradeWindow and assigns it to the UpgradeWindow field.
func (o *InlineResponse20095) SetUpgradeWindow(v InlineResponse20095UpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineResponse20095) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineResponse20095) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineResponse20095) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineResponse20095) GetProducts() InlineResponse20095Products {
	if o == nil || isNil(o.Products) {
		var ret InlineResponse20095Products
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095) GetProductsOk() (*InlineResponse20095Products, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineResponse20095) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given InlineResponse20095Products and assigns it to the Products field.
func (o *InlineResponse20095) SetProducts(v InlineResponse20095Products) {
	o.Products = &v
}

func (o InlineResponse20095) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20095 struct {
	value *InlineResponse20095
	isSet bool
}

func (v NullableInlineResponse20095) Get() *InlineResponse20095 {
	return v.value
}

func (v *NullableInlineResponse20095) Set(val *InlineResponse20095) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095(val *InlineResponse20095) *NullableInlineResponse20095 {
	return &NullableInlineResponse20095{value: val, isSet: true}
}

func (v NullableInlineResponse20095) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


