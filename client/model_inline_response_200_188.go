/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200188 struct for InlineResponse200188
type InlineResponse200188 struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Network clients list URL
	Url *string `json:"url,omitempty"`
	// Network tags
	Tags []string `json:"tags,omitempty"`
	Clients *OrganizationsOrganizationIdSummaryTopNetworksByStatusClients `json:"clients,omitempty"`
	Statuses *OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses `json:"statuses,omitempty"`
	Devices *OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices `json:"devices,omitempty"`
	// Product types in network
	ProductTypes []string `json:"productTypes,omitempty"`
}

// NewInlineResponse200188 instantiates a new InlineResponse200188 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200188() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// NewInlineResponse200188WithDefaults instantiates a new InlineResponse200188 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200188WithDefaults() *InlineResponse200188 {
	this := InlineResponse200188{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200188) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200188) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200188) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200188) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200188) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200188) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200188) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200188) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200188) SetUrl(v string) {
	o.Url = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200188) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200188) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200188) SetTags(v []string) {
	o.Tags = v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200188) GetClients() OrganizationsOrganizationIdSummaryTopNetworksByStatusClients {
	if o == nil || isNil(o.Clients) {
		var ret OrganizationsOrganizationIdSummaryTopNetworksByStatusClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200188) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given OrganizationsOrganizationIdSummaryTopNetworksByStatusClients and assigns it to the Clients field.
func (o *InlineResponse200188) SetClients(v OrganizationsOrganizationIdSummaryTopNetworksByStatusClients) {
	o.Clients = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *InlineResponse200188) GetStatuses() OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses {
	if o == nil || isNil(o.Statuses) {
		var ret OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetStatusesOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses, bool) {
	if o == nil || isNil(o.Statuses) {
    return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *InlineResponse200188) HasStatuses() bool {
	if o != nil && !isNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses and assigns it to the Statuses field.
func (o *InlineResponse200188) SetStatuses(v OrganizationsOrganizationIdSummaryTopNetworksByStatusStatuses) {
	o.Statuses = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *InlineResponse200188) GetDevices() OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices {
	if o == nil || isNil(o.Devices) {
		var ret OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetDevicesOk() (*OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *InlineResponse200188) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices and assigns it to the Devices field.
func (o *InlineResponse200188) SetDevices(v OrganizationsOrganizationIdSummaryTopNetworksByStatusDevices) {
	o.Devices = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse200188) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200188) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse200188) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse200188) SetProductTypes(v []string) {
	o.ProductTypes = v
}

func (o InlineResponse200188) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !isNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200188 struct {
	value *InlineResponse200188
	isSet bool
}

func (v NullableInlineResponse200188) Get() *InlineResponse200188 {
	return v.value
}

func (v *NullableInlineResponse200188) Set(val *InlineResponse200188) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200188) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200188) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200188(val *InlineResponse200188) *NullableInlineResponse200188 {
	return &NullableInlineResponse200188{value: val, isSet: true}
}

func (v NullableInlineResponse200188) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200188) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


