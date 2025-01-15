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

// InlineResponse20059AllowedFiles struct for InlineResponse20059AllowedFiles
type InlineResponse20059AllowedFiles struct {
	// The sha256 digest of allowed file
	Sha256 *string `json:"sha256,omitempty"`
	// Comment about the allowed file
	Comment *string `json:"comment,omitempty"`
}

// NewInlineResponse20059AllowedFiles instantiates a new InlineResponse20059AllowedFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20059AllowedFiles() *InlineResponse20059AllowedFiles {
	this := InlineResponse20059AllowedFiles{}
	return &this
}

// NewInlineResponse20059AllowedFilesWithDefaults instantiates a new InlineResponse20059AllowedFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20059AllowedFilesWithDefaults() *InlineResponse20059AllowedFiles {
	this := InlineResponse20059AllowedFiles{}
	return &this
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *InlineResponse20059AllowedFiles) GetSha256() string {
	if o == nil || isNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059AllowedFiles) GetSha256Ok() (*string, bool) {
	if o == nil || isNil(o.Sha256) {
    return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *InlineResponse20059AllowedFiles) HasSha256() bool {
	if o != nil && !isNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *InlineResponse20059AllowedFiles) SetSha256(v string) {
	o.Sha256 = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20059AllowedFiles) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20059AllowedFiles) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20059AllowedFiles) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20059AllowedFiles) SetComment(v string) {
	o.Comment = &v
}

func (o InlineResponse20059AllowedFiles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20059AllowedFiles struct {
	value *InlineResponse20059AllowedFiles
	isSet bool
}

func (v NullableInlineResponse20059AllowedFiles) Get() *InlineResponse20059AllowedFiles {
	return v.value
}

func (v *NullableInlineResponse20059AllowedFiles) Set(val *InlineResponse20059AllowedFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20059AllowedFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20059AllowedFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20059AllowedFiles(val *InlineResponse20059AllowedFiles) *NullableInlineResponse20059AllowedFiles {
	return &NullableInlineResponse20059AllowedFiles{value: val, isSet: true}
}

func (v NullableInlineResponse20059AllowedFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20059AllowedFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

