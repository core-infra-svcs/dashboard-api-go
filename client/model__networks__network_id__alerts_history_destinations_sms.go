/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDestinationsSms sms destinations for this alert
type NetworksNetworkIdAlertsHistoryDestinationsSms struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDestinationsSms instantiates a new NetworksNetworkIdAlertsHistoryDestinationsSms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDestinationsSms() *NetworksNetworkIdAlertsHistoryDestinationsSms {
	this := NetworksNetworkIdAlertsHistoryDestinationsSms{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDestinationsSmsWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDestinationsSms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDestinationsSmsWithDefaults() *NetworksNetworkIdAlertsHistoryDestinationsSms {
	this := NetworksNetworkIdAlertsHistoryDestinationsSms{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinationsSms) GetSentAt() string {
	if o == nil || isNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsSms) GetSentAtOk() (*string, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsSms) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *NetworksNetworkIdAlertsHistoryDestinationsSms) SetSentAt(v string) {
	o.SentAt = &v
}

func (o NetworksNetworkIdAlertsHistoryDestinationsSms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDestinationsSms struct {
	value *NetworksNetworkIdAlertsHistoryDestinationsSms
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsSms) Get() *NetworksNetworkIdAlertsHistoryDestinationsSms {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsSms) Set(val *NetworksNetworkIdAlertsHistoryDestinationsSms) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsSms) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsSms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDestinationsSms(val *NetworksNetworkIdAlertsHistoryDestinationsSms) *NullableNetworksNetworkIdAlertsHistoryDestinationsSms {
	return &NullableNetworksNetworkIdAlertsHistoryDestinationsSms{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsSms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsSms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


