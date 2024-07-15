/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat Tower Signal Status
type OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat struct {
	// Reference Signal Received Power
	Rsrp *string `json:"rsrp,omitempty"`
	// Reference Signal Received Quality
	Rsrq *string `json:"rsrq,omitempty"`
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat{}
	return &this
}

// NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStatWithDefaults instantiates a new OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStatWithDefaults() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	this := OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat{}
	return &this
}

// GetRsrp returns the Rsrp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) GetRsrp() string {
	if o == nil || isNil(o.Rsrp) {
		var ret string
		return ret
	}
	return *o.Rsrp
}

// GetRsrpOk returns a tuple with the Rsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) GetRsrpOk() (*string, bool) {
	if o == nil || isNil(o.Rsrp) {
    return nil, false
	}
	return o.Rsrp, true
}

// HasRsrp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) HasRsrp() bool {
	if o != nil && !isNil(o.Rsrp) {
		return true
	}

	return false
}

// SetRsrp gets a reference to the given string and assigns it to the Rsrp field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) SetRsrp(v string) {
	o.Rsrp = &v
}

// GetRsrq returns the Rsrq field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) GetRsrq() string {
	if o == nil || isNil(o.Rsrq) {
		var ret string
		return ret
	}
	return *o.Rsrq
}

// GetRsrqOk returns a tuple with the Rsrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) GetRsrqOk() (*string, bool) {
	if o == nil || isNil(o.Rsrq) {
    return nil, false
	}
	return o.Rsrq, true
}

// HasRsrq returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) HasRsrq() bool {
	if o != nil && !isNil(o.Rsrq) {
		return true
	}

	return false
}

// SetRsrq gets a reference to the given string and assigns it to the Rsrq field.
func (o *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) SetRsrq(v string) {
	o.Rsrq = &v
}

func (o OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rsrp) {
		toSerialize["rsrp"] = o.Rsrp
	}
	if !isNil(o.Rsrq) {
		toSerialize["rsrq"] = o.Rsrq
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat struct {
	value *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) Get() *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) Set(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat(val *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	return &NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


