/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdClaimLicenses struct for OrganizationsOrganizationIdClaimLicenses
type OrganizationsOrganizationIdClaimLicenses struct {
	// The key of the license
	Key string `json:"key"`
	// Either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. This parameter is legacy and does not apply to organizations with per-device licensing enabled.
	Mode *string `json:"mode,omitempty"`
}

// NewOrganizationsOrganizationIdClaimLicenses instantiates a new OrganizationsOrganizationIdClaimLicenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdClaimLicenses(key string) *OrganizationsOrganizationIdClaimLicenses {
	this := OrganizationsOrganizationIdClaimLicenses{}
	this.Key = key
	return &this
}

// NewOrganizationsOrganizationIdClaimLicensesWithDefaults instantiates a new OrganizationsOrganizationIdClaimLicenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdClaimLicensesWithDefaults() *OrganizationsOrganizationIdClaimLicenses {
	this := OrganizationsOrganizationIdClaimLicenses{}
	return &this
}

// GetKey returns the Key field value
func (o *OrganizationsOrganizationIdClaimLicenses) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdClaimLicenses) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OrganizationsOrganizationIdClaimLicenses) SetKey(v string) {
	o.Key = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdClaimLicenses) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdClaimLicenses) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdClaimLicenses) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OrganizationsOrganizationIdClaimLicenses) SetMode(v string) {
	o.Mode = &v
}

func (o OrganizationsOrganizationIdClaimLicenses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdClaimLicenses struct {
	value *OrganizationsOrganizationIdClaimLicenses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdClaimLicenses) Get() *OrganizationsOrganizationIdClaimLicenses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdClaimLicenses) Set(val *OrganizationsOrganizationIdClaimLicenses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdClaimLicenses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdClaimLicenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdClaimLicenses(val *OrganizationsOrganizationIdClaimLicenses) *NullableOrganizationsOrganizationIdClaimLicenses {
	return &NullableOrganizationsOrganizationIdClaimLicenses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdClaimLicenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdClaimLicenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


