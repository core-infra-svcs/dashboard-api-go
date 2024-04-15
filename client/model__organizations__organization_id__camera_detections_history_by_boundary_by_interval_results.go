/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults The analytics data
type OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults struct {
	// The period start time
	StartTime *string `json:"startTime,omitempty"`
	// The period end time
	EndTime *string `json:"endTime,omitempty"`
	// The detection type
	ObjectType *string `json:"objectType,omitempty"`
	// The number of detections entered
	In *int32 `json:"in,omitempty"`
	// The number of detections exited
	Out *int32 `json:"out,omitempty"`
}

// NewOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults instantiates a new OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults() *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	this := OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults{}
	return &this
}

// NewOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResultsWithDefaults instantiates a new OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResultsWithDefaults() *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	this := OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetStartTime() string {
	if o == nil || isNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetStartTimeOk() (*string, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetEndTime() string {
	if o == nil || isNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetEndTimeOk() (*string, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) SetEndTime(v string) {
	o.EndTime = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
    return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetIn() int32 {
	if o == nil || isNil(o.In) {
		var ret int32
		return ret
	}
	return *o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetInOk() (*int32, bool) {
	if o == nil || isNil(o.In) {
    return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) HasIn() bool {
	if o != nil && !isNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given int32 and assigns it to the In field.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) SetIn(v int32) {
	o.In = &v
}

// GetOut returns the Out field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetOut() int32 {
	if o == nil || isNil(o.Out) {
		var ret int32
		return ret
	}
	return *o.Out
}

// GetOutOk returns a tuple with the Out field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) GetOutOk() (*int32, bool) {
	if o == nil || isNil(o.Out) {
    return nil, false
	}
	return o.Out, true
}

// HasOut returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) HasOut() bool {
	if o != nil && !isNil(o.Out) {
		return true
	}

	return false
}

// SetOut gets a reference to the given int32 and assigns it to the Out field.
func (o *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) SetOut(v int32) {
	o.Out = &v
}

func (o OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.In) {
		toSerialize["in"] = o.In
	}
	if !isNil(o.Out) {
		toSerialize["out"] = o.Out
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults struct {
	value *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) Get() *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) Set(val *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults(val *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) *NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	return &NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


