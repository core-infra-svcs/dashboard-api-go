/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20022Entries struct for InlineResponse20022Entries
type InlineResponse20022Entries struct {
	// The IP address of the ARP table entry
	Ip *string `json:"ip,omitempty"`
	// The MAC address of the ARP table entry
	Mac *string `json:"mac,omitempty"`
	// The VLAN ID of the ARP table entry
	VlanId *int32 `json:"vlanId,omitempty"`
	// Time of the last update of the ARP table entry
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
}

// NewInlineResponse20022Entries instantiates a new InlineResponse20022Entries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022Entries() *InlineResponse20022Entries {
	this := InlineResponse20022Entries{}
	return &this
}

// NewInlineResponse20022EntriesWithDefaults instantiates a new InlineResponse20022Entries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022EntriesWithDefaults() *InlineResponse20022Entries {
	this := InlineResponse20022Entries{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20022Entries) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022Entries) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20022Entries) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20022Entries) SetIp(v string) {
	o.Ip = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20022Entries) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022Entries) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20022Entries) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20022Entries) SetMac(v string) {
	o.Mac = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse20022Entries) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022Entries) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse20022Entries) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineResponse20022Entries) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse20022Entries) GetLastUpdatedAt() time.Time {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022Entries) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse20022Entries) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *InlineResponse20022Entries) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

func (o InlineResponse20022Entries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022Entries struct {
	value *InlineResponse20022Entries
	isSet bool
}

func (v NullableInlineResponse20022Entries) Get() *InlineResponse20022Entries {
	return v.value
}

func (v *NullableInlineResponse20022Entries) Set(val *InlineResponse20022Entries) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022Entries) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022Entries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022Entries(val *InlineResponse20022Entries) *NullableInlineResponse20022Entries {
	return &NullableInlineResponse20022Entries{value: val, isSet: true}
}

func (v NullableInlineResponse20022Entries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022Entries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


