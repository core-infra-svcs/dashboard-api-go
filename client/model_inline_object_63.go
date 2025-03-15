/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject63 struct for InlineObject63
type InlineObject63 struct {
	// The name of the SSID.
	Name *string `json:"name,omitempty"`
	// Whether or not the SSID is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode.
	DefaultVlanId *int32 `json:"defaultVlanId,omitempty"`
	// The association control method for the SSID ('open', 'psk', '8021x-meraki' or '8021x-radius').
	AuthMode *string `json:"authMode,omitempty"`
	// The passkey for the SSID. This param is only valid if the authMode is 'psk'.
	Psk *string `json:"psk,omitempty"`
	// The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is '8021x-radius'.
	RadiusServers []NetworksNetworkIdApplianceSsidsNumberRadiusServers `json:"radiusServers,omitempty"`
	// The psk encryption mode for the SSID ('wep' or 'wpa'). This param is only valid if the authMode is 'psk'.
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// The types of WPA encryption. ('WPA1 and WPA2', 'WPA2 only', 'WPA3 Transition Mode' or 'WPA3 only'). This param is only valid if (1) the authMode is 'psk' & the encryptionMode is 'wpa' OR (2) the authMode is '8021x-meraki' OR (3) the authMode is '8021x-radius'
	WpaEncryptionMode *string `json:"wpaEncryptionMode,omitempty"`
	// Boolean indicating whether the MX should advertise or hide this SSID.
	Visible *bool `json:"visible,omitempty"`
	DhcpEnforcedDeauthentication *NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication `json:"dhcpEnforcedDeauthentication,omitempty"`
	Dot11w *NetworksNetworkIdApplianceSsidsNumberDot11w `json:"dot11w,omitempty"`
}

// NewInlineObject63 instantiates a new InlineObject63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject63() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// NewInlineObject63WithDefaults instantiates a new InlineObject63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject63WithDefaults() *InlineObject63 {
	this := InlineObject63{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject63) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject63) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject63) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject63) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject63) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject63) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultVlanId returns the DefaultVlanId field value if set, zero value otherwise.
func (o *InlineObject63) GetDefaultVlanId() int32 {
	if o == nil || isNil(o.DefaultVlanId) {
		var ret int32
		return ret
	}
	return *o.DefaultVlanId
}

// GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetDefaultVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.DefaultVlanId) {
    return nil, false
	}
	return o.DefaultVlanId, true
}

// HasDefaultVlanId returns a boolean if a field has been set.
func (o *InlineObject63) HasDefaultVlanId() bool {
	if o != nil && !isNil(o.DefaultVlanId) {
		return true
	}

	return false
}

// SetDefaultVlanId gets a reference to the given int32 and assigns it to the DefaultVlanId field.
func (o *InlineObject63) SetDefaultVlanId(v int32) {
	o.DefaultVlanId = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *InlineObject63) GetAuthMode() string {
	if o == nil || isNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetAuthModeOk() (*string, bool) {
	if o == nil || isNil(o.AuthMode) {
    return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *InlineObject63) HasAuthMode() bool {
	if o != nil && !isNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *InlineObject63) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetPsk returns the Psk field value if set, zero value otherwise.
func (o *InlineObject63) GetPsk() string {
	if o == nil || isNil(o.Psk) {
		var ret string
		return ret
	}
	return *o.Psk
}

// GetPskOk returns a tuple with the Psk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetPskOk() (*string, bool) {
	if o == nil || isNil(o.Psk) {
    return nil, false
	}
	return o.Psk, true
}

// HasPsk returns a boolean if a field has been set.
func (o *InlineObject63) HasPsk() bool {
	if o != nil && !isNil(o.Psk) {
		return true
	}

	return false
}

// SetPsk gets a reference to the given string and assigns it to the Psk field.
func (o *InlineObject63) SetPsk(v string) {
	o.Psk = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *InlineObject63) GetRadiusServers() []NetworksNetworkIdApplianceSsidsNumberRadiusServers {
	if o == nil || isNil(o.RadiusServers) {
		var ret []NetworksNetworkIdApplianceSsidsNumberRadiusServers
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetRadiusServersOk() ([]NetworksNetworkIdApplianceSsidsNumberRadiusServers, bool) {
	if o == nil || isNil(o.RadiusServers) {
    return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *InlineObject63) HasRadiusServers() bool {
	if o != nil && !isNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []NetworksNetworkIdApplianceSsidsNumberRadiusServers and assigns it to the RadiusServers field.
func (o *InlineObject63) SetRadiusServers(v []NetworksNetworkIdApplianceSsidsNumberRadiusServers) {
	o.RadiusServers = v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *InlineObject63) GetEncryptionMode() string {
	if o == nil || isNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionMode) {
    return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *InlineObject63) HasEncryptionMode() bool {
	if o != nil && !isNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *InlineObject63) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetWpaEncryptionMode returns the WpaEncryptionMode field value if set, zero value otherwise.
func (o *InlineObject63) GetWpaEncryptionMode() string {
	if o == nil || isNil(o.WpaEncryptionMode) {
		var ret string
		return ret
	}
	return *o.WpaEncryptionMode
}

// GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetWpaEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.WpaEncryptionMode) {
    return nil, false
	}
	return o.WpaEncryptionMode, true
}

// HasWpaEncryptionMode returns a boolean if a field has been set.
func (o *InlineObject63) HasWpaEncryptionMode() bool {
	if o != nil && !isNil(o.WpaEncryptionMode) {
		return true
	}

	return false
}

// SetWpaEncryptionMode gets a reference to the given string and assigns it to the WpaEncryptionMode field.
func (o *InlineObject63) SetWpaEncryptionMode(v string) {
	o.WpaEncryptionMode = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *InlineObject63) GetVisible() bool {
	if o == nil || isNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetVisibleOk() (*bool, bool) {
	if o == nil || isNil(o.Visible) {
    return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *InlineObject63) HasVisible() bool {
	if o != nil && !isNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *InlineObject63) SetVisible(v bool) {
	o.Visible = &v
}

// GetDhcpEnforcedDeauthentication returns the DhcpEnforcedDeauthentication field value if set, zero value otherwise.
func (o *InlineObject63) GetDhcpEnforcedDeauthentication() NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication {
	if o == nil || isNil(o.DhcpEnforcedDeauthentication) {
		var ret NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication
		return ret
	}
	return *o.DhcpEnforcedDeauthentication
}

// GetDhcpEnforcedDeauthenticationOk returns a tuple with the DhcpEnforcedDeauthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetDhcpEnforcedDeauthenticationOk() (*NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication, bool) {
	if o == nil || isNil(o.DhcpEnforcedDeauthentication) {
    return nil, false
	}
	return o.DhcpEnforcedDeauthentication, true
}

// HasDhcpEnforcedDeauthentication returns a boolean if a field has been set.
func (o *InlineObject63) HasDhcpEnforcedDeauthentication() bool {
	if o != nil && !isNil(o.DhcpEnforcedDeauthentication) {
		return true
	}

	return false
}

// SetDhcpEnforcedDeauthentication gets a reference to the given NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication and assigns it to the DhcpEnforcedDeauthentication field.
func (o *InlineObject63) SetDhcpEnforcedDeauthentication(v NetworksNetworkIdApplianceSsidsNumberDhcpEnforcedDeauthentication) {
	o.DhcpEnforcedDeauthentication = &v
}

// GetDot11w returns the Dot11w field value if set, zero value otherwise.
func (o *InlineObject63) GetDot11w() NetworksNetworkIdApplianceSsidsNumberDot11w {
	if o == nil || isNil(o.Dot11w) {
		var ret NetworksNetworkIdApplianceSsidsNumberDot11w
		return ret
	}
	return *o.Dot11w
}

// GetDot11wOk returns a tuple with the Dot11w field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject63) GetDot11wOk() (*NetworksNetworkIdApplianceSsidsNumberDot11w, bool) {
	if o == nil || isNil(o.Dot11w) {
    return nil, false
	}
	return o.Dot11w, true
}

// HasDot11w returns a boolean if a field has been set.
func (o *InlineObject63) HasDot11w() bool {
	if o != nil && !isNil(o.Dot11w) {
		return true
	}

	return false
}

// SetDot11w gets a reference to the given NetworksNetworkIdApplianceSsidsNumberDot11w and assigns it to the Dot11w field.
func (o *InlineObject63) SetDot11w(v NetworksNetworkIdApplianceSsidsNumberDot11w) {
	o.Dot11w = &v
}

func (o InlineObject63) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DefaultVlanId) {
		toSerialize["defaultVlanId"] = o.DefaultVlanId
	}
	if !isNil(o.AuthMode) {
		toSerialize["authMode"] = o.AuthMode
	}
	if !isNil(o.Psk) {
		toSerialize["psk"] = o.Psk
	}
	if !isNil(o.RadiusServers) {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !isNil(o.EncryptionMode) {
		toSerialize["encryptionMode"] = o.EncryptionMode
	}
	if !isNil(o.WpaEncryptionMode) {
		toSerialize["wpaEncryptionMode"] = o.WpaEncryptionMode
	}
	if !isNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !isNil(o.DhcpEnforcedDeauthentication) {
		toSerialize["dhcpEnforcedDeauthentication"] = o.DhcpEnforcedDeauthentication
	}
	if !isNil(o.Dot11w) {
		toSerialize["dot11w"] = o.Dot11w
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject63 struct {
	value *InlineObject63
	isSet bool
}

func (v NullableInlineObject63) Get() *InlineObject63 {
	return v.value
}

func (v *NullableInlineObject63) Set(val *InlineObject63) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject63) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject63(val *InlineObject63) *NullableInlineObject63 {
	return &NullableInlineObject63{value: val, isSet: true}
}

func (v NullableInlineObject63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


