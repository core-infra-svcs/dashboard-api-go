/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject162 struct for InlineObject162
type InlineObject162 struct {
	// Is the alert config enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The alert type
	Type *string `json:"type,omitempty"`
	AlertCondition *OrganizationsOrganizationIdAlertsProfilesAlertCondition `json:"alertCondition,omitempty"`
	Recipients *OrganizationsOrganizationIdAlertsProfilesRecipients `json:"recipients,omitempty"`
	// Networks with these tags will be monitored for the alert
	NetworkTags *[]string `json:"networkTags,omitempty"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewInlineObject162 instantiates a new InlineObject162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject162() *InlineObject162 {
	this := InlineObject162{}
	return &this
}

// NewInlineObject162WithDefaults instantiates a new InlineObject162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject162WithDefaults() *InlineObject162 {
	this := InlineObject162{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject162) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject162) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject162) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject162) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject162) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject162) SetType(v string) {
	o.Type = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *InlineObject162) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition {
	if o == nil || o.AlertCondition == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesAlertCondition
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool) {
	if o == nil || o.AlertCondition == nil {
		return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *InlineObject162) HasAlertCondition() bool {
	if o != nil && o.AlertCondition != nil {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given OrganizationsOrganizationIdAlertsProfilesAlertCondition and assigns it to the AlertCondition field.
func (o *InlineObject162) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition) {
	o.AlertCondition = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InlineObject162) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients {
	if o == nil || o.Recipients == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool) {
	if o == nil || o.Recipients == nil {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InlineObject162) HasRecipients() bool {
	if o != nil && o.Recipients != nil {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given OrganizationsOrganizationIdAlertsProfilesRecipients and assigns it to the Recipients field.
func (o *InlineObject162) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients) {
	o.Recipients = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *InlineObject162) GetNetworkTags() []string {
	if o == nil || o.NetworkTags == nil {
		var ret []string
		return ret
	}
	return *o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetNetworkTagsOk() (*[]string, bool) {
	if o == nil || o.NetworkTags == nil {
		return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *InlineObject162) HasNetworkTags() bool {
	if o != nil && o.NetworkTags != nil {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *InlineObject162) SetNetworkTags(v []string) {
	o.NetworkTags = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject162) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject162) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject162) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject162) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject162) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.AlertCondition != nil {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	if o.NetworkTags != nil {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject162 struct {
	value *InlineObject162
	isSet bool
}

func (v NullableInlineObject162) Get() *InlineObject162 {
	return v.value
}

func (v *NullableInlineObject162) Set(val *InlineObject162) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject162) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject162) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject162(val *InlineObject162) *NullableInlineObject162 {
	return &NullableInlineObject162{value: val, isSet: true}
}

func (v NullableInlineObject162) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject162) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


