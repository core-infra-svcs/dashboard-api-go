/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200105 struct for InlineResponse200105
type InlineResponse200105 struct {
	// Start of the timespan over which sensor alerts are counted
	StartTs *time.Time `json:"startTs,omitempty"`
	// End of the timespan over which sensor alerts are counted
	EndTs *time.Time `json:"endTs,omitempty"`
	Counts *NetworksNetworkIdSensorAlertsOverviewByMetricCounts `json:"counts,omitempty"`
}

// NewInlineResponse200105 instantiates a new InlineResponse200105 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200105() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// NewInlineResponse200105WithDefaults instantiates a new InlineResponse200105 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200105WithDefaults() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200105) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200105) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200105) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200105) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200105) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200105) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200105) GetCounts() NetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	if o == nil || isNil(o.Counts) {
		var ret NetworksNetworkIdSensorAlertsOverviewByMetricCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetCountsOk() (*NetworksNetworkIdSensorAlertsOverviewByMetricCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200105) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given NetworksNetworkIdSensorAlertsOverviewByMetricCounts and assigns it to the Counts field.
func (o *InlineResponse200105) SetCounts(v NetworksNetworkIdSensorAlertsOverviewByMetricCounts) {
	o.Counts = &v
}

func (o InlineResponse200105) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200105 struct {
	value *InlineResponse200105
	isSet bool
}

func (v NullableInlineResponse200105) Get() *InlineResponse200105 {
	return v.value
}

func (v *NullableInlineResponse200105) Set(val *InlineResponse200105) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200105) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200105) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200105(val *InlineResponse200105) *NullableInlineResponse200105 {
	return &NullableInlineResponse200105{value: val, isSet: true}
}

func (v NullableInlineResponse200105) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200105) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


