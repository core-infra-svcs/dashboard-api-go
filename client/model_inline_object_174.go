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

// InlineObject174 struct for InlineObject174
type InlineObject174 struct {
	// The name of the configuration template
	Name *string `json:"name,omitempty"`
	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewInlineObject174 instantiates a new InlineObject174 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject174() *InlineObject174 {
	this := InlineObject174{}
	return &this
}

// NewInlineObject174WithDefaults instantiates a new InlineObject174 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject174WithDefaults() *InlineObject174 {
	this := InlineObject174{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject174) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject174) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject174) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject174) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineObject174) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject174) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineObject174) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineObject174) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o InlineObject174) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject174 struct {
	value *InlineObject174
	isSet bool
}

func (v NullableInlineObject174) Get() *InlineObject174 {
	return v.value
}

func (v *NullableInlineObject174) Set(val *InlineObject174) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject174) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject174) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject174(val *InlineObject174) *NullableInlineObject174 {
	return &NullableInlineObject174{value: val, isSet: true}
}

func (v NullableInlineObject174) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject174) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


