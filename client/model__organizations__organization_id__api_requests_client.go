/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApiRequestsClient Client information
type OrganizationsOrganizationIdApiRequestsClient struct {
	// ID for the client which made the request, if applicable.
	Id *string `json:"id,omitempty"`
	// Type of client which made the request, if applicable. Available options are: oauth, api_key
	Type *string `json:"type,omitempty"`
}

// NewOrganizationsOrganizationIdApiRequestsClient instantiates a new OrganizationsOrganizationIdApiRequestsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApiRequestsClient() *OrganizationsOrganizationIdApiRequestsClient {
	this := OrganizationsOrganizationIdApiRequestsClient{}
	return &this
}

// NewOrganizationsOrganizationIdApiRequestsClientWithDefaults instantiates a new OrganizationsOrganizationIdApiRequestsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApiRequestsClientWithDefaults() *OrganizationsOrganizationIdApiRequestsClient {
	this := OrganizationsOrganizationIdApiRequestsClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApiRequestsClient) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApiRequestsClient) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApiRequestsClient) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdApiRequestsClient) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApiRequestsClient) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApiRequestsClient) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApiRequestsClient) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationsOrganizationIdApiRequestsClient) SetType(v string) {
	o.Type = &v
}

func (o OrganizationsOrganizationIdApiRequestsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApiRequestsClient struct {
	value *OrganizationsOrganizationIdApiRequestsClient
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApiRequestsClient) Get() *OrganizationsOrganizationIdApiRequestsClient {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApiRequestsClient) Set(val *OrganizationsOrganizationIdApiRequestsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApiRequestsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApiRequestsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApiRequestsClient(val *OrganizationsOrganizationIdApiRequestsClient) *NullableOrganizationsOrganizationIdApiRequestsClient {
	return &NullableOrganizationsOrganizationIdApiRequestsClient{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApiRequestsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApiRequestsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


