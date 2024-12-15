/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdAlertsProfilesRecipients List of recipients that will recieve the alert.
type OrganizationsOrganizationIdAlertsProfilesRecipients struct {
	// A list of emails that will receive information about the alert
	Emails []string `json:"emails,omitempty"`
	// A list base64 encoded urls of webhook endpoints that will receive information about the alert
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewOrganizationsOrganizationIdAlertsProfilesRecipients instantiates a new OrganizationsOrganizationIdAlertsProfilesRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdAlertsProfilesRecipients() *OrganizationsOrganizationIdAlertsProfilesRecipients {
	this := OrganizationsOrganizationIdAlertsProfilesRecipients{}
	return &this
}

// NewOrganizationsOrganizationIdAlertsProfilesRecipientsWithDefaults instantiates a new OrganizationsOrganizationIdAlertsProfilesRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdAlertsProfilesRecipientsWithDefaults() *OrganizationsOrganizationIdAlertsProfilesRecipients {
	this := OrganizationsOrganizationIdAlertsProfilesRecipients{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) SetEmails(v []string) {
	o.Emails = v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) GetHttpServerIds() []string {
	if o == nil || isNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.HttpServerIds) {
    return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) HasHttpServerIds() bool {
	if o != nil && !isNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *OrganizationsOrganizationIdAlertsProfilesRecipients) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o OrganizationsOrganizationIdAlertsProfilesRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !isNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdAlertsProfilesRecipients struct {
	value *OrganizationsOrganizationIdAlertsProfilesRecipients
	isSet bool
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesRecipients) Get() *OrganizationsOrganizationIdAlertsProfilesRecipients {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesRecipients) Set(val *OrganizationsOrganizationIdAlertsProfilesRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdAlertsProfilesRecipients(val *OrganizationsOrganizationIdAlertsProfilesRecipients) *NullableOrganizationsOrganizationIdAlertsProfilesRecipients {
	return &NullableOrganizationsOrganizationIdAlertsProfilesRecipients{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdAlertsProfilesRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdAlertsProfilesRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


