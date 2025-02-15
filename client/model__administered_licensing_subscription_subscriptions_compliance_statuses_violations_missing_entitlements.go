/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements struct for AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements struct {
	// SKU of the required product
	Sku *string `json:"sku,omitempty"`
	// Number required
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlementsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlementsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) SetSku(v string) {
	o.Sku = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) GetQuantity() int32 {
	if o == nil || isNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) GetQuantityOk() (*int32, bool) {
	if o == nil || isNil(o.Quantity) {
    return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) HasQuantity() bool {
	if o != nil && !isNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements struct {
	value *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) Get() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) Set(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


