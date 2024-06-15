/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2012 struct for InlineResponse2012
type InlineResponse2012 struct {
	// Id of the cable test request. Used to check the status of the request.
	CableTestId *string `json:"cableTestId,omitempty"`
	// GET this url to check the status of your cable test request.
	Url *string `json:"url,omitempty"`
	Request *InlineResponse2012Request `json:"request,omitempty"`
	// Status of the cable test request.
	Status *string `json:"status,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2012 instantiates a new InlineResponse2012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2012() *InlineResponse2012 {
	this := InlineResponse2012{}
	return &this
}

// NewInlineResponse2012WithDefaults instantiates a new InlineResponse2012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2012WithDefaults() *InlineResponse2012 {
	this := InlineResponse2012{}
	return &this
}

// GetCableTestId returns the CableTestId field value if set, zero value otherwise.
func (o *InlineResponse2012) GetCableTestId() string {
	if o == nil || isNil(o.CableTestId) {
		var ret string
		return ret
	}
	return *o.CableTestId
}

// GetCableTestIdOk returns a tuple with the CableTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012) GetCableTestIdOk() (*string, bool) {
	if o == nil || isNil(o.CableTestId) {
    return nil, false
	}
	return o.CableTestId, true
}

// HasCableTestId returns a boolean if a field has been set.
func (o *InlineResponse2012) HasCableTestId() bool {
	if o != nil && !isNil(o.CableTestId) {
		return true
	}

	return false
}

// SetCableTestId gets a reference to the given string and assigns it to the CableTestId field.
func (o *InlineResponse2012) SetCableTestId(v string) {
	o.CableTestId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2012) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2012) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2012) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2012) GetRequest() InlineResponse2012Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2012Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012) GetRequestOk() (*InlineResponse2012Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2012) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2012Request and assigns it to the Request field.
func (o *InlineResponse2012) SetRequest(v InlineResponse2012Request) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2012) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2012) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2012) SetStatus(v string) {
	o.Status = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2012) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2012) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2012) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2012) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CableTestId) {
		toSerialize["cableTestId"] = o.CableTestId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2012 struct {
	value *InlineResponse2012
	isSet bool
}

func (v NullableInlineResponse2012) Get() *InlineResponse2012 {
	return v.value
}

func (v *NullableInlineResponse2012) Set(val *InlineResponse2012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2012(val *InlineResponse2012) *NullableInlineResponse2012 {
	return &NullableInlineResponse2012{value: val, isSet: true}
}

func (v NullableInlineResponse2012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


