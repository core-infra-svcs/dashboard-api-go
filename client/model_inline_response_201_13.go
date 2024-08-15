/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20113 struct for InlineResponse20113
type InlineResponse20113 struct {
	// Custom analytics artifact ID
	ArtifactId *string `json:"artifactId,omitempty"`
	// Organization ID
	OrganizationId *string `json:"organizationId,omitempty"`
	// Custom analytics artifact name
	Name *string `json:"name,omitempty"`
	Status *OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus `json:"status,omitempty"`
	// Upload ID
	UploadId *string `json:"uploadId,omitempty"`
	// Upload URL
	UploadUrl *string `json:"uploadUrl,omitempty"`
	// Upload URL expiry time
	UploadUrlExpiry *time.Time `json:"uploadUrlExpiry,omitempty"`
}

// NewInlineResponse20113 instantiates a new InlineResponse20113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20113() *InlineResponse20113 {
	this := InlineResponse20113{}
	return &this
}

// NewInlineResponse20113WithDefaults instantiates a new InlineResponse20113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20113WithDefaults() *InlineResponse20113 {
	this := InlineResponse20113{}
	return &this
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *InlineResponse20113) GetArtifactId() string {
	if o == nil || isNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetArtifactIdOk() (*string, bool) {
	if o == nil || isNil(o.ArtifactId) {
    return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *InlineResponse20113) HasArtifactId() bool {
	if o != nil && !isNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *InlineResponse20113) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse20113) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse20113) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse20113) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20113) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20113) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20113) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20113) GetStatus() OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus {
	if o == nil || isNil(o.Status) {
		var ret OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetStatusOk() (*OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20113) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus and assigns it to the Status field.
func (o *InlineResponse20113) SetStatus(v OrganizationsOrganizationIdCameraCustomAnalyticsArtifactsStatus) {
	o.Status = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *InlineResponse20113) GetUploadId() string {
	if o == nil || isNil(o.UploadId) {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetUploadIdOk() (*string, bool) {
	if o == nil || isNil(o.UploadId) {
    return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *InlineResponse20113) HasUploadId() bool {
	if o != nil && !isNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *InlineResponse20113) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *InlineResponse20113) GetUploadUrl() string {
	if o == nil || isNil(o.UploadUrl) {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetUploadUrlOk() (*string, bool) {
	if o == nil || isNil(o.UploadUrl) {
    return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *InlineResponse20113) HasUploadUrl() bool {
	if o != nil && !isNil(o.UploadUrl) {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *InlineResponse20113) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadUrlExpiry returns the UploadUrlExpiry field value if set, zero value otherwise.
func (o *InlineResponse20113) GetUploadUrlExpiry() time.Time {
	if o == nil || isNil(o.UploadUrlExpiry) {
		var ret time.Time
		return ret
	}
	return *o.UploadUrlExpiry
}

// GetUploadUrlExpiryOk returns a tuple with the UploadUrlExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20113) GetUploadUrlExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.UploadUrlExpiry) {
    return nil, false
	}
	return o.UploadUrlExpiry, true
}

// HasUploadUrlExpiry returns a boolean if a field has been set.
func (o *InlineResponse20113) HasUploadUrlExpiry() bool {
	if o != nil && !isNil(o.UploadUrlExpiry) {
		return true
	}

	return false
}

// SetUploadUrlExpiry gets a reference to the given time.Time and assigns it to the UploadUrlExpiry field.
func (o *InlineResponse20113) SetUploadUrlExpiry(v time.Time) {
	o.UploadUrlExpiry = &v
}

func (o InlineResponse20113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ArtifactId) {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.UploadId) {
		toSerialize["uploadId"] = o.UploadId
	}
	if !isNil(o.UploadUrl) {
		toSerialize["uploadUrl"] = o.UploadUrl
	}
	if !isNil(o.UploadUrlExpiry) {
		toSerialize["uploadUrlExpiry"] = o.UploadUrlExpiry
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20113 struct {
	value *InlineResponse20113
	isSet bool
}

func (v NullableInlineResponse20113) Get() *InlineResponse20113 {
	return v.value
}

func (v *NullableInlineResponse20113) Set(val *InlineResponse20113) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20113) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20113(val *InlineResponse20113) *NullableInlineResponse20113 {
	return &NullableInlineResponse20113{value: val, isSet: true}
}

func (v NullableInlineResponse20113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


