/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject91 struct for InlineObject91
type InlineObject91 struct {
	UpgradeWindow *InlineResponse20088UpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *NetworksNetworkIdFirmwareUpgradesProducts `json:"products,omitempty"`
}

// NewInlineObject91 instantiates a new InlineObject91 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject91() *InlineObject91 {
	this := InlineObject91{}
	return &this
}

// NewInlineObject91WithDefaults instantiates a new InlineObject91 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject91WithDefaults() *InlineObject91 {
	this := InlineObject91{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *InlineObject91) GetUpgradeWindow() InlineResponse20088UpgradeWindow {
	if o == nil || isNil(o.UpgradeWindow) {
		var ret InlineResponse20088UpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetUpgradeWindowOk() (*InlineResponse20088UpgradeWindow, bool) {
	if o == nil || isNil(o.UpgradeWindow) {
    return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *InlineObject91) HasUpgradeWindow() bool {
	if o != nil && !isNil(o.UpgradeWindow) {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given InlineResponse20088UpgradeWindow and assigns it to the UpgradeWindow field.
func (o *InlineObject91) SetUpgradeWindow(v InlineResponse20088UpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineObject91) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineObject91) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineObject91) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineObject91) GetProducts() NetworksNetworkIdFirmwareUpgradesProducts {
	if o == nil || isNil(o.Products) {
		var ret NetworksNetworkIdFirmwareUpgradesProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject91) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesProducts, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineObject91) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given NetworksNetworkIdFirmwareUpgradesProducts and assigns it to the Products field.
func (o *InlineObject91) SetProducts(v NetworksNetworkIdFirmwareUpgradesProducts) {
	o.Products = &v
}

func (o InlineObject91) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject91 struct {
	value *InlineObject91
	isSet bool
}

func (v NullableInlineObject91) Get() *InlineObject91 {
	return v.value
}

func (v *NullableInlineObject91) Set(val *InlineObject91) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject91) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject91) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject91(val *InlineObject91) *NullableInlineObject91 {
	return &NullableInlineObject91{value: val, isSet: true}
}

func (v NullableInlineObject91) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject91) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


