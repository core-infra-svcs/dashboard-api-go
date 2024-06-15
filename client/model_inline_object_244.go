/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject244 struct for InlineObject244
type InlineObject244 struct {
	Destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination `json:"destination"`
	// The list of licenses to move
	Licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses `json:"licenses"`
}

// NewInlineObject244 instantiates a new InlineObject244 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject244(destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *InlineObject244 {
	this := InlineObject244{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewInlineObject244WithDefaults instantiates a new InlineObject244 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject244WithDefaults() *InlineObject244 {
	this := InlineObject244{}
	return &this
}

// GetDestination returns the Destination field value
func (o *InlineObject244) GetDestination() OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	if o == nil {
		var ret OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineObject244) GetDestinationOk() (*OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineObject244) SetDestination(v OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *InlineObject244) GetLicenses() []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *InlineObject244) GetLicensesOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses, bool) {
	if o == nil {
    return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *InlineObject244) SetLicenses(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	o.Licenses = v
}

func (o InlineObject244) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject244 struct {
	value *InlineObject244
	isSet bool
}

func (v NullableInlineObject244) Get() *InlineObject244 {
	return v.value
}

func (v *NullableInlineObject244) Set(val *InlineObject244) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject244) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject244) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject244(val *InlineObject244) *NullableInlineObject244 {
	return &NullableInlineObject244{value: val, isSet: true}
}

func (v NullableInlineObject244) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject244) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


