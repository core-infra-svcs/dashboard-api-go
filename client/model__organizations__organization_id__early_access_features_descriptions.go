/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions Descriptions of the early access feature
type OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions struct {
	// Short description
	Short *string `json:"short,omitempty"`
	// Long description
	Long *string `json:"long,omitempty"`
}

// NewOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions instantiates a new OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions() *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions {
	this := OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions{}
	return &this
}

// NewOrganizationsOrganizationIdEarlyAccessFeaturesDescriptionsWithDefaults instantiates a new OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdEarlyAccessFeaturesDescriptionsWithDefaults() *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions {
	this := OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions{}
	return &this
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) GetShort() string {
	if o == nil || isNil(o.Short) {
		var ret string
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) GetShortOk() (*string, bool) {
	if o == nil || isNil(o.Short) {
    return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) HasShort() bool {
	if o != nil && !isNil(o.Short) {
		return true
	}

	return false
}

// SetShort gets a reference to the given string and assigns it to the Short field.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) SetShort(v string) {
	o.Short = &v
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) GetLong() string {
	if o == nil || isNil(o.Long) {
		var ret string
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) GetLongOk() (*string, bool) {
	if o == nil || isNil(o.Long) {
    return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) HasLong() bool {
	if o != nil && !isNil(o.Long) {
		return true
	}

	return false
}

// SetLong gets a reference to the given string and assigns it to the Long field.
func (o *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) SetLong(v string) {
	o.Long = &v
}

func (o OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Short) {
		toSerialize["short"] = o.Short
	}
	if !isNil(o.Long) {
		toSerialize["long"] = o.Long
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions struct {
	value *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions
	isSet bool
}

func (v NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) Get() *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) Set(val *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions(val *OrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) *NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions {
	return &NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdEarlyAccessFeaturesDescriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


