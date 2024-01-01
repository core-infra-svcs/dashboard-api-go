/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200173 struct for InlineResponse200173
type InlineResponse200173 struct {
	// Name of the manufacturer
	Name *string `json:"name,omitempty"`
	Clients *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients `json:"clients,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage `json:"usage,omitempty"`
}

// NewInlineResponse200173 instantiates a new InlineResponse200173 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200173() *InlineResponse200173 {
	this := InlineResponse200173{}
	return &this
}

// NewInlineResponse200173WithDefaults instantiates a new InlineResponse200173 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200173WithDefaults() *InlineResponse200173 {
	this := InlineResponse200173{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200173) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200173) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200173) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200173) SetName(v string) {
	o.Name = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200173) GetClients() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	if o == nil || isNil(o.Clients) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200173) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200173) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients and assigns it to the Clients field.
func (o *InlineResponse200173) SetClients(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) {
	o.Clients = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200173) GetUsage() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200173) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200173) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200173) SetUsage(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) {
	o.Usage = &v
}

func (o InlineResponse200173) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200173 struct {
	value *InlineResponse200173
	isSet bool
}

func (v NullableInlineResponse200173) Get() *InlineResponse200173 {
	return v.value
}

func (v *NullableInlineResponse200173) Set(val *InlineResponse200173) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200173) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200173) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200173(val *InlineResponse200173) *NullableInlineResponse200173 {
	return &NullableInlineResponse200173{value: val, isSet: true}
}

func (v NullableInlineResponse200173) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200173) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


