/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026 struct for InlineResponse20026
type InlineResponse20026 struct {
	// ID of the Wake-on-LAN job
	WakeOnLanId *string `json:"wakeOnLanId,omitempty"`
	// GET this url to check the status of your ping request
	Url *string `json:"url,omitempty"`
	// Status of the Wake-on-LAN request
	Status *string `json:"status,omitempty"`
	Request *InlineResponse2017Request `json:"request,omitempty"`
	// An error message for a failed execution
	Error *string `json:"error,omitempty"`
}

// NewInlineResponse20026 instantiates a new InlineResponse20026 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// NewInlineResponse20026WithDefaults instantiates a new InlineResponse20026 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026WithDefaults() *InlineResponse20026 {
	this := InlineResponse20026{}
	return &this
}

// GetWakeOnLanId returns the WakeOnLanId field value if set, zero value otherwise.
func (o *InlineResponse20026) GetWakeOnLanId() string {
	if o == nil || isNil(o.WakeOnLanId) {
		var ret string
		return ret
	}
	return *o.WakeOnLanId
}

// GetWakeOnLanIdOk returns a tuple with the WakeOnLanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetWakeOnLanIdOk() (*string, bool) {
	if o == nil || isNil(o.WakeOnLanId) {
    return nil, false
	}
	return o.WakeOnLanId, true
}

// HasWakeOnLanId returns a boolean if a field has been set.
func (o *InlineResponse20026) HasWakeOnLanId() bool {
	if o != nil && !isNil(o.WakeOnLanId) {
		return true
	}

	return false
}

// SetWakeOnLanId gets a reference to the given string and assigns it to the WakeOnLanId field.
func (o *InlineResponse20026) SetWakeOnLanId(v string) {
	o.WakeOnLanId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20026) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20026) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20026) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20026) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20026) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20026) SetStatus(v string) {
	o.Status = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse20026) GetRequest() InlineResponse2017Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2017Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetRequestOk() (*InlineResponse2017Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse20026) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2017Request and assigns it to the Request field.
func (o *InlineResponse20026) SetRequest(v InlineResponse2017Request) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20026) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20026) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse20026) SetError(v string) {
	o.Error = &v
}

func (o InlineResponse20026) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026 struct {
	value *InlineResponse20026
	isSet bool
}

func (v NullableInlineResponse20026) Get() *InlineResponse20026 {
	return v.value
}

func (v *NullableInlineResponse20026) Set(val *InlineResponse20026) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026(val *InlineResponse20026) *NullableInlineResponse20026 {
	return &NullableInlineResponse20026{value: val, isSet: true}
}

func (v NullableInlineResponse20026) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


