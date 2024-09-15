/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200227 struct for InlineResponse200227
type InlineResponse200227 struct {
	// List of eSIM Devices
	Items []OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta `json:"meta,omitempty"`
}

// NewInlineResponse200227 instantiates a new InlineResponse200227 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200227() *InlineResponse200227 {
	this := InlineResponse200227{}
	return &this
}

// NewInlineResponse200227WithDefaults instantiates a new InlineResponse200227 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200227WithDefaults() *InlineResponse200227 {
	this := InlineResponse200227{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200227) GetItems() []OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227) GetItemsOk() ([]OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200227) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems and assigns it to the Items field.
func (o *InlineResponse200227) SetItems(v []OrganizationsOrganizationIdCellularGatewayEsimsInventoryItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200227) GetMeta() OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227) GetMetaOk() (*OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200227) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta and assigns it to the Meta field.
func (o *InlineResponse200227) SetMeta(v OrganizationsOrganizationIdCellularGatewayEsimsInventoryMeta) {
	o.Meta = &v
}

func (o InlineResponse200227) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200227 struct {
	value *InlineResponse200227
	isSet bool
}

func (v NullableInlineResponse200227) Get() *InlineResponse200227 {
	return v.value
}

func (v *NullableInlineResponse200227) Set(val *InlineResponse200227) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200227) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200227) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200227(val *InlineResponse200227) *NullableInlineResponse200227 {
	return &NullableInlineResponse200227{value: val, isSet: true}
}

func (v NullableInlineResponse200227) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200227) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


