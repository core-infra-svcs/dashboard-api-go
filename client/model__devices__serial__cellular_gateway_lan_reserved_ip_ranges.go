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

// DevicesSerialCellularGatewayLanReservedIpRanges struct for DevicesSerialCellularGatewayLanReservedIpRanges
type DevicesSerialCellularGatewayLanReservedIpRanges struct {
	// Starting IP included in the reserved range of IPs
	Start string `json:"start"`
	// Ending IP included in the reserved range of IPs
	End string `json:"end"`
	// Comment explaining the reserved IP range
	Comment string `json:"comment"`
}

// NewDevicesSerialCellularGatewayLanReservedIpRanges instantiates a new DevicesSerialCellularGatewayLanReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialCellularGatewayLanReservedIpRanges(start string, end string, comment string) *DevicesSerialCellularGatewayLanReservedIpRanges {
	this := DevicesSerialCellularGatewayLanReservedIpRanges{}
	this.Start = start
	this.End = end
	this.Comment = comment
	return &this
}

// NewDevicesSerialCellularGatewayLanReservedIpRangesWithDefaults instantiates a new DevicesSerialCellularGatewayLanReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialCellularGatewayLanReservedIpRangesWithDefaults() *DevicesSerialCellularGatewayLanReservedIpRanges {
	this := DevicesSerialCellularGatewayLanReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *DevicesSerialCellularGatewayLanReservedIpRanges) SetComment(v string) {
	o.Comment = v
}

func (o DevicesSerialCellularGatewayLanReservedIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialCellularGatewayLanReservedIpRanges struct {
	value *DevicesSerialCellularGatewayLanReservedIpRanges
	isSet bool
}

func (v NullableDevicesSerialCellularGatewayLanReservedIpRanges) Get() *DevicesSerialCellularGatewayLanReservedIpRanges {
	return v.value
}

func (v *NullableDevicesSerialCellularGatewayLanReservedIpRanges) Set(val *DevicesSerialCellularGatewayLanReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialCellularGatewayLanReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialCellularGatewayLanReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialCellularGatewayLanReservedIpRanges(val *DevicesSerialCellularGatewayLanReservedIpRanges) *NullableDevicesSerialCellularGatewayLanReservedIpRanges {
	return &NullableDevicesSerialCellularGatewayLanReservedIpRanges{value: val, isSet: true}
}

func (v NullableDevicesSerialCellularGatewayLanReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialCellularGatewayLanReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


