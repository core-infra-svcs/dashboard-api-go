/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAssuranceAlertsScopeLldp Port of affected device
type OrganizationsOrganizationIdAssuranceAlertsScopeLldp struct {
	// Port of affect device
	Port *string `json:"port,omitempty"`
}

// NewOrganizationsOrganizationIdAssuranceAlertsScopeLldp instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeLldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAssuranceAlertsScopeLldp() *OrganizationsOrganizationIdAssuranceAlertsScopeLldp {
	this := OrganizationsOrganizationIdAssuranceAlertsScopeLldp{}
	return &this
}

// NewOrganizationsOrganizationIdAssuranceAlertsScopeLldpWithDefaults instantiates a new OrganizationsOrganizationIdAssuranceAlertsScopeLldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAssuranceAlertsScopeLldpWithDefaults() *OrganizationsOrganizationIdAssuranceAlertsScopeLldp {
	this := OrganizationsOrganizationIdAssuranceAlertsScopeLldp{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) SetPort(v string) {
	o.Port = &v
}

func (o OrganizationsOrganizationIdAssuranceAlertsScopeLldp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp struct {
	value *OrganizationsOrganizationIdAssuranceAlertsScopeLldp
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) Get() *OrganizationsOrganizationIdAssuranceAlertsScopeLldp {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) Set(val *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp(val *OrganizationsOrganizationIdAssuranceAlertsScopeLldp) *NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp {
	return &NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAssuranceAlertsScopeLldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


