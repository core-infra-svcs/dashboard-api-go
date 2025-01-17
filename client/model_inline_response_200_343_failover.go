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

// InlineResponse200343Failover Wireless LAN controller failover information
type InlineResponse200343Failover struct {
	Last *InlineResponse200343FailoverLast `json:"last,omitempty"`
	Counts *InlineResponse200343FailoverCounts `json:"counts,omitempty"`
}

// NewInlineResponse200343Failover instantiates a new InlineResponse200343Failover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200343Failover() *InlineResponse200343Failover {
	this := InlineResponse200343Failover{}
	return &this
}

// NewInlineResponse200343FailoverWithDefaults instantiates a new InlineResponse200343Failover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200343FailoverWithDefaults() *InlineResponse200343Failover {
	this := InlineResponse200343Failover{}
	return &this
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *InlineResponse200343Failover) GetLast() InlineResponse200343FailoverLast {
	if o == nil || isNil(o.Last) {
		var ret InlineResponse200343FailoverLast
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343Failover) GetLastOk() (*InlineResponse200343FailoverLast, bool) {
	if o == nil || isNil(o.Last) {
    return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *InlineResponse200343Failover) HasLast() bool {
	if o != nil && !isNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given InlineResponse200343FailoverLast and assigns it to the Last field.
func (o *InlineResponse200343Failover) SetLast(v InlineResponse200343FailoverLast) {
	o.Last = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200343Failover) GetCounts() InlineResponse200343FailoverCounts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200343FailoverCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200343Failover) GetCountsOk() (*InlineResponse200343FailoverCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200343Failover) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200343FailoverCounts and assigns it to the Counts field.
func (o *InlineResponse200343Failover) SetCounts(v InlineResponse200343FailoverCounts) {
	o.Counts = &v
}

func (o InlineResponse200343Failover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200343Failover struct {
	value *InlineResponse200343Failover
	isSet bool
}

func (v NullableInlineResponse200343Failover) Get() *InlineResponse200343Failover {
	return v.value
}

func (v *NullableInlineResponse200343Failover) Set(val *InlineResponse200343Failover) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200343Failover) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200343Failover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200343Failover(val *InlineResponse200343Failover) *NullableInlineResponse200343Failover {
	return &NullableInlineResponse200343Failover{value: val, isSet: true}
}

func (v NullableInlineResponse200343Failover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200343Failover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


