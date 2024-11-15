/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta Meta details about the result
type OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta struct {
	Counts *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts `json:"counts,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta {
	this := OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) GetCounts() OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts {
	if o == nil || isNil(o.Counts) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) GetCountsOk() (*OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts and assigns it to the Counts field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) SetCounts(v OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMetaCounts) {
	o.Counts = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) Get() *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta(val *OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


