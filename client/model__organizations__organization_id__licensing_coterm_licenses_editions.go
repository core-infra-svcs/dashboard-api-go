/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdLicensingCotermLicensesEditions struct for OrganizationsOrganizationIdLicensingCotermLicensesEditions
type OrganizationsOrganizationIdLicensingCotermLicensesEditions struct {
	// The name of the license edition
	Edition *string `json:"edition,omitempty"`
	// The product type of the license edition
	ProductType *string `json:"productType,omitempty"`
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesEditions instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesEditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensingCotermLicensesEditions() *OrganizationsOrganizationIdLicensingCotermLicensesEditions {
	this := OrganizationsOrganizationIdLicensingCotermLicensesEditions{}
	return &this
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesEditionsWithDefaults instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesEditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensingCotermLicensesEditionsWithDefaults() *OrganizationsOrganizationIdLicensingCotermLicensesEditions {
	this := OrganizationsOrganizationIdLicensingCotermLicensesEditions{}
	return &this
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) GetEdition() string {
	if o == nil || isNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) GetEditionOk() (*string, bool) {
	if o == nil || isNil(o.Edition) {
    return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) HasEdition() bool {
	if o != nil && !isNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) SetEdition(v string) {
	o.Edition = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesEditions) SetProductType(v string) {
	o.ProductType = &v
}

func (o OrganizationsOrganizationIdLicensingCotermLicensesEditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions struct {
	value *OrganizationsOrganizationIdLicensingCotermLicensesEditions
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) Get() *OrganizationsOrganizationIdLicensingCotermLicensesEditions {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) Set(val *OrganizationsOrganizationIdLicensingCotermLicensesEditions) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensingCotermLicensesEditions(val *OrganizationsOrganizationIdLicensingCotermLicensesEditions) *NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions {
	return &NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesEditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


