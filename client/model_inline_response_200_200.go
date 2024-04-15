/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200200 struct for InlineResponse200200
type InlineResponse200200 struct {
	// The alert config ID
	Id *string `json:"id,omitempty"`
	// The alert type
	Type *string `json:"type,omitempty"`
	// Is the alert config enabled
	Enabled *bool `json:"enabled,omitempty"`
	AlertCondition *OrganizationsOrganizationIdAlertsProfilesAlertCondition `json:"alertCondition,omitempty"`
	Recipients *OrganizationsOrganizationIdAlertsProfilesRecipients `json:"recipients,omitempty"`
	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags,omitempty"`
	// User supplied description of the alert
	Description *string `json:"description,omitempty"`
}

// NewInlineResponse200200 instantiates a new InlineResponse200200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// NewInlineResponse200200WithDefaults instantiates a new InlineResponse200200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200WithDefaults() *InlineResponse200200 {
	this := InlineResponse200200{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200200) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200200) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200200) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200200) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200200) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200200) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200200) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200200) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200200) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise.
func (o *InlineResponse200200) GetAlertCondition() OrganizationsOrganizationIdAlertsProfilesAlertCondition {
	if o == nil || isNil(o.AlertCondition) {
		var ret OrganizationsOrganizationIdAlertsProfilesAlertCondition
		return ret
	}
	return *o.AlertCondition
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetAlertConditionOk() (*OrganizationsOrganizationIdAlertsProfilesAlertCondition, bool) {
	if o == nil || isNil(o.AlertCondition) {
    return nil, false
	}
	return o.AlertCondition, true
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *InlineResponse200200) HasAlertCondition() bool {
	if o != nil && !isNil(o.AlertCondition) {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given OrganizationsOrganizationIdAlertsProfilesAlertCondition and assigns it to the AlertCondition field.
func (o *InlineResponse200200) SetAlertCondition(v OrganizationsOrganizationIdAlertsProfilesAlertCondition) {
	o.AlertCondition = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InlineResponse200200) GetRecipients() OrganizationsOrganizationIdAlertsProfilesRecipients {
	if o == nil || isNil(o.Recipients) {
		var ret OrganizationsOrganizationIdAlertsProfilesRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetRecipientsOk() (*OrganizationsOrganizationIdAlertsProfilesRecipients, bool) {
	if o == nil || isNil(o.Recipients) {
    return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InlineResponse200200) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given OrganizationsOrganizationIdAlertsProfilesRecipients and assigns it to the Recipients field.
func (o *InlineResponse200200) SetRecipients(v OrganizationsOrganizationIdAlertsProfilesRecipients) {
	o.Recipients = &v
}

// GetNetworkTags returns the NetworkTags field value if set, zero value otherwise.
func (o *InlineResponse200200) GetNetworkTags() []string {
	if o == nil || isNil(o.NetworkTags) {
		var ret []string
		return ret
	}
	return o.NetworkTags
}

// GetNetworkTagsOk returns a tuple with the NetworkTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetNetworkTagsOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkTags) {
    return nil, false
	}
	return o.NetworkTags, true
}

// HasNetworkTags returns a boolean if a field has been set.
func (o *InlineResponse200200) HasNetworkTags() bool {
	if o != nil && !isNil(o.NetworkTags) {
		return true
	}

	return false
}

// SetNetworkTags gets a reference to the given []string and assigns it to the NetworkTags field.
func (o *InlineResponse200200) SetNetworkTags(v []string) {
	o.NetworkTags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200200) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200200) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200200) SetDescription(v string) {
	o.Description = &v
}

func (o InlineResponse200200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
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

type NullableInlineResponse200200 struct {
	value *InlineResponse200200
	isSet bool
}

func (v NullableInlineResponse200200) Get() *InlineResponse200200 {
	return v.value
}

func (v *NullableInlineResponse200200) Set(val *InlineResponse200200) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200(val *InlineResponse200200) *NullableInlineResponse200200 {
	return &NullableInlineResponse200200{value: val, isSet: true}
}

func (v NullableInlineResponse200200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


