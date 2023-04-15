/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20034 struct for InlineResponse20034
type InlineResponse20034 struct {
	// Name of client
	Name *string `json:"name,omitempty"`
	// ID of client
	ClientId *string `json:"clientId,omitempty"`
	// Assigned policies
	Assigned []NetworksNetworkIdPoliciesByClientAssigned `json:"assigned,omitempty"`
}

// NewInlineResponse20034 instantiates a new InlineResponse20034 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20034() *InlineResponse20034 {
	this := InlineResponse20034{}
	return &this
}

// NewInlineResponse20034WithDefaults instantiates a new InlineResponse20034 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20034WithDefaults() *InlineResponse20034 {
	this := InlineResponse20034{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20034) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20034) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20034) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20034) SetName(v string) {
	o.Name = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse20034) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20034) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse20034) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse20034) SetClientId(v string) {
	o.ClientId = &v
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *InlineResponse20034) GetAssigned() []NetworksNetworkIdPoliciesByClientAssigned {
	if o == nil || isNil(o.Assigned) {
		var ret []NetworksNetworkIdPoliciesByClientAssigned
		return ret
	}
	return o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20034) GetAssignedOk() ([]NetworksNetworkIdPoliciesByClientAssigned, bool) {
	if o == nil || isNil(o.Assigned) {
    return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *InlineResponse20034) HasAssigned() bool {
	if o != nil && !isNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given []NetworksNetworkIdPoliciesByClientAssigned and assigns it to the Assigned field.
func (o *InlineResponse20034) SetAssigned(v []NetworksNetworkIdPoliciesByClientAssigned) {
	o.Assigned = v
}

func (o InlineResponse20034) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20034 struct {
	value *InlineResponse20034
	isSet bool
}

func (v NullableInlineResponse20034) Get() *InlineResponse20034 {
	return v.value
}

func (v *NullableInlineResponse20034) Set(val *InlineResponse20034) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20034) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20034) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20034(val *InlineResponse20034) *NullableInlineResponse20034 {
	return &NullableInlineResponse20034{value: val, isSet: true}
}

func (v NullableInlineResponse20034) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20034) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


