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

// OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries struct for OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries
type OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries struct {
	// Timestamp for this data point
	Ts *time.Time `json:"ts,omitempty"`
	// Loss percentage
	LossPercent *float32 `json:"lossPercent,omitempty"`
	// Latency in milliseconds
	LatencyMs *float32 `json:"latencyMs,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries instantiates a new OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries() *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	this := OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeriesWithDefaults instantiates a new OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeriesWithDefaults() *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	this := OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) SetTs(v time.Time) {
	o.Ts = &v
}

// GetLossPercent returns the LossPercent field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetLossPercent() float32 {
	if o == nil || isNil(o.LossPercent) {
		var ret float32
		return ret
	}
	return *o.LossPercent
}

// GetLossPercentOk returns a tuple with the LossPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetLossPercentOk() (*float32, bool) {
	if o == nil || isNil(o.LossPercent) {
    return nil, false
	}
	return o.LossPercent, true
}

// HasLossPercent returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) HasLossPercent() bool {
	if o != nil && !isNil(o.LossPercent) {
		return true
	}

	return false
}

// SetLossPercent gets a reference to the given float32 and assigns it to the LossPercent field.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) SetLossPercent(v float32) {
	o.LossPercent = &v
}

// GetLatencyMs returns the LatencyMs field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetLatencyMs() float32 {
	if o == nil || isNil(o.LatencyMs) {
		var ret float32
		return ret
	}
	return *o.LatencyMs
}

// GetLatencyMsOk returns a tuple with the LatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) GetLatencyMsOk() (*float32, bool) {
	if o == nil || isNil(o.LatencyMs) {
    return nil, false
	}
	return o.LatencyMs, true
}

// HasLatencyMs returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) HasLatencyMs() bool {
	if o != nil && !isNil(o.LatencyMs) {
		return true
	}

	return false
}

// SetLatencyMs gets a reference to the given float32 and assigns it to the LatencyMs field.
func (o *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) SetLatencyMs(v float32) {
	o.LatencyMs = &v
}

func (o OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.LossPercent) {
		toSerialize["lossPercent"] = o.LossPercent
	}
	if !isNil(o.LatencyMs) {
		toSerialize["latencyMs"] = o.LatencyMs
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries struct {
	value *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) Get() *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) Set(val *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries(val *OrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) *NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries {
	return &NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksLossAndLatencyTimeSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


