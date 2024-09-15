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

// InlineResponse200225 struct for InlineResponse200225
type InlineResponse200225 struct {
	// The boundary id
	BoundaryId *string `json:"boundaryId,omitempty"`
	// The boundary type
	Type *string `json:"type,omitempty"`
	Results *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults `json:"results,omitempty"`
}

// NewInlineResponse200225 instantiates a new InlineResponse200225 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200225() *InlineResponse200225 {
	this := InlineResponse200225{}
	return &this
}

// NewInlineResponse200225WithDefaults instantiates a new InlineResponse200225 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200225WithDefaults() *InlineResponse200225 {
	this := InlineResponse200225{}
	return &this
}

// GetBoundaryId returns the BoundaryId field value if set, zero value otherwise.
func (o *InlineResponse200225) GetBoundaryId() string {
	if o == nil || isNil(o.BoundaryId) {
		var ret string
		return ret
	}
	return *o.BoundaryId
}

// GetBoundaryIdOk returns a tuple with the BoundaryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200225) GetBoundaryIdOk() (*string, bool) {
	if o == nil || isNil(o.BoundaryId) {
    return nil, false
	}
	return o.BoundaryId, true
}

// HasBoundaryId returns a boolean if a field has been set.
func (o *InlineResponse200225) HasBoundaryId() bool {
	if o != nil && !isNil(o.BoundaryId) {
		return true
	}

	return false
}

// SetBoundaryId gets a reference to the given string and assigns it to the BoundaryId field.
func (o *InlineResponse200225) SetBoundaryId(v string) {
	o.BoundaryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200225) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200225) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200225) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200225) SetType(v string) {
	o.Type = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse200225) GetResults() OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	if o == nil || isNil(o.Results) {
		var ret OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200225) GetResultsOk() (*OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse200225) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults and assigns it to the Results field.
func (o *InlineResponse200225) SetResults(v OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) {
	o.Results = &v
}

func (o InlineResponse200225) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BoundaryId) {
		toSerialize["boundaryId"] = o.BoundaryId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200225 struct {
	value *InlineResponse200225
	isSet bool
}

func (v NullableInlineResponse200225) Get() *InlineResponse200225 {
	return v.value
}

func (v *NullableInlineResponse200225) Set(val *InlineResponse200225) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200225) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200225) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200225(val *InlineResponse200225) *NullableInlineResponse200225 {
	return &NullableInlineResponse200225{value: val, isSet: true}
}

func (v NullableInlineResponse200225) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200225) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


