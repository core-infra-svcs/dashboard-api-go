/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject98 struct for InlineObject98
type InlineObject98 struct {
	// Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)
	LocalStatusPageEnabled *bool `json:"localStatusPageEnabled,omitempty"`
	// Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true
	RemoteStatusPageEnabled *bool `json:"remoteStatusPageEnabled,omitempty"`
	SecureConnect *NetworksNetworkIdSettingsSecureConnect `json:"secureConnect,omitempty"`
	LocalStatusPage *NetworksNetworkIdSettingsLocalStatusPage `json:"localStatusPage,omitempty"`
}

// NewInlineObject98 instantiates a new InlineObject98 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject98() *InlineObject98 {
	this := InlineObject98{}
	return &this
}

// NewInlineObject98WithDefaults instantiates a new InlineObject98 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject98WithDefaults() *InlineObject98 {
	this := InlineObject98{}
	return &this
}

// GetLocalStatusPageEnabled returns the LocalStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineObject98) GetLocalStatusPageEnabled() bool {
	if o == nil || o.LocalStatusPageEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LocalStatusPageEnabled
}

// GetLocalStatusPageEnabledOk returns a tuple with the LocalStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject98) GetLocalStatusPageEnabledOk() (*bool, bool) {
	if o == nil || o.LocalStatusPageEnabled == nil {
		return nil, false
	}
	return o.LocalStatusPageEnabled, true
}

// HasLocalStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineObject98) HasLocalStatusPageEnabled() bool {
	if o != nil && o.LocalStatusPageEnabled != nil {
		return true
	}

	return false
}

// SetLocalStatusPageEnabled gets a reference to the given bool and assigns it to the LocalStatusPageEnabled field.
func (o *InlineObject98) SetLocalStatusPageEnabled(v bool) {
	o.LocalStatusPageEnabled = &v
}

// GetRemoteStatusPageEnabled returns the RemoteStatusPageEnabled field value if set, zero value otherwise.
func (o *InlineObject98) GetRemoteStatusPageEnabled() bool {
	if o == nil || o.RemoteStatusPageEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteStatusPageEnabled
}

// GetRemoteStatusPageEnabledOk returns a tuple with the RemoteStatusPageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject98) GetRemoteStatusPageEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteStatusPageEnabled == nil {
		return nil, false
	}
	return o.RemoteStatusPageEnabled, true
}

// HasRemoteStatusPageEnabled returns a boolean if a field has been set.
func (o *InlineObject98) HasRemoteStatusPageEnabled() bool {
	if o != nil && o.RemoteStatusPageEnabled != nil {
		return true
	}

	return false
}

// SetRemoteStatusPageEnabled gets a reference to the given bool and assigns it to the RemoteStatusPageEnabled field.
func (o *InlineObject98) SetRemoteStatusPageEnabled(v bool) {
	o.RemoteStatusPageEnabled = &v
}

// GetSecureConnect returns the SecureConnect field value if set, zero value otherwise.
func (o *InlineObject98) GetSecureConnect() NetworksNetworkIdSettingsSecureConnect {
	if o == nil || o.SecureConnect == nil {
		var ret NetworksNetworkIdSettingsSecureConnect
		return ret
	}
	return *o.SecureConnect
}

// GetSecureConnectOk returns a tuple with the SecureConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject98) GetSecureConnectOk() (*NetworksNetworkIdSettingsSecureConnect, bool) {
	if o == nil || o.SecureConnect == nil {
		return nil, false
	}
	return o.SecureConnect, true
}

// HasSecureConnect returns a boolean if a field has been set.
func (o *InlineObject98) HasSecureConnect() bool {
	if o != nil && o.SecureConnect != nil {
		return true
	}

	return false
}

// SetSecureConnect gets a reference to the given NetworksNetworkIdSettingsSecureConnect and assigns it to the SecureConnect field.
func (o *InlineObject98) SetSecureConnect(v NetworksNetworkIdSettingsSecureConnect) {
	o.SecureConnect = &v
}

// GetLocalStatusPage returns the LocalStatusPage field value if set, zero value otherwise.
func (o *InlineObject98) GetLocalStatusPage() NetworksNetworkIdSettingsLocalStatusPage {
	if o == nil || o.LocalStatusPage == nil {
		var ret NetworksNetworkIdSettingsLocalStatusPage
		return ret
	}
	return *o.LocalStatusPage
}

// GetLocalStatusPageOk returns a tuple with the LocalStatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject98) GetLocalStatusPageOk() (*NetworksNetworkIdSettingsLocalStatusPage, bool) {
	if o == nil || o.LocalStatusPage == nil {
		return nil, false
	}
	return o.LocalStatusPage, true
}

// HasLocalStatusPage returns a boolean if a field has been set.
func (o *InlineObject98) HasLocalStatusPage() bool {
	if o != nil && o.LocalStatusPage != nil {
		return true
	}

	return false
}

// SetLocalStatusPage gets a reference to the given NetworksNetworkIdSettingsLocalStatusPage and assigns it to the LocalStatusPage field.
func (o *InlineObject98) SetLocalStatusPage(v NetworksNetworkIdSettingsLocalStatusPage) {
	o.LocalStatusPage = &v
}

func (o InlineObject98) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalStatusPageEnabled != nil {
		toSerialize["localStatusPageEnabled"] = o.LocalStatusPageEnabled
	}
	if o.RemoteStatusPageEnabled != nil {
		toSerialize["remoteStatusPageEnabled"] = o.RemoteStatusPageEnabled
	}
	if o.SecureConnect != nil {
		toSerialize["secureConnect"] = o.SecureConnect
	}
	if o.LocalStatusPage != nil {
		toSerialize["localStatusPage"] = o.LocalStatusPage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject98 struct {
	value *InlineObject98
	isSet bool
}

func (v NullableInlineObject98) Get() *InlineObject98 {
	return v.value
}

func (v *NullableInlineObject98) Set(val *InlineObject98) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject98) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject98) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject98(val *InlineObject98) *NullableInlineObject98 {
	return &NullableInlineObject98{value: val, isSet: true}
}

func (v NullableInlineObject98) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject98) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


