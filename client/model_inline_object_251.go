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

// InlineObject251 struct for InlineObject251
type InlineObject251 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId string `json:"destOrganizationId"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds"`
}

// NewInlineObject251 instantiates a new InlineObject251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject251(destOrganizationId string, licenseIds []string) *InlineObject251 {
	this := InlineObject251{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseIds = licenseIds
	return &this
}

// NewInlineObject251WithDefaults instantiates a new InlineObject251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject251WithDefaults() *InlineObject251 {
	this := InlineObject251{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *InlineObject251) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *InlineObject251) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseIds returns the LicenseIds field value
func (o *InlineObject251) GetLicenseIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject251) GetLicenseIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LicenseIds, true
}

// SetLicenseIds sets field value
func (o *InlineObject251) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineObject251) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject251 struct {
	value *InlineObject251
	isSet bool
}

func (v NullableInlineObject251) Get() *InlineObject251 {
	return v.value
}

func (v *NullableInlineObject251) Set(val *InlineObject251) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject251) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject251(val *InlineObject251) *NullableInlineObject251 {
	return &NullableInlineObject251{value: val, isSet: true}
}

func (v NullableInlineObject251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


