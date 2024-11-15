/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDestinationsPush push destinations for this alert
type NetworksNetworkIdAlertsHistoryDestinationsPush struct {
	// time when the alert was sent to the user(s) for this channel
	SentAt *string `json:"sentAt,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDestinationsPush instantiates a new NetworksNetworkIdAlertsHistoryDestinationsPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDestinationsPush() *NetworksNetworkIdAlertsHistoryDestinationsPush {
	this := NetworksNetworkIdAlertsHistoryDestinationsPush{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDestinationsPushWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDestinationsPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDestinationsPushWithDefaults() *NetworksNetworkIdAlertsHistoryDestinationsPush {
	this := NetworksNetworkIdAlertsHistoryDestinationsPush{}
	return &this
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinationsPush) GetSentAt() string {
	if o == nil || isNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsPush) GetSentAtOk() (*string, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinationsPush) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *NetworksNetworkIdAlertsHistoryDestinationsPush) SetSentAt(v string) {
	o.SentAt = &v
}

func (o NetworksNetworkIdAlertsHistoryDestinationsPush) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDestinationsPush struct {
	value *NetworksNetworkIdAlertsHistoryDestinationsPush
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsPush) Get() *NetworksNetworkIdAlertsHistoryDestinationsPush {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsPush) Set(val *NetworksNetworkIdAlertsHistoryDestinationsPush) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsPush) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDestinationsPush(val *NetworksNetworkIdAlertsHistoryDestinationsPush) *NullableNetworksNetworkIdAlertsHistoryDestinationsPush {
	return &NullableNetworksNetworkIdAlertsHistoryDestinationsPush{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinationsPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinationsPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


