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

// OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower Power details object
type OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower struct {
	// The PoE power mode for the AP. Can be 'full' or 'low'
	Mode *string `json:"mode,omitempty"`
	Ac *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc `json:"ac,omitempty"`
	Poe *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe `json:"poe,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerWithDefaults() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower {
	this := OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) SetMode(v string) {
	o.Mode = &v
}

// GetAc returns the Ac field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetAc() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc {
	if o == nil || isNil(o.Ac) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc
		return ret
	}
	return *o.Ac
}

// GetAcOk returns a tuple with the Ac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetAcOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc, bool) {
	if o == nil || isNil(o.Ac) {
    return nil, false
	}
	return o.Ac, true
}

// HasAc returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) HasAc() bool {
	if o != nil && !isNil(o.Ac) {
		return true
	}

	return false
}

// SetAc gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc and assigns it to the Ac field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) SetAc(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerAc) {
	o.Ac = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetPoe() OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe {
	if o == nil || isNil(o.Poe) {
		var ret OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) GetPoeOk() (*OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe and assigns it to the Poe field.
func (o *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) SetPoe(v OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPowerPoe) {
	o.Poe = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Ac) {
		toSerialize["ac"] = o.Ac
	}
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower struct {
	value *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) Get() *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) Set(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower(val *OrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower {
	return &NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesEthernetStatusesPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


