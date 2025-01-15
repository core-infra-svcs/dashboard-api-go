/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200274LicenseTypes struct for InlineResponse200274LicenseTypes
type InlineResponse200274LicenseTypes struct {
	// License type
	LicenseType *string `json:"licenseType,omitempty"`
	Counts *InlineResponse200274Counts `json:"counts,omitempty"`
}

// NewInlineResponse200274LicenseTypes instantiates a new InlineResponse200274LicenseTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274LicenseTypes() *InlineResponse200274LicenseTypes {
	this := InlineResponse200274LicenseTypes{}
	return &this
}

// NewInlineResponse200274LicenseTypesWithDefaults instantiates a new InlineResponse200274LicenseTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274LicenseTypesWithDefaults() *InlineResponse200274LicenseTypes {
	this := InlineResponse200274LicenseTypes{}
	return &this
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *InlineResponse200274LicenseTypes) GetLicenseType() string {
	if o == nil || isNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274LicenseTypes) GetLicenseTypeOk() (*string, bool) {
	if o == nil || isNil(o.LicenseType) {
    return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *InlineResponse200274LicenseTypes) HasLicenseType() bool {
	if o != nil && !isNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *InlineResponse200274LicenseTypes) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200274LicenseTypes) GetCounts() InlineResponse200274Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200274Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274LicenseTypes) GetCountsOk() (*InlineResponse200274Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200274LicenseTypes) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200274Counts and assigns it to the Counts field.
func (o *InlineResponse200274LicenseTypes) SetCounts(v InlineResponse200274Counts) {
	o.Counts = &v
}

func (o InlineResponse200274LicenseTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274LicenseTypes struct {
	value *InlineResponse200274LicenseTypes
	isSet bool
}

func (v NullableInlineResponse200274LicenseTypes) Get() *InlineResponse200274LicenseTypes {
	return v.value
}

func (v *NullableInlineResponse200274LicenseTypes) Set(val *InlineResponse200274LicenseTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274LicenseTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274LicenseTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274LicenseTypes(val *InlineResponse200274LicenseTypes) *NullableInlineResponse200274LicenseTypes {
	return &NullableInlineResponse200274LicenseTypes{value: val, isSet: true}
}

func (v NullableInlineResponse200274LicenseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274LicenseTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

