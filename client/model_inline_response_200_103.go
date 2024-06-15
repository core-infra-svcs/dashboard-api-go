/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200103 struct for InlineResponse200103
type InlineResponse200103 struct {
	// The network or organization identifier
	Id *string `json:"id,omitempty"`
	// If the data returned is organization-wide. False indicates the data is network-wide.
	OrganizationWide *bool `json:"organizationWide,omitempty"`
	// The network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// The type of PII request
	Type *string `json:"type,omitempty"`
	// The MAC address of the PII request
	Mac *string `json:"mac,omitempty"`
	// The stringified array of datasets related to the provided key that should be deleted.
	Datasets *string `json:"datasets,omitempty"`
	// The status of the PII request
	Status *string `json:"status,omitempty"`
	// The request's creation time
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// The request's completion time
	CompletedAt *int32 `json:"completedAt,omitempty"`
}

// NewInlineResponse200103 instantiates a new InlineResponse200103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200103() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200103WithDefaults() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200103) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200103) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200103) SetId(v string) {
	o.Id = &v
}

// GetOrganizationWide returns the OrganizationWide field value if set, zero value otherwise.
func (o *InlineResponse200103) GetOrganizationWide() bool {
	if o == nil || isNil(o.OrganizationWide) {
		var ret bool
		return ret
	}
	return *o.OrganizationWide
}

// GetOrganizationWideOk returns a tuple with the OrganizationWide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetOrganizationWideOk() (*bool, bool) {
	if o == nil || isNil(o.OrganizationWide) {
    return nil, false
	}
	return o.OrganizationWide, true
}

// HasOrganizationWide returns a boolean if a field has been set.
func (o *InlineResponse200103) HasOrganizationWide() bool {
	if o != nil && !isNil(o.OrganizationWide) {
		return true
	}

	return false
}

// SetOrganizationWide gets a reference to the given bool and assigns it to the OrganizationWide field.
func (o *InlineResponse200103) SetOrganizationWide(v bool) {
	o.OrganizationWide = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200103) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200103) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200103) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200103) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200103) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200103) SetType(v string) {
	o.Type = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200103) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200103) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200103) SetMac(v string) {
	o.Mac = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *InlineResponse200103) GetDatasets() string {
	if o == nil || isNil(o.Datasets) {
		var ret string
		return ret
	}
	return *o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetDatasetsOk() (*string, bool) {
	if o == nil || isNil(o.Datasets) {
    return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *InlineResponse200103) HasDatasets() bool {
	if o != nil && !isNil(o.Datasets) {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given string and assigns it to the Datasets field.
func (o *InlineResponse200103) SetDatasets(v string) {
	o.Datasets = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200103) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200103) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200103) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200103) GetCreatedAt() int32 {
	if o == nil || isNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetCreatedAtOk() (*int32, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200103) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *InlineResponse200103) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *InlineResponse200103) GetCompletedAt() int32 {
	if o == nil || isNil(o.CompletedAt) {
		var ret int32
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetCompletedAtOk() (*int32, bool) {
	if o == nil || isNil(o.CompletedAt) {
    return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *InlineResponse200103) HasCompletedAt() bool {
	if o != nil && !isNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given int32 and assigns it to the CompletedAt field.
func (o *InlineResponse200103) SetCompletedAt(v int32) {
	o.CompletedAt = &v
}

func (o InlineResponse200103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationWide) {
		toSerialize["organizationWide"] = o.OrganizationWide
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Datasets) {
		toSerialize["datasets"] = o.Datasets
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200103 struct {
	value *InlineResponse200103
	isSet bool
}

func (v NullableInlineResponse200103) Get() *InlineResponse200103 {
	return v.value
}

func (v *NullableInlineResponse200103) Set(val *InlineResponse200103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200103(val *InlineResponse200103) *NullableInlineResponse200103 {
	return &NullableInlineResponse200103{value: val, isSet: true}
}

func (v NullableInlineResponse200103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


