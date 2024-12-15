/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations Violations
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations struct {
	// List of violations by product class that are not compliance
	ByProductClass []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass `json:"byProductClass,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations{}
	return &this
}

// GetByProductClass returns the ByProductClass field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) GetByProductClass() []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
	if o == nil || isNil(o.ByProductClass) {
		var ret []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass
		return ret
	}
	return o.ByProductClass
}

// GetByProductClassOk returns a tuple with the ByProductClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) GetByProductClassOk() ([]AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass, bool) {
	if o == nil || isNil(o.ByProductClass) {
    return nil, false
	}
	return o.ByProductClass, true
}

// HasByProductClass returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) HasByProductClass() bool {
	if o != nil && !isNil(o.ByProductClass) {
		return true
	}

	return false
}

// SetByProductClass gets a reference to the given []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass and assigns it to the ByProductClass field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) SetByProductClass(v []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass) {
	o.ByProductClass = v
}

func (o AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByProductClass) {
		toSerialize["byProductClass"] = o.ByProductClass
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations struct {
	value *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) Get() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) Set(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


