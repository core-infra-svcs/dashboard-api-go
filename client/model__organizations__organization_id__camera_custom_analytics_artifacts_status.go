/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus Custom analytics artifact status
type OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus struct {
	// Status type
	Type *string `json:"type,omitempty"`
	// Status message
	Message *string `json:"message,omitempty"`
}

// NewOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus instantiates a new OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus() *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus {
	this := OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus{}
	return &this
}

// NewOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatusWithDefaults instantiates a new OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatusWithDefaults() *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus {
	this := OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) SetMessage(v string) {
	o.Message = &v
}

func (o OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus struct {
	value *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus
	isSet bool
}

func (v NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) Get() *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) Set(val *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus(val *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) *NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus {
	return &NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


