/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2016 struct for InlineResponse2016
type InlineResponse2016 struct {
	// ID of throughput test job
	ThroughputTestId *string `json:"throughputTestId,omitempty"`
	// GET this url to check the status of your throughput test request
	Url *string `json:"url,omitempty"`
	// Status of the throughput test request
	Status *string `json:"status,omitempty"`
	Result *InlineResponse2016Result `json:"result,omitempty"`
	Request *InlineResponse2016Request `json:"request,omitempty"`
	// Description of the error.
	Error *string `json:"error,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2016 instantiates a new InlineResponse2016 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2016() *InlineResponse2016 {
	this := InlineResponse2016{}
	return &this
}

// NewInlineResponse2016WithDefaults instantiates a new InlineResponse2016 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2016WithDefaults() *InlineResponse2016 {
	this := InlineResponse2016{}
	return &this
}

// GetThroughputTestId returns the ThroughputTestId field value if set, zero value otherwise.
func (o *InlineResponse2016) GetThroughputTestId() string {
	if o == nil || isNil(o.ThroughputTestId) {
		var ret string
		return ret
	}
	return *o.ThroughputTestId
}

// GetThroughputTestIdOk returns a tuple with the ThroughputTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetThroughputTestIdOk() (*string, bool) {
	if o == nil || isNil(o.ThroughputTestId) {
    return nil, false
	}
	return o.ThroughputTestId, true
}

// HasThroughputTestId returns a boolean if a field has been set.
func (o *InlineResponse2016) HasThroughputTestId() bool {
	if o != nil && !isNil(o.ThroughputTestId) {
		return true
	}

	return false
}

// SetThroughputTestId gets a reference to the given string and assigns it to the ThroughputTestId field.
func (o *InlineResponse2016) SetThroughputTestId(v string) {
	o.ThroughputTestId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2016) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2016) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2016) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2016) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2016) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2016) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InlineResponse2016) GetResult() InlineResponse2016Result {
	if o == nil || isNil(o.Result) {
		var ret InlineResponse2016Result
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetResultOk() (*InlineResponse2016Result, bool) {
	if o == nil || isNil(o.Result) {
    return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InlineResponse2016) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given InlineResponse2016Result and assigns it to the Result field.
func (o *InlineResponse2016) SetResult(v InlineResponse2016Result) {
	o.Result = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2016) GetRequest() InlineResponse2016Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2016Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetRequestOk() (*InlineResponse2016Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2016) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2016Request and assigns it to the Request field.
func (o *InlineResponse2016) SetRequest(v InlineResponse2016Request) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse2016) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse2016) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse2016) SetError(v string) {
	o.Error = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2016) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2016) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2016) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2016) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2016) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2016 struct {
	value *InlineResponse2016
	isSet bool
}

func (v NullableInlineResponse2016) Get() *InlineResponse2016 {
	return v.value
}

func (v *NullableInlineResponse2016) Set(val *InlineResponse2016) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2016) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2016) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2016(val *InlineResponse2016) *NullableInlineResponse2016 {
	return &NullableInlineResponse2016{value: val, isSet: true}
}

func (v NullableInlineResponse2016) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2016) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


