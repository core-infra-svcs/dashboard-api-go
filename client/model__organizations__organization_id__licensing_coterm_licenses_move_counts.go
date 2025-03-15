/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts struct for OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts
type OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts struct {
	// The license model type to move counts of
	Model string `json:"model"`
	// The number of counts to move
	Count int32 `json:"count"`
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts(model string, count int32) *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts{}
	this.Model = model
	this.Count = count
	return &this
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveCountsWithDefaults instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveCountsWithDefaults() *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts{}
	return &this
}

// GetModel returns the Model field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) SetModel(v string) {
	o.Model = v
}

// GetCount returns the Count field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) GetCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) SetCount(v int32) {
	o.Count = v
}

func (o OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts struct {
	value *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) Get() *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) Set(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts {
	return &NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


