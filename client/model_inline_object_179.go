/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject179 struct for InlineObject179
type InlineObject179 struct {
	// Is the alert config enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The alert type
	Type *string `json:"type,omitempty"`
	AlertCondition *OrganizationsOrganizationIdAlertsProfilesAlertCondition `json:"alertCondition,omitempty"`
	Recipients *OrganizationsOrganizationIdAlertsProfilesRecipients `json:"recipients,omitempty"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags,omitempty"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewInlineObject179 instantiates a new InlineObject179 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject179() *InlineObject179 {
	this := InlineObject179{}
	return &this
}

// NewInlineObject179WithDefaults instantiates a new InlineObject179 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject179WithDefaults() *InlineObject179 {
	this := InlineObject179{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject179) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject179) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject179) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject179) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject179) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject179) SetType(v string) {
	o.Type = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *InlineObject179) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition {
	if o == nil || isNil(o.AlertCondition) {
		var ret OrganizationsOrganizationIdAlertsProfilesAlertCondition
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool) {
	if o == nil || isNil(o.AlertCondition) {
    return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *InlineObject179) HasAlertCondition() bool {
	if o != nil && !isNil(o.AlertCondition) {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given OrganizationsOrganizationIdAlertsProfilesAlertCondition and assigns it to the AlertCondition field.
func (o *InlineObject179) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition) {
	o.AlertCondition = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InlineObject179) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients {
	if o == nil || isNil(o.Recipients) {
		var ret OrganizationsOrganizationIdAlertsProfilesRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool) {
	if o == nil || isNil(o.Recipients) {
    return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InlineObject179) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given OrganizationsOrganizationIdAlertsProfilesRecipients and assigns it to the Recipients field.
func (o *InlineObject179) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients) {
	o.Recipients = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *InlineObject179) GetNetworkTags() []string {
	if o == nil || isNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkTags) {
    return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *InlineObject179) HasNetworkTags() bool {
	if o != nil && !isNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *InlineObject179) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject179) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject179) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject179) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject179) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject179) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.AlertCondition) {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if !isNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !isNil(o.NetworkTags) {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject179 struct {
	value *InlineObject179
	isSet bool
}

func (v NullableInlineObject179) Get() *InlineObject179 {
	return v.value
}

func (v *NullableInlineObject179) Set(val *InlineObject179) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject179) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject179) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject179(val *InlineObject179) *NullableInlineObject179 {
	return &NullableInlineObject179{value: val, isSet: true}
}

func (v NullableInlineObject179) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject179) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


