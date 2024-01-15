/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryClaimLicenses struct for OrganizationsOrganizationIdInventoryClaimLicenses
type OrganizationsOrganizationIdInventoryClaimLicenses struct {
	// The key of the license
	Key string `json:"key"`
	// Co-term licensing only: either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.
	Mode *string `json:"mode,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryClaimLicenses instantiates a new OrganizationsOrganizationIdInventoryClaimLicenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryClaimLicenses(key string) *OrganizationsOrganizationIdInventoryClaimLicenses {
	this := OrganizationsOrganizationIdInventoryClaimLicenses{}
	this.Key = key
	return &this
}

// NewOrganizationsOrganizationIdInventoryClaimLicensesWithDefaults instantiates a new OrganizationsOrganizationIdInventoryClaimLicenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryClaimLicensesWithDefaults() *OrganizationsOrganizationIdInventoryClaimLicenses {
	this := OrganizationsOrganizationIdInventoryClaimLicenses{}
	return &this
}

// GetKey returns the Key field value
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) SetKey(v string) {
	o.Key = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OrganizationsOrganizationIdInventoryClaimLicenses) SetMode(v string) {
	o.Mode = &v
}

func (o OrganizationsOrganizationIdInventoryClaimLicenses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryClaimLicenses struct {
	value *OrganizationsOrganizationIdInventoryClaimLicenses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryClaimLicenses) Get() *OrganizationsOrganizationIdInventoryClaimLicenses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryClaimLicenses) Set(val *OrganizationsOrganizationIdInventoryClaimLicenses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryClaimLicenses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryClaimLicenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryClaimLicenses(val *OrganizationsOrganizationIdInventoryClaimLicenses) *NullableOrganizationsOrganizationIdInventoryClaimLicenses {
	return &NullableOrganizationsOrganizationIdInventoryClaimLicenses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryClaimLicenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryClaimLicenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


