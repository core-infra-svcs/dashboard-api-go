/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20017 struct for InlineResponse20017
type InlineResponse20017 struct {
	// Name of the MG.
	DeviceName *string `json:"deviceName,omitempty"`
	// Lan IP of the MG
	DeviceLanIp *string `json:"deviceLanIp,omitempty"`
	// Subnet configuration of the MG.
	DeviceSubnet *string `json:"deviceSubnet,omitempty"`
	// list of all fixed IP assignments for a single MG
	FixedIpAssignments []InlineResponse20017FixedIpAssignments `json:"fixedIpAssignments,omitempty"`
	// list of all reserved IP ranges for a single MG
	ReservedIpRanges []InlineResponse20017ReservedIpRanges `json:"reservedIpRanges,omitempty"`
}

// NewInlineResponse20017 instantiates a new InlineResponse20017 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017() *InlineResponse20017 {
	this := InlineResponse20017{}
	return &this
}

// NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017WithDefaults() *InlineResponse20017 {
	this := InlineResponse20017{}
	return &this
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *InlineResponse20017) GetDeviceName() string {
	if o == nil || isNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.DeviceName) {
    return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *InlineResponse20017) HasDeviceName() bool {
	if o != nil && !isNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *InlineResponse20017) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceLanIp returns the DeviceLanIp field value if set, zero value otherwise.
func (o *InlineResponse20017) GetDeviceLanIp() string {
	if o == nil || isNil(o.DeviceLanIp) {
		var ret string
		return ret
	}
	return *o.DeviceLanIp
}

// GetDeviceLanIpOk returns a tuple with the DeviceLanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetDeviceLanIpOk() (*string, bool) {
	if o == nil || isNil(o.DeviceLanIp) {
    return nil, false
	}
	return o.DeviceLanIp, true
}

// HasDeviceLanIp returns a boolean if a field has been set.
func (o *InlineResponse20017) HasDeviceLanIp() bool {
	if o != nil && !isNil(o.DeviceLanIp) {
		return true
	}

	return false
}

// SetDeviceLanIp gets a reference to the given string and assigns it to the DeviceLanIp field.
func (o *InlineResponse20017) SetDeviceLanIp(v string) {
	o.DeviceLanIp = &v
}

// GetDeviceSubnet returns the DeviceSubnet field value if set, zero value otherwise.
func (o *InlineResponse20017) GetDeviceSubnet() string {
	if o == nil || isNil(o.DeviceSubnet) {
		var ret string
		return ret
	}
	return *o.DeviceSubnet
}

// GetDeviceSubnetOk returns a tuple with the DeviceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetDeviceSubnetOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSubnet) {
    return nil, false
	}
	return o.DeviceSubnet, true
}

// HasDeviceSubnet returns a boolean if a field has been set.
func (o *InlineResponse20017) HasDeviceSubnet() bool {
	if o != nil && !isNil(o.DeviceSubnet) {
		return true
	}

	return false
}

// SetDeviceSubnet gets a reference to the given string and assigns it to the DeviceSubnet field.
func (o *InlineResponse20017) SetDeviceSubnet(v string) {
	o.DeviceSubnet = &v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *InlineResponse20017) GetFixedIpAssignments() []InlineResponse20017FixedIpAssignments {
	if o == nil || isNil(o.FixedIpAssignments) {
		var ret []InlineResponse20017FixedIpAssignments
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetFixedIpAssignmentsOk() ([]InlineResponse20017FixedIpAssignments, bool) {
	if o == nil || isNil(o.FixedIpAssignments) {
    return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *InlineResponse20017) HasFixedIpAssignments() bool {
	if o != nil && !isNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []InlineResponse20017FixedIpAssignments and assigns it to the FixedIpAssignments field.
func (o *InlineResponse20017) SetFixedIpAssignments(v []InlineResponse20017FixedIpAssignments) {
	o.FixedIpAssignments = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *InlineResponse20017) GetReservedIpRanges() []InlineResponse20017ReservedIpRanges {
	if o == nil || isNil(o.ReservedIpRanges) {
		var ret []InlineResponse20017ReservedIpRanges
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017) GetReservedIpRangesOk() ([]InlineResponse20017ReservedIpRanges, bool) {
	if o == nil || isNil(o.ReservedIpRanges) {
    return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *InlineResponse20017) HasReservedIpRanges() bool {
	if o != nil && !isNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []InlineResponse20017ReservedIpRanges and assigns it to the ReservedIpRanges field.
func (o *InlineResponse20017) SetReservedIpRanges(v []InlineResponse20017ReservedIpRanges) {
	o.ReservedIpRanges = v
}

func (o InlineResponse20017) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !isNil(o.DeviceLanIp) {
		toSerialize["deviceLanIp"] = o.DeviceLanIp
	}
	if !isNil(o.DeviceSubnet) {
		toSerialize["deviceSubnet"] = o.DeviceSubnet
	}
	if !isNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	if !isNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017 struct {
	value *InlineResponse20017
	isSet bool
}

func (v NullableInlineResponse20017) Get() *InlineResponse20017 {
	return v.value
}

func (v *NullableInlineResponse20017) Set(val *InlineResponse20017) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017(val *InlineResponse20017) *NullableInlineResponse20017 {
	return &NullableInlineResponse20017{value: val, isSet: true}
}

func (v NullableInlineResponse20017) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


