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

// OrganizationsOrganizationIdActionBatchesActions1 struct for OrganizationsOrganizationIdActionBatchesActions1
type OrganizationsOrganizationIdActionBatchesActions1 struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used
	Operation string `json:"operation"`
	// The body of the action
	Body map[string]interface{} `json:"body,omitempty"`
}

// NewOrganizationsOrganizationIdActionBatchesActions1 instantiates a new OrganizationsOrganizationIdActionBatchesActions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesActions1(resource string, operation string) *OrganizationsOrganizationIdActionBatchesActions1 {
	this := OrganizationsOrganizationIdActionBatchesActions1{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesActions1WithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesActions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesActions1WithDefaults() *OrganizationsOrganizationIdActionBatchesActions1 {
	this := OrganizationsOrganizationIdActionBatchesActions1{}
	return &this
}

// GetResource returns the Resource field value
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetResourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *OrganizationsOrganizationIdActionBatchesActions1) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetOperationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *OrganizationsOrganizationIdActionBatchesActions1) SetOperation(v string) {
	o.Operation = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetBody() map[string]interface{} {
	if o == nil || isNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions1) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Body) {
    return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions1) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *OrganizationsOrganizationIdActionBatchesActions1) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o OrganizationsOrganizationIdActionBatchesActions1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdActionBatchesActions1 struct {
	value *OrganizationsOrganizationIdActionBatchesActions1
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions1) Get() *OrganizationsOrganizationIdActionBatchesActions1 {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions1) Set(val *OrganizationsOrganizationIdActionBatchesActions1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesActions1(val *OrganizationsOrganizationIdActionBatchesActions1) *NullableOrganizationsOrganizationIdActionBatchesActions1 {
	return &NullableOrganizationsOrganizationIdActionBatchesActions1{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


