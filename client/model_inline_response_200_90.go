/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20090 struct for InlineResponse20090
type InlineResponse20090 struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds,omitempty"`
}

// NewInlineResponse20090 instantiates a new InlineResponse20090 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090() *InlineResponse20090 {
	this := InlineResponse20090{}
	return &this
}

// NewInlineResponse20090WithDefaults instantiates a new InlineResponse20090 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090WithDefaults() *InlineResponse20090 {
	this := InlineResponse20090{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *InlineResponse20090) GetDestOrganizationId() string {
	if o == nil || isNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.DestOrganizationId) {
    return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse20090) HasDestOrganizationId() bool {
	if o != nil && !isNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *InlineResponse20090) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseIds returns the LicenseIds field value if set, zero value otherwise.
func (o *InlineResponse20090) GetLicenseIds() []string {
	if o == nil || isNil(o.LicenseIds) {
		var ret []string
		return ret
	}
	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20090) GetLicenseIdsOk() ([]string, bool) {
	if o == nil || isNil(o.LicenseIds) {
    return nil, false
	}
	return o.LicenseIds, true
}

// HasLicenseIds returns a boolean if a field has been set.
func (o *InlineResponse20090) HasLicenseIds() bool {
	if o != nil && !isNil(o.LicenseIds) {
		return true
	}

	return false
}

// SetLicenseIds gets a reference to the given []string and assigns it to the LicenseIds field.
func (o *InlineResponse20090) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o InlineResponse20090) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !isNil(o.LicenseIds) {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090 struct {
	value *InlineResponse20090
	isSet bool
}

func (v NullableInlineResponse20090) Get() *InlineResponse20090 {
	return v.value
}

func (v *NullableInlineResponse20090) Set(val *InlineResponse20090) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090(val *InlineResponse20090) *NullableInlineResponse20090 {
	return &NullableInlineResponse20090{value: val, isSet: true}
}

func (v NullableInlineResponse20090) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

