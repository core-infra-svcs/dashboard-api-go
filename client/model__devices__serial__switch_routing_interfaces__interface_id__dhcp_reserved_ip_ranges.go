/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges struct for DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges
type DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges struct {
	// The starting IP address of the reserved IP range
	Start string `json:"start"`
	// The ending IP address of the reserved IP range
	End string `json:"end"`
	// The comment for the reserved IP range
	Comment *string `json:"comment,omitempty"`
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges(start string, end string) *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges{}
	this.Start = start
	this.End = end
	return &this
}

// NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRangesWithDefaults instantiates a new DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRangesWithDefaults() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges {
	this := DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) SetEnd(v string) {
	o.End = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) SetComment(v string) {
	o.Comment = &v
}

func (o DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges struct {
	value *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges
	isSet bool
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) Get() *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges {
	return v.value
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) Set(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges(val *DevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges {
	return &NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges{value: val, isSet: true}
}

func (v NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialSwitchRoutingInterfacesInterfaceIdDhcpReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


