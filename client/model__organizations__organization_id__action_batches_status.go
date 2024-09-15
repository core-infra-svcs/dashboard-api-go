/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdActionBatchesStatus Status of action batch
type OrganizationsOrganizationIdActionBatchesStatus struct {
	// Flag describing whether all actions in the action batch have completed
	Completed *bool `json:"completed,omitempty"`
	// Flag describing whether any actions in the action batch failed
	Failed *bool `json:"failed,omitempty"`
	// List of errors encountered when running actions in the action batch
	Errors []string `json:"errors,omitempty"`
	// Resources created as a result of this action batch
	CreatedResources []OrganizationsOrganizationIdActionBatchesStatusCreatedResources `json:"createdResources"`
}

// NewOrganizationsOrganizationIdActionBatchesStatus instantiates a new OrganizationsOrganizationIdActionBatchesStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesStatus(createdResources []OrganizationsOrganizationIdActionBatchesStatusCreatedResources) *OrganizationsOrganizationIdActionBatchesStatus {
	this := OrganizationsOrganizationIdActionBatchesStatus{}
	this.CreatedResources = createdResources
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesStatusWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesStatusWithDefaults() *OrganizationsOrganizationIdActionBatchesStatus {
	this := OrganizationsOrganizationIdActionBatchesStatus{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCompleted() bool {
	if o == nil || isNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.Completed) {
    return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *OrganizationsOrganizationIdActionBatchesStatus) SetCompleted(v bool) {
	o.Completed = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetFailed() bool {
	if o == nil || isNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetFailedOk() (*bool, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *OrganizationsOrganizationIdActionBatchesStatus) SetFailed(v bool) {
	o.Failed = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *OrganizationsOrganizationIdActionBatchesStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedResources returns the CreatedResources field value
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCreatedResources() []OrganizationsOrganizationIdActionBatchesStatusCreatedResources {
	if o == nil {
		var ret []OrganizationsOrganizationIdActionBatchesStatusCreatedResources
		return ret
	}

	return o.CreatedResources
}

// GetCreatedResourcesOk returns a tuple with the CreatedResources field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesStatus) GetCreatedResourcesOk() ([]OrganizationsOrganizationIdActionBatchesStatusCreatedResources, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedResources, true
}

// SetCreatedResources sets field value
func (o *OrganizationsOrganizationIdActionBatchesStatus) SetCreatedResources(v []OrganizationsOrganizationIdActionBatchesStatusCreatedResources) {
	o.CreatedResources = v
}

func (o OrganizationsOrganizationIdActionBatchesStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["createdResources"] = o.CreatedResources
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdActionBatchesStatus struct {
	value *OrganizationsOrganizationIdActionBatchesStatus
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatus) Get() *OrganizationsOrganizationIdActionBatchesStatus {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatus) Set(val *OrganizationsOrganizationIdActionBatchesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesStatus(val *OrganizationsOrganizationIdActionBatchesStatus) *NullableOrganizationsOrganizationIdActionBatchesStatus {
	return &NullableOrganizationsOrganizationIdActionBatchesStatus{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


