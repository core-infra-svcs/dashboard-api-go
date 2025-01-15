/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdministeredLicensingSubscriptionSubscriptionsEntitlements struct for AdministeredLicensingSubscriptionSubscriptionsEntitlements
type AdministeredLicensingSubscriptionSubscriptionsEntitlements struct {
	// SKU of the required product
	Sku *string `json:"sku,omitempty"`
	// Web order line ID
	WebOrderLineId *string `json:"webOrderLineId,omitempty"`
	Seats *AdministeredLicensingSubscriptionSubscriptionsSeats `json:"seats,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsEntitlements instantiates a new AdministeredLicensingSubscriptionSubscriptionsEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsEntitlements() *AdministeredLicensingSubscriptionSubscriptionsEntitlements {
	this := AdministeredLicensingSubscriptionSubscriptionsEntitlements{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsEntitlementsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsEntitlementsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsEntitlements {
	this := AdministeredLicensingSubscriptionSubscriptionsEntitlements{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) SetSku(v string) {
	o.Sku = &v
}

// GetWebOrderLineId returns the WebOrderLineId field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetWebOrderLineId() string {
	if o == nil || isNil(o.WebOrderLineId) {
		var ret string
		return ret
	}
	return *o.WebOrderLineId
}

// GetWebOrderLineIdOk returns a tuple with the WebOrderLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetWebOrderLineIdOk() (*string, bool) {
	if o == nil || isNil(o.WebOrderLineId) {
    return nil, false
	}
	return o.WebOrderLineId, true
}

// HasWebOrderLineId returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) HasWebOrderLineId() bool {
	if o != nil && !isNil(o.WebOrderLineId) {
		return true
	}

	return false
}

// SetWebOrderLineId gets a reference to the given string and assigns it to the WebOrderLineId field.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) SetWebOrderLineId(v string) {
	o.WebOrderLineId = &v
}

// GetSeats returns the Seats field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetSeats() AdministeredLicensingSubscriptionSubscriptionsSeats {
	if o == nil || isNil(o.Seats) {
		var ret AdministeredLicensingSubscriptionSubscriptionsSeats
		return ret
	}
	return *o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) GetSeatsOk() (*AdministeredLicensingSubscriptionSubscriptionsSeats, bool) {
	if o == nil || isNil(o.Seats) {
    return nil, false
	}
	return o.Seats, true
}

// HasSeats returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) HasSeats() bool {
	if o != nil && !isNil(o.Seats) {
		return true
	}

	return false
}

// SetSeats gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsSeats and assigns it to the Seats field.
func (o *AdministeredLicensingSubscriptionSubscriptionsEntitlements) SetSeats(v AdministeredLicensingSubscriptionSubscriptionsSeats) {
	o.Seats = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.WebOrderLineId) {
		toSerialize["webOrderLineId"] = o.WebOrderLineId
	}
	if !isNil(o.Seats) {
		toSerialize["seats"] = o.Seats
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements struct {
	value *AdministeredLicensingSubscriptionSubscriptionsEntitlements
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) Get() *AdministeredLicensingSubscriptionSubscriptionsEntitlements {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) Set(val *AdministeredLicensingSubscriptionSubscriptionsEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsEntitlements(val *AdministeredLicensingSubscriptionSubscriptionsEntitlements) *NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


