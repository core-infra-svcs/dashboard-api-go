/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2013 struct for InlineResponse2013
type InlineResponse2013 struct {
	// Id to check the status of your ping request.
	PingId *string `json:"pingId,omitempty"`
	// GET this url to check the status of your ping request.
	Url *string `json:"url,omitempty"`
	Request *InlineResponse2013Request `json:"request,omitempty"`
	// Status of the ping request.
	Status *string `json:"status,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2013 instantiates a new InlineResponse2013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2013() *InlineResponse2013 {
	this := InlineResponse2013{}
	return &this
}

// NewInlineResponse2013WithDefaults instantiates a new InlineResponse2013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2013WithDefaults() *InlineResponse2013 {
	this := InlineResponse2013{}
	return &this
}

// GetPingId returns the PingId field value if set, zero value otherwise.
func (o *InlineResponse2013) GetPingId() string {
	if o == nil || isNil(o.PingId) {
		var ret string
		return ret
	}
	return *o.PingId
}

// GetPingIdOk returns a tuple with the PingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetPingIdOk() (*string, bool) {
	if o == nil || isNil(o.PingId) {
    return nil, false
	}
	return o.PingId, true
}

// HasPingId returns a boolean if a field has been set.
func (o *InlineResponse2013) HasPingId() bool {
	if o != nil && !isNil(o.PingId) {
		return true
	}

	return false
}

// SetPingId gets a reference to the given string and assigns it to the PingId field.
func (o *InlineResponse2013) SetPingId(v string) {
	o.PingId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2013) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2013) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2013) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2013) GetRequest() InlineResponse2013Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2013Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetRequestOk() (*InlineResponse2013Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2013) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2013Request and assigns it to the Request field.
func (o *InlineResponse2013) SetRequest(v InlineResponse2013Request) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2013) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2013) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2013) SetStatus(v string) {
	o.Status = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2013) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2013) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2013) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2013) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PingId) {
		toSerialize["pingId"] = o.PingId
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

type NullableInlineResponse2013 struct {
	value *InlineResponse2013
	isSet bool
}

func (v NullableInlineResponse2013) Get() *InlineResponse2013 {
	return v.value
}

func (v *NullableInlineResponse2013) Set(val *InlineResponse2013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2013(val *InlineResponse2013) *NullableInlineResponse2013 {
	return &NullableInlineResponse2013{value: val, isSet: true}
}

func (v NullableInlineResponse2013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


