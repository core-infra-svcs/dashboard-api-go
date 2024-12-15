/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject102 struct for InlineObject102
type InlineObject102 struct {
	// The list of auto locate jobs to be scheduled. Up to 100 jobs can be provided in a request.
	Jobs []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs `json:"jobs"`
}

// NewInlineObject102 instantiates a new InlineObject102 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject102(jobs []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) *InlineObject102 {
	this := InlineObject102{}
	this.Jobs = jobs
	return &this
}

// NewInlineObject102WithDefaults instantiates a new InlineObject102 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject102WithDefaults() *InlineObject102 {
	this := InlineObject102{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *InlineObject102) GetJobs() []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	if o == nil {
		var ret []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *InlineObject102) GetJobsOk() ([]NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs, bool) {
	if o == nil {
    return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *InlineObject102) SetJobs(v []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) {
	o.Jobs = v
}

func (o InlineObject102) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject102 struct {
	value *InlineObject102
	isSet bool
}

func (v NullableInlineObject102) Get() *InlineObject102 {
	return v.value
}

func (v *NullableInlineObject102) Set(val *InlineObject102) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject102) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject102) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject102(val *InlineObject102) *NullableInlineObject102 {
	return &NullableInlineObject102{value: val, isSet: true}
}

func (v NullableInlineObject102) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject102) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


