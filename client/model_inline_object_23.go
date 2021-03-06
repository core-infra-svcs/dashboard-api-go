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

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	// The name of the network
	Name *string `json:"name,omitempty"`
	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone *string `json:"timeZone,omitempty"`
	// A list of tags to be applied to the network
	Tags *[]string `json:"tags,omitempty"`
	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break.
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// Add any notes or additional information about this network here.
	Notes *string `json:"notes,omitempty"`
}

// NewInlineObject23 instantiates a new InlineObject23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject23() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// NewInlineObject23WithDefaults instantiates a new InlineObject23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject23WithDefaults() *InlineObject23 {
	this := InlineObject23{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject23) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject23) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject23) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject23) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject23) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject23) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject23) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject23) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject23) SetTags(v []string) {
	o.Tags = &v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineObject23) GetEnrollmentString() string {
	if o == nil || o.EnrollmentString == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || o.EnrollmentString == nil {
		return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineObject23) HasEnrollmentString() bool {
	if o != nil && o.EnrollmentString != nil {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineObject23) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject23) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject23) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject23) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject23) SetNotes(v string) {
	o.Notes = &v
}

func (o InlineObject23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.EnrollmentString != nil {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject23 struct {
	value *InlineObject23
	isSet bool
}

func (v NullableInlineObject23) Get() *InlineObject23 {
	return v.value
}

func (v *NullableInlineObject23) Set(val *InlineObject23) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject23) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject23(val *InlineObject23) *NullableInlineObject23 {
	return &NullableInlineObject23{value: val, isSet: true}
}

func (v NullableInlineObject23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


