/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200277 struct for InlineResponse200277
type InlineResponse200277 struct {
	// theme id
	Id *string `json:"id,omitempty"`
	// theme name
	Name *string `json:"name,omitempty"`
	// list of theme assets
	ThemeAssets []OrganizationsOrganizationIdSplashThemesThemeAssets `json:"themeAssets,omitempty"`
}

// NewInlineResponse200277 instantiates a new InlineResponse200277 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200277() *InlineResponse200277 {
	this := InlineResponse200277{}
	return &this
}

// NewInlineResponse200277WithDefaults instantiates a new InlineResponse200277 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200277WithDefaults() *InlineResponse200277 {
	this := InlineResponse200277{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200277) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200277) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200277) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200277) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200277) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200277) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200277) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200277) SetName(v string) {
	o.Name = &v
}

// GetThemeAssets returns the ThemeAssets field value if set, zero value otherwise.
func (o *InlineResponse200277) GetThemeAssets() []OrganizationsOrganizationIdSplashThemesThemeAssets {
	if o == nil || isNil(o.ThemeAssets) {
		var ret []OrganizationsOrganizationIdSplashThemesThemeAssets
		return ret
	}
	return o.ThemeAssets
}

// GetThemeAssetsOk returns a tuple with the ThemeAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200277) GetThemeAssetsOk() ([]OrganizationsOrganizationIdSplashThemesThemeAssets, bool) {
	if o == nil || isNil(o.ThemeAssets) {
    return nil, false
	}
	return o.ThemeAssets, true
}

// HasThemeAssets returns a boolean if a field has been set.
func (o *InlineResponse200277) HasThemeAssets() bool {
	if o != nil && !isNil(o.ThemeAssets) {
		return true
	}

	return false
}

// SetThemeAssets gets a reference to the given []OrganizationsOrganizationIdSplashThemesThemeAssets and assigns it to the ThemeAssets field.
func (o *InlineResponse200277) SetThemeAssets(v []OrganizationsOrganizationIdSplashThemesThemeAssets) {
	o.ThemeAssets = v
}

func (o InlineResponse200277) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ThemeAssets) {
		toSerialize["themeAssets"] = o.ThemeAssets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200277 struct {
	value *InlineResponse200277
	isSet bool
}

func (v NullableInlineResponse200277) Get() *InlineResponse200277 {
	return v.value
}

func (v *NullableInlineResponse200277) Set(val *InlineResponse200277) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200277) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200277) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200277(val *InlineResponse200277) *NullableInlineResponse200277 {
	return &NullableInlineResponse200277{value: val, isSet: true}
}

func (v NullableInlineResponse200277) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200277) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


