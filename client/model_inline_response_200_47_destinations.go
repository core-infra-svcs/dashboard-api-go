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

// InlineResponse20047Destinations struct for InlineResponse20047Destinations
type InlineResponse20047Destinations struct {
	// The IP address to test connectivity with
	Ip *string `json:"ip,omitempty"`
	// Description of the testing destination. Optional, defaults to an empty string
	Description *string `json:"description,omitempty"`
	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`
}

// NewInlineResponse20047Destinations instantiates a new InlineResponse20047Destinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047Destinations() *InlineResponse20047Destinations {
	this := InlineResponse20047Destinations{}
	return &this
}

// NewInlineResponse20047DestinationsWithDefaults instantiates a new InlineResponse20047Destinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047DestinationsWithDefaults() *InlineResponse20047Destinations {
	this := InlineResponse20047Destinations{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InlineResponse20047Destinations) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047Destinations) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InlineResponse20047Destinations) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InlineResponse20047Destinations) SetIp(v string) {
	o.Ip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20047Destinations) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047Destinations) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20047Destinations) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20047Destinations) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *InlineResponse20047Destinations) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047Destinations) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *InlineResponse20047Destinations) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *InlineResponse20047Destinations) SetDefault(v bool) {
	o.Default = &v
}

func (o InlineResponse20047Destinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047Destinations struct {
	value *InlineResponse20047Destinations
	isSet bool
}

func (v NullableInlineResponse20047Destinations) Get() *InlineResponse20047Destinations {
	return v.value
}

func (v *NullableInlineResponse20047Destinations) Set(val *InlineResponse20047Destinations) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047Destinations) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047Destinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047Destinations(val *InlineResponse20047Destinations) *NullableInlineResponse20047Destinations {
	return &NullableInlineResponse20047Destinations{value: val, isSet: true}
}

func (v NullableInlineResponse20047Destinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047Destinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

