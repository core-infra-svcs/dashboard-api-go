/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdAlertsHistoryDestinations the destinations this alert is configured to be delivered to
type NetworksNetworkIdAlertsHistoryDestinations struct {
	Email *NetworksNetworkIdAlertsHistoryDestinationsEmail `json:"email,omitempty"`
	Push *NetworksNetworkIdAlertsHistoryDestinationsPush `json:"push,omitempty"`
	Sms *NetworksNetworkIdAlertsHistoryDestinationsSms `json:"sms,omitempty"`
	Webhook *NetworksNetworkIdAlertsHistoryDestinationsWebhook `json:"webhook,omitempty"`
}

// NewNetworksNetworkIdAlertsHistoryDestinations instantiates a new NetworksNetworkIdAlertsHistoryDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsHistoryDestinations() *NetworksNetworkIdAlertsHistoryDestinations {
	this := NetworksNetworkIdAlertsHistoryDestinations{}
	return &this
}

// NewNetworksNetworkIdAlertsHistoryDestinationsWithDefaults instantiates a new NetworksNetworkIdAlertsHistoryDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsHistoryDestinationsWithDefaults() *NetworksNetworkIdAlertsHistoryDestinations {
	this := NetworksNetworkIdAlertsHistoryDestinations{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetEmail() NetworksNetworkIdAlertsHistoryDestinationsEmail {
	if o == nil || isNil(o.Email) {
		var ret NetworksNetworkIdAlertsHistoryDestinationsEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetEmailOk() (*NetworksNetworkIdAlertsHistoryDestinationsEmail, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NetworksNetworkIdAlertsHistoryDestinationsEmail and assigns it to the Email field.
func (o *NetworksNetworkIdAlertsHistoryDestinations) SetEmail(v NetworksNetworkIdAlertsHistoryDestinationsEmail) {
	o.Email = &v
}

// GetPush returns the Push field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetPush() NetworksNetworkIdAlertsHistoryDestinationsPush {
	if o == nil || isNil(o.Push) {
		var ret NetworksNetworkIdAlertsHistoryDestinationsPush
		return ret
	}
	return *o.Push
}

// GetPushOk returns a tuple with the Push field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetPushOk() (*NetworksNetworkIdAlertsHistoryDestinationsPush, bool) {
	if o == nil || isNil(o.Push) {
    return nil, false
	}
	return o.Push, true
}

// HasPush returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) HasPush() bool {
	if o != nil && !isNil(o.Push) {
		return true
	}

	return false
}

// SetPush gets a reference to the given NetworksNetworkIdAlertsHistoryDestinationsPush and assigns it to the Push field.
func (o *NetworksNetworkIdAlertsHistoryDestinations) SetPush(v NetworksNetworkIdAlertsHistoryDestinationsPush) {
	o.Push = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetSms() NetworksNetworkIdAlertsHistoryDestinationsSms {
	if o == nil || isNil(o.Sms) {
		var ret NetworksNetworkIdAlertsHistoryDestinationsSms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetSmsOk() (*NetworksNetworkIdAlertsHistoryDestinationsSms, bool) {
	if o == nil || isNil(o.Sms) {
    return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) HasSms() bool {
	if o != nil && !isNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given NetworksNetworkIdAlertsHistoryDestinationsSms and assigns it to the Sms field.
func (o *NetworksNetworkIdAlertsHistoryDestinations) SetSms(v NetworksNetworkIdAlertsHistoryDestinationsSms) {
	o.Sms = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetWebhook() NetworksNetworkIdAlertsHistoryDestinationsWebhook {
	if o == nil || isNil(o.Webhook) {
		var ret NetworksNetworkIdAlertsHistoryDestinationsWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) GetWebhookOk() (*NetworksNetworkIdAlertsHistoryDestinationsWebhook, bool) {
	if o == nil || isNil(o.Webhook) {
    return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsHistoryDestinations) HasWebhook() bool {
	if o != nil && !isNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NetworksNetworkIdAlertsHistoryDestinationsWebhook and assigns it to the Webhook field.
func (o *NetworksNetworkIdAlertsHistoryDestinations) SetWebhook(v NetworksNetworkIdAlertsHistoryDestinationsWebhook) {
	o.Webhook = &v
}

func (o NetworksNetworkIdAlertsHistoryDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Push) {
		toSerialize["push"] = o.Push
	}
	if !isNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}
	if !isNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsHistoryDestinations struct {
	value *NetworksNetworkIdAlertsHistoryDestinations
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinations) Get() *NetworksNetworkIdAlertsHistoryDestinations {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinations) Set(val *NetworksNetworkIdAlertsHistoryDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsHistoryDestinations(val *NetworksNetworkIdAlertsHistoryDestinations) *NullableNetworksNetworkIdAlertsHistoryDestinations {
	return &NullableNetworksNetworkIdAlertsHistoryDestinations{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsHistoryDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsHistoryDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


