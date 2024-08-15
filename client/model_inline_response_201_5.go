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

// InlineResponse2015 struct for InlineResponse2015
type InlineResponse2015 struct {
	// ID of the Wake-on-LAN job
	WakeOnLanId *string `json:"wakeOnLanId,omitempty"`
	// GET this url to check the status of your ping request
	Url *string `json:"url,omitempty"`
	// Status of the Wake-on-LAN request
	Status *string `json:"status,omitempty"`
	Request *InlineResponse2015Request `json:"request,omitempty"`
	// An error message for a failed execution
	Error *string `json:"error,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2015 instantiates a new InlineResponse2015 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015() *InlineResponse2015 {
	this := InlineResponse2015{}
	return &this
}

// NewInlineResponse2015WithDefaults instantiates a new InlineResponse2015 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015WithDefaults() *InlineResponse2015 {
	this := InlineResponse2015{}
	return &this
}

// GetWakeOnLanId returns the WakeOnLanId field value if set, zero value otherwise.
func (o *InlineResponse2015) GetWakeOnLanId() string {
	if o == nil || isNil(o.WakeOnLanId) {
		var ret string
		return ret
	}
	return *o.WakeOnLanId
}

// GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetWakeOnLanIdOk() (*string, bool) {
	if o == nil || isNil(o.WakeOnLanId) {
    return nil, false
	}
	return o.WakeOnLanId, true
}

// HasWakeOnLanId returns a boolean if a field has been set.
func (o *InlineResponse2015) HasWakeOnLanId() bool {
	if o != nil && !isNil(o.WakeOnLanId) {
		return true
	}

	return false
}

// SetWakeOnLanId gets a reference to the given string and assigns it to the WakeOnLanId field.
func (o *InlineResponse2015) SetWakeOnLanId(v string) {
	o.WakeOnLanId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2015) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2015) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2015) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2015) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2015) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2015) SetStatus(v string) {
	o.Status = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2015) GetRequest() InlineResponse2015Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2015Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetRequestOk() (*InlineResponse2015Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2015) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2015Request and assigns it to the Request field.
func (o *InlineResponse2015) SetRequest(v InlineResponse2015Request) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse2015) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse2015) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse2015) SetError(v string) {
	o.Error = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2015) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2015) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2015) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2015) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WakeOnLanId) {
		toSerialize["wakeOnLanId"] = o.WakeOnLanId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015 struct {
	value *InlineResponse2015
	isSet bool
}

func (v NullableInlineResponse2015) Get() *InlineResponse2015 {
	return v.value
}

func (v *NullableInlineResponse2015) Set(val *InlineResponse2015) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015(val *InlineResponse2015) *NullableInlineResponse2015 {
	return &NullableInlineResponse2015{value: val, isSet: true}
}

func (v NullableInlineResponse2015) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


