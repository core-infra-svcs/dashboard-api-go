/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 June, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject186 struct for InlineObject186
type InlineObject186 struct {
	// The name of the new network
	Name string `json:"name"`
	// The product type(s) of the new network. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, environmental. If more than one type is included, the network will be a combined network.
	ProductTypes []string `json:"productTypes"`
	// A list of tags to be applied to the network
	Tags *[]string `json:"tags,omitempty"`
	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone *string `json:"timeZone,omitempty"`
	// The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.
	CopyFromNetworkId *string `json:"copyFromNetworkId,omitempty"`
	// Add any notes or additional information about this network here.
	Notes *string `json:"notes,omitempty"`
}

// NewInlineObject186 instantiates a new InlineObject186 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject186(name string, productTypes []string) *InlineObject186 {
	this := InlineObject186{}
	this.Name = name
	this.ProductTypes = productTypes
	return &this
}

// NewInlineObject186WithDefaults instantiates a new InlineObject186 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject186WithDefaults() *InlineObject186 {
	this := InlineObject186{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject186) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject186) SetName(v string) {
	o.Name = v
}

// GetProductTypes returns the ProductTypes field value
func (o *InlineObject186) GetProductTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetProductTypesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductTypes, true
}

// SetProductTypes sets field value
func (o *InlineObject186) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject186) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject186) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject186) SetTags(v []string) {
	o.Tags = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject186) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject186) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject186) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCopyFromNetworkId returns the CopyFromNetworkId field value if set, zero value otherwise.
func (o *InlineObject186) GetCopyFromNetworkId() string {
	if o == nil || o.CopyFromNetworkId == nil {
		var ret string
		return ret
	}
	return *o.CopyFromNetworkId
}

// GetCopyFromNetworkIdOk returns a tuple with the CopyFromNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetCopyFromNetworkIdOk() (*string, bool) {
	if o == nil || o.CopyFromNetworkId == nil {
		return nil, false
	}
	return o.CopyFromNetworkId, true
}

// HasCopyFromNetworkId returns a boolean if a field has been set.
func (o *InlineObject186) HasCopyFromNetworkId() bool {
	if o != nil && o.CopyFromNetworkId != nil {
		return true
	}

	return false
}

// SetCopyFromNetworkId gets a reference to the given string and assigns it to the CopyFromNetworkId field.
func (o *InlineObject186) SetCopyFromNetworkId(v string) {
	o.CopyFromNetworkId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject186) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject186) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject186) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject186) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineObject186) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.CopyFromNetworkId != nil {
		toSerialize["copyFromNetworkId"] = o.CopyFromNetworkId
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject186 struct {
	value *InlineObject186
	isSet bool
}

func (v NullableInlineObject186) Get() *InlineObject186 {
	return v.value
}

func (v *NullableInlineObject186) Set(val *InlineObject186) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject186) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject186) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject186(val *InlineObject186) *NullableInlineObject186 {
	return &NullableInlineObject186{value: val, isSet: true}
}

func (v NullableInlineObject186) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject186) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


