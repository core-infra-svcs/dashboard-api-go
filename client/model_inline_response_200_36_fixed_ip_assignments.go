/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20036FixedIpAssignments struct for InlineResponse20036FixedIpAssignments
type InlineResponse20036FixedIpAssignments struct {
	// The name of the client which has fixed IP address
	Name *string `json:"name,omitempty"`
	// The MAC address of the client which has fixed IP address
	Mac *string `json:"mac,omitempty"`
	// The IP address of the client which has fixed IP address assigned to it
	Ip *string `json:"ip,omitempty"`
}

// NewInlineResponse20036FixedIpAssignments instantiates a new InlineResponse20036FixedIpAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20036FixedIpAssignments() *InlineResponse20036FixedIpAssignments {
	this := InlineResponse20036FixedIpAssignments{}
	return &this
}

// NewInlineResponse20036FixedIpAssignmentsWithDefaults instantiates a new InlineResponse20036FixedIpAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20036FixedIpAssignmentsWithDefaults() *InlineResponse20036FixedIpAssignments {
	this := InlineResponse20036FixedIpAssignments{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20036FixedIpAssignments) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036FixedIpAssignments) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20036FixedIpAssignments) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20036FixedIpAssignments) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20036FixedIpAssignments) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036FixedIpAssignments) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20036FixedIpAssignments) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20036FixedIpAssignments) SetMac(v string) {
	o.Mac = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20036FixedIpAssignments) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20036FixedIpAssignments) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20036FixedIpAssignments) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20036FixedIpAssignments) SetIp(v string) {
	o.Ip = &v
}

func (o InlineResponse20036FixedIpAssignments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20036FixedIpAssignments struct {
	value *InlineResponse20036FixedIpAssignments
	isSet bool
}

func (v NullableInlineResponse20036FixedIpAssignments) Get() *InlineResponse20036FixedIpAssignments {
	return v.value
}

func (v *NullableInlineResponse20036FixedIpAssignments) Set(val *InlineResponse20036FixedIpAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20036FixedIpAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20036FixedIpAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20036FixedIpAssignments(val *InlineResponse20036FixedIpAssignments) *NullableInlineResponse20036FixedIpAssignments {
	return &NullableInlineResponse20036FixedIpAssignments{value: val, isSet: true}
}

func (v NullableInlineResponse20036FixedIpAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20036FixedIpAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


