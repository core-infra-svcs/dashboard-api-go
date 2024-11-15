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

// InlineResponse200336 struct for InlineResponse200336
type InlineResponse200336 struct {
	// Wireless LAN controller HA failover events
	Items []OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems `json:"items,omitempty"`
	Meta *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta `json:"meta,omitempty"`
}

// NewInlineResponse200336 instantiates a new InlineResponse200336 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200336() *InlineResponse200336 {
	this := InlineResponse200336{}
	return &this
}

// NewInlineResponse200336WithDefaults instantiates a new InlineResponse200336 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200336WithDefaults() *InlineResponse200336 {
	this := InlineResponse200336{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InlineResponse200336) GetItems() []OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems {
	if o == nil || isNil(o.Items) {
		var ret []OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336) GetItemsOk() ([]OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InlineResponse200336) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems and assigns it to the Items field.
func (o *InlineResponse200336) SetItems(v []OrganizationsOrganizationIdWirelessControllerDevicesRedundancyFailoverHistoryItems) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse200336) GetMeta() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta {
	if o == nil || isNil(o.Meta) {
		var ret OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200336) GetMetaOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse200336) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta and assigns it to the Meta field.
func (o *InlineResponse200336) SetMeta(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesMeta) {
	o.Meta = &v
}

func (o InlineResponse200336) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200336 struct {
	value *InlineResponse200336
	isSet bool
}

func (v NullableInlineResponse200336) Get() *InlineResponse200336 {
	return v.value
}

func (v *NullableInlineResponse200336) Set(val *InlineResponse200336) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200336) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200336) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200336(val *InlineResponse200336) *NullableInlineResponse200336 {
	return &NullableInlineResponse200336{value: val, isSet: true}
}

func (v NullableInlineResponse200336) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200336) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

