/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject233 struct for InlineObject233
type InlineObject233 struct {
	// The type of log event this is recording, e.g. download or opening a banner
	LogEvent string `json:"logEvent"`
	// A JavaScript UTC datetime stamp for when the even occurred
	Timestamp int32 `json:"timestamp"`
	// The name of the onboarding distro being downloaded
	TargetOS *string `json:"targetOS,omitempty"`
	// Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used
	Request *string `json:"request,omitempty"`
}

// NewInlineObject233 instantiates a new InlineObject233 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject233(logEvent string, timestamp int32) *InlineObject233 {
	this := InlineObject233{}
	this.LogEvent = logEvent
	this.Timestamp = timestamp
	return &this
}

// NewInlineObject233WithDefaults instantiates a new InlineObject233 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject233WithDefaults() *InlineObject233 {
	this := InlineObject233{}
	return &this
}

// GetLogEvent returns the LogEvent field value
func (o *InlineObject233) GetLogEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogEvent
}

// GetLogEventOk returns a tuple with the LogEvent field value
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetLogEventOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LogEvent, true
}

// SetLogEvent sets field value
func (o *InlineObject233) SetLogEvent(v string) {
	o.LogEvent = v
}

// GetTimestamp returns the Timestamp field value
func (o *InlineObject233) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetTimestampOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *InlineObject233) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTargetOS returns the TargetOS field value if set, zero value otherwise.
func (o *InlineObject233) GetTargetOS() string {
	if o == nil || isNil(o.TargetOS) {
		var ret string
		return ret
	}
	return *o.TargetOS
}

// GetTargetOSOk returns a tuple with the TargetOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetTargetOSOk() (*string, bool) {
	if o == nil || isNil(o.TargetOS) {
    return nil, false
	}
	return o.TargetOS, true
}

// HasTargetOS returns a boolean if a field has been set.
func (o *InlineObject233) HasTargetOS() bool {
	if o != nil && !isNil(o.TargetOS) {
		return true
	}

	return false
}

// SetTargetOS gets a reference to the given string and assigns it to the TargetOS field.
func (o *InlineObject233) SetTargetOS(v string) {
	o.TargetOS = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineObject233) GetRequest() string {
	if o == nil || isNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject233) GetRequestOk() (*string, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineObject233) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *InlineObject233) SetRequest(v string) {
	o.Request = &v
}

func (o InlineObject233) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logEvent"] = o.LogEvent
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.TargetOS) {
		toSerialize["targetOS"] = o.TargetOS
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject233 struct {
	value *InlineObject233
	isSet bool
}

func (v NullableInlineObject233) Get() *InlineObject233 {
	return v.value
}

func (v *NullableInlineObject233) Set(val *InlineObject233) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject233) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject233) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject233(val *InlineObject233) *NullableInlineObject233 {
	return &NullableInlineObject233{value: val, isSet: true}
}

func (v NullableInlineObject233) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject233) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


