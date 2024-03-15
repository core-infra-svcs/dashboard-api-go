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

// InlineObject231 struct for InlineObject231
type InlineObject231 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId string `json:"destOrganizationId"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds"`
}

// NewInlineObject231 instantiates a new InlineObject231 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject231(destOrganizationId string, licenseIds []string) *InlineObject231 {
	this := InlineObject231{}
	this.DestOrganizationId = destOrganizationId
	this.LicenseIds = licenseIds
	return &this
}

// NewInlineObject231WithDefaults instantiates a new InlineObject231 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject231WithDefaults() *InlineObject231 {
	this := InlineObject231{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value
func (o *InlineObject231) GetDestOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value
// and a boolean to check if the value has been set.
func (o *InlineObject231) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestOrganizationId, true
}

// SetDestOrganizationId sets field value
func (o *InlineObject231) SetDestOrganizationId(v string) {
	o.DestOrganizationId = v
}

// GetLicenseIds returns the LicenseIds field value
func (o *InlineObject231) GetLicenseIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value
// and a boolean to check if the value has been set.
func (o *InlineObject231) GetLicenseIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LicenseIds, true
}

// SetLicenseIds sets field value
func (o *InlineObject231) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineObject231) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if true {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject231 struct {
	value *InlineObject231
	isSet bool
}

func (v NullableInlineObject231) Get() *InlineObject231 {
	return v.value
}

func (v *NullableInlineObject231) Set(val *InlineObject231) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject231) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject231) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject231(val *InlineObject231) *NullableInlineObject231 {
	return &NullableInlineObject231{value: val, isSet: true}
}

func (v NullableInlineObject231) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject231) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


