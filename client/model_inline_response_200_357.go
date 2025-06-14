/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200357 struct for InlineResponse200357
type InlineResponse200357 struct {
	// List of Certificate Authorities
	Items []OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta `json:"meta,omitempty"`
}

// NewInlineResponse200357 instantiates a new InlineResponse200357 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200357() *InlineResponse200357 {
	this := InlineResponse200357{}
	return &this
}

// NewInlineResponse200357WithDefaults instantiates a new InlineResponse200357 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200357WithDefaults() *InlineResponse200357 {
	this := InlineResponse200357{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200357) GetItems() []OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200357) GetItemsOk() ([]OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200357) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems and assigns it to the Items field.
func (o *InlineResponse200357) SetItems(v []OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200357) GetMeta() OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200357) GetMetaOk() (*OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200357) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta and assigns it to the Meta field.
func (o *InlineResponse200357) SetMeta(v OrganizationsOrganizationIdWirelessDevicesRadsecCertificatesAuthoritiesMeta) {
	o.Meta = &v
}

func (o InlineResponse200357) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200357 struct {
	value *InlineResponse200357
	isSet bool
}

func (v NullableInlineResponse200357) Get() *InlineResponse200357 {
	return v.value
}

func (v *NullableInlineResponse200357) Set(val *InlineResponse200357) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200357) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200357) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200357(val *InlineResponse200357) *NullableInlineResponse200357 {
	return &NullableInlineResponse200357{value: val, isSet: true}
}

func (v NullableInlineResponse200357) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200357) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


