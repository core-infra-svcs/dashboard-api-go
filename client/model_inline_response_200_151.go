/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200151 struct for InlineResponse200151
type InlineResponse200151 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds,omitempty"`
}

// NewInlineResponse200151 instantiates a new InlineResponse200151 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200151() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// NewInlineResponse200151WithDefaults instantiates a new InlineResponse200151 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200151WithDefaults() *InlineResponse200151 {
	this := InlineResponse200151{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200151) GetDestOrganizationId() string {
	if o == nil || isNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.DestOrganizationId) {
    return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200151) HasDestOrganizationId() bool {
	if o != nil && !isNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *InlineResponse200151) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseIds returns the LicenseIds field value if set, zero value otherwise.
func (o *InlineResponse200151) GetLicenseIds() []string {
	if o == nil || isNil(o.LicenseIds) {
		var ret []string
		return ret
	}
	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200151) GetLicenseIdsOk() ([]string, bool) {
	if o == nil || isNil(o.LicenseIds) {
    return nil, false
	}
	return o.LicenseIds, true
}

// HasLicenseIds returns a boolean if a field has been set.
func (o *InlineResponse200151) HasLicenseIds() bool {
	if o != nil && !isNil(o.LicenseIds) {
		return true
	}

	return false
}

// SetLicenseIds gets a reference to the given []string and assigns it to the LicenseIds field.
func (o *InlineResponse200151) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineResponse200151) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !isNil(o.LicenseIds) {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200151 struct {
	value *InlineResponse200151
	isSet bool
}

func (v NullableInlineResponse200151) Get() *InlineResponse200151 {
	return v.value
}

func (v *NullableInlineResponse200151) Set(val *InlineResponse200151) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200151) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200151) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200151(val *InlineResponse200151) *NullableInlineResponse200151 {
	return &NullableInlineResponse200151{value: val, isSet: true}
}

func (v NullableInlineResponse200151) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200151) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


