/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20063 struct for InlineResponse20063
type InlineResponse20063 struct {
	// Current status of malware prevention
	Mode *string `json:"mode,omitempty"`
	// URLs permitted by the malware detection engine
	AllowedUrls []InlineResponse20063AllowedUrls `json:"allowedUrls,omitempty"`
	// Sha256 digests of files permitted by the malware detection engine
	AllowedFiles []InlineResponse20063AllowedFiles `json:"allowedFiles,omitempty"`
}

// NewInlineResponse20063 instantiates a new InlineResponse20063 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063WithDefaults() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse20063) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse20063) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse20063) SetMode(v string) {
	o.Mode = &v
}

// GetAllowedUrls returns the AllowedUrls field value if set, zero value otherwise.
func (o *InlineResponse20063) GetAllowedUrls() []InlineResponse20063AllowedUrls {
	if o == nil || isNil(o.AllowedUrls) {
		var ret []InlineResponse20063AllowedUrls
		return ret
	}
	return o.AllowedUrls
}

// GetAllowedUrlsOk returns a tuple with the AllowedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetAllowedUrlsOk() ([]InlineResponse20063AllowedUrls, bool) {
	if o == nil || isNil(o.AllowedUrls) {
    return nil, false
	}
	return o.AllowedUrls, true
}

// HasAllowedUrls returns a boolean if a field has been set.
func (o *InlineResponse20063) HasAllowedUrls() bool {
	if o != nil && !isNil(o.AllowedUrls) {
		return true
	}

	return false
}

// SetAllowedUrls gets a reference to the given []InlineResponse20063AllowedUrls and assigns it to the AllowedUrls field.
func (o *InlineResponse20063) SetAllowedUrls(v []InlineResponse20063AllowedUrls) {
	o.AllowedUrls = v
}

// GetAllowedFiles returns the AllowedFiles field value if set, zero value otherwise.
func (o *InlineResponse20063) GetAllowedFiles() []InlineResponse20063AllowedFiles {
	if o == nil || isNil(o.AllowedFiles) {
		var ret []InlineResponse20063AllowedFiles
		return ret
	}
	return o.AllowedFiles
}

// GetAllowedFilesOk returns a tuple with the AllowedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetAllowedFilesOk() ([]InlineResponse20063AllowedFiles, bool) {
	if o == nil || isNil(o.AllowedFiles) {
    return nil, false
	}
	return o.AllowedFiles, true
}

// HasAllowedFiles returns a boolean if a field has been set.
func (o *InlineResponse20063) HasAllowedFiles() bool {
	if o != nil && !isNil(o.AllowedFiles) {
		return true
	}

	return false
}

// SetAllowedFiles gets a reference to the given []InlineResponse20063AllowedFiles and assigns it to the AllowedFiles field.
func (o *InlineResponse20063) SetAllowedFiles(v []InlineResponse20063AllowedFiles) {
	o.AllowedFiles = v
}

func (o InlineResponse20063) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20063 struct {
	value *InlineResponse20063
	isSet bool
}

func (v NullableInlineResponse20063) Get() *InlineResponse20063 {
	return v.value
}

func (v *NullableInlineResponse20063) Set(val *InlineResponse20063) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063(val *InlineResponse20063) *NullableInlineResponse20063 {
	return &NullableInlineResponse20063{value: val, isSet: true}
}

func (v NullableInlineResponse20063) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


