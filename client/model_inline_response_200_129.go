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

// InlineResponse200129 struct for InlineResponse200129
type InlineResponse200129 struct {
	// Name of the manufacturer
	Name *string `json:"name,omitempty"`
	Clients *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients `json:"clients,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage `json:"usage,omitempty"`
}

// NewInlineResponse200129 instantiates a new InlineResponse200129 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200129() *InlineResponse200129 {
	this := InlineResponse200129{}
	return &this
}

// NewInlineResponse200129WithDefaults instantiates a new InlineResponse200129 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200129WithDefaults() *InlineResponse200129 {
	this := InlineResponse200129{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200129) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200129) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200129) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200129) SetName(v string) {
	o.Name = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200129) GetClients() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients {
	if o == nil || isNil(o.Clients) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200129) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200129) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients and assigns it to the Clients field.
func (o *InlineResponse200129) SetClients(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageClients) {
	o.Clients = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200129) GetUsage() OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200129) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200129) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200129) SetUsage(v OrganizationsOrganizationIdSummaryTopClientsManufacturersByUsageUsage) {
	o.Usage = &v
}

func (o InlineResponse200129) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200129 struct {
	value *InlineResponse200129
	isSet bool
}

func (v NullableInlineResponse200129) Get() *InlineResponse200129 {
	return v.value
}

func (v *NullableInlineResponse200129) Set(val *InlineResponse200129) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200129) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200129) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200129(val *InlineResponse200129) *NullableInlineResponse200129 {
	return &NullableInlineResponse200129{value: val, isSet: true}
}

func (v NullableInlineResponse200129) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200129) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


