/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20057 struct for InlineResponse20057
type InlineResponse20057 struct {
	// Current status of malware prevention
	Mode *string `json:"mode,omitempty"`
	// URLs permitted by the malware detection engine
	AllowedUrls []InlineResponse20057AllowedUrls `json:"allowedUrls,omitempty"`
	// Sha256 digests of files permitted by the malware detection engine
	AllowedFiles []InlineResponse20057AllowedFiles `json:"allowedFiles,omitempty"`
}

// NewInlineResponse20057 instantiates a new InlineResponse20057 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20057() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// NewInlineResponse20057WithDefaults instantiates a new InlineResponse20057 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20057WithDefaults() *InlineResponse20057 {
	this := InlineResponse20057{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse20057) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse20057) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse20057) SetMode(v string) {
	o.Mode = &v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise.
func (o *InlineResponse20057) GetAllowedUrls() []InlineResponse20057AllowedUrls {
	if o == nil || isNil(o.AllowedUrls) {
		var ret []InlineResponse20057AllowedUrls
		return ret
	}
	return o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetAllowedUrlsOk() ([]InlineResponse20057AllowedUrls, bool) {
	if o == nil || isNil(o.AllowedUrls) {
    return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *InlineResponse20057) HasAllowedUrls() bool {
	if o != nil && !isNil(o.AllowedUrls) {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []InlineResponse20057AllowedUrls and assigns it to the AllowedUrls field.
func (o *InlineResponse20057) SetAllowedUrls(v []InlineResponse20057AllowedUrls) {
	o.AllowedUrls = v
}

// GetAllowedFiles returns the AllowedFiles field value if set, zero value otherwise.
func (o *InlineResponse20057) GetAllowedFiles() []InlineResponse20057AllowedFiles {
	if o == nil || isNil(o.AllowedFiles) {
		var ret []InlineResponse20057AllowedFiles
		return ret
	}
	return o.AllowedFiles
}

// GetAllowedFilesOk returns a tuple with the AllowedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20057) GetAllowedFilesOk() ([]InlineResponse20057AllowedFiles, bool) {
	if o == nil || isNil(o.AllowedFiles) {
    return nil, false
	}
	return o.AllowedFiles, true
}

// HasAllowedFiles returns a boolean if a field has been set.
func (o *InlineResponse20057) HasAllowedFiles() bool {
	if o != nil && !isNil(o.AllowedFiles) {
		return true
	}

	return false
}

// SetAllowedFiles gets a reference to the given []InlineResponse20057AllowedFiles and assigns it to the AllowedFiles field.
func (o *InlineResponse20057) SetAllowedFiles(v []InlineResponse20057AllowedFiles) {
	o.AllowedFiles = v
}

func (o InlineResponse20057) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
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

type NullableInlineResponse20057 struct {
	value *InlineResponse20057
	isSet bool
}

func (v NullableInlineResponse20057) Get() *InlineResponse20057 {
	return v.value
}

func (v *NullableInlineResponse20057) Set(val *InlineResponse20057) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20057) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20057) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20057(val *InlineResponse20057) *NullableInlineResponse20057 {
	return &NullableInlineResponse20057{value: val, isSet: true}
}

func (v NullableInlineResponse20057) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20057) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


