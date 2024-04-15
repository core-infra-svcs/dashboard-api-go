/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200167 struct for InlineResponse200167
type InlineResponse200167 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Total channel utilization
	UtilizationTotal *float32 `json:"utilizationTotal,omitempty"`
	// Average wifi utilization
	Utilization80211 *float32 `json:"utilization80211,omitempty"`
	// Average signal interference
	UtilizationNon80211 *float32 `json:"utilizationNon80211,omitempty"`
}

// NewInlineResponse200167 instantiates a new InlineResponse200167 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200167() *InlineResponse200167 {
	this := InlineResponse200167{}
	return &this
}

// NewInlineResponse200167WithDefaults instantiates a new InlineResponse200167 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200167WithDefaults() *InlineResponse200167 {
	this := InlineResponse200167{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200167) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200167) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200167) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200167) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200167) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200167) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetUtilizationTotal returns the UtilizationTotal field value if set, zero value otherwise.
func (o *InlineResponse200167) GetUtilizationTotal() float32 {
	if o == nil || isNil(o.UtilizationTotal) {
		var ret float32
		return ret
	}
	return *o.UtilizationTotal
}

// GetUtilizationTotalOk returns a tuple with the UtilizationTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167) GetUtilizationTotalOk() (*float32, bool) {
	if o == nil || isNil(o.UtilizationTotal) {
    return nil, false
	}
	return o.UtilizationTotal, true
}

// HasUtilizationTotal returns a boolean if a field has been set.
func (o *InlineResponse200167) HasUtilizationTotal() bool {
	if o != nil && !isNil(o.UtilizationTotal) {
		return true
	}

	return false
}

// SetUtilizationTotal gets a reference to the given float32 and assigns it to the UtilizationTotal field.
func (o *InlineResponse200167) SetUtilizationTotal(v float32) {
	o.UtilizationTotal = &v
}

// GetUtilization80211 returns the Utilization80211 field value if set, zero value otherwise.
func (o *InlineResponse200167) GetUtilization80211() float32 {
	if o == nil || isNil(o.Utilization80211) {
		var ret float32
		return ret
	}
	return *o.Utilization80211
}

// GetUtilization80211Ok returns a tuple with the Utilization80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167) GetUtilization80211Ok() (*float32, bool) {
	if o == nil || isNil(o.Utilization80211) {
    return nil, false
	}
	return o.Utilization80211, true
}

// HasUtilization80211 returns a boolean if a field has been set.
func (o *InlineResponse200167) HasUtilization80211() bool {
	if o != nil && !isNil(o.Utilization80211) {
		return true
	}

	return false
}

// SetUtilization80211 gets a reference to the given float32 and assigns it to the Utilization80211 field.
func (o *InlineResponse200167) SetUtilization80211(v float32) {
	o.Utilization80211 = &v
}

// GetUtilizationNon80211 returns the UtilizationNon80211 field value if set, zero value otherwise.
func (o *InlineResponse200167) GetUtilizationNon80211() float32 {
	if o == nil || isNil(o.UtilizationNon80211) {
		var ret float32
		return ret
	}
	return *o.UtilizationNon80211
}

// GetUtilizationNon80211Ok returns a tuple with the UtilizationNon80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167) GetUtilizationNon80211Ok() (*float32, bool) {
	if o == nil || isNil(o.UtilizationNon80211) {
    return nil, false
	}
	return o.UtilizationNon80211, true
}

// HasUtilizationNon80211 returns a boolean if a field has been set.
func (o *InlineResponse200167) HasUtilizationNon80211() bool {
	if o != nil && !isNil(o.UtilizationNon80211) {
		return true
	}

	return false
}

// SetUtilizationNon80211 gets a reference to the given float32 and assigns it to the UtilizationNon80211 field.
func (o *InlineResponse200167) SetUtilizationNon80211(v float32) {
	o.UtilizationNon80211 = &v
}

func (o InlineResponse200167) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.UtilizationTotal) {
		toSerialize["utilizationTotal"] = o.UtilizationTotal
	}
	if !isNil(o.Utilization80211) {
		toSerialize["utilization80211"] = o.Utilization80211
	}
	if !isNil(o.UtilizationNon80211) {
		toSerialize["utilizationNon80211"] = o.UtilizationNon80211
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200167 struct {
	value *InlineResponse200167
	isSet bool
}

func (v NullableInlineResponse200167) Get() *InlineResponse200167 {
	return v.value
}

func (v *NullableInlineResponse200167) Set(val *InlineResponse200167) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200167) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200167) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200167(val *InlineResponse200167) *NullableInlineResponse200167 {
	return &NullableInlineResponse200167{value: val, isSet: true}
}

func (v NullableInlineResponse200167) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200167) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


