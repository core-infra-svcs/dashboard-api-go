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

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo The logo used in the splash page.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo struct {
	// The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the logo file.
	Extension *string `json:"extension,omitempty"`
	Image *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage `json:"image,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetExtension() string {
	if o == nil || isNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetExtensionOk() (*string, bool) {
	if o == nil || isNil(o.Extension) {
    return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) HasExtension() bool {
	if o != nil && !isNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) SetExtension(v string) {
	o.Extension = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetImage() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage {
	if o == nil || isNil(o.Image) {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) GetImageOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage and assigns it to the Image field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) SetImage(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) {
	o.Image = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !isNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


