/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject108 struct for InlineObject108
type InlineObject108 struct {
	// The wifiMacs of the devices to be modified.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the devices to be modified.
	Ids []string `json:"ids,omitempty"`
	// The serials of the devices to be modified.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.
	Scope []string `json:"scope,omitempty"`
	// The tags to be added, deleted, or updated.
	Tags []string `json:"tags"`
	// One of add, delete, or update. Only devices that have been modified will be returned.
	UpdateAction string `json:"updateAction"`
}

// NewInlineObject108 instantiates a new InlineObject108 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject108(tags []string, updateAction string) *InlineObject108 {
	this := InlineObject108{}
	this.Tags = tags
	this.UpdateAction = updateAction
	return &this
}

// NewInlineObject108WithDefaults instantiates a new InlineObject108 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject108WithDefaults() *InlineObject108 {
	this := InlineObject108{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *InlineObject108) GetWifiMacs() []string {
	if o == nil || isNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetWifiMacsOk() ([]string, bool) {
	if o == nil || isNil(o.WifiMacs) {
    return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *InlineObject108) HasWifiMacs() bool {
	if o != nil && !isNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *InlineObject108) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *InlineObject108) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
    return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *InlineObject108) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *InlineObject108) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineObject108) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineObject108) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineObject108) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineObject108) GetScope() []string {
	if o == nil || isNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetScopeOk() ([]string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject108) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *InlineObject108) SetScope(v []string) {
	o.Scope = v
}

// GetTags returns the Tags field value
func (o *InlineObject108) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetTagsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *InlineObject108) SetTags(v []string) {
	o.Tags = v
}

// GetUpdateAction returns the UpdateAction field value
func (o *InlineObject108) GetUpdateAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateAction
}

// GetUpdateActionOk returns a tuple with the UpdateAction field value
// and a boolean to check if the value has been set.
func (o *InlineObject108) GetUpdateActionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdateAction, true
}

// SetUpdateAction sets field value
func (o *InlineObject108) SetUpdateAction(v string) {
	o.UpdateAction = v
}

func (o InlineObject108) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WifiMacs) {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["updateAction"] = o.UpdateAction
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject108 struct {
	value *InlineObject108
	isSet bool
}

func (v NullableInlineObject108) Get() *InlineObject108 {
	return v.value
}

func (v *NullableInlineObject108) Set(val *InlineObject108) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject108) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject108) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject108(val *InlineObject108) *NullableInlineObject108 {
	return &NullableInlineObject108{value: val, isSet: true}
}

func (v NullableInlineObject108) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject108) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


