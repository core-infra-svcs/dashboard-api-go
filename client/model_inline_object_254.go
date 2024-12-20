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

// InlineObject254 struct for InlineObject254
type InlineObject254 struct {
	Destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination `json:"destination"`
	// The list of licenses to move
	Licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses `json:"licenses"`
}

// NewInlineObject254 instantiates a new InlineObject254 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject254(destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *InlineObject254 {
	this := InlineObject254{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewInlineObject254WithDefaults instantiates a new InlineObject254 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject254WithDefaults() *InlineObject254 {
	this := InlineObject254{}
	return &this
}

// GetDestination returns the Destination field value
func (o *InlineObject254) GetDestination() OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	if o == nil {
		var ret OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineObject254) GetDestinationOk() (*OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineObject254) SetDestination(v OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *InlineObject254) GetLicenses() []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *InlineObject254) GetLicensesOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses, bool) {
	if o == nil {
    return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *InlineObject254) SetLicenses(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	o.Licenses = v
}

func (o InlineObject254) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject254 struct {
	value *InlineObject254
	isSet bool
}

func (v NullableInlineObject254) Get() *InlineObject254 {
	return v.value
}

func (v *NullableInlineObject254) Set(val *InlineObject254) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject254) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject254) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject254(val *InlineObject254) *NullableInlineObject254 {
	return &NullableInlineObject254{value: val, isSet: true}
}

func (v NullableInlineObject254) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject254) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


