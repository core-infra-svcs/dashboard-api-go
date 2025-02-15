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

// AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount Smart Account data
type AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount struct {
	// Smart Account ID
	Id *string `json:"id,omitempty"`
	// The name of the smart account
	Name *string `json:"name,omitempty"`
	// The domain of the Smart Account
	Domain *string `json:"domain,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount instantiates a new AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount() *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount {
	this := AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccountWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccountWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount {
	this := AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) SetName(v string) {
	o.Name = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetDomain() string {
	if o == nil || isNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) GetDomainOk() (*string, bool) {
	if o == nil || isNil(o.Domain) {
    return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) HasDomain() bool {
	if o != nil && !isNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) SetDomain(v string) {
	o.Domain = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount struct {
	value *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) Get() *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) Set(val *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount(val *AdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) *NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSmartAccountAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


