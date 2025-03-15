/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200279 struct for InlineResponse200279
type InlineResponse200279 struct {
	// Application identifier
	ApplicationId *string `json:"applicationId,omitempty"`
	// Application name
	Name *string `json:"name,omitempty"`
	Thresholds *OrganizationsOrganizationIdInsightApplicationsThresholds `json:"thresholds,omitempty"`
}

// NewInlineResponse200279 instantiates a new InlineResponse200279 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200279() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// NewInlineResponse200279WithDefaults instantiates a new InlineResponse200279 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200279WithDefaults() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *InlineResponse200279) GetApplicationId() string {
	if o == nil || isNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetApplicationIdOk() (*string, bool) {
	if o == nil || isNil(o.ApplicationId) {
    return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *InlineResponse200279) HasApplicationId() bool {
	if o != nil && !isNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *InlineResponse200279) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200279) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200279) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200279) SetName(v string) {
	o.Name = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *InlineResponse200279) GetThresholds() OrganizationsOrganizationIdInsightApplicationsThresholds {
	if o == nil || isNil(o.Thresholds) {
		var ret OrganizationsOrganizationIdInsightApplicationsThresholds
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetThresholdsOk() (*OrganizationsOrganizationIdInsightApplicationsThresholds, bool) {
	if o == nil || isNil(o.Thresholds) {
    return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *InlineResponse200279) HasThresholds() bool {
	if o != nil && !isNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given OrganizationsOrganizationIdInsightApplicationsThresholds and assigns it to the Thresholds field.
func (o *InlineResponse200279) SetThresholds(v OrganizationsOrganizationIdInsightApplicationsThresholds) {
	o.Thresholds = &v
}

func (o InlineResponse200279) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Thresholds) {
		toSerialize["thresholds"] = o.Thresholds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200279 struct {
	value *InlineResponse200279
	isSet bool
}

func (v NullableInlineResponse200279) Get() *InlineResponse200279 {
	return v.value
}

func (v *NullableInlineResponse200279) Set(val *InlineResponse200279) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200279) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200279) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200279(val *InlineResponse200279) *NullableInlineResponse200279 {
	return &NullableInlineResponse200279{value: val, isSet: true}
}

func (v NullableInlineResponse200279) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200279) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


