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

// InlineResponse200296 struct for InlineResponse200296
type InlineResponse200296 struct {
	// Name of client
	Name *string `json:"name,omitempty"`
	// MAC address of client
	Mac *string `json:"mac,omitempty"`
	// ID of client
	Id *string `json:"id,omitempty"`
	Network *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork `json:"network,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage `json:"usage,omitempty"`
}

// NewInlineResponse200296 instantiates a new InlineResponse200296 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200296() *InlineResponse200296 {
	this := InlineResponse200296{}
	return &this
}

// NewInlineResponse200296WithDefaults instantiates a new InlineResponse200296 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200296WithDefaults() *InlineResponse200296 {
	this := InlineResponse200296{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200296) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200296) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200296) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200296) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200296) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200296) SetMac(v string) {
	o.Mac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200296) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200296) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200296) SetId(v string) {
	o.Id = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200296) GetNetwork() OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200296) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork and assigns it to the Network field.
func (o *InlineResponse200296) SetNetwork(v OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) {
	o.Network = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200296) GetUsage() OrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopClientsByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200296) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopClientsByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200296) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200296) SetUsage(v OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) {
	o.Usage = &v
}

func (o InlineResponse200296) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200296 struct {
	value *InlineResponse200296
	isSet bool
}

func (v NullableInlineResponse200296) Get() *InlineResponse200296 {
	return v.value
}

func (v *NullableInlineResponse200296) Set(val *InlineResponse200296) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200296) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200296) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200296(val *InlineResponse200296) *NullableInlineResponse200296 {
	return &NullableInlineResponse200296{value: val, isSet: true}
}

func (v NullableInlineResponse200296) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200296) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


