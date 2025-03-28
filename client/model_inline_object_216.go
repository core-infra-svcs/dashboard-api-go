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

// InlineObject216 struct for InlineObject216
type InlineObject216 struct {
	// The alert type
	Type string `json:"type"`
	AlertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition1 `json:"alertCondition"`
	Recipients OrganizationsOrganizationIdAlertsProfilesRecipients `json:"recipients"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewInlineObject216 instantiates a new InlineObject216 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject216(type_ string, alertCondition OrganizationsOrganizationIdAlertsProfilesAlertCondition1, recipients OrganizationsOrganizationIdAlertsProfilesRecipients, networkTags []string) *InlineObject216 {
	this := InlineObject216{}
	this.Type = type_
	this.AlertCondition = alertCondition
	this.Recipients = recipients
	this.NetworkTags = networkTags
	return &this
}

// NewInlineObject216WithDefaults instantiates a new InlineObject216 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject216WithDefaults() *InlineObject216 {
	this := InlineObject216{}
	return &this
}

// GetType returns the Type field value
func (o *InlineObject216) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject216) SetType(v string) {
	o.Type = v
}

// GetAlertCondition returns the AlertCondition field value
func (o *InlineObject216) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition1 {
	if o == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesAlertCondition1
		return ret
	}

	return o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlertCondition, true
}

// SetAlertCondition sets field value
func (o *InlineObject216) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition1) {
	o.AlertCondition = v
}

// GetRecipients returns the Recipients field value
func (o *InlineObject216) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients {
	if o == nil {
		var ret OrganizationsOrganizationIdAlertsProfilesRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *InlineObject216) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients) {
	o.Recipients = v
}

// GetNetworkTags returns the NetworkTags field value
func (o *InlineObject216) GetNetworkTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetNetworkTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NetworkTags, true
}

// SetNetworkTags sets field value
func (o *InlineObject216) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject216) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject216) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject216) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject216) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject216) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject216 struct {
	value *InlineObject216
	isSet bool
}

func (v NullableInlineObject216) Get() *InlineObject216 {
	return v.value
}

func (v *NullableInlineObject216) Set(val *InlineObject216) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject216) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject216) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject216(val *InlineObject216) *NullableInlineObject216 {
	return &NullableInlineObject216{value: val, isSet: true}
}

func (v NullableInlineObject216) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject216) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


