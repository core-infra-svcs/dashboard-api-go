/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200287 struct for InlineResponse200287
type InlineResponse200287 struct {
	// Name of the SSID
	Name *string `json:"name,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage `json:"usage,omitempty"`
	Clients *OrganizationsOrganizationIdSummaryTopSsidsByUsageClients `json:"clients,omitempty"`
}

// NewInlineResponse200287 instantiates a new InlineResponse200287 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200287() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// NewInlineResponse200287WithDefaults instantiates a new InlineResponse200287 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200287WithDefaults() *InlineResponse200287 {
	this := InlineResponse200287{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200287) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200287) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200287) SetName(v string) {
	o.Name = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200287) GetUsage() OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200287) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200287) SetUsage(v OrganizationsOrganizationIdSummaryTopSsidsByUsageUsage) {
	o.Usage = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200287) GetClients() OrganizationsOrganizationIdSummaryTopSsidsByUsageClients {
	if o == nil || isNil(o.Clients) {
		var ret OrganizationsOrganizationIdSummaryTopSsidsByUsageClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200287) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopSsidsByUsageClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200287) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given OrganizationsOrganizationIdSummaryTopSsidsByUsageClients and assigns it to the Clients field.
func (o *InlineResponse200287) SetClients(v OrganizationsOrganizationIdSummaryTopSsidsByUsageClients) {
	o.Clients = &v
}

func (o InlineResponse200287) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200287 struct {
	value *InlineResponse200287
	isSet bool
}

func (v NullableInlineResponse200287) Get() *InlineResponse200287 {
	return v.value
}

func (v *NullableInlineResponse200287) Set(val *InlineResponse200287) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200287) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200287) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200287(val *InlineResponse200287) *NullableInlineResponse200287 {
	return &NullableInlineResponse200287{value: val, isSet: true}
}

func (v NullableInlineResponse200287) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200287) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

