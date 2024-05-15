/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination Destination data for the license move
type OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination struct {
	// The organization to move the license to
	OrganizationId *string `json:"organizationId,omitempty"`
	// The claim mode of the moved license
	Mode *string `json:"mode,omitempty"`
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination() *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination{}
	return &this
}

// NewOrganizationsOrganizationIdLicensingCotermLicensesMoveDestinationWithDefaults instantiates a new OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdLicensingCotermLicensesMoveDestinationWithDefaults() *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	this := OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) SetMode(v string) {
	o.Mode = &v
}

func (o OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination struct {
	value *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
	isSet bool
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) Get() *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) Set(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination(val *OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	return &NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


