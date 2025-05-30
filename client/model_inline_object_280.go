/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject280 struct for InlineObject280
type InlineObject280 struct {
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

// NewInlineObject280 instantiates a new InlineObject280 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject280(name string, productTypes []string) *InlineObject280 {
	this := InlineObject280{}
	this.Name = name
	this.ProductTypes = productTypes
	return &this
}

// NewInlineObject280WithDefaults instantiates a new InlineObject280 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject280WithDefaults() *InlineObject280 {
	this := InlineObject280{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject280) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject280) SetName(v string) {
	o.Name = v
}

// GetProductTypes returns the ProductTypes field value
func (o *InlineObject280) GetProductTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetProductTypesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProductTypes, true
}

// SetProductTypes sets field value
func (o *InlineObject280) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject280) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject280) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject280) SetTags(v []string) {
	o.Tags = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject280) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject280) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject280) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCopyFromNetworkId returns the CopyFromNetworkId field value if set, zero value otherwise.
func (o *InlineObject280) GetCopyFromNetworkId() string {
	if o == nil || isNil(o.CopyFromNetworkId) {
		var ret string
		return ret
	}
	return *o.CopyFromNetworkId
}

// GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetCopyFromNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyFromNetworkId) {
    return nil, false
	}
	return o.CopyFromNetworkId, true
}

// HasCopyFromNetworkId returns a boolean if a field has been set.
func (o *InlineObject280) HasCopyFromNetworkId() bool {
	if o != nil && !isNil(o.CopyFromNetworkId) {
		return true
	}

	return false
}

// SetCopyFromNetworkId gets a reference to the given string and assigns it to the CopyFromNetworkId field.
func (o *InlineObject280) SetCopyFromNetworkId(v string) {
	o.CopyFromNetworkId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject280) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject280) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject280) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject280) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineObject280) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject280 struct {
	value *InlineObject280
	isSet bool
}

func (v NullableInlineObject280) Get() *InlineObject280 {
	return v.value
}

func (v *NullableInlineObject280) Set(val *InlineObject280) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject280) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject280) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject280(val *InlineObject280) *NullableInlineObject280 {
	return &NullableInlineObject280{value: val, isSet: true}
}

func (v NullableInlineObject280) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject280) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


