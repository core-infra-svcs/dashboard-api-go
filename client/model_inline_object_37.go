/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	// Set mode to 'enabled' to enable malware prevention, otherwise 'disabled'
	Mode string `json:"mode"`
	// The urls that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedUrls *[]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls `json:"allowedUrls,omitempty"`
	// The sha256 digests of files that should be permitted by the malware detection engine. If omitted, the current config will remain unchanged. This is available only if your network supports AMP allow listing
	AllowedFiles *[]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles `json:"allowedFiles,omitempty"`
}

// NewInlineObject37 instantiates a new InlineObject37 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject37(mode string) *InlineObject37 {
	this := InlineObject37{}
	this.Mode = mode
	return &this
}

// NewInlineObject37WithDefaults instantiates a new InlineObject37 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject37WithDefaults() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// GetMode returns the Mode field value
func (o *InlineObject37) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InlineObject37) SetMode(v string) {
	o.Mode = v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise.
func (o *InlineObject37) GetAllowedUrls() []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls {
	if o == nil || o.AllowedUrls == nil {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls
		return ret
	}
	return *o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetAllowedUrlsOk() (*[]NetworksNetworkIdApplianceSecurityMalwareAllowedUrls, bool) {
	if o == nil || o.AllowedUrls == nil {
		return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *InlineObject37) HasAllowedUrls() bool {
	if o != nil && o.AllowedUrls != nil {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls and assigns it to the AllowedUrls field.
func (o *InlineObject37) SetAllowedUrls(v []NetworksNetworkIdApplianceSecurityMalwareAllowedUrls) {
	o.AllowedUrls = &v
}

// GetAllowedFiles returns the AllowedFiles field value if set, zero value otherwise.
func (o *InlineObject37) GetAllowedFiles() []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles {
	if o == nil || o.AllowedFiles == nil {
		var ret []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles
		return ret
	}
	return *o.AllowedFiles
}

// GetAllowedFilesOk returns a tuple with the AllowedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetAllowedFilesOk() (*[]NetworksNetworkIdApplianceSecurityMalwareAllowedFiles, bool) {
	if o == nil || o.AllowedFiles == nil {
		return nil, false
	}
	return o.AllowedFiles, true
}

// HasAllowedFiles returns a boolean if a field has been set.
func (o *InlineObject37) HasAllowedFiles() bool {
	if o != nil && o.AllowedFiles != nil {
		return true
	}

	return false
}

// SetAllowedFiles gets a reference to the given []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles and assigns it to the AllowedFiles field.
func (o *InlineObject37) SetAllowedFiles(v []NetworksNetworkIdApplianceSecurityMalwareAllowedFiles) {
	o.AllowedFiles = &v
}

func (o InlineObject37) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.AllowedUrls != nil {
		toSerialize["allowedUrls"] = o.AllowedUrls
	}
	if o.AllowedFiles != nil {
		toSerialize["allowedFiles"] = o.AllowedFiles
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject37 struct {
	value *InlineObject37
	isSet bool
}

func (v NullableInlineObject37) Get() *InlineObject37 {
	return v.value
}

func (v *NullableInlineObject37) Set(val *InlineObject37) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject37) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject37) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject37(val *InlineObject37) *NullableInlineObject37 {
	return &NullableInlineObject37{value: val, isSet: true}
}

func (v NullableInlineObject37) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject37) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


