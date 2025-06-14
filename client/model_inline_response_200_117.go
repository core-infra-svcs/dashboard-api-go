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

// InlineResponse200117 struct for InlineResponse200117
type InlineResponse200117 struct {
	// Start of the timespan over which sensor alerts are counted
	StartTs *time.Time `json:"startTs,omitempty"`
	// End of the timespan over which sensor alerts are counted
	EndTs *time.Time `json:"endTs,omitempty"`
	Counts *NetworksNetworkIdSensorAlertsOverviewByMetricCounts `json:"counts,omitempty"`
}

// NewInlineResponse200117 instantiates a new InlineResponse200117 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200117() *InlineResponse200117 {
	this := InlineResponse200117{}
	return &this
}

// NewInlineResponse200117WithDefaults instantiates a new InlineResponse200117 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200117WithDefaults() *InlineResponse200117 {
	this := InlineResponse200117{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200117) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200117) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200117) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200117) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200117) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200117) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200117) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200117) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200117) GetCounts() NetworksNetworkIdSensorAlertsOverviewByMetricCounts {
	if o == nil || isNil(o.Counts) {
		var ret NetworksNetworkIdSensorAlertsOverviewByMetricCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200117) GetCountsOk() (*NetworksNetworkIdSensorAlertsOverviewByMetricCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200117) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given NetworksNetworkIdSensorAlertsOverviewByMetricCounts and assigns it to the Counts field.
func (o *InlineResponse200117) SetCounts(v NetworksNetworkIdSensorAlertsOverviewByMetricCounts) {
	o.Counts = &v
}

func (o InlineResponse200117) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200117 struct {
	value *InlineResponse200117
	isSet bool
}

func (v NullableInlineResponse200117) Get() *InlineResponse200117 {
	return v.value
}

func (v *NullableInlineResponse200117) Set(val *InlineResponse200117) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200117) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200117) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200117(val *InlineResponse200117) *NullableInlineResponse200117 {
	return &NullableInlineResponse200117{value: val, isSet: true}
}

func (v NullableInlineResponse200117) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200117) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


