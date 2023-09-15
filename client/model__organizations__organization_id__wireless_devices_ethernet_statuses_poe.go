/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe PoE details object for the port
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe struct {
	// The PoE Standard for the port. Can be '802.3at', '802.3af', '802.3bt', or null
	Standard *string `json:"standard,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoeWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoeWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe{}
	return &this
}

// GetStandard returns the Standard field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) GetStandard() string {
	if o == nil || isNil(o.Standard) {
		var ret string
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) GetStandardOk() (*string, bool) {
	if o == nil || isNil(o.Standard) {
    return nil, false
	}
	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) HasStandard() bool {
	if o != nil && !isNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given string and assigns it to the Standard field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) SetStandard(v string) {
	o.Standard = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Standard) {
		toSerialize["standard"] = o.Standard
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


