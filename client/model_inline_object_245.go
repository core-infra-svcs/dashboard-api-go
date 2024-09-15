/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject245 struct for InlineObject245
type InlineObject245 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId string `json:"destOrganizationId"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds"`
}

// NewInlineObject245 instantiates a new InlineObject245 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject245(destOrganizationId string, licenseIds []string) *InlineObject245 {
	this := InlineObject245{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseIds = licenseIds
	return &this
}

// NewInlineObject245WithDefaults instantiates a new InlineObject245 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject245WithDefaults() *InlineObject245 {
	this := InlineObject245{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *InlineObject245) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *InlineObject245) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *InlineObject245) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseIds returns the LicenseIds field value
func (o *InlineObject245) GetLicenseIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject245) GetLicenseIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LicenseIds, true
}

// SetLicenseIds sets field value
func (o *InlineObject245) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineObject245) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject245 struct {
	value *InlineObject245
	isSet bool
}

func (v NullableInlineObject245) Get() *InlineObject245 {
	return v.value
}

func (v *NullableInlineObject245) Set(val *InlineObject245) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject245) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject245) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject245(val *InlineObject245) *NullableInlineObject245 {
	return &NullableInlineObject245{value: val, isSet: true}
}

func (v NullableInlineObject245) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject245) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


