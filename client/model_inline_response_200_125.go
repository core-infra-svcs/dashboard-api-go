/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200125 struct for InlineResponse200125
type InlineResponse200125 struct {
	// The start time of the access period
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the access period
	EndTs *time.Time `json:"endTs,omitempty"`
	// list of response codes and a count of how many requests had that code in the given time period
	Counts []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts `json:"counts,omitempty"`
}

// NewInlineResponse200125 instantiates a new InlineResponse200125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200125() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// NewInlineResponse200125WithDefaults instantiates a new InlineResponse200125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200125WithDefaults() *InlineResponse200125 {
	this := InlineResponse200125{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200125) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200125) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200125) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse200125) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200125) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200125) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200125) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse200125) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200125) GetCounts() []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts {
	if o == nil || isNil(o.Counts) {
		var ret []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200125) GetCountsOk() ([]OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200125) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts and assigns it to the Counts field.
func (o *InlineResponse200125) SetCounts(v []OrganizationsOrganizationIdApiRequestsOverviewResponseCodesByIntervalCounts) {
	o.Counts = v
}

func (o InlineResponse200125) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200125 struct {
	value *InlineResponse200125
	isSet bool
}

func (v NullableInlineResponse200125) Get() *InlineResponse200125 {
	return v.value
}

func (v *NullableInlineResponse200125) Set(val *InlineResponse200125) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200125) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200125(val *InlineResponse200125) *NullableInlineResponse200125 {
	return &NullableInlineResponse200125{value: val, isSet: true}
}

func (v NullableInlineResponse200125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


