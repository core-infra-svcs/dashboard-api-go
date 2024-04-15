/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20023 struct for InlineResponse20023
type InlineResponse20023 struct {
	// ID of throughput test job
	ThroughputTestId *string `json:"throughputTestId,omitempty"`
	// GET this url to check the status of your throughput test request
	Url *string `json:"url,omitempty"`
	// Status of the throughput test request
	Status *string `json:"status,omitempty"`
	Result *InlineResponse2014Result `json:"result,omitempty"`
	Request *InlineResponse2014Request `json:"request,omitempty"`
	// Description of the error.
	Error *string `json:"error,omitempty"`
}

// NewInlineResponse20023 instantiates a new InlineResponse20023 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20023() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20023WithDefaults() *InlineResponse20023 {
	this := InlineResponse20023{}
	return &this
}

// GetThroughputTestId returns the ThroughputTestId field value if set, zero value otherwise.
func (o *InlineResponse20023) GetThroughputTestId() string {
	if o == nil || isNil(o.ThroughputTestId) {
		var ret string
		return ret
	}
	return *o.ThroughputTestId
}

// GetThroughputTestIdOk returns a tuple with the ThroughputTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetThroughputTestIdOk() (*string, bool) {
	if o == nil || isNil(o.ThroughputTestId) {
    return nil, false
	}
	return o.ThroughputTestId, true
}

// HasThroughputTestId returns a boolean if a field has been set.
func (o *InlineResponse20023) HasThroughputTestId() bool {
	if o != nil && !isNil(o.ThroughputTestId) {
		return true
	}

	return false
}

// SetThroughputTestId gets a reference to the given string and assigns it to the ThroughputTestId field.
func (o *InlineResponse20023) SetThroughputTestId(v string) {
	o.ThroughputTestId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20023) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20023) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20023) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20023) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20023) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20023) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InlineResponse20023) GetResult() InlineResponse2014Result {
	if o == nil || isNil(o.Result) {
		var ret InlineResponse2014Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetResultOk() (*InlineResponse2014Result, bool) {
	if o == nil || isNil(o.Result) {
    return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InlineResponse20023) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given InlineResponse2014Result and assigns it to the Result field.
func (o *InlineResponse20023) SetResult(v InlineResponse2014Result) {
	o.Result = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse20023) GetRequest() InlineResponse2014Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2014Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetRequestOk() (*InlineResponse2014Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse20023) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2014Request and assigns it to the Request field.
func (o *InlineResponse20023) SetRequest(v InlineResponse2014Request) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20023) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20023) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20023) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse20023) SetError(v string) {
	o.Error = &v
}

func (o InlineResponse20023) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ThroughputTestId) {
		toSerialize["throughputTestId"] = o.ThroughputTestId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20023 struct {
	value *InlineResponse20023
	isSet bool
}

func (v NullableInlineResponse20023) Get() *InlineResponse20023 {
	return v.value
}

func (v *NullableInlineResponse20023) Set(val *InlineResponse20023) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20023) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20023) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20023(val *InlineResponse20023) *NullableInlineResponse20023 {
	return &NullableInlineResponse20023{value: val, isSet: true}
}

func (v NullableInlineResponse20023) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20023) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


