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

// InlineObject40 struct for InlineObject40
type InlineObject40 struct {
	// A list of URL patterns that are allowed
	AllowedUrlPatterns []string `json:"allowedUrlPatterns,omitempty"`
	// A list of URL patterns that are blocked
	BlockedUrlPatterns []string `json:"blockedUrlPatterns,omitempty"`
	// A list of URL categories to block
	BlockedUrlCategories []string `json:"blockedUrlCategories,omitempty"`
	// URL category list size which is either 'topSites' or 'fullList'
	UrlCategoryListSize *string `json:"urlCategoryListSize,omitempty"`
}

// NewInlineObject40 instantiates a new InlineObject40 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject40() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// NewInlineObject40WithDefaults instantiates a new InlineObject40 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject40WithDefaults() *InlineObject40 {
	this := InlineObject40{}
	return &this
}

// GetAllowedUrlPatterns returns the AllowedUrlPatterns field value if set, zero value otherwise.
func (o *InlineObject40) GetAllowedUrlPatterns() []string {
	if o == nil || isNil(o.AllowedUrlPatterns) {
		var ret []string
		return ret
	}
	return o.AllowedUrlPatterns
}

// GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetAllowedUrlPatternsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedUrlPatterns) {
    return nil, false
	}
	return o.AllowedUrlPatterns, true
}

// HasAllowedUrlPatterns returns a boolean if a field has been set.
func (o *InlineObject40) HasAllowedUrlPatterns() bool {
	if o != nil && !isNil(o.AllowedUrlPatterns) {
		return true
	}

	return false
}

// SetAllowedUrlPatterns gets a reference to the given []string and assigns it to the AllowedUrlPatterns field.
func (o *InlineObject40) SetAllowedUrlPatterns(v []string) {
	o.AllowedUrlPatterns = v
}

// GetBlockedUrlPatterns returns the BlockedUrlPatterns field value if set, zero value otherwise.
func (o *InlineObject40) GetBlockedUrlPatterns() []string {
	if o == nil || isNil(o.BlockedUrlPatterns) {
		var ret []string
		return ret
	}
	return o.BlockedUrlPatterns
}

// GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetBlockedUrlPatternsOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedUrlPatterns) {
    return nil, false
	}
	return o.BlockedUrlPatterns, true
}

// HasBlockedUrlPatterns returns a boolean if a field has been set.
func (o *InlineObject40) HasBlockedUrlPatterns() bool {
	if o != nil && !isNil(o.BlockedUrlPatterns) {
		return true
	}

	return false
}

// SetBlockedUrlPatterns gets a reference to the given []string and assigns it to the BlockedUrlPatterns field.
func (o *InlineObject40) SetBlockedUrlPatterns(v []string) {
	o.BlockedUrlPatterns = v
}

// GetBlockedUrlCategories returns the BlockedUrlCategories field value if set, zero value otherwise.
func (o *InlineObject40) GetBlockedUrlCategories() []string {
	if o == nil || isNil(o.BlockedUrlCategories) {
		var ret []string
		return ret
	}
	return o.BlockedUrlCategories
}

// GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetBlockedUrlCategoriesOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedUrlCategories) {
    return nil, false
	}
	return o.BlockedUrlCategories, true
}

// HasBlockedUrlCategories returns a boolean if a field has been set.
func (o *InlineObject40) HasBlockedUrlCategories() bool {
	if o != nil && !isNil(o.BlockedUrlCategories) {
		return true
	}

	return false
}

// SetBlockedUrlCategories gets a reference to the given []string and assigns it to the BlockedUrlCategories field.
func (o *InlineObject40) SetBlockedUrlCategories(v []string) {
	o.BlockedUrlCategories = v
}

// GetUrlCategoryListSize returns the UrlCategoryListSize field value if set, zero value otherwise.
func (o *InlineObject40) GetUrlCategoryListSize() string {
	if o == nil || isNil(o.UrlCategoryListSize) {
		var ret string
		return ret
	}
	return *o.UrlCategoryListSize
}

// GetUrlCategoryListSizeOk returns a tuple with the UrlCategoryListSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject40) GetUrlCategoryListSizeOk() (*string, bool) {
	if o == nil || isNil(o.UrlCategoryListSize) {
    return nil, false
	}
	return o.UrlCategoryListSize, true
}

// HasUrlCategoryListSize returns a boolean if a field has been set.
func (o *InlineObject40) HasUrlCategoryListSize() bool {
	if o != nil && !isNil(o.UrlCategoryListSize) {
		return true
	}

	return false
}

// SetUrlCategoryListSize gets a reference to the given string and assigns it to the UrlCategoryListSize field.
func (o *InlineObject40) SetUrlCategoryListSize(v string) {
	o.UrlCategoryListSize = &v
}

func (o InlineObject40) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedUrlPatterns) {
		toSerialize["allowedUrlPatterns"] = o.AllowedUrlPatterns
	}
	if !isNil(o.BlockedUrlPatterns) {
		toSerialize["blockedUrlPatterns"] = o.BlockedUrlPatterns
	}
	if !isNil(o.BlockedUrlCategories) {
		toSerialize["blockedUrlCategories"] = o.BlockedUrlCategories
	}
	if !isNil(o.UrlCategoryListSize) {
		toSerialize["urlCategoryListSize"] = o.UrlCategoryListSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject40 struct {
	value *InlineObject40
	isSet bool
}

func (v NullableInlineObject40) Get() *InlineObject40 {
	return v.value
}

func (v *NullableInlineObject40) Set(val *InlineObject40) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject40) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject40) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject40(val *InlineObject40) *NullableInlineObject40 {
	return &NullableInlineObject40{value: val, isSet: true}
}

func (v NullableInlineObject40) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject40) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


