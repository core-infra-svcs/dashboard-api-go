/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage Properties for setting a new image.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage struct {
	// The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
	// The file contents (a base 64 encoded string) of your new prepaid front.
	Contents *string `json:"contents,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImageWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImageWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) SetFormat(v string) {
	o.Format = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) GetContents() string {
	if o == nil || isNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) GetContentsOk() (*string, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) SetContents(v string) {
	o.Contents = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashPrepaidFrontImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


