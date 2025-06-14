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

// InlineObject39 struct for InlineObject39
type InlineObject39 struct {
	// The name of the network
	Name *string `json:"name,omitempty"`
	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone *string `json:"timeZone,omitempty"`
	// A list of tags to be applied to the network
	Tags []string `json:"tags,omitempty"`
	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break.
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// Add any notes or additional information about this network here.
	Notes NullableString `json:"notes,omitempty"`
}

// NewInlineObject39 instantiates a new InlineObject39 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject39() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// NewInlineObject39WithDefaults instantiates a new InlineObject39 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject39WithDefaults() *InlineObject39 {
	this := InlineObject39{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject39) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject39) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject39) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject39) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject39) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject39) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject39) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject39) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject39) SetTags(v []string) {
	o.Tags = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineObject39) GetEnrollmentString() string {
	if o == nil || isNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject39) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || isNil(o.EnrollmentString) {
    return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineObject39) HasEnrollmentString() bool {
	if o != nil && !isNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineObject39) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject39) GetNotes() string {
	if o == nil || isNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject39) GetNotesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject39) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *InlineObject39) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *InlineObject39) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *InlineObject39) UnsetNotes() {
	o.Notes.Unset()
}

func (o InlineObject39) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject39 struct {
	value *InlineObject39
	isSet bool
}

func (v NullableInlineObject39) Get() *InlineObject39 {
	return v.value
}

func (v *NullableInlineObject39) Set(val *InlineObject39) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject39) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject39) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject39(val *InlineObject39) *NullableInlineObject39 {
	return &NullableInlineObject39{value: val, isSet: true}
}

func (v NullableInlineObject39) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject39) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


