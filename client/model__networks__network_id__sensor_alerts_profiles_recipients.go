/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesRecipients List of recipients that will receive the alert.
type NetworksNetworkIdSensorAlertsProfilesRecipients struct {
	// A list of emails that will receive information about the alert.
	Emails []string `json:"emails,omitempty"`
	// A list of SMS numbers that will receive information about the alert.
	SmsNumbers []string `json:"smsNumbers,omitempty"`
	// A list of webhook endpoint IDs that will receive information about the alert.
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesRecipients instantiates a new NetworksNetworkIdSensorAlertsProfilesRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesRecipients() *NetworksNetworkIdSensorAlertsProfilesRecipients {
	this := NetworksNetworkIdSensorAlertsProfilesRecipients{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesRecipientsWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesRecipientsWithDefaults() *NetworksNetworkIdSensorAlertsProfilesRecipients {
	this := NetworksNetworkIdSensorAlertsProfilesRecipients{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) SetEmails(v []string) {
	o.Emails = v
}

// GetSmsNumbers returns the SmsNumbers field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetSmsNumbers() []string {
	if o == nil || isNil(o.SmsNumbers) {
		var ret []string
		return ret
	}
	return o.SmsNumbers
}

// GetSmsNumbersOk returns a tuple with the SmsNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetSmsNumbersOk() ([]string, bool) {
	if o == nil || isNil(o.SmsNumbers) {
    return nil, false
	}
	return o.SmsNumbers, true
}

// HasSmsNumbers returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) HasSmsNumbers() bool {
	if o != nil && !isNil(o.SmsNumbers) {
		return true
	}

	return false
}

// SetSmsNumbers gets a reference to the given []string and assigns it to the SmsNumbers field.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) SetSmsNumbers(v []string) {
	o.SmsNumbers = v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetHttpServerIds() []string {
	if o == nil || isNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.HttpServerIds) {
    return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) HasHttpServerIds() bool {
	if o != nil && !isNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *NetworksNetworkIdSensorAlertsProfilesRecipients) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o NetworksNetworkIdSensorAlertsProfilesRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !isNil(o.SmsNumbers) {
		toSerialize["smsNumbers"] = o.SmsNumbers
	}
	if !isNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesRecipients struct {
	value *NetworksNetworkIdSensorAlertsProfilesRecipients
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesRecipients) Get() *NetworksNetworkIdSensorAlertsProfilesRecipients {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesRecipients) Set(val *NetworksNetworkIdSensorAlertsProfilesRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesRecipients(val *NetworksNetworkIdSensorAlertsProfilesRecipients) *NullableNetworksNetworkIdSensorAlertsProfilesRecipients {
	return &NullableNetworksNetworkIdSensorAlertsProfilesRecipients{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


