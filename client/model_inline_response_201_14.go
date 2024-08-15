/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20114 struct for InlineResponse20114
type InlineResponse20114 struct {
	// Cloud monitor import status
	Status *string `json:"status,omitempty"`
	// Unique id associated with the import of the device
	ImportId *string `json:"importId,omitempty"`
	// Response method
	Message *string `json:"message,omitempty"`
}

// NewInlineResponse20114 instantiates a new InlineResponse20114 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20114() *InlineResponse20114 {
	this := InlineResponse20114{}
	return &this
}

// NewInlineResponse20114WithDefaults instantiates a new InlineResponse20114 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20114WithDefaults() *InlineResponse20114 {
	this := InlineResponse20114{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20114) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20114) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20114) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20114) SetStatus(v string) {
	o.Status = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *InlineResponse20114) GetImportId() string {
	if o == nil || isNil(o.ImportId) {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20114) GetImportIdOk() (*string, bool) {
	if o == nil || isNil(o.ImportId) {
    return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *InlineResponse20114) HasImportId() bool {
	if o != nil && !isNil(o.ImportId) {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *InlineResponse20114) SetImportId(v string) {
	o.ImportId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20114) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20114) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20114) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20114) SetMessage(v string) {
	o.Message = &v
}

func (o InlineResponse20114) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ImportId) {
		toSerialize["importId"] = o.ImportId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20114 struct {
	value *InlineResponse20114
	isSet bool
}

func (v NullableInlineResponse20114) Get() *InlineResponse20114 {
	return v.value
}

func (v *NullableInlineResponse20114) Set(val *InlineResponse20114) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20114) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20114) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20114(val *InlineResponse20114) *NullableInlineResponse20114 {
	return &NullableInlineResponse20114{value: val, isSet: true}
}

func (v NullableInlineResponse20114) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20114) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


