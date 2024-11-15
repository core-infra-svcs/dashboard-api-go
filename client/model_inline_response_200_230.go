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

// InlineResponse200230 struct for InlineResponse200230
type InlineResponse200230 struct {
	// The boundary id
	BoundaryId *string `json:"boundaryId,omitempty"`
	// The boundary type
	Type *string `json:"type,omitempty"`
	Results *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults `json:"results,omitempty"`
}

// NewInlineResponse200230 instantiates a new InlineResponse200230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200230() *InlineResponse200230 {
	this := InlineResponse200230{}
	return &this
}

// NewInlineResponse200230WithDefaults instantiates a new InlineResponse200230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200230WithDefaults() *InlineResponse200230 {
	this := InlineResponse200230{}
	return &this
}

// GetBoundaryId returns the BoundaryId field value if set, zero value otherwise.
func (o *InlineResponse200230) GetBoundaryId() string {
	if o == nil || isNil(o.BoundaryId) {
		var ret string
		return ret
	}
	return *o.BoundaryId
}

// GetBoundaryIdOk returns a tuple with the BoundaryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230) GetBoundaryIdOk() (*string, bool) {
	if o == nil || isNil(o.BoundaryId) {
    return nil, false
	}
	return o.BoundaryId, true
}

// HasBoundaryId returns a boolean if a field has been set.
func (o *InlineResponse200230) HasBoundaryId() bool {
	if o != nil && !isNil(o.BoundaryId) {
		return true
	}

	return false
}

// SetBoundaryId gets a reference to the given string and assigns it to the BoundaryId field.
func (o *InlineResponse200230) SetBoundaryId(v string) {
	o.BoundaryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200230) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200230) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200230) SetType(v string) {
	o.Type = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse200230) GetResults() OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	if o == nil || isNil(o.Results) {
		var ret OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200230) GetResultsOk() (*OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse200230) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults and assigns it to the Results field.
func (o *InlineResponse200230) SetResults(v OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) {
	o.Results = &v
}

func (o InlineResponse200230) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200230 struct {
	value *InlineResponse200230
	isSet bool
}

func (v NullableInlineResponse200230) Get() *InlineResponse200230 {
	return v.value
}

func (v *NullableInlineResponse200230) Set(val *InlineResponse200230) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200230) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200230(val *InlineResponse200230) *NullableInlineResponse200230 {
	return &NullableInlineResponse200230{value: val, isSet: true}
}

func (v NullableInlineResponse200230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


