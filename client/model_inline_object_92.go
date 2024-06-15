/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject92 struct for InlineObject92
type InlineObject92 struct {
	UpgradeWindow *InlineResponse20089UpgradeWindow `json:"upgradeWindow,omitempty"`
	// The timezone for the network
	Timezone *string `json:"timezone,omitempty"`
	Products *NetworksNetworkIdFirmwareUpgradesProducts `json:"products,omitempty"`
}

// NewInlineObject92 instantiates a new InlineObject92 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject92() *InlineObject92 {
	this := InlineObject92{}
	return &this
}

// NewInlineObject92WithDefaults instantiates a new InlineObject92 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject92WithDefaults() *InlineObject92 {
	this := InlineObject92{}
	return &this
}

// GetUpgradeWindow returns the UpgradeWindow field value if set, zero value otherwise.
func (o *InlineObject92) GetUpgradeWindow() InlineResponse20089UpgradeWindow {
	if o == nil || isNil(o.UpgradeWindow) {
		var ret InlineResponse20089UpgradeWindow
		return ret
	}
	return *o.UpgradeWindow
}

// GetUpgradeWindowOk returns a tuple with the UpgradeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetUpgradeWindowOk() (*InlineResponse20089UpgradeWindow, bool) {
	if o == nil || isNil(o.UpgradeWindow) {
    return nil, false
	}
	return o.UpgradeWindow, true
}

// HasUpgradeWindow returns a boolean if a field has been set.
func (o *InlineObject92) HasUpgradeWindow() bool {
	if o != nil && !isNil(o.UpgradeWindow) {
		return true
	}

	return false
}

// SetUpgradeWindow gets a reference to the given InlineResponse20089UpgradeWindow and assigns it to the UpgradeWindow field.
func (o *InlineObject92) SetUpgradeWindow(v InlineResponse20089UpgradeWindow) {
	o.UpgradeWindow = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *InlineObject92) GetTimezone() string {
	if o == nil || isNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetTimezoneOk() (*string, bool) {
	if o == nil || isNil(o.Timezone) {
    return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *InlineObject92) HasTimezone() bool {
	if o != nil && !isNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *InlineObject92) SetTimezone(v string) {
	o.Timezone = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InlineObject92) GetProducts() NetworksNetworkIdFirmwareUpgradesProducts {
	if o == nil || isNil(o.Products) {
		var ret NetworksNetworkIdFirmwareUpgradesProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject92) GetProductsOk() (*NetworksNetworkIdFirmwareUpgradesProducts, bool) {
	if o == nil || isNil(o.Products) {
    return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InlineObject92) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given NetworksNetworkIdFirmwareUpgradesProducts and assigns it to the Products field.
func (o *InlineObject92) SetProducts(v NetworksNetworkIdFirmwareUpgradesProducts) {
	o.Products = &v
}

func (o InlineObject92) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject92 struct {
	value *InlineObject92
	isSet bool
}

func (v NullableInlineObject92) Get() *InlineObject92 {
	return v.value
}

func (v *NullableInlineObject92) Set(val *InlineObject92) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject92) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject92) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject92(val *InlineObject92) *NullableInlineObject92 {
	return &NullableInlineObject92{value: val, isSet: true}
}

func (v NullableInlineObject92) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject92) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


