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

// InlineResponse200353Items struct for InlineResponse200353Items
type InlineResponse200353Items struct {
	// Network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Network Name
	Name *string `json:"name,omitempty"`
	// Whether to enable collection of location and scanning analytics
	Enabled *bool `json:"enabled,omitempty"`
	Api *InlineResponse200194Api `json:"api,omitempty"`
}

// NewInlineResponse200353Items instantiates a new InlineResponse200353Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200353Items() *InlineResponse200353Items {
	this := InlineResponse200353Items{}
	return &this
}

// NewInlineResponse200353ItemsWithDefaults instantiates a new InlineResponse200353Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200353ItemsWithDefaults() *InlineResponse200353Items {
	this := InlineResponse200353Items{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200353Items) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Items) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200353Items) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200353Items) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200353Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200353Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200353Items) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200353Items) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Items) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200353Items) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200353Items) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *InlineResponse200353Items) GetApi() InlineResponse200194Api {
	if o == nil || isNil(o.Api) {
		var ret InlineResponse200194Api
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200353Items) GetApiOk() (*InlineResponse200194Api, bool) {
	if o == nil || isNil(o.Api) {
    return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *InlineResponse200353Items) HasApi() bool {
	if o != nil && !isNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given InlineResponse200194Api and assigns it to the Api field.
func (o *InlineResponse200353Items) SetApi(v InlineResponse200194Api) {
	o.Api = &v
}

func (o InlineResponse200353Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200353Items struct {
	value *InlineResponse200353Items
	isSet bool
}

func (v NullableInlineResponse200353Items) Get() *InlineResponse200353Items {
	return v.value
}

func (v *NullableInlineResponse200353Items) Set(val *InlineResponse200353Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200353Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200353Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200353Items(val *InlineResponse200353Items) *NullableInlineResponse200353Items {
	return &NullableInlineResponse200353Items{value: val, isSet: true}
}

func (v NullableInlineResponse200353Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200353Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


