/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdministeredLicensingSubscriptionSubscriptionsCountsSeats Seat distribution
type AdministeredLicensingSubscriptionSubscriptionsCountsSeats struct {
	// Number of seats in use
	Assigned *int32 `json:"assigned,omitempty"`
	// Number of seats available for use
	Available *int32 `json:"available,omitempty"`
	// Total number of seats provided by this subscription
	Limit *int32 `json:"limit,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsCountsSeats instantiates a new AdministeredLicensingSubscriptionSubscriptionsCountsSeats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsCountsSeats() *AdministeredLicensingSubscriptionSubscriptionsCountsSeats {
	this := AdministeredLicensingSubscriptionSubscriptionsCountsSeats{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsCountsSeatsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsCountsSeats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsCountsSeatsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsCountsSeats {
	this := AdministeredLicensingSubscriptionSubscriptionsCountsSeats{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetAssigned() int32 {
	if o == nil || isNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetAssignedOk() (*int32, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetAvailable() int32 {
	if o == nil || isNil(o.Available) {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetAvailableOk() (*int32, bool) {
	if o == nil || isNil(o.Available) {
    return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) SetAvailable(v int32) {
	o.Available = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) SetLimit(v int32) {
	o.Limit = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsCountsSeats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !isNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats struct {
	value *AdministeredLicensingSubscriptionSubscriptionsCountsSeats
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) Get() *AdministeredLicensingSubscriptionSubscriptionsCountsSeats {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) Set(val *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats(val *AdministeredLicensingSubscriptionSubscriptionsCountsSeats) *NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsCountsSeats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


