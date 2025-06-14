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

// AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription Subscription details
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription struct {
	// Subscription's ID
	Id *string `json:"id,omitempty"`
	// Friendly name to identify the subscription
	Name *string `json:"name,omitempty"`
	// One of the following: \"inactive\" | \"active\" | \"out_of_compliance\" | \"expired\" | \"canceled\"
	Status *string `json:"status,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscriptionWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscriptionWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription {
	this := AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) SetStatus(v string) {
	o.Status = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription struct {
	value *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) Get() *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) Set(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription(val *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


