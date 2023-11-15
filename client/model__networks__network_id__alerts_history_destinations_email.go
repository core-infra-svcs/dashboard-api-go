/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDestinationsEmail email destinations for this alert
type NetworksNetworkIdAlertsHistoryDestinationsEmail struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDestinationsEmail instantiates a new NetworksNetworkIdAlertsHistoryDestinationsEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDestinationsEmail() *NetworksNetworkIdAlertsHistoryDestinationsEmail {
	this := NetworksNetworkIdAlertsHistoryDestinationsEmail{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDestinationsEmailWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDestinationsEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDestinationsEmailWithDefaults() *NetworksNetworkIdAlertsHistoryDestinationsEmail {
	this := NetworksNetworkIdAlertsHistoryDestinationsEmail{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinationsEmail) GetSentAt() string {
	if o == nil || isNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsEmail) GetSentAtOk() (*string, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsEmail) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *NetworksNetworkIdAlertsHistoryDestinationsEmail) SetSentAt(v string) {
	o.SentAt = &v
}

func (o NetworksNetworkIdAlertsHistoryDestinationsEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDestinationsEmail struct {
	value *NetworksNetworkIdAlertsHistoryDestinationsEmail
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) Get() *NetworksNetworkIdAlertsHistoryDestinationsEmail {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) Set(val *NetworksNetworkIdAlertsHistoryDestinationsEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDestinationsEmail(val *NetworksNetworkIdAlertsHistoryDestinationsEmail) *NullableNetworksNetworkIdAlertsHistoryDestinationsEmail {
	return &NullableNetworksNetworkIdAlertsHistoryDestinationsEmail{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


