/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20025 struct for InlineResponse20025
type InlineResponse20025 struct {
	// Client tracking method of a network
	ClientTrackingMethod *string `json:"clientTrackingMethod,omitempty"`
	// Deployment mode of a network
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	DynamicDns *InlineResponse20025DynamicDns `json:"dynamicDns,omitempty"`
}

// NewInlineResponse20025 instantiates a new InlineResponse20025 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// NewInlineResponse20025WithDefaults instantiates a new InlineResponse20025 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025WithDefaults() *InlineResponse20025 {
	this := InlineResponse20025{}
	return &this
}

// GetClientTrackingMethod returns the ClientTrackingMethod field value if set, zero value otherwise.
func (o *InlineResponse20025) GetClientTrackingMethod() string {
	if o == nil || isNil(o.ClientTrackingMethod) {
		var ret string
		return ret
	}
	return *o.ClientTrackingMethod
}

// GetClientTrackingMethodOk returns a tuple with the ClientTrackingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetClientTrackingMethodOk() (*string, bool) {
	if o == nil || isNil(o.ClientTrackingMethod) {
    return nil, false
	}
	return o.ClientTrackingMethod, true
}

// HasClientTrackingMethod returns a boolean if a field has been set.
func (o *InlineResponse20025) HasClientTrackingMethod() bool {
	if o != nil && !isNil(o.ClientTrackingMethod) {
		return true
	}

	return false
}

// SetClientTrackingMethod gets a reference to the given string and assigns it to the ClientTrackingMethod field.
func (o *InlineResponse20025) SetClientTrackingMethod(v string) {
	o.ClientTrackingMethod = &v
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *InlineResponse20025) GetDeploymentMode() string {
	if o == nil || isNil(o.DeploymentMode) {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetDeploymentModeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentMode) {
    return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *InlineResponse20025) HasDeploymentMode() bool {
	if o != nil && !isNil(o.DeploymentMode) {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *InlineResponse20025) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetDynamicDns returns the DynamicDns field value if set, zero value otherwise.
func (o *InlineResponse20025) GetDynamicDns() InlineResponse20025DynamicDns {
	if o == nil || isNil(o.DynamicDns) {
		var ret InlineResponse20025DynamicDns
		return ret
	}
	return *o.DynamicDns
}

// GetDynamicDnsOk returns a tuple with the DynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025) GetDynamicDnsOk() (*InlineResponse20025DynamicDns, bool) {
	if o == nil || isNil(o.DynamicDns) {
    return nil, false
	}
	return o.DynamicDns, true
}

// HasDynamicDns returns a boolean if a field has been set.
func (o *InlineResponse20025) HasDynamicDns() bool {
	if o != nil && !isNil(o.DynamicDns) {
		return true
	}

	return false
}

// SetDynamicDns gets a reference to the given InlineResponse20025DynamicDns and assigns it to the DynamicDns field.
func (o *InlineResponse20025) SetDynamicDns(v InlineResponse20025DynamicDns) {
	o.DynamicDns = &v
}

func (o InlineResponse20025) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientTrackingMethod) {
		toSerialize["clientTrackingMethod"] = o.ClientTrackingMethod
	}
	if !isNil(o.DeploymentMode) {
		toSerialize["deploymentMode"] = o.DeploymentMode
	}
	if !isNil(o.DynamicDns) {
		toSerialize["dynamicDns"] = o.DynamicDns
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025 struct {
	value *InlineResponse20025
	isSet bool
}

func (v NullableInlineResponse20025) Get() *InlineResponse20025 {
	return v.value
}

func (v *NullableInlineResponse20025) Set(val *InlineResponse20025) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025(val *InlineResponse20025) *NullableInlineResponse20025 {
	return &NullableInlineResponse20025{value: val, isSet: true}
}

func (v NullableInlineResponse20025) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


