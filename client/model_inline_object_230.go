/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject230 struct for InlineObject230
type InlineObject230 struct {
	// The name of the new network
	Name string `json:"name"`
	// The product type(s) of the new network. If more than one type is included, the network will be a combined network.
	ProductTypes []string `json:"productTypes"`
	// A list of tags to be applied to the network
	Tags []string `json:"tags,omitempty"`
	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone *string `json:"timeZone,omitempty"`
	// The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.
	CopyFromNetworkId *string `json:"copyFromNetworkId,omitempty"`
	// Add any notes or additional information about this network here.
	Notes *string `json:"notes,omitempty"`
}

// NewInlineObject230 instantiates a new InlineObject230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject230(name string, productTypes []string) *InlineObject230 {
	this := InlineObject230{}
	this.Name = name
	this.ProductTypes = productTypes
	return &this
}

// NewInlineObject230WithDefaults instantiates a new InlineObject230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject230WithDefaults() *InlineObject230 {
	this := InlineObject230{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject230) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject230) SetName(v string) {
	o.Name = v
}

// GetProductTypes returns the ProductTypes field value
func (o *InlineObject230) GetProductTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetProductTypesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProductTypes, true
}

// SetProductTypes sets field value
func (o *InlineObject230) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject230) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject230) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject230) SetTags(v []string) {
	o.Tags = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject230) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject230) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject230) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCopyFromNetworkId returns the CopyFromNetworkId field value if set, zero value otherwise.
func (o *InlineObject230) GetCopyFromNetworkId() string {
	if o == nil || isNil(o.CopyFromNetworkId) {
		var ret string
		return ret
	}
	return *o.CopyFromNetworkId
}

// GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetCopyFromNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyFromNetworkId) {
    return nil, false
	}
	return o.CopyFromNetworkId, true
}

// HasCopyFromNetworkId returns a boolean if a field has been set.
func (o *InlineObject230) HasCopyFromNetworkId() bool {
	if o != nil && !isNil(o.CopyFromNetworkId) {
		return true
	}

	return false
}

// SetCopyFromNetworkId gets a reference to the given string and assigns it to the CopyFromNetworkId field.
func (o *InlineObject230) SetCopyFromNetworkId(v string) {
	o.CopyFromNetworkId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject230) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject230) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject230) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineObject230) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.CopyFromNetworkId) {
		toSerialize["copyFromNetworkId"] = o.CopyFromNetworkId
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject230 struct {
	value *InlineObject230
	isSet bool
}

func (v NullableInlineObject230) Get() *InlineObject230 {
	return v.value
}

func (v *NullableInlineObject230) Set(val *InlineObject230) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject230) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject230(val *InlineObject230) *NullableInlineObject230 {
	return &NullableInlineObject230{value: val, isSet: true}
}

func (v NullableInlineObject230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


