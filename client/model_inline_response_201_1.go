/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2011 struct for InlineResponse2011
type InlineResponse2011 struct {
	// Id of the ARP table request. Used to check the status of the request.
	ArpTableId *string `json:"arpTableId,omitempty"`
	// GET this url to check the status of your ARP table request.
	Url *string `json:"url,omitempty"`
	Request *InlineResponse2011Request `json:"request,omitempty"`
	// Status of the ARP table request.
	Status *string `json:"status,omitempty"`
	Callback *InlineResponse2011Callback `json:"callback,omitempty"`
}

// NewInlineResponse2011 instantiates a new InlineResponse2011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2011() *InlineResponse2011 {
	this := InlineResponse2011{}
	return &this
}

// NewInlineResponse2011WithDefaults instantiates a new InlineResponse2011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2011WithDefaults() *InlineResponse2011 {
	this := InlineResponse2011{}
	return &this
}

// GetArpTableId returns the ArpTableId field value if set, zero value otherwise.
func (o *InlineResponse2011) GetArpTableId() string {
	if o == nil || isNil(o.ArpTableId) {
		var ret string
		return ret
	}
	return *o.ArpTableId
}

// GetArpTableIdOk returns a tuple with the ArpTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetArpTableIdOk() (*string, bool) {
	if o == nil || isNil(o.ArpTableId) {
    return nil, false
	}
	return o.ArpTableId, true
}

// HasArpTableId returns a boolean if a field has been set.
func (o *InlineResponse2011) HasArpTableId() bool {
	if o != nil && !isNil(o.ArpTableId) {
		return true
	}

	return false
}

// SetArpTableId gets a reference to the given string and assigns it to the ArpTableId field.
func (o *InlineResponse2011) SetArpTableId(v string) {
	o.ArpTableId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2011) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2011) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2011) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineResponse2011) GetRequest() InlineResponse2011Request {
	if o == nil || isNil(o.Request) {
		var ret InlineResponse2011Request
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetRequestOk() (*InlineResponse2011Request, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineResponse2011) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineResponse2011Request and assigns it to the Request field.
func (o *InlineResponse2011) SetRequest(v InlineResponse2011Request) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2011) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2011) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2011) SetStatus(v string) {
	o.Status = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineResponse2011) GetCallback() InlineResponse2011Callback {
	if o == nil || isNil(o.Callback) {
		var ret InlineResponse2011Callback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2011) GetCallbackOk() (*InlineResponse2011Callback, bool) {
	if o == nil || isNil(o.Callback) {
    return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineResponse2011) HasCallback() bool {
	if o != nil && !isNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given InlineResponse2011Callback and assigns it to the Callback field.
func (o *InlineResponse2011) SetCallback(v InlineResponse2011Callback) {
	o.Callback = &v
}

func (o InlineResponse2011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ArpTableId) {
		toSerialize["arpTableId"] = o.ArpTableId
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

type NullableInlineResponse2011 struct {
	value *InlineResponse2011
	isSet bool
}

func (v NullableInlineResponse2011) Get() *InlineResponse2011 {
	return v.value
}

func (v *NullableInlineResponse2011) Set(val *InlineResponse2011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2011(val *InlineResponse2011) *NullableInlineResponse2011 {
	return &NullableInlineResponse2011{value: val, isSet: true}
}

func (v NullableInlineResponse2011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


