/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200221 struct for InlineResponse200221
type InlineResponse200221 struct {
	// ID of the health alert
	Id string `json:"id"`
	// Category type that the health alert belongs to
	CategoryType string `json:"categoryType"`
	Network OrganizationsOrganizationIdAssuranceAlertsNetwork `json:"network"`
	// Time when the alert started
	StartedAt time.Time `json:"startedAt"`
	// Time when the alert was resolved
	ResolvedAt *time.Time `json:"resolvedAt,omitempty"`
	// Time when the alert was dismissed
	DismissedAt *time.Time `json:"dismissedAt,omitempty"`
	// Device Type that the alert occurred on
	DeviceType *string `json:"deviceType,omitempty"`
	// Alert Type
	Type string `json:"type"`
	// Human Readable Title for Alert type
	Title string `json:"title"`
	// Description of the alert
	Description *string `json:"description,omitempty"`
	// Alert severity
	Severity string `json:"severity"`
	Scope *OrganizationsOrganizationIdAssuranceAlertsScope `json:"scope,omitempty"`
}

// NewInlineResponse200221 instantiates a new InlineResponse200221 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200221(id string, categoryType string, network OrganizationsOrganizationIdAssuranceAlertsNetwork, startedAt time.Time, type_ string, title string, severity string) *InlineResponse200221 {
	this := InlineResponse200221{}
	this.Id = id
	this.CategoryType = categoryType
	this.Network = network
	this.StartedAt = startedAt
	this.Type = type_
	this.Title = title
	this.Severity = severity
	return &this
}

// NewInlineResponse200221WithDefaults instantiates a new InlineResponse200221 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200221WithDefaults() *InlineResponse200221 {
	this := InlineResponse200221{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse200221) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse200221) SetId(v string) {
	o.Id = v
}

// GetCategoryType returns the CategoryType field value
func (o *InlineResponse200221) GetCategoryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryType
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetCategoryTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CategoryType, true
}

// SetCategoryType sets field value
func (o *InlineResponse200221) SetCategoryType(v string) {
	o.CategoryType = v
}

// GetNetwork returns the Network field value
func (o *InlineResponse200221) GetNetwork() OrganizationsOrganizationIdAssuranceAlertsNetwork {
	if o == nil {
		var ret OrganizationsOrganizationIdAssuranceAlertsNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetNetworkOk() (*OrganizationsOrganizationIdAssuranceAlertsNetwork, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *InlineResponse200221) SetNetwork(v OrganizationsOrganizationIdAssuranceAlertsNetwork) {
	o.Network = v
}

// GetStartedAt returns the StartedAt field value
func (o *InlineResponse200221) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *InlineResponse200221) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise.
func (o *InlineResponse200221) GetResolvedAt() time.Time {
	if o == nil || isNil(o.ResolvedAt) {
		var ret time.Time
		return ret
	}
	return *o.ResolvedAt
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ResolvedAt) {
    return nil, false
	}
	return o.ResolvedAt, true
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *InlineResponse200221) HasResolvedAt() bool {
	if o != nil && !isNil(o.ResolvedAt) {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given time.Time and assigns it to the ResolvedAt field.
func (o *InlineResponse200221) SetResolvedAt(v time.Time) {
	o.ResolvedAt = &v
}

// GetDismissedAt returns the DismissedAt field value if set, zero value otherwise.
func (o *InlineResponse200221) GetDismissedAt() time.Time {
	if o == nil || isNil(o.DismissedAt) {
		var ret time.Time
		return ret
	}
	return *o.DismissedAt
}

// GetDismissedAtOk returns a tuple with the DismissedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetDismissedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.DismissedAt) {
    return nil, false
	}
	return o.DismissedAt, true
}

// HasDismissedAt returns a boolean if a field has been set.
func (o *InlineResponse200221) HasDismissedAt() bool {
	if o != nil && !isNil(o.DismissedAt) {
		return true
	}

	return false
}

// SetDismissedAt gets a reference to the given time.Time and assigns it to the DismissedAt field.
func (o *InlineResponse200221) SetDismissedAt(v time.Time) {
	o.DismissedAt = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *InlineResponse200221) GetDeviceType() string {
	if o == nil || isNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetDeviceTypeOk() (*string, bool) {
	if o == nil || isNil(o.DeviceType) {
    return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *InlineResponse200221) HasDeviceType() bool {
	if o != nil && !isNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *InlineResponse200221) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetType returns the Type field value
func (o *InlineResponse200221) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse200221) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *InlineResponse200221) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InlineResponse200221) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200221) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200221) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200221) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value
func (o *InlineResponse200221) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetSeverityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *InlineResponse200221) SetSeverity(v string) {
	o.Severity = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200221) GetScope() OrganizationsOrganizationIdAssuranceAlertsScope {
	if o == nil || isNil(o.Scope) {
		var ret OrganizationsOrganizationIdAssuranceAlertsScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200221) GetScopeOk() (*OrganizationsOrganizationIdAssuranceAlertsScope, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200221) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given OrganizationsOrganizationIdAssuranceAlertsScope and assigns it to the Scope field.
func (o *InlineResponse200221) SetScope(v OrganizationsOrganizationIdAssuranceAlertsScope) {
	o.Scope = &v
}

func (o InlineResponse200221) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["categoryType"] = o.CategoryType
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !isNil(o.ResolvedAt) {
		toSerialize["resolvedAt"] = o.ResolvedAt
	}
	if !isNil(o.DismissedAt) {
		toSerialize["dismissedAt"] = o.DismissedAt
	}
	if !isNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200221 struct {
	value *InlineResponse200221
	isSet bool
}

func (v NullableInlineResponse200221) Get() *InlineResponse200221 {
	return v.value
}

func (v *NullableInlineResponse200221) Set(val *InlineResponse200221) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200221) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200221) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200221(val *InlineResponse200221) *NullableInlineResponse200221 {
	return &NullableInlineResponse200221{value: val, isSet: true}
}

func (v NullableInlineResponse200221) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200221) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


