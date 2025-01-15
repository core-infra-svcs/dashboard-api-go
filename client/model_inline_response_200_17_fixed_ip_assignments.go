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

// InlineResponse20017FixedIpAssignments struct for InlineResponse20017FixedIpAssignments
type InlineResponse20017FixedIpAssignments struct {
	// A descriptive name of the assignment
	Name *string `json:"name,omitempty"`
	// The IP address you want to assign to a specific server or device
	Ip *string `json:"ip,omitempty"`
	// The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	Mac *string `json:"mac,omitempty"`
}

// NewInlineResponse20017FixedIpAssignments instantiates a new InlineResponse20017FixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017FixedIpAssignments() *InlineResponse20017FixedIpAssignments {
	this := InlineResponse20017FixedIpAssignments{}
	return &this
}

// NewInlineResponse20017FixedIpAssignmentsWithDefaults instantiates a new InlineResponse20017FixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017FixedIpAssignmentsWithDefaults() *InlineResponse20017FixedIpAssignments {
	this := InlineResponse20017FixedIpAssignments{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20017FixedIpAssignments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017FixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20017FixedIpAssignments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20017FixedIpAssignments) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20017FixedIpAssignments) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017FixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20017FixedIpAssignments) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20017FixedIpAssignments) SetIp(v string) {
	o.Ip = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20017FixedIpAssignments) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20017FixedIpAssignments) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20017FixedIpAssignments) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20017FixedIpAssignments) SetMac(v string) {
	o.Mac = &v
}

func (o InlineResponse20017FixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017FixedIpAssignments struct {
	value *InlineResponse20017FixedIpAssignments
	isSet bool
}

func (v NullableInlineResponse20017FixedIpAssignments) Get() *InlineResponse20017FixedIpAssignments {
	return v.value
}

func (v *NullableInlineResponse20017FixedIpAssignments) Set(val *InlineResponse20017FixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017FixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017FixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017FixedIpAssignments(val *InlineResponse20017FixedIpAssignments) *NullableInlineResponse20017FixedIpAssignments {
	return &NullableInlineResponse20017FixedIpAssignments{value: val, isSet: true}
}

func (v NullableInlineResponse20017FixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017FixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


