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

// InlineResponse200274 struct for InlineResponse200274
type InlineResponse200274 struct {
	// Id of packet capture file
	CaptureId *string `json:"captureId,omitempty"`
	// Download URL of captured packet file -- Depracated
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// Download URL of captured packet file
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse200274 instantiates a new InlineResponse200274 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// NewInlineResponse200274WithDefaults instantiates a new InlineResponse200274 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274WithDefaults() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// GetCaptureId returns the CaptureId field value if set, zero value otherwise.
func (o *InlineResponse200274) GetCaptureId() string {
	if o == nil || isNil(o.CaptureId) {
		var ret string
		return ret
	}
	return *o.CaptureId
}

// GetCaptureIdOk returns a tuple with the CaptureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetCaptureIdOk() (*string, bool) {
	if o == nil || isNil(o.CaptureId) {
    return nil, false
	}
	return o.CaptureId, true
}

// HasCaptureId returns a boolean if a field has been set.
func (o *InlineResponse200274) HasCaptureId() bool {
	if o != nil && !isNil(o.CaptureId) {
		return true
	}

	return false
}

// SetCaptureId gets a reference to the given string and assigns it to the CaptureId field.
func (o *InlineResponse200274) SetCaptureId(v string) {
	o.CaptureId = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *InlineResponse200274) GetDownloadUrl() string {
	if o == nil || isNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetDownloadUrlOk() (*string, bool) {
	if o == nil || isNil(o.DownloadUrl) {
    return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *InlineResponse200274) HasDownloadUrl() bool {
	if o != nil && !isNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *InlineResponse200274) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200274) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200274) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200274) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse200274) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaptureId) {
		toSerialize["captureId"] = o.CaptureId
	}
	if !isNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274 struct {
	value *InlineResponse200274
	isSet bool
}

func (v NullableInlineResponse200274) Get() *InlineResponse200274 {
	return v.value
}

func (v *NullableInlineResponse200274) Set(val *InlineResponse200274) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274(val *InlineResponse200274) *NullableInlineResponse200274 {
	return &NullableInlineResponse200274{value: val, isSet: true}
}

func (v NullableInlineResponse200274) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


