/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2011Callback Information for callback used to send back results
type InlineResponse2011Callback struct {
	// The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
	Id *string `json:"id,omitempty"`
	// The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver
	Url *string `json:"url,omitempty"`
	// The status of the callback
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse2011Callback instantiates a new InlineResponse2011Callback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2011Callback() *InlineResponse2011Callback {
	this := InlineResponse2011Callback{}
	return &this
}

// NewInlineResponse2011CallbackWithDefaults instantiates a new InlineResponse2011Callback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2011CallbackWithDefaults() *InlineResponse2011Callback {
	this := InlineResponse2011Callback{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2011Callback) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Callback) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2011Callback) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2011Callback) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2011Callback) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Callback) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2011Callback) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2011Callback) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2011Callback) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011Callback) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2011Callback) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2011Callback) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse2011Callback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2011Callback struct {
	value *InlineResponse2011Callback
	isSet bool
}

func (v NullableInlineResponse2011Callback) Get() *InlineResponse2011Callback {
	return v.value
}

func (v *NullableInlineResponse2011Callback) Set(val *InlineResponse2011Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2011Callback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2011Callback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2011Callback(val *InlineResponse2011Callback) *NullableInlineResponse2011Callback {
	return &NullableInlineResponse2011Callback{value: val, isSet: true}
}

func (v NullableInlineResponse2011Callback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2011Callback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


