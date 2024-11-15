/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs struct for NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs
type NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs struct {
	// The ID of the floor plan to run auto locate for
	FloorPlanId string `json:"floorPlanId"`
	// The types of location data that should be refreshed for this job. The list must either contain both 'gnss' and 'ranging' or be empty, as we currently only support refreshing both 'gnss' and 'ranging', or neither.
	Refresh []string `json:"refresh,omitempty"`
	// Timestamp in ISO8601 format which indicates when the auto locate job should be run. If omitted, the auto locate job will start immediately.
	ScheduledAt *time.Time `json:"scheduledAt,omitempty"`
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs(floorPlanId string) *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{}
	this.FloorPlanId = floorPlanId
	return &this
}

// NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobsWithDefaults instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobsWithDefaults() *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	this := NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{}
	return &this
}

// GetFloorPlanId returns the FloorPlanId field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetFloorPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetFloorPlanIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FloorPlanId, true
}

// SetFloorPlanId sets field value
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetFloorPlanId(v string) {
	o.FloorPlanId = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetRefresh() []string {
	if o == nil || isNil(o.Refresh) {
		var ret []string
		return ret
	}
	return o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetRefreshOk() ([]string, bool) {
	if o == nil || isNil(o.Refresh) {
    return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) HasRefresh() bool {
	if o != nil && !isNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given []string and assigns it to the Refresh field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetRefresh(v []string) {
	o.Refresh = v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetScheduledAt() time.Time {
	if o == nil || isNil(o.ScheduledAt) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledAt) {
    return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) HasScheduledAt() bool {
	if o != nil && !isNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

func (o NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	if !isNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	if !isNil(o.ScheduledAt) {
		toSerialize["scheduledAt"] = o.ScheduledAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs struct {
	value *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs
	isSet bool
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) Get() *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	return v.value
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) Set(val *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs(val *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) *NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs {
	return &NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

