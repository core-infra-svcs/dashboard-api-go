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

// OrganizationsLicensing Licensing related settings
type OrganizationsLicensing struct {
	// Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
	Model *string `json:"model,omitempty"`
}

// NewOrganizationsLicensing instantiates a new OrganizationsLicensing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsLicensing() *OrganizationsLicensing {
	this := OrganizationsLicensing{}
	return &this
}

// NewOrganizationsLicensingWithDefaults instantiates a new OrganizationsLicensing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsLicensingWithDefaults() *OrganizationsLicensing {
	this := OrganizationsLicensing{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *OrganizationsLicensing) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsLicensing) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *OrganizationsLicensing) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *OrganizationsLicensing) SetModel(v string) {
	o.Model = &v
}

func (o OrganizationsLicensing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsLicensing struct {
	value *OrganizationsLicensing
	isSet bool
}

func (v NullableOrganizationsLicensing) Get() *OrganizationsLicensing {
	return v.value
}

func (v *NullableOrganizationsLicensing) Set(val *OrganizationsLicensing) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsLicensing) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsLicensing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsLicensing(val *OrganizationsLicensing) *NullableOrganizationsLicensing {
	return &NullableOrganizationsLicensing{value: val, isSet: true}
}

func (v NullableOrganizationsLicensing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsLicensing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


