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

// InlineObject101 struct for InlineObject101
type InlineObject101 struct {
	// The list of auto locate jobs to be scheduled. Up to 100 jobs can be provided in a request.
	Jobs []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs `json:"jobs"`
}

// NewInlineObject101 instantiates a new InlineObject101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject101(jobs []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) *InlineObject101 {
	this := InlineObject101{}
	this.Jobs = jobs
	return &this
}

// NewInlineObject101WithDefaults instantiates a new InlineObject101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject101WithDefaults() *InlineObject101 {
	this := InlineObject101{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *InlineObject101) GetJobs() []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	if o == nil {
		var ret []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *InlineObject101) GetJobsOk() ([]NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs, bool) {
	if o == nil {
    return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *InlineObject101) SetJobs(v []NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) {
	o.Jobs = v
}

func (o InlineObject101) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject101 struct {
	value *InlineObject101
	isSet bool
}

func (v NullableInlineObject101) Get() *InlineObject101 {
	return v.value
}

func (v *NullableInlineObject101) Set(val *InlineObject101) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject101) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject101) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject101(val *InlineObject101) *NullableInlineObject101 {
	return &NullableInlineObject101{value: val, isSet: true}
}

func (v NullableInlineObject101) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject101) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


