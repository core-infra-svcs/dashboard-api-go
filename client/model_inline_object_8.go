/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	// list of all reserved IP ranges for a single MG
	ReservedIpRanges *[]DevicesSerialCellularGatewayLanReservedIpRanges `json:"reservedIpRanges,omitempty"`
	// list of all fixed IP assignments for a single MG
	FixedIpAssignments *[]DevicesSerialCellularGatewayLanFixedIpAssignments `json:"fixedIpAssignments,omitempty"`
}

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineObject8) GetReservedIpRanges() []DevicesSerialCellularGatewayLanReservedIpRanges {
	if o == nil || o.ReservedIpRanges == nil {
		var ret []DevicesSerialCellularGatewayLanReservedIpRanges
		return ret
	}
	return *o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetReservedIpRangesOk() (*[]DevicesSerialCellularGatewayLanReservedIpRanges, bool) {
	if o == nil || o.ReservedIpRanges == nil {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineObject8) HasReservedIpRanges() bool {
	if o != nil && o.ReservedIpRanges != nil {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []DevicesSerialCellularGatewayLanReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineObject8) SetReservedIpRanges(v []DevicesSerialCellularGatewayLanReservedIpRanges) {
	o.ReservedIpRanges = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineObject8) GetFixedIpAssignments() []DevicesSerialCellularGatewayLanFixedIpAssignments {
	if o == nil || o.FixedIpAssignments == nil {
		var ret []DevicesSerialCellularGatewayLanFixedIpAssignments
		return ret
	}
	return *o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetFixedIpAssignmentsOk() (*[]DevicesSerialCellularGatewayLanFixedIpAssignments, bool) {
	if o == nil || o.FixedIpAssignments == nil {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineObject8) HasFixedIpAssignments() bool {
	if o != nil && o.FixedIpAssignments != nil {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []DevicesSerialCellularGatewayLanFixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineObject8) SetFixedIpAssignments(v []DevicesSerialCellularGatewayLanFixedIpAssignments) {
	o.FixedIpAssignments = &v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReservedIpRanges != nil {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if o.FixedIpAssignments != nil {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


