/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs struct for OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs
type OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs struct {
	// Auto locate job ID
	Id *string `json:"id,omitempty"`
	// Auto locate job status. Possible values: 'scheduled', 'in progress', 'canceling', 'error', 'finished', 'published', 'canceled'
	Status *string `json:"status,omitempty"`
	// Scheduled start time for auto locate job
	ScheduledAt *time.Time `json:"scheduledAt,omitempty"`
	Completed *InlineResponse200101Completed `json:"completed,omitempty"`
	Ranging *InlineResponse200101Ranging `json:"ranging,omitempty"`
	Gnss *InlineResponse200101Gnss `json:"gnss,omitempty"`
	// List of errors that occurred during a failed run of auto locate
	Errors []InlineResponse200101Errors `json:"errors,omitempty"`
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs{}
	return &this
}

// NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobsWithDefaults instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobsWithDefaults() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs {
	this := OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetStatus(v string) {
	o.Status = &v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetScheduledAt() time.Time {
	if o == nil || isNil(o.ScheduledAt) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ScheduledAt) {
    return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasScheduledAt() bool {
	if o != nil && !isNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given time.Time and assigns it to the ScheduledAt field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetScheduledAt(v time.Time) {
	o.ScheduledAt = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetCompleted() InlineResponse200101Completed {
	if o == nil || isNil(o.Completed) {
		var ret InlineResponse200101Completed
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetCompletedOk() (*InlineResponse200101Completed, bool) {
	if o == nil || isNil(o.Completed) {
    return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given InlineResponse200101Completed and assigns it to the Completed field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetCompleted(v InlineResponse200101Completed) {
	o.Completed = &v
}

// GetRanging returns the Ranging field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetRanging() InlineResponse200101Ranging {
	if o == nil || isNil(o.Ranging) {
		var ret InlineResponse200101Ranging
		return ret
	}
	return *o.Ranging
}

// GetRangingOk returns a tuple with the Ranging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetRangingOk() (*InlineResponse200101Ranging, bool) {
	if o == nil || isNil(o.Ranging) {
    return nil, false
	}
	return o.Ranging, true
}

// HasRanging returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasRanging() bool {
	if o != nil && !isNil(o.Ranging) {
		return true
	}

	return false
}

// SetRanging gets a reference to the given InlineResponse200101Ranging and assigns it to the Ranging field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetRanging(v InlineResponse200101Ranging) {
	o.Ranging = &v
}

// GetGnss returns the Gnss field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetGnss() InlineResponse200101Gnss {
	if o == nil || isNil(o.Gnss) {
		var ret InlineResponse200101Gnss
		return ret
	}
	return *o.Gnss
}

// GetGnssOk returns a tuple with the Gnss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetGnssOk() (*InlineResponse200101Gnss, bool) {
	if o == nil || isNil(o.Gnss) {
    return nil, false
	}
	return o.Gnss, true
}

// HasGnss returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasGnss() bool {
	if o != nil && !isNil(o.Gnss) {
		return true
	}

	return false
}

// SetGnss gets a reference to the given InlineResponse200101Gnss and assigns it to the Gnss field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetGnss(v InlineResponse200101Gnss) {
	o.Gnss = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetErrors() []InlineResponse200101Errors {
	if o == nil || isNil(o.Errors) {
		var ret []InlineResponse200101Errors
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetErrorsOk() ([]InlineResponse200101Errors, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []InlineResponse200101Errors and assigns it to the Errors field.
func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetErrors(v []InlineResponse200101Errors) {
	o.Errors = v
}

func (o OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ScheduledAt) {
		toSerialize["scheduledAt"] = o.ScheduledAt
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.Ranging) {
		toSerialize["ranging"] = o.Ranging
	}
	if !isNil(o.Gnss) {
		toSerialize["gnss"] = o.Gnss
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs struct {
	value *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) Get() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) Set(val *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs(val *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs {
	return &NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


