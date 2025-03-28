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

// InlineResponse20082 struct for InlineResponse20082
type InlineResponse20082 struct {
	// Deployment mode for the cellular gateways in the network. (Passthrough/Routed)
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	// CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.
	Cidr *string `json:"cidr,omitempty"`
	// Mask used for the subnet of all MGs in  this network.
	Mask *int32 `json:"mask,omitempty"`
	// List of subnets of all MGs in this network.
	Subnets []InlineResponse20082Subnets `json:"subnets,omitempty"`
}

// NewInlineResponse20082 instantiates a new InlineResponse20082 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20082() *InlineResponse20082 {
	this := InlineResponse20082{}
	return &this
}

// NewInlineResponse20082WithDefaults instantiates a new InlineResponse20082 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20082WithDefaults() *InlineResponse20082 {
	this := InlineResponse20082{}
	return &this
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *InlineResponse20082) GetDeploymentMode() string {
	if o == nil || isNil(o.DeploymentMode) {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetDeploymentModeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentMode) {
    return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *InlineResponse20082) HasDeploymentMode() bool {
	if o != nil && !isNil(o.DeploymentMode) {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *InlineResponse20082) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20082) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20082) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20082) SetCidr(v string) {
	o.Cidr = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *InlineResponse20082) GetMask() int32 {
	if o == nil || isNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetMaskOk() (*int32, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *InlineResponse20082) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *InlineResponse20082) SetMask(v int32) {
	o.Mask = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *InlineResponse20082) GetSubnets() []InlineResponse20082Subnets {
	if o == nil || isNil(o.Subnets) {
		var ret []InlineResponse20082Subnets
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082) GetSubnetsOk() ([]InlineResponse20082Subnets, bool) {
	if o == nil || isNil(o.Subnets) {
    return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *InlineResponse20082) HasSubnets() bool {
	if o != nil && !isNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []InlineResponse20082Subnets and assigns it to the Subnets field.
func (o *InlineResponse20082) SetSubnets(v []InlineResponse20082Subnets) {
	o.Subnets = v
}

func (o InlineResponse20082) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeploymentMode) {
		toSerialize["deploymentMode"] = o.DeploymentMode
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20082 struct {
	value *InlineResponse20082
	isSet bool
}

func (v NullableInlineResponse20082) Get() *InlineResponse20082 {
	return v.value
}

func (v *NullableInlineResponse20082) Set(val *InlineResponse20082) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20082) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20082) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20082(val *InlineResponse20082) *NullableInlineResponse20082 {
	return &NullableInlineResponse20082{value: val, isSet: true}
}

func (v NullableInlineResponse20082) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20082) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


