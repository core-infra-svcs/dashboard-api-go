/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200242 struct for InlineResponse200242
type InlineResponse200242 struct {
	// The boundary id
	BoundaryId *string `json:"boundaryId,omitempty"`
	// The boundary type
	Type *string `json:"type,omitempty"`
	Results *OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults `json:"results,omitempty"`
}

// NewInlineResponse200242 instantiates a new InlineResponse200242 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200242() *InlineResponse200242 {
	this := InlineResponse200242{}
	return &this
}

// NewInlineResponse200242WithDefaults instantiates a new InlineResponse200242 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200242WithDefaults() *InlineResponse200242 {
	this := InlineResponse200242{}
	return &this
}

// GetBoundaryId returns the BoundaryId field value if set, zero value otherwise.
func (o *InlineResponse200242) GetBoundaryId() string {
	if o == nil || isNil(o.BoundaryId) {
		var ret string
		return ret
	}
	return *o.BoundaryId
}

// GetBoundaryIdOk returns a tuple with the BoundaryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200242) GetBoundaryIdOk() (*string, bool) {
	if o == nil || isNil(o.BoundaryId) {
    return nil, false
	}
	return o.BoundaryId, true
}

// HasBoundaryId returns a boolean if a field has been set.
func (o *InlineResponse200242) HasBoundaryId() bool {
	if o != nil && !isNil(o.BoundaryId) {
		return true
	}

	return false
}

// SetBoundaryId gets a reference to the given string and assigns it to the BoundaryId field.
func (o *InlineResponse200242) SetBoundaryId(v string) {
	o.BoundaryId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200242) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200242) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200242) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200242) SetType(v string) {
	o.Type = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse200242) GetResults() OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults {
	if o == nil || isNil(o.Results) {
		var ret OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200242) GetResultsOk() (*OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse200242) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults and assigns it to the Results field.
func (o *InlineResponse200242) SetResults(v OrganizationsOrganizationIdCameraDetectionsHistoryByBoundaryByIntervalResults) {
	o.Results = &v
}

func (o InlineResponse200242) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200242 struct {
	value *InlineResponse200242
	isSet bool
}

func (v NullableInlineResponse200242) Get() *InlineResponse200242 {
	return v.value
}

func (v *NullableInlineResponse200242) Set(val *InlineResponse200242) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200242) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200242) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200242(val *InlineResponse200242) *NullableInlineResponse200242 {
	return &NullableInlineResponse200242{value: val, isSet: true}
}

func (v NullableInlineResponse200242) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200242) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


