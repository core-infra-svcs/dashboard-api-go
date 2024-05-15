/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200256 struct for InlineResponse200256
type InlineResponse200256 struct {
	// The key of the license
	Key *string `json:"key,omitempty"`
	// The ID of the organization that the license is claimed in
	OrganizationId *string `json:"organizationId,omitempty"`
	// The duration (term length) of the license, measured in days
	Duration *int32 `json:"duration,omitempty"`
	// The operation mode of the license when it was claimed
	Mode *string `json:"mode,omitempty"`
	// When the license's term began (approximately the date when the license was created)
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// When the license was claimed into the organization
	ClaimedAt *time.Time `json:"claimedAt,omitempty"`
	// Flag to indicated that the license is invalidated
	Invalidated *bool `json:"invalidated,omitempty"`
	// When the license was invalidated. Will be null for active licenses
	InvalidatedAt *time.Time `json:"invalidatedAt,omitempty"`
	// Flag to indicate if the license is expired
	Expired *bool `json:"expired,omitempty"`
	// The editions of the license for each relevant product type
	Editions []OrganizationsOrganizationIdLicensingCotermLicensesEditions `json:"editions,omitempty"`
	// The counts of the license by model type
	Counts []OrganizationsOrganizationIdLicensingCotermLicensesCounts `json:"counts,omitempty"`
}

// NewInlineResponse200256 instantiates a new InlineResponse200256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200256() *InlineResponse200256 {
	this := InlineResponse200256{}
	return &this
}

// NewInlineResponse200256WithDefaults instantiates a new InlineResponse200256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200256WithDefaults() *InlineResponse200256 {
	this := InlineResponse200256{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineResponse200256) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineResponse200256) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InlineResponse200256) SetKey(v string) {
	o.Key = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200256) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200256) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse200256) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse200256) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse200256) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse200256) SetDuration(v int32) {
	o.Duration = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse200256) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse200256) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse200256) SetMode(v string) {
	o.Mode = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *InlineResponse200256) GetStartedAt() time.Time {
	if o == nil || isNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartedAt) {
    return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *InlineResponse200256) HasStartedAt() bool {
	if o != nil && !isNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *InlineResponse200256) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetClaimedAt returns the ClaimedAt field value if set, zero value otherwise.
func (o *InlineResponse200256) GetClaimedAt() time.Time {
	if o == nil || isNil(o.ClaimedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClaimedAt
}

// GetClaimedAtOk returns a tuple with the ClaimedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetClaimedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ClaimedAt) {
    return nil, false
	}
	return o.ClaimedAt, true
}

// HasClaimedAt returns a boolean if a field has been set.
func (o *InlineResponse200256) HasClaimedAt() bool {
	if o != nil && !isNil(o.ClaimedAt) {
		return true
	}

	return false
}

// SetClaimedAt gets a reference to the given time.Time and assigns it to the ClaimedAt field.
func (o *InlineResponse200256) SetClaimedAt(v time.Time) {
	o.ClaimedAt = &v
}

// GetInvalidated returns the Invalidated field value if set, zero value otherwise.
func (o *InlineResponse200256) GetInvalidated() bool {
	if o == nil || isNil(o.Invalidated) {
		var ret bool
		return ret
	}
	return *o.Invalidated
}

// GetInvalidatedOk returns a tuple with the Invalidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetInvalidatedOk() (*bool, bool) {
	if o == nil || isNil(o.Invalidated) {
    return nil, false
	}
	return o.Invalidated, true
}

// HasInvalidated returns a boolean if a field has been set.
func (o *InlineResponse200256) HasInvalidated() bool {
	if o != nil && !isNil(o.Invalidated) {
		return true
	}

	return false
}

// SetInvalidated gets a reference to the given bool and assigns it to the Invalidated field.
func (o *InlineResponse200256) SetInvalidated(v bool) {
	o.Invalidated = &v
}

// GetInvalidatedAt returns the InvalidatedAt field value if set, zero value otherwise.
func (o *InlineResponse200256) GetInvalidatedAt() time.Time {
	if o == nil || isNil(o.InvalidatedAt) {
		var ret time.Time
		return ret
	}
	return *o.InvalidatedAt
}

// GetInvalidatedAtOk returns a tuple with the InvalidatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetInvalidatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.InvalidatedAt) {
    return nil, false
	}
	return o.InvalidatedAt, true
}

// HasInvalidatedAt returns a boolean if a field has been set.
func (o *InlineResponse200256) HasInvalidatedAt() bool {
	if o != nil && !isNil(o.InvalidatedAt) {
		return true
	}

	return false
}

// SetInvalidatedAt gets a reference to the given time.Time and assigns it to the InvalidatedAt field.
func (o *InlineResponse200256) SetInvalidatedAt(v time.Time) {
	o.InvalidatedAt = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *InlineResponse200256) GetExpired() bool {
	if o == nil || isNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetExpiredOk() (*bool, bool) {
	if o == nil || isNil(o.Expired) {
    return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *InlineResponse200256) HasExpired() bool {
	if o != nil && !isNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *InlineResponse200256) SetExpired(v bool) {
	o.Expired = &v
}

// GetEditions returns the Editions field value if set, zero value otherwise.
func (o *InlineResponse200256) GetEditions() []OrganizationsOrganizationIdLicensingCotermLicensesEditions {
	if o == nil || isNil(o.Editions) {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesEditions
		return ret
	}
	return o.Editions
}

// GetEditionsOk returns a tuple with the Editions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetEditionsOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesEditions, bool) {
	if o == nil || isNil(o.Editions) {
    return nil, false
	}
	return o.Editions, true
}

// HasEditions returns a boolean if a field has been set.
func (o *InlineResponse200256) HasEditions() bool {
	if o != nil && !isNil(o.Editions) {
		return true
	}

	return false
}

// SetEditions gets a reference to the given []OrganizationsOrganizationIdLicensingCotermLicensesEditions and assigns it to the Editions field.
func (o *InlineResponse200256) SetEditions(v []OrganizationsOrganizationIdLicensingCotermLicensesEditions) {
	o.Editions = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200256) GetCounts() []OrganizationsOrganizationIdLicensingCotermLicensesCounts {
	if o == nil || isNil(o.Counts) {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesCounts
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200256) GetCountsOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200256) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []OrganizationsOrganizationIdLicensingCotermLicensesCounts and assigns it to the Counts field.
func (o *InlineResponse200256) SetCounts(v []OrganizationsOrganizationIdLicensingCotermLicensesCounts) {
	o.Counts = v
}

func (o InlineResponse200256) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !isNil(o.ClaimedAt) {
		toSerialize["claimedAt"] = o.ClaimedAt
	}
	if !isNil(o.Invalidated) {
		toSerialize["invalidated"] = o.Invalidated
	}
	if !isNil(o.InvalidatedAt) {
		toSerialize["invalidatedAt"] = o.InvalidatedAt
	}
	if !isNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !isNil(o.Editions) {
		toSerialize["editions"] = o.Editions
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200256 struct {
	value *InlineResponse200256
	isSet bool
}

func (v NullableInlineResponse200256) Get() *InlineResponse200256 {
	return v.value
}

func (v *NullableInlineResponse200256) Set(val *InlineResponse200256) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200256) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200256(val *InlineResponse200256) *NullableInlineResponse200256 {
	return &NullableInlineResponse200256{value: val, isSet: true}
}

func (v NullableInlineResponse200256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


