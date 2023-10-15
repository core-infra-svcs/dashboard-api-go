/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDestinationsWebhook webhook destinations for this alert
type NetworksNetworkIdAlertsHistoryDestinationsWebhook struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDestinationsWebhook instantiates a new NetworksNetworkIdAlertsHistoryDestinationsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDestinationsWebhook() *NetworksNetworkIdAlertsHistoryDestinationsWebhook {
	this := NetworksNetworkIdAlertsHistoryDestinationsWebhook{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDestinationsWebhookWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDestinationsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDestinationsWebhookWithDefaults() *NetworksNetworkIdAlertsHistoryDestinationsWebhook {
	this := NetworksNetworkIdAlertsHistoryDestinationsWebhook{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinationsWebhook) GetSentAt() string {
	if o == nil || isNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsWebhook) GetSentAtOk() (*string, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsWebhook) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *NetworksNetworkIdAlertsHistoryDestinationsWebhook) SetSentAt(v string) {
	o.SentAt = &v
}

func (o NetworksNetworkIdAlertsHistoryDestinationsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook struct {
	value *NetworksNetworkIdAlertsHistoryDestinationsWebhook
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) Get() *NetworksNetworkIdAlertsHistoryDestinationsWebhook {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) Set(val *NetworksNetworkIdAlertsHistoryDestinationsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDestinationsWebhook(val *NetworksNetworkIdAlertsHistoryDestinationsWebhook) *NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook {
	return &NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


