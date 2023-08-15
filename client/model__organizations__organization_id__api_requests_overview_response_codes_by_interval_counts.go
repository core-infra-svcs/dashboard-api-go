/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts struct for OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts
type OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts struct {
	// Response status code of the API response
	Code *int32 `json:"code,omitempty"`
	// Number of records that match the status code
	Count *int32 `json:"count,omitempty"`
}

// NewOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts instantiates a new OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts() *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts {
	this := OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts{}
	return &this
}

// NewOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCountsWithDefaults instantiates a new OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCountsWithDefaults() *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts {
	this := OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) GetCode() int32 {
	if o == nil || isNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) GetCodeOk() (*int32, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) SetCode(v int32) {
	o.Code = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) SetCount(v int32) {
	o.Count = &v
}

func (o OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts struct {
	value *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) Get() *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) Set(val *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts(val *OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) *NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts {
	return &NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


