/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses struct for OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses
type OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses struct {
	// Permanently queued license ID
	Id *string `json:"id,omitempty"`
	// License type
	LicenseType *string `json:"licenseType,omitempty"`
	// License key
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Order number
	OrderNumber *string `json:"orderNumber,omitempty"`
	// The duration of the individual license
	DurationInDays *int32 `json:"durationInDays,omitempty"`
}

// NewOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses instantiates a new OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses() *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses {
	this := OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses{}
	return &this
}

// NewOrganizationsOrganizationIdLicensesPermanentlyQueuedLicensesWithDefaults instantiates a new OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensesPermanentlyQueuedLicensesWithDefaults() *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses {
	this := OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) SetId(v string) {
	o.Id = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetLicenseType() string {
	if o == nil || isNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetLicenseTypeOk() (*string, bool) {
	if o == nil || isNil(o.LicenseType) {
    return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) HasLicenseType() bool {
	if o != nil && !isNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetLicenseKey() string {
	if o == nil || isNil(o.LicenseKey) {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetLicenseKeyOk() (*string, bool) {
	if o == nil || isNil(o.LicenseKey) {
    return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) HasLicenseKey() bool {
	if o != nil && !isNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetOrderNumber() string {
	if o == nil || isNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetOrderNumberOk() (*string, bool) {
	if o == nil || isNil(o.OrderNumber) {
    return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) HasOrderNumber() bool {
	if o != nil && !isNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetDurationInDays returns the DurationInDays field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetDurationInDays() int32 {
	if o == nil || isNil(o.DurationInDays) {
		var ret int32
		return ret
	}
	return *o.DurationInDays
}

// GetDurationInDaysOk returns a tuple with the DurationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) GetDurationInDaysOk() (*int32, bool) {
	if o == nil || isNil(o.DurationInDays) {
    return nil, false
	}
	return o.DurationInDays, true
}

// HasDurationInDays returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) HasDurationInDays() bool {
	if o != nil && !isNil(o.DurationInDays) {
		return true
	}

	return false
}

// SetDurationInDays gets a reference to the given int32 and assigns it to the DurationInDays field.
func (o *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) SetDurationInDays(v int32) {
	o.DurationInDays = &v
}

func (o OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !isNil(o.LicenseKey) {
		toSerialize["licenseKey"] = o.LicenseKey
	}
	if !isNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !isNil(o.DurationInDays) {
		toSerialize["durationInDays"] = o.DurationInDays
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses struct {
	value *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) Get() *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) Set(val *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses(val *OrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) *NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses {
	return &NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensesPermanentlyQueuedLicenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


