/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20110Clients struct for InlineResponse20110Clients
type InlineResponse20110Clients struct {
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// The identifier of the client
	ClientId *string `json:"clientId,omitempty"`
	// The name of the client
	Name *string `json:"name,omitempty"`
	// The client's display message if its group policy is 'Blocked'
	Message *string `json:"message,omitempty"`
}

// NewInlineResponse20110Clients instantiates a new InlineResponse20110Clients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20110Clients() *InlineResponse20110Clients {
	this := InlineResponse20110Clients{}
	return &this
}

// NewInlineResponse20110ClientsWithDefaults instantiates a new InlineResponse20110Clients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20110ClientsWithDefaults() *InlineResponse20110Clients {
	this := InlineResponse20110Clients{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20110Clients) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Clients) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20110Clients) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20110Clients) SetMac(v string) {
	o.Mac = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse20110Clients) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Clients) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse20110Clients) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse20110Clients) SetClientId(v string) {
	o.ClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20110Clients) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Clients) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20110Clients) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20110Clients) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20110Clients) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20110Clients) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20110Clients) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20110Clients) SetMessage(v string) {
	o.Message = &v
}

func (o InlineResponse20110Clients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20110Clients struct {
	value *InlineResponse20110Clients
	isSet bool
}

func (v NullableInlineResponse20110Clients) Get() *InlineResponse20110Clients {
	return v.value
}

func (v *NullableInlineResponse20110Clients) Set(val *InlineResponse20110Clients) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20110Clients) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20110Clients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20110Clients(val *InlineResponse20110Clients) *NullableInlineResponse20110Clients {
	return &NullableInlineResponse20110Clients{value: val, isSet: true}
}

func (v NullableInlineResponse20110Clients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20110Clients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


