/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject41 struct for InlineObject41
type InlineObject41 struct {
	// A list of URL patterns that are allowed
	AllowedUrlPatterns []string `json:"allowedUrlPatterns,omitempty"`
	// A list of URL patterns that are blocked
	BlockedUrlPatterns []string `json:"blockedUrlPatterns,omitempty"`
	// A list of URL categories to block
	BlockedUrlCategories []string `json:"blockedUrlCategories,omitempty"`
	// URL category list size which is either 'topSites' or 'fullList'
	UrlCategoryListSize *string `json:"urlCategoryListSize,omitempty"`
}

// NewInlineObject41 instantiates a new InlineObject41 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject41() *InlineObject41 {
	this := InlineObject41{}
	return &this
}

// NewInlineObject41WithDefaults instantiates a new InlineObject41 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject41WithDefaults() *InlineObject41 {
	this := InlineObject41{}
	return &this
}

// GetAllowedUrlPatterns returns the AllowedUrlPatterns field value if set, zero value otherwise.
func (o *InlineObject41) GetAllowedUrlPatterns() []string {
	if o == nil || isNil(o.AllowedUrlPatterns) {
		var ret []string
		return ret
	}
	return o.AllowedUrlPatterns
}

// GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetAllowedUrlPatternsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedUrlPatterns) {
    return nil, false
	}
	return o.AllowedUrlPatterns, true
}

// HasAllowedUrlPatterns returns a boolean if a field has been set.
func (o *InlineObject41) HasAllowedUrlPatterns() bool {
	if o != nil && !isNil(o.AllowedUrlPatterns) {
		return true
	}

	return false
}

// SetAllowedUrlPatterns gets a reference to the given []string and assigns it to the AllowedUrlPatterns field.
func (o *InlineObject41) SetAllowedUrlPatterns(v []string) {
	o.AllowedUrlPatterns = v
}

// GetBlockedUrlPatterns returns the BlockedUrlPatterns field value if set, zero value otherwise.
func (o *InlineObject41) GetBlockedUrlPatterns() []string {
	if o == nil || isNil(o.BlockedUrlPatterns) {
		var ret []string
		return ret
	}
	return o.BlockedUrlPatterns
}

// GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetBlockedUrlPatternsOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedUrlPatterns) {
    return nil, false
	}
	return o.BlockedUrlPatterns, true
}

// HasBlockedUrlPatterns returns a boolean if a field has been set.
func (o *InlineObject41) HasBlockedUrlPatterns() bool {
	if o != nil && !isNil(o.BlockedUrlPatterns) {
		return true
	}

	return false
}

// SetBlockedUrlPatterns gets a reference to the given []string and assigns it to the BlockedUrlPatterns field.
func (o *InlineObject41) SetBlockedUrlPatterns(v []string) {
	o.BlockedUrlPatterns = v
}

// GetBlockedUrlCategories returns the BlockedUrlCategories field value if set, zero value otherwise.
func (o *InlineObject41) GetBlockedUrlCategories() []string {
	if o == nil || isNil(o.BlockedUrlCategories) {
		var ret []string
		return ret
	}
	return o.BlockedUrlCategories
}

// GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetBlockedUrlCategoriesOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedUrlCategories) {
    return nil, false
	}
	return o.BlockedUrlCategories, true
}

// HasBlockedUrlCategories returns a boolean if a field has been set.
func (o *InlineObject41) HasBlockedUrlCategories() bool {
	if o != nil && !isNil(o.BlockedUrlCategories) {
		return true
	}

	return false
}

// SetBlockedUrlCategories gets a reference to the given []string and assigns it to the BlockedUrlCategories field.
func (o *InlineObject41) SetBlockedUrlCategories(v []string) {
	o.BlockedUrlCategories = v
}

// GetUrlCategoryListSize returns the UrlCategoryListSize field value if set, zero value otherwise.
func (o *InlineObject41) GetUrlCategoryListSize() string {
	if o == nil || isNil(o.UrlCategoryListSize) {
		var ret string
		return ret
	}
	return *o.UrlCategoryListSize
}

// GetUrlCategoryListSizeOk returns a tuple with the UrlCategoryListSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject41) GetUrlCategoryListSizeOk() (*string, bool) {
	if o == nil || isNil(o.UrlCategoryListSize) {
    return nil, false
	}
	return o.UrlCategoryListSize, true
}

// HasUrlCategoryListSize returns a boolean if a field has been set.
func (o *InlineObject41) HasUrlCategoryListSize() bool {
	if o != nil && !isNil(o.UrlCategoryListSize) {
		return true
	}

	return false
}

// SetUrlCategoryListSize gets a reference to the given string and assigns it to the UrlCategoryListSize field.
func (o *InlineObject41) SetUrlCategoryListSize(v string) {
	o.UrlCategoryListSize = &v
}

func (o InlineObject41) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject41 struct {
	value *InlineObject41
	isSet bool
}

func (v NullableInlineObject41) Get() *InlineObject41 {
	return v.value
}

func (v *NullableInlineObject41) Set(val *InlineObject41) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject41) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject41) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject41(val *InlineObject41) *NullableInlineObject41 {
	return &NullableInlineObject41{value: val, isSet: true}
}

func (v NullableInlineObject41) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject41) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


