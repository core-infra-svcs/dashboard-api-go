/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20051 struct for InlineResponse20051
type InlineResponse20051 struct {
	// Appliance service name
	Service *string `json:"service,omitempty"`
	// A string indicating the rule for which IPs are allowed to use the specified service
	Access *string `json:"access,omitempty"`
	// An array of allowed IPs that can access the service
	AllowedIps []string `json:"allowedIps,omitempty"`
}

// NewInlineResponse20051 instantiates a new InlineResponse20051 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20051() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// NewInlineResponse20051WithDefaults instantiates a new InlineResponse20051 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20051WithDefaults() *InlineResponse20051 {
	this := InlineResponse20051{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *InlineResponse20051) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *InlineResponse20051) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *InlineResponse20051) SetService(v string) {
	o.Service = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse20051) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse20051) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse20051) SetAccess(v string) {
	o.Access = &v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *InlineResponse20051) GetAllowedIps() []string {
	if o == nil || isNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20051) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedIps) {
    return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *InlineResponse20051) HasAllowedIps() bool {
	if o != nil && !isNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *InlineResponse20051) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

func (o InlineResponse20051) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20051 struct {
	value *InlineResponse20051
	isSet bool
}

func (v NullableInlineResponse20051) Get() *InlineResponse20051 {
	return v.value
}

func (v *NullableInlineResponse20051) Set(val *InlineResponse20051) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20051) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20051) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20051(val *InlineResponse20051) *NullableInlineResponse20051 {
	return &NullableInlineResponse20051{value: val, isSet: true}
}

func (v NullableInlineResponse20051) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20051) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


