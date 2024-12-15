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

// InlineResponse20062 struct for InlineResponse20062
type InlineResponse20062 struct {
	// The number of the SSID.
	Number *int32 `json:"number,omitempty"`
	// The name of the SSID.
	Name *string `json:"name,omitempty"`
	// Whether or not the SSID is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The VLAN ID of the VLAN associated to this SSID.
	DefaultVlanId *int32 `json:"defaultVlanId,omitempty"`
	// The association control method for the SSID.
	AuthMode *string `json:"authMode,omitempty"`
	// The RADIUS 802.1x servers to be used for authentication.
	RadiusServers []NetworksNetworkIdApplianceSsidsRadiusServers `json:"radiusServers,omitempty"`
	// The psk encryption mode for the SSID.
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// WPA encryption mode for the SSID.
	WpaEncryptionMode *string `json:"wpaEncryptionMode,omitempty"`
	// Boolean indicating whether the MX should advertise or hide this SSID.
	Visible *bool `json:"visible,omitempty"`
}

// NewInlineResponse20062 instantiates a new InlineResponse20062 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20062() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// NewInlineResponse20062WithDefaults instantiates a new InlineResponse20062 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20062WithDefaults() *InlineResponse20062 {
	this := InlineResponse20062{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *InlineResponse20062) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineResponse20062) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *InlineResponse20062) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20062) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20062) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20062) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20062) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20062) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20062) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultVlanId returns the DefaultVlanId field value if set, zero value otherwise.
func (o *InlineResponse20062) GetDefaultVlanId() int32 {
	if o == nil || isNil(o.DefaultVlanId) {
		var ret int32
		return ret
	}
	return *o.DefaultVlanId
}

// GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetDefaultVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.DefaultVlanId) {
    return nil, false
	}
	return o.DefaultVlanId, true
}

// HasDefaultVlanId returns a boolean if a field has been set.
func (o *InlineResponse20062) HasDefaultVlanId() bool {
	if o != nil && !isNil(o.DefaultVlanId) {
		return true
	}

	return false
}

// SetDefaultVlanId gets a reference to the given int32 and assigns it to the DefaultVlanId field.
func (o *InlineResponse20062) SetDefaultVlanId(v int32) {
	o.DefaultVlanId = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *InlineResponse20062) GetAuthMode() string {
	if o == nil || isNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetAuthModeOk() (*string, bool) {
	if o == nil || isNil(o.AuthMode) {
    return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *InlineResponse20062) HasAuthMode() bool {
	if o != nil && !isNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *InlineResponse20062) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *InlineResponse20062) GetRadiusServers() []NetworksNetworkIdApplianceSsidsRadiusServers {
	if o == nil || isNil(o.RadiusServers) {
		var ret []NetworksNetworkIdApplianceSsidsRadiusServers
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetRadiusServersOk() ([]NetworksNetworkIdApplianceSsidsRadiusServers, bool) {
	if o == nil || isNil(o.RadiusServers) {
    return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *InlineResponse20062) HasRadiusServers() bool {
	if o != nil && !isNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []NetworksNetworkIdApplianceSsidsRadiusServers and assigns it to the RadiusServers field.
func (o *InlineResponse20062) SetRadiusServers(v []NetworksNetworkIdApplianceSsidsRadiusServers) {
	o.RadiusServers = v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *InlineResponse20062) GetEncryptionMode() string {
	if o == nil || isNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionMode) {
    return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *InlineResponse20062) HasEncryptionMode() bool {
	if o != nil && !isNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *InlineResponse20062) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetWpaEncryptionMode returns the WpaEncryptionMode field value if set, zero value otherwise.
func (o *InlineResponse20062) GetWpaEncryptionMode() string {
	if o == nil || isNil(o.WpaEncryptionMode) {
		var ret string
		return ret
	}
	return *o.WpaEncryptionMode
}

// GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetWpaEncryptionModeOk() (*string, bool) {
	if o == nil || isNil(o.WpaEncryptionMode) {
    return nil, false
	}
	return o.WpaEncryptionMode, true
}

// HasWpaEncryptionMode returns a boolean if a field has been set.
func (o *InlineResponse20062) HasWpaEncryptionMode() bool {
	if o != nil && !isNil(o.WpaEncryptionMode) {
		return true
	}

	return false
}

// SetWpaEncryptionMode gets a reference to the given string and assigns it to the WpaEncryptionMode field.
func (o *InlineResponse20062) SetWpaEncryptionMode(v string) {
	o.WpaEncryptionMode = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *InlineResponse20062) GetVisible() bool {
	if o == nil || isNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20062) GetVisibleOk() (*bool, bool) {
	if o == nil || isNil(o.Visible) {
    return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *InlineResponse20062) HasVisible() bool {
	if o != nil && !isNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *InlineResponse20062) SetVisible(v bool) {
	o.Visible = &v
}

func (o InlineResponse20062) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20062 struct {
	value *InlineResponse20062
	isSet bool
}

func (v NullableInlineResponse20062) Get() *InlineResponse20062 {
	return v.value
}

func (v *NullableInlineResponse20062) Set(val *InlineResponse20062) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20062) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20062) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20062(val *InlineResponse20062) *NullableInlineResponse20062 {
	return &NullableInlineResponse20062{value: val, isSet: true}
}

func (v NullableInlineResponse20062) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20062) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


