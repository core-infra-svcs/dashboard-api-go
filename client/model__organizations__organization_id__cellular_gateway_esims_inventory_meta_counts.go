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

// OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts Counts of involved entities
type OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts struct {
	Items *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems `json:"items,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsWithDefaults() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts {
	this := OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) GetItems() OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems {
	if o == nil || isNil(o.Items) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) GetItemsOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems and assigns it to the Items field.
func (o *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) SetItems(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCountsItems) {
	o.Items = &v
}

func (o OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts struct {
	value *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) Get() *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) Set(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts(val *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts {
	return &NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayEsimsInventoryMetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


