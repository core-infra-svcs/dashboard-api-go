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

// InlineResponse200272 struct for InlineResponse200272
type InlineResponse200272 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds,omitempty"`
}

// NewInlineResponse200272 instantiates a new InlineResponse200272 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200272() *InlineResponse200272 {
	this := InlineResponse200272{}
	return &this
}

// NewInlineResponse200272WithDefaults instantiates a new InlineResponse200272 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200272WithDefaults() *InlineResponse200272 {
	this := InlineResponse200272{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200272) GetDestOrganizationId() string {
	if o == nil || isNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.DestOrganizationId) {
    return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200272) HasDestOrganizationId() bool {
	if o != nil && !isNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *InlineResponse200272) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseIds returns the LicenseIds field value if set, zero value otherwise.
func (o *InlineResponse200272) GetLicenseIds() []string {
	if o == nil || isNil(o.LicenseIds) {
		var ret []string
		return ret
	}
	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200272) GetLicenseIdsOk() ([]string, bool) {
	if o == nil || isNil(o.LicenseIds) {
    return nil, false
	}
	return o.LicenseIds, true
}

// HasLicenseIds returns a boolean if a field has been set.
func (o *InlineResponse200272) HasLicenseIds() bool {
	if o != nil && !isNil(o.LicenseIds) {
		return true
	}

	return false
}

// SetLicenseIds gets a reference to the given []string and assigns it to the LicenseIds field.
func (o *InlineResponse200272) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineResponse200272) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !isNil(o.LicenseIds) {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200272 struct {
	value *InlineResponse200272
	isSet bool
}

func (v NullableInlineResponse200272) Get() *InlineResponse200272 {
	return v.value
}

func (v *NullableInlineResponse200272) Set(val *InlineResponse200272) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200272) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200272) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200272(val *InlineResponse200272) *NullableInlineResponse200272 {
	return &NullableInlineResponse200272{value: val, isSet: true}
}

func (v NullableInlineResponse200272) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200272) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


