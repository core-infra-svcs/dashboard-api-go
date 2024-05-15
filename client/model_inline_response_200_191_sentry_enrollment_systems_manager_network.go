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

// InlineResponse200191SentryEnrollmentSystemsManagerNetwork Systems Manager network targeted for sentry enrollment.
type InlineResponse200191SentryEnrollmentSystemsManagerNetwork struct {
	// The network ID of the Systems Manager network.
	Id *string `json:"id,omitempty"`
}

// NewInlineResponse200191SentryEnrollmentSystemsManagerNetwork instantiates a new InlineResponse200191SentryEnrollmentSystemsManagerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200191SentryEnrollmentSystemsManagerNetwork() *InlineResponse200191SentryEnrollmentSystemsManagerNetwork {
	this := InlineResponse200191SentryEnrollmentSystemsManagerNetwork{}
	return &this
}

// NewInlineResponse200191SentryEnrollmentSystemsManagerNetworkWithDefaults instantiates a new InlineResponse200191SentryEnrollmentSystemsManagerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200191SentryEnrollmentSystemsManagerNetworkWithDefaults() *InlineResponse200191SentryEnrollmentSystemsManagerNetwork {
	this := InlineResponse200191SentryEnrollmentSystemsManagerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) SetId(v string) {
	o.Id = &v
}

func (o InlineResponse200191SentryEnrollmentSystemsManagerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork struct {
	value *InlineResponse200191SentryEnrollmentSystemsManagerNetwork
	isSet bool
}

func (v NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) Get() *InlineResponse200191SentryEnrollmentSystemsManagerNetwork {
	return v.value
}

func (v *NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) Set(val *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork(val *InlineResponse200191SentryEnrollmentSystemsManagerNetwork) *NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork {
	return &NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork{value: val, isSet: true}
}

func (v NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200191SentryEnrollmentSystemsManagerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


