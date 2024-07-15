/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200258LicenseTypes struct for InlineResponse200258LicenseTypes
type InlineResponse200258LicenseTypes struct {
	// License type
	LicenseType *string `json:"licenseType,omitempty"`
	Counts *InlineResponse200258Counts `json:"counts,omitempty"`
}

// NewInlineResponse200258LicenseTypes instantiates a new InlineResponse200258LicenseTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200258LicenseTypes() *InlineResponse200258LicenseTypes {
	this := InlineResponse200258LicenseTypes{}
	return &this
}

// NewInlineResponse200258LicenseTypesWithDefaults instantiates a new InlineResponse200258LicenseTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200258LicenseTypesWithDefaults() *InlineResponse200258LicenseTypes {
	this := InlineResponse200258LicenseTypes{}
	return &this
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *InlineResponse200258LicenseTypes) GetLicenseType() string {
	if o == nil || isNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258LicenseTypes) GetLicenseTypeOk() (*string, bool) {
	if o == nil || isNil(o.LicenseType) {
    return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *InlineResponse200258LicenseTypes) HasLicenseType() bool {
	if o != nil && !isNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *InlineResponse200258LicenseTypes) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200258LicenseTypes) GetCounts() InlineResponse200258Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200258Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258LicenseTypes) GetCountsOk() (*InlineResponse200258Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200258LicenseTypes) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200258Counts and assigns it to the Counts field.
func (o *InlineResponse200258LicenseTypes) SetCounts(v InlineResponse200258Counts) {
	o.Counts = &v
}

func (o InlineResponse200258LicenseTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200258LicenseTypes struct {
	value *InlineResponse200258LicenseTypes
	isSet bool
}

func (v NullableInlineResponse200258LicenseTypes) Get() *InlineResponse200258LicenseTypes {
	return v.value
}

func (v *NullableInlineResponse200258LicenseTypes) Set(val *InlineResponse200258LicenseTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200258LicenseTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200258LicenseTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200258LicenseTypes(val *InlineResponse200258LicenseTypes) *NullableInlineResponse200258LicenseTypes {
	return &NullableInlineResponse200258LicenseTypes{value: val, isSet: true}
}

func (v NullableInlineResponse200258LicenseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200258LicenseTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

