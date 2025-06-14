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

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts struct for OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts struct {
	// Label of the port
	Name *string `json:"name,omitempty"`
	Poe *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe `json:"poe,omitempty"`
	LinkNegotiation *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation `json:"linkNegotiation,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPortsWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPortsWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) SetName(v string) {
	o.Name = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetPoe() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe {
	if o == nil || isNil(o.Poe) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetPoeOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe and assigns it to the Poe field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) SetPoe(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) {
	o.Poe = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetLinkNegotiation() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation {
	if o == nil || isNil(o.LinkNegotiation) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) GetLinkNegotiationOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation, bool) {
	if o == nil || isNil(o.LinkNegotiation) {
    return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) HasLinkNegotiation() bool {
	if o != nil && !isNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation and assigns it to the LinkNegotiation field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) SetLinkNegotiation(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesLinkNegotiation) {
	o.LinkNegotiation = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	if !isNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


