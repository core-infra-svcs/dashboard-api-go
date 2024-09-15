/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage The image used in the splash page.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage struct {
	// The MD5 value of the image file. Setting this to null will remove the image from the splash page.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the image file.
	Extension *string `json:"extension,omitempty"`
	Image *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage `json:"image,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetExtension() string {
	if o == nil || isNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetExtensionOk() (*string, bool) {
	if o == nil || isNil(o.Extension) {
    return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) HasExtension() bool {
	if o != nil && !isNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) SetExtension(v string) {
	o.Extension = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetImage() NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage {
	if o == nil || isNil(o.Image) {
		var ret NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) GetImageOk() (*NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage and assigns it to the Image field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) SetImage(v NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) {
	o.Image = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


