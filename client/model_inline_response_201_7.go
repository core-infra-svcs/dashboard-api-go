/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2017 struct for InlineResponse2017
type InlineResponse2017 struct {
	// ID of the Wake-on-LAN job
	WakeOnLanId *string `json:"wakeOnLanId,omitempty"`
	// GET this url to check the status of your ping request
	Url *string `json:"url,omitempty"`
	// Status of the Wake-on-LAN request
	Status *string `json:"status,omitempty"`
	Request *InlineResponse2017Request `json:"request,omitempty"`
	// An error message for a failed execution
	Error *string `json:"error,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2017 instantiates a new InlineResponse2017 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2017() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// NewInlineResponse2017WithDefaults instantiates a new InlineResponse2017 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2017WithDefaults() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// GetWakeOnLanId returns the WakeOnLanId field value if set, zero value otherwise.
func (o *InlineResponse2017) GetWakeOnLanId() string {
	if o == nil || isNil(o.WakeOnLanId) {
		var ret string
		return ret
	}
	return *o.WakeOnLanId
}

// GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetWakeOnLanIdOk() (*string, bool) {
	if o == nil || isNil(o.WakeOnLanId) {
    return nil, false
	}
	return o.WakeOnLanId, true
}

// HasWakeOnLanId returns a boolean if a field has been set.
func (o *InlineResponse2017) HasWakeOnLanId() bool {
	if o != nil && !isNil(o.WakeOnLanId) {
		return true
	}

	return false
}

// SetWakeOnLanId gets a reference to the given string and assigns it to the WakeOnLanId field.
func (o *InlineResponse2017) SetWakeOnLanId(v string) {
	o.WakeOnLanId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2017) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2017) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2017) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2017) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2017) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2017) SetStatus(v string) {
	o.Status = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2017) GetRequest() InlineResponse2017Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2017Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetRequestOk() (*InlineResponse2017Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2017) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2017Request and assigns it to the Request field.
func (o *InlineResponse2017) SetRequest(v InlineResponse2017Request) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse2017) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse2017) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse2017) SetError(v string) {
	o.Error = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2017) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2017) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2017) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2017) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2017 struct {
	value *InlineResponse2017
	isSet bool
}

func (v NullableInlineResponse2017) Get() *InlineResponse2017 {
	return v.value
}

func (v *NullableInlineResponse2017) Set(val *InlineResponse2017) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2017) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2017) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2017(val *InlineResponse2017) *NullableInlineResponse2017 {
	return &NullableInlineResponse2017{value: val, isSet: true}
}

func (v NullableInlineResponse2017) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2017) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


