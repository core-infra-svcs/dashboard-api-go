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

// InlineResponse20063 struct for InlineResponse20063
type InlineResponse20063 struct {
	// Client tracking method of a network
	ClientTrackingMethod *string `json:"clientTrackingMethod,omitempty"`
	// Deployment mode of a network
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	DynamicDns *InlineResponse20063DynamicDns `json:"dynamicDns,omitempty"`
}

// NewInlineResponse20063 instantiates a new InlineResponse20063 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20063() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// NewInlineResponse20063WithDefaults instantiates a new InlineResponse20063 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20063WithDefaults() *InlineResponse20063 {
	this := InlineResponse20063{}
	return &this
}

// GetClientTrackingMethod returns the ClientTrackingMethod field value if set, zero value otherwise.
func (o *InlineResponse20063) GetClientTrackingMethod() string {
	if o == nil || isNil(o.ClientTrackingMethod) {
		var ret string
		return ret
	}
	return *o.ClientTrackingMethod
}

// GetClientTrackingMethodOk returns a tuple with the ClientTrackingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetClientTrackingMethodOk() (*string, bool) {
	if o == nil || isNil(o.ClientTrackingMethod) {
    return nil, false
	}
	return o.ClientTrackingMethod, true
}

// HasClientTrackingMethod returns a boolean if a field has been set.
func (o *InlineResponse20063) HasClientTrackingMethod() bool {
	if o != nil && !isNil(o.ClientTrackingMethod) {
		return true
	}

	return false
}

// SetClientTrackingMethod gets a reference to the given string and assigns it to the ClientTrackingMethod field.
func (o *InlineResponse20063) SetClientTrackingMethod(v string) {
	o.ClientTrackingMethod = &v
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *InlineResponse20063) GetDeploymentMode() string {
	if o == nil || isNil(o.DeploymentMode) {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetDeploymentModeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentMode) {
    return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *InlineResponse20063) HasDeploymentMode() bool {
	if o != nil && !isNil(o.DeploymentMode) {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *InlineResponse20063) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetDynamicDns returns the DynamicDns field value if set, zero value otherwise.
func (o *InlineResponse20063) GetDynamicDns() InlineResponse20063DynamicDns {
	if o == nil || isNil(o.DynamicDns) {
		var ret InlineResponse20063DynamicDns
		return ret
	}
	return *o.DynamicDns
}

// GetDynamicDnsOk returns a tuple with the DynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20063) GetDynamicDnsOk() (*InlineResponse20063DynamicDns, bool) {
	if o == nil || isNil(o.DynamicDns) {
    return nil, false
	}
	return o.DynamicDns, true
}

// HasDynamicDns returns a boolean if a field has been set.
func (o *InlineResponse20063) HasDynamicDns() bool {
	if o != nil && !isNil(o.DynamicDns) {
		return true
	}

	return false
}

// SetDynamicDns gets a reference to the given InlineResponse20063DynamicDns and assigns it to the DynamicDns field.
func (o *InlineResponse20063) SetDynamicDns(v InlineResponse20063DynamicDns) {
	o.DynamicDns = &v
}

func (o InlineResponse20063) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20063 struct {
	value *InlineResponse20063
	isSet bool
}

func (v NullableInlineResponse20063) Get() *InlineResponse20063 {
	return v.value
}

func (v *NullableInlineResponse20063) Set(val *InlineResponse20063) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20063) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20063) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20063(val *InlineResponse20063) *NullableInlineResponse20063 {
	return &NullableInlineResponse20063{value: val, isSet: true}
}

func (v NullableInlineResponse20063) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20063) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


