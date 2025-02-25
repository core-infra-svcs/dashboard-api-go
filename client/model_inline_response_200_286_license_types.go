/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200286LicenseTypes struct for InlineResponse200286LicenseTypes
type InlineResponse200286LicenseTypes struct {
	// License type
	LicenseType *string `json:"licenseType,omitempty"`
	Counts *InlineResponse200286Counts `json:"counts,omitempty"`
}

// NewInlineResponse200286LicenseTypes instantiates a new InlineResponse200286LicenseTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200286LicenseTypes() *InlineResponse200286LicenseTypes {
	this := InlineResponse200286LicenseTypes{}
	return &this
}

// NewInlineResponse200286LicenseTypesWithDefaults instantiates a new InlineResponse200286LicenseTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200286LicenseTypesWithDefaults() *InlineResponse200286LicenseTypes {
	this := InlineResponse200286LicenseTypes{}
	return &this
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *InlineResponse200286LicenseTypes) GetLicenseType() string {
	if o == nil || isNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286LicenseTypes) GetLicenseTypeOk() (*string, bool) {
	if o == nil || isNil(o.LicenseType) {
    return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *InlineResponse200286LicenseTypes) HasLicenseType() bool {
	if o != nil && !isNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *InlineResponse200286LicenseTypes) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200286LicenseTypes) GetCounts() InlineResponse200286Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200286Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200286LicenseTypes) GetCountsOk() (*InlineResponse200286Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200286LicenseTypes) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200286Counts and assigns it to the Counts field.
func (o *InlineResponse200286LicenseTypes) SetCounts(v InlineResponse200286Counts) {
	o.Counts = &v
}

func (o InlineResponse200286LicenseTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200286LicenseTypes struct {
	value *InlineResponse200286LicenseTypes
	isSet bool
}

func (v NullableInlineResponse200286LicenseTypes) Get() *InlineResponse200286LicenseTypes {
	return v.value
}

func (v *NullableInlineResponse200286LicenseTypes) Set(val *InlineResponse200286LicenseTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200286LicenseTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200286LicenseTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200286LicenseTypes(val *InlineResponse200286LicenseTypes) *NullableInlineResponse200286LicenseTypes {
	return &NullableInlineResponse200286LicenseTypes{value: val, isSet: true}
}

func (v NullableInlineResponse200286LicenseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200286LicenseTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


