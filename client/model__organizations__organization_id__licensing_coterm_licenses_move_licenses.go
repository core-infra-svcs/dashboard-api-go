/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses struct for OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
type OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses struct {
	// The license key to move counts from
	Key string `json:"key"`
	// The counts to move from the license by model type
	Counts []OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts `json:"counts"`
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses(key string, counts []OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses{}
	this.Key = key
	this.Counts = counts
	return &this
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveLicensesWithDefaults instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveLicensesWithDefaults() *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses{}
	return &this
}

// GetKey returns the Key field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) SetKey(v string) {
	o.Key = v
}

// GetCounts returns the Counts field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) GetCounts() []OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) GetCountsOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts, bool) {
	if o == nil {
    return nil, false
	}
	return o.Counts, true
}

// SetCounts sets field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) SetCounts(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) {
	o.Counts = v
}

func (o OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses struct {
	value *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) Get() *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) Set(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	return &NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


