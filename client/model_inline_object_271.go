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

// InlineObject271 struct for InlineObject271
type InlineObject271 struct {
	// List of replacments to perform
	Swaps []OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps `json:"swaps"`
}

// NewInlineObject271 instantiates a new InlineObject271 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject271(swaps []OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) *InlineObject271 {
	this := InlineObject271{}
	this.Swaps = swaps
	return &this
}

// NewInlineObject271WithDefaults instantiates a new InlineObject271 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject271WithDefaults() *InlineObject271 {
	this := InlineObject271{}
	return &this
}

// GetSwaps returns the Swaps field value
func (o *InlineObject271) GetSwaps() []OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps {
	if o == nil {
		var ret []OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps
		return ret
	}

	return o.Swaps
}

// GetSwapsOk returns a tuple with the Swaps field value
// and a boolean to check if the value has been set.
func (o *InlineObject271) GetSwapsOk() ([]OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps, bool) {
	if o == nil {
    return nil, false
	}
	return o.Swaps, true
}

// SetSwaps sets field value
func (o *InlineObject271) SetSwaps(v []OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps) {
	o.Swaps = v
}

func (o InlineObject271) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["swaps"] = o.Swaps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject271 struct {
	value *InlineObject271
	isSet bool
}

func (v NullableInlineObject271) Get() *InlineObject271 {
	return v.value
}

func (v *NullableInlineObject271) Set(val *InlineObject271) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject271) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject271) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject271(val *InlineObject271) *NullableInlineObject271 {
	return &NullableInlineObject271{value: val, isSet: true}
}

func (v NullableInlineObject271) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject271) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


