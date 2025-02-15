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

// InlineObject59 struct for InlineObject59
type InlineObject59 struct {
	// Set mode to 'enabled' to enable malware prevention, otherwise 'disabled'
	Mode string `json:"mode"`
	// The urls that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedUrls []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls `json:"allowedUrls,omitempty"`
	// The sha256 digests of files that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedFiles []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"`
}

// NewInlineObject59 instantiates a new InlineObject59 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject59(mode string) *InlineObject59 {
	this := InlineObject59{}
	this.Mode = mode
	return &this
}

// NewInlineObject59WithDefaults instantiates a new InlineObject59 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject59WithDefaults() *InlineObject59 {
	this := InlineObject59{}
	return &this
}

// GetMode returns the Mode field value
func (o *InlineObject59) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InlineObject59) SetMode(v string) {
	o.Mode = v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise.
func (o *InlineObject59) GetAllowedUrls() []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls {
	if o == nil || isNil(o.AllowedUrls) {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls
		return ret
	}
	return o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetAllowedUrlsOk() ([]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls, bool) {
	if o == nil || isNil(o.AllowedUrls) {
    return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *InlineObject59) HasAllowedUrls() bool {
	if o != nil && !isNil(o.AllowedUrls) {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls and assigns it to the AllowedUrls field.
func (o *InlineObject59) SetAllowedUrls(v []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls) {
	o.AllowedUrls = v
}

// GetAllowedFiles returns the AllowedFiles field value if set, zero value otherwise.
func (o *InlineObject59) GetAllowedFiles() []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles {
	if o == nil || isNil(o.AllowedFiles) {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles
		return ret
	}
	return o.AllowedFiles
}

// GetAllowedFilesOk returns a tuple with the AllowedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject59) GetAllowedFilesOk() ([]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles, bool) {
	if o == nil || isNil(o.AllowedFiles) {
    return nil, false
	}
	return o.AllowedFiles, true
}

// HasAllowedFiles returns a boolean if a field has been set.
func (o *InlineObject59) HasAllowedFiles() bool {
	if o != nil && !isNil(o.AllowedFiles) {
		return true
	}

	return false
}

// SetAllowedFiles gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles and assigns it to the AllowedFiles field.
func (o *InlineObject59) SetAllowedFiles(v []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles) {
	o.AllowedFiles = v
}

func (o InlineObject59) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.AllowedUrls) {
		toSerialize["allowedUrls"] = o.AllowedUrls
	}
	if !isNil(o.AllowedFiles) {
		toSerialize["allowedFiles"] = o.AllowedFiles
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject59 struct {
	value *InlineObject59
	isSet bool
}

func (v NullableInlineObject59) Get() *InlineObject59 {
	return v.value
}

func (v *NullableInlineObject59) Set(val *InlineObject59) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject59) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject59) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject59(val *InlineObject59) *NullableInlineObject59 {
	return &NullableInlineObject59{value: val, isSet: true}
}

func (v NullableInlineObject59) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject59) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


