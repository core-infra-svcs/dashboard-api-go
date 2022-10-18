/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20092 struct for InlineResponse20092
type InlineResponse20092 struct {
	// Name of client
	Name *string `json:"name,omitempty"`
	// MAC address of client
	Mac *string `json:"mac,omitempty"`
	// ID of client
	Id *string `json:"id,omitempty"`
	Network *OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork `json:"network,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopClientsByUsageUsage `json:"usage,omitempty"`
}

// NewInlineResponse20092 instantiates a new InlineResponse20092 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20092() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20092WithDefaults() *InlineResponse20092 {
	this := InlineResponse20092{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20092) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20092) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20092) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20092) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20092) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20092) SetMac(v string) {
	o.Mac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20092) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20092) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20092) SetId(v string) {
	o.Id = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse20092) GetNetwork() OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork {
	if o == nil || o.Network == nil {
		var ret OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse20092) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork and assigns it to the Network field.
func (o *InlineResponse20092) SetNetwork(v OrganizationsOrganizationIdSummaryTopClientsByUsageNetwork) {
	o.Network = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse20092) GetUsage() OrganizationsOrganizationIdSummaryTopClientsByUsageUsage {
	if o == nil || o.Usage == nil {
		var ret OrganizationsOrganizationIdSummaryTopClientsByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20092) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopClientsByUsageUsage, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse20092) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse20092) SetUsage(v OrganizationsOrganizationIdSummaryTopClientsByUsageUsage) {
	o.Usage = &v
}

func (o InlineResponse20092) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20092 struct {
	value *InlineResponse20092
	isSet bool
}

func (v NullableInlineResponse20092) Get() *InlineResponse20092 {
	return v.value
}

func (v *NullableInlineResponse20092) Set(val *InlineResponse20092) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20092) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20092) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20092(val *InlineResponse20092) *NullableInlineResponse20092 {
	return &NullableInlineResponse20092{value: val, isSet: true}
}

func (v NullableInlineResponse20092) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20092) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


