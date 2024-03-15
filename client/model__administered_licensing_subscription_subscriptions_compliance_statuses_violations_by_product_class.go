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

// AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass struct for AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass struct {
	// Name of the product class
	ProductClass *string `json:"productClass,omitempty"`
	// End date of the grace period in ISO 8601 format
	GracePeriodEndsAt *string `json:"gracePeriodEndsAt,omitempty"`
	Missing *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing `json:"missing,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass{}
	return &this
}

// GetProductClass returns the ProductClass field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetProductClass() string {
	if o == nil || isNil(o.ProductClass) {
		var ret string
		return ret
	}
	return *o.ProductClass
}

// GetProductClassOk returns a tuple with the ProductClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetProductClassOk() (*string, bool) {
	if o == nil || isNil(o.ProductClass) {
    return nil, false
	}
	return o.ProductClass, true
}

// HasProductClass returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) HasProductClass() bool {
	if o != nil && !isNil(o.ProductClass) {
		return true
	}

	return false
}

// SetProductClass gets a reference to the given string and assigns it to the ProductClass field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) SetProductClass(v string) {
	o.ProductClass = &v
}

// GetGracePeriodEndsAt returns the GracePeriodEndsAt field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetGracePeriodEndsAt() string {
	if o == nil || isNil(o.GracePeriodEndsAt) {
		var ret string
		return ret
	}
	return *o.GracePeriodEndsAt
}

// GetGracePeriodEndsAtOk returns a tuple with the GracePeriodEndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetGracePeriodEndsAtOk() (*string, bool) {
	if o == nil || isNil(o.GracePeriodEndsAt) {
    return nil, false
	}
	return o.GracePeriodEndsAt, true
}

// HasGracePeriodEndsAt returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) HasGracePeriodEndsAt() bool {
	if o != nil && !isNil(o.GracePeriodEndsAt) {
		return true
	}

	return false
}

// SetGracePeriodEndsAt gets a reference to the given string and assigns it to the GracePeriodEndsAt field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) SetGracePeriodEndsAt(v string) {
	o.GracePeriodEndsAt = &v
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetMissing() AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing {
	if o == nil || isNil(o.Missing) {
		var ret AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing
		return ret
	}
	return *o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) GetMissingOk() (*AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing, bool) {
	if o == nil || isNil(o.Missing) {
    return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) HasMissing() bool {
	if o != nil && !isNil(o.Missing) {
		return true
	}

	return false
}

// SetMissing gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing and assigns it to the Missing field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) SetMissing(v AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) {
	o.Missing = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProductClass) {
		toSerialize["productClass"] = o.ProductClass
	}
	if !isNil(o.GracePeriodEndsAt) {
		toSerialize["gracePeriodEndsAt"] = o.GracePeriodEndsAt
	}
	if !isNil(o.Missing) {
		toSerialize["missing"] = o.Missing
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass struct {
	value *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) Get() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) Set(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


