/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdministeredLicensingSubscriptionSubscriptionsSeats Seat distribution
type AdministeredLicensingSubscriptionSubscriptionsSeats struct {
	// Number of seats in use
	Assigned *int32 `json:"assigned,omitempty"`
	// Number of seats available for use
	Available *int32 `json:"available,omitempty"`
	// Total number of seats provided by this subscription for this sku
	Limit *int32 `json:"limit,omitempty"`
}

// NewAdministeredLicensingSubscriptionSubscriptionsSeats instantiates a new AdministeredLicensingSubscriptionSubscriptionsSeats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministeredLicensingSubscriptionSubscriptionsSeats() *AdministeredLicensingSubscriptionSubscriptionsSeats {
	this := AdministeredLicensingSubscriptionSubscriptionsSeats{}
	return &this
}

// NewAdministeredLicensingSubscriptionSubscriptionsSeatsWithDefaults instantiates a new AdministeredLicensingSubscriptionSubscriptionsSeats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministeredLicensingSubscriptionSubscriptionsSeatsWithDefaults() *AdministeredLicensingSubscriptionSubscriptionsSeats {
	this := AdministeredLicensingSubscriptionSubscriptionsSeats{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetAssigned() int32 {
	if o == nil || isNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetAssignedOk() (*int32, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetAvailable() int32 {
	if o == nil || isNil(o.Available) {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetAvailableOk() (*int32, bool) {
	if o == nil || isNil(o.Available) {
    return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) HasAvailable() bool {
	if o != nil && !isNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) SetAvailable(v int32) {
	o.Available = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *AdministeredLicensingSubscriptionSubscriptionsSeats) SetLimit(v int32) {
	o.Limit = &v
}

func (o AdministeredLicensingSubscriptionSubscriptionsSeats) MarshalJSON() ([]byte, error) {
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

type NullableAdministeredLicensingSubscriptionSubscriptionsSeats struct {
	value *AdministeredLicensingSubscriptionSubscriptionsSeats
	isSet bool
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSeats) Get() *AdministeredLicensingSubscriptionSubscriptionsSeats {
	return v.value
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSeats) Set(val *AdministeredLicensingSubscriptionSubscriptionsSeats) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSeats) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSeats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministeredLicensingSubscriptionSubscriptionsSeats(val *AdministeredLicensingSubscriptionSubscriptionsSeats) *NullableAdministeredLicensingSubscriptionSubscriptionsSeats {
	return &NullableAdministeredLicensingSubscriptionSubscriptionsSeats{value: val, isSet: true}
}

func (v NullableAdministeredLicensingSubscriptionSubscriptionsSeats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministeredLicensingSubscriptionSubscriptionsSeats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


