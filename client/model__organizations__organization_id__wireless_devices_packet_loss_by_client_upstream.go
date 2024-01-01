/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream Packets sent from a client to an AP.
type OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream struct {
	// Total packets sent by a client to an AP.
	Total *int32 `json:"total,omitempty"`
	// Total packets sent by a client and did not reach the AP.
	Lost *int32 `json:"lost,omitempty"`
	// Percentage of lost packets.
	LossPercentage *float32 `json:"lossPercentage,omitempty"`
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream{}
	return &this
}

// NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstreamWithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstreamWithDefaults() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	this := OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) SetTotal(v int32) {
	o.Total = &v
}

// GetLost returns the Lost field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetLost() int32 {
	if o == nil || isNil(o.Lost) {
		var ret int32
		return ret
	}
	return *o.Lost
}

// GetLostOk returns a tuple with the Lost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetLostOk() (*int32, bool) {
	if o == nil || isNil(o.Lost) {
    return nil, false
	}
	return o.Lost, true
}

// HasLost returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) HasLost() bool {
	if o != nil && !isNil(o.Lost) {
		return true
	}

	return false
}

// SetLost gets a reference to the given int32 and assigns it to the Lost field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) SetLost(v int32) {
	o.Lost = &v
}

// GetLossPercentage returns the LossPercentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetLossPercentage() float32 {
	if o == nil || isNil(o.LossPercentage) {
		var ret float32
		return ret
	}
	return *o.LossPercentage
}

// GetLossPercentageOk returns a tuple with the LossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) GetLossPercentageOk() (*float32, bool) {
	if o == nil || isNil(o.LossPercentage) {
    return nil, false
	}
	return o.LossPercentage, true
}

// HasLossPercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) HasLossPercentage() bool {
	if o != nil && !isNil(o.LossPercentage) {
		return true
	}

	return false
}

// SetLossPercentage gets a reference to the given float32 and assigns it to the LossPercentage field.
func (o *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) SetLossPercentage(v float32) {
	o.LossPercentage = &v
}

func (o OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Lost) {
		toSerialize["lost"] = o.Lost
	}
	if !isNil(o.LossPercentage) {
		toSerialize["lossPercentage"] = o.LossPercentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream struct {
	value *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream
	isSet bool
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) Get() *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) Set(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream(val *OrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream {
	return &NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdWirelessDevicesPacketLossByClientUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


