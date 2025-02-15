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

// AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing Missing entitlements details
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing struct {
	// List of missing entitlements
	Entitlements []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements `json:"entitlements,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) GetEntitlements() []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	if o == nil || isNil(o.Entitlements) {
		var ret []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) GetEntitlementsOk() ([]AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements, bool) {
	if o == nil || isNil(o.Entitlements) {
    return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements and assigns it to the Entitlements field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) SetEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) {
	o.Entitlements = v
}

func (o AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing struct {
	value *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) Get() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) Set(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


