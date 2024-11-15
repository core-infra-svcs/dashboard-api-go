/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject253 struct for InlineObject253
type InlineObject253 struct {
	Destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination `json:"destination"`
	// The list of licenses to move
	Licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses `json:"licenses"`
}

// NewInlineObject253 instantiates a new InlineObject253 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject253(destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *InlineObject253 {
	this := InlineObject253{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewInlineObject253WithDefaults instantiates a new InlineObject253 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject253WithDefaults() *InlineObject253 {
	this := InlineObject253{}
	return &this
}

// GetDestination returns the Destination field value
func (o *InlineObject253) GetDestination() OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	if o == nil {
		var ret OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineObject253) GetDestinationOk() (*OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineObject253) SetDestination(v OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *InlineObject253) GetLicenses() []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *InlineObject253) GetLicensesOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses, bool) {
	if o == nil {
    return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *InlineObject253) SetLicenses(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	o.Licenses = v
}

func (o InlineObject253) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject253 struct {
	value *InlineObject253
	isSet bool
}

func (v NullableInlineObject253) Get() *InlineObject253 {
	return v.value
}

func (v *NullableInlineObject253) Set(val *InlineObject253) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject253) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject253) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject253(val *InlineObject253) *NullableInlineObject253 {
	return &NullableInlineObject253{value: val, isSet: true}
}

func (v NullableInlineObject253) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject253) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


