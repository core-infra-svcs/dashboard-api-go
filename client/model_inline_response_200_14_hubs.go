/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20014Hubs struct for InlineResponse20014Hubs
type InlineResponse20014Hubs struct {
	// The network ID of the hub.
	HubId *string `json:"hubId,omitempty"`
	// Indicates whether default route traffic should be sent to this hub.
	UseDefaultRoute *bool `json:"useDefaultRoute,omitempty"`
}

// NewInlineResponse20014Hubs instantiates a new InlineResponse20014Hubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014Hubs() *InlineResponse20014Hubs {
	this := InlineResponse20014Hubs{}
	return &this
}

// NewInlineResponse20014HubsWithDefaults instantiates a new InlineResponse20014Hubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014HubsWithDefaults() *InlineResponse20014Hubs {
	this := InlineResponse20014Hubs{}
	return &this
}

// GetHubId returns the HubId field value if set, zero value otherwise.
func (o *InlineResponse20014Hubs) GetHubId() string {
	if o == nil || isNil(o.HubId) {
		var ret string
		return ret
	}
	return *o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Hubs) GetHubIdOk() (*string, bool) {
	if o == nil || isNil(o.HubId) {
    return nil, false
	}
	return o.HubId, true
}

// HasHubId returns a boolean if a field has been set.
func (o *InlineResponse20014Hubs) HasHubId() bool {
	if o != nil && !isNil(o.HubId) {
		return true
	}

	return false
}

// SetHubId gets a reference to the given string and assigns it to the HubId field.
func (o *InlineResponse20014Hubs) SetHubId(v string) {
	o.HubId = &v
}

// GetUseDefaultRoute returns the UseDefaultRoute field value if set, zero value otherwise.
func (o *InlineResponse20014Hubs) GetUseDefaultRoute() bool {
	if o == nil || isNil(o.UseDefaultRoute) {
		var ret bool
		return ret
	}
	return *o.UseDefaultRoute
}

// GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014Hubs) GetUseDefaultRouteOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefaultRoute) {
    return nil, false
	}
	return o.UseDefaultRoute, true
}

// HasUseDefaultRoute returns a boolean if a field has been set.
func (o *InlineResponse20014Hubs) HasUseDefaultRoute() bool {
	if o != nil && !isNil(o.UseDefaultRoute) {
		return true
	}

	return false
}

// SetUseDefaultRoute gets a reference to the given bool and assigns it to the UseDefaultRoute field.
func (o *InlineResponse20014Hubs) SetUseDefaultRoute(v bool) {
	o.UseDefaultRoute = &v
}

func (o InlineResponse20014Hubs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HubId) {
		toSerialize["hubId"] = o.HubId
	}
	if !isNil(o.UseDefaultRoute) {
		toSerialize["useDefaultRoute"] = o.UseDefaultRoute
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20014Hubs struct {
	value *InlineResponse20014Hubs
	isSet bool
}

func (v NullableInlineResponse20014Hubs) Get() *InlineResponse20014Hubs {
	return v.value
}

func (v *NullableInlineResponse20014Hubs) Set(val *InlineResponse20014Hubs) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20014Hubs) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20014Hubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20014Hubs(val *InlineResponse20014Hubs) *NullableInlineResponse20014Hubs {
	return &NullableInlineResponse20014Hubs{value: val, isSet: true}
}

func (v NullableInlineResponse20014Hubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20014Hubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

