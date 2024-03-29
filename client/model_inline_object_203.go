/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject203 struct for InlineObject203
type InlineObject203 struct {
	// The alert type
	Type string `json:"type"`
	AlertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition `json:"alertCondition"`
	Recipients OrganizationsOrganizationIdAlertsProfilesRecipients `json:"recipients"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewInlineObject203 instantiates a new InlineObject203 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject203(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string) *InlineObject203 {
	this := InlineObject203{}
	this.Type = type_
	this.AlertCondition = alertCondition
	this.Recipients = recipients
	this.NetworkTags = networkTags
	return &this
}

// NewInlineObject203WithDefaults instantiates a new InlineObject203 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject203WithDefaults() *InlineObject203 {
	this := InlineObject203{}
	return &this
}

// GetType returns the Type field value
func (o *InlineObject203) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject203) SetType(v string) {
	o.Type = v
}

// GetAlertCondition returns the AlertCondition field value
func (o *InlineObject203) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition {
	if o == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesAlertCondition
		return ret
	}

	return o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlertCondition, true
}

// SetAlertCondition sets field value
func (o *InlineObject203) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition) {
	o.AlertCondition = v
}

// GetRecipients returns the Recipients field value
func (o *InlineObject203) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients {
	if o == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *InlineObject203) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients) {
	o.Recipients = v
}

// GetNetworkTags returns the NetworkTags field value
func (o *InlineObject203) GetNetworkTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetNetworkTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NetworkTags, true
}

// SetNetworkTags sets field value
func (o *InlineObject203) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject203) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject203) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject203) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject203) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["alertCondition"] = o.AlertCondition
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["networkTags"] = o.NetworkTags
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject203 struct {
	value *InlineObject203
	isSet bool
}

func (v NullableInlineObject203) Get() *InlineObject203 {
	return v.value
}

func (v *NullableInlineObject203) Set(val *InlineObject203) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject203) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject203) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject203(val *InlineObject203) *NullableInlineObject203 {
	return &NullableInlineObject203{value: val, isSet: true}
}

func (v NullableInlineObject203) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject203) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


