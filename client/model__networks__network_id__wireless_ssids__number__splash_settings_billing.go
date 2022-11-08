/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling Details associated with billing splash.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling struct {
	FreeAccess *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess `json:"freeAccess,omitempty"`
	// Whether or not billing uses the fast login prepaid access option.
	PrepaidAccessFastLoginEnabled *bool `json:"prepaidAccessFastLoginEnabled,omitempty"`
	// The email address that receives replies from clients.
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling{}
	return &this
}

// GetFreeAccess returns the FreeAccess field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetFreeAccess() NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess {
	if o == nil || isNil(o.FreeAccess) {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess
		return ret
	}
	return *o.FreeAccess
}

// GetFreeAccessOk returns a tuple with the FreeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetFreeAccessOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess, bool) {
	if o == nil || isNil(o.FreeAccess) {
    return nil, false
	}
	return o.FreeAccess, true
}

// HasFreeAccess returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) HasFreeAccess() bool {
	if o != nil && !isNil(o.FreeAccess) {
		return true
	}

	return false
}

// SetFreeAccess gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess and assigns it to the FreeAccess field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) SetFreeAccess(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsBillingFreeAccess) {
	o.FreeAccess = &v
}

// GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetPrepaidAccessFastLoginEnabled() bool {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.PrepaidAccessFastLoginEnabled
}

// GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
    return nil, false
	}
	return o.PrepaidAccessFastLoginEnabled, true
}

// HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) HasPrepaidAccessFastLoginEnabled() bool {
	if o != nil && !isNil(o.PrepaidAccessFastLoginEnabled) {
		return true
	}

	return false
}

// SetPrepaidAccessFastLoginEnabled gets a reference to the given bool and assigns it to the PrepaidAccessFastLoginEnabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) SetPrepaidAccessFastLoginEnabled(v bool) {
	o.PrepaidAccessFastLoginEnabled = &v
}

// GetReplyToEmailAddress returns the ReplyToEmailAddress field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetReplyToEmailAddress() string {
	if o == nil || isNil(o.ReplyToEmailAddress) {
		var ret string
		return ret
	}
	return *o.ReplyToEmailAddress
}

// GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) GetReplyToEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.ReplyToEmailAddress) {
    return nil, false
	}
	return o.ReplyToEmailAddress, true
}

// HasReplyToEmailAddress returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) HasReplyToEmailAddress() bool {
	if o != nil && !isNil(o.ReplyToEmailAddress) {
		return true
	}

	return false
}

// SetReplyToEmailAddress gets a reference to the given string and assigns it to the ReplyToEmailAddress field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) SetReplyToEmailAddress(v string) {
	o.ReplyToEmailAddress = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FreeAccess) {
		toSerialize["freeAccess"] = o.FreeAccess
	}
	if !isNil(o.PrepaidAccessFastLoginEnabled) {
		toSerialize["prepaidAccessFastLoginEnabled"] = o.PrepaidAccessFastLoginEnabled
	}
	if !isNil(o.ReplyToEmailAddress) {
		toSerialize["replyToEmailAddress"] = o.ReplyToEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


