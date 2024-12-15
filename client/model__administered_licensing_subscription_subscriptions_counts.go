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

// AdministeredLicensingSubscriptionSubscriptionsCounts Numeric breakdown of network, organizations, entitlement counts
type AdministeredLicensingSubscriptionSubscriptionsCounts struct {
	Seats *AdministeredLicensingSubscriptionSubscriptionsCountsSeats `json:"seats,omitempty"`
	// Number of networks bound to this subscription
	Networks *int32 `json:"networks,omitempty"`
	// Number of organizations bound to this subscription
	Organizations *int32 `json:"organizations,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsCounts instantiates a new AdministeredLicensingSubscriptionSubscriptionsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsCounts() *AdministeredLicensingSubscriptionSubscriptionsCounts {
	this := AdministeredLicensingSubscriptionSubscriptionsCounts{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsCountsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsCountsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsCounts {
	this := AdministeredLicensingSubscriptionSubscriptionsCounts{}
	return &this
}

// GetSeats returns the Seats field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetSeats() AdministeredLicensingSubscriptionSubscriptionsCountsSeats {
	if o == nil || isNil(o.Seats) {
		var ret AdministeredLicensingSubscriptionSubscriptionsCountsSeats
		return ret
	}
	return *o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetSeatsOk() (*AdministeredLicensingSubscriptionSubscriptionsCountsSeats, bool) {
	if o == nil || isNil(o.Seats) {
    return nil, false
	}
	return o.Seats, true
}

// HasSeats returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) HasSeats() bool {
	if o != nil && !isNil(o.Seats) {
		return true
	}

	return false
}

// SetSeats gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsCountsSeats and assigns it to the Seats field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) SetSeats(v AdministeredLicensingSubscriptionSubscriptionsCountsSeats) {
	o.Seats = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetNetworks() int32 {
	if o == nil || isNil(o.Networks) {
		var ret int32
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetNetworksOk() (*int32, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given int32 and assigns it to the Networks field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) SetNetworks(v int32) {
	o.Networks = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetOrganizations() int32 {
	if o == nil || isNil(o.Organizations) {
		var ret int32
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) GetOrganizationsOk() (*int32, bool) {
	if o == nil || isNil(o.Organizations) {
    return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) HasOrganizations() bool {
	if o != nil && !isNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given int32 and assigns it to the Organizations field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCounts) SetOrganizations(v int32) {
	o.Organizations = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Seats) {
		toSerialize["seats"] = o.Seats
	}
	if !isNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !isNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsCounts struct {
	value *AdministeredLicensingSubscriptionSubscriptionsCounts
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCounts) Get() *AdministeredLicensingSubscriptionSubscriptionsCounts {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCounts) Set(val *AdministeredLicensingSubscriptionSubscriptionsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsCounts(val *AdministeredLicensingSubscriptionSubscriptionsCounts) *NullableAdministeredLicensingSubscriptionSubscriptionsCounts {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsCounts{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


