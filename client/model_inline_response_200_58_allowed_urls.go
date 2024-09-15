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

// InlineResponse20058AllowedUrls struct for InlineResponse20058AllowedUrls
type InlineResponse20058AllowedUrls struct {
	// The allowed URL
	Url *string `json:"url,omitempty"`
	// Comment about the allowed URL
	Comment *string `json:"comment,omitempty"`
}

// NewInlineResponse20058AllowedUrls instantiates a new InlineResponse20058AllowedUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20058AllowedUrls() *InlineResponse20058AllowedUrls {
	this := InlineResponse20058AllowedUrls{}
	return &this
}

// NewInlineResponse20058AllowedUrlsWithDefaults instantiates a new InlineResponse20058AllowedUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20058AllowedUrlsWithDefaults() *InlineResponse20058AllowedUrls {
	this := InlineResponse20058AllowedUrls{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20058AllowedUrls) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20058AllowedUrls) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20058AllowedUrls) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20058AllowedUrls) SetUrl(v string) {
	o.Url = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20058AllowedUrls) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20058AllowedUrls) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20058AllowedUrls) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20058AllowedUrls) SetComment(v string) {
	o.Comment = &v
}

func (o InlineResponse20058AllowedUrls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20058AllowedUrls struct {
	value *InlineResponse20058AllowedUrls
	isSet bool
}

func (v NullableInlineResponse20058AllowedUrls) Get() *InlineResponse20058AllowedUrls {
	return v.value
}

func (v *NullableInlineResponse20058AllowedUrls) Set(val *InlineResponse20058AllowedUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20058AllowedUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20058AllowedUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20058AllowedUrls(val *InlineResponse20058AllowedUrls) *NullableInlineResponse20058AllowedUrls {
	return &NullableInlineResponse20058AllowedUrls{value: val, isSet: true}
}

func (v NullableInlineResponse20058AllowedUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20058AllowedUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


