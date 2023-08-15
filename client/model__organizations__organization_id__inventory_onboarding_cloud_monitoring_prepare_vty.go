/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty VTY Related Parameters
type OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty struct {
	// Starting line VTY number
	StartLineNumber *int32 `json:"startLineNumber,omitempty"`
	// Ending line VTY number
	EndLineNumber *int32 `json:"endLineNumber,omitempty"`
	Authentication *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication `json:"authentication,omitempty"`
	Authorization *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization `json:"authorization,omitempty"`
	AccessList *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList `json:"accessList,omitempty"`
	// SSH rotary number
	RotaryNumber *int32 `json:"rotaryNumber,omitempty"`
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty{}
	return &this
}

// NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyWithDefaults instantiates a new OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyWithDefaults() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty {
	this := OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty{}
	return &this
}

// GetStartLineNumber returns the StartLineNumber field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetStartLineNumber() int32 {
	if o == nil || isNil(o.StartLineNumber) {
		var ret int32
		return ret
	}
	return *o.StartLineNumber
}

// GetStartLineNumberOk returns a tuple with the StartLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetStartLineNumberOk() (*int32, bool) {
	if o == nil || isNil(o.StartLineNumber) {
    return nil, false
	}
	return o.StartLineNumber, true
}

// HasStartLineNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasStartLineNumber() bool {
	if o != nil && !isNil(o.StartLineNumber) {
		return true
	}

	return false
}

// SetStartLineNumber gets a reference to the given int32 and assigns it to the StartLineNumber field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetStartLineNumber(v int32) {
	o.StartLineNumber = &v
}

// GetEndLineNumber returns the EndLineNumber field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetEndLineNumber() int32 {
	if o == nil || isNil(o.EndLineNumber) {
		var ret int32
		return ret
	}
	return *o.EndLineNumber
}

// GetEndLineNumberOk returns a tuple with the EndLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetEndLineNumberOk() (*int32, bool) {
	if o == nil || isNil(o.EndLineNumber) {
    return nil, false
	}
	return o.EndLineNumber, true
}

// HasEndLineNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasEndLineNumber() bool {
	if o != nil && !isNil(o.EndLineNumber) {
		return true
	}

	return false
}

// SetEndLineNumber gets a reference to the given int32 and assigns it to the EndLineNumber field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetEndLineNumber(v int32) {
	o.EndLineNumber = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAuthentication() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAuthenticationOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication and assigns it to the Authentication field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetAuthentication(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthentication) {
	o.Authentication = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAuthorization() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization {
	if o == nil || isNil(o.Authorization) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAuthorizationOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization, bool) {
	if o == nil || isNil(o.Authorization) {
    return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasAuthorization() bool {
	if o != nil && !isNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization and assigns it to the Authorization field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetAuthorization(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAuthorization) {
	o.Authorization = &v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAccessList() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList {
	if o == nil || isNil(o.AccessList) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList
		return ret
	}
	return *o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetAccessListOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList, bool) {
	if o == nil || isNil(o.AccessList) {
    return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasAccessList() bool {
	if o != nil && !isNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList and assigns it to the AccessList field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetAccessList(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVtyAccessList) {
	o.AccessList = &v
}

// GetRotaryNumber returns the RotaryNumber field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetRotaryNumber() int32 {
	if o == nil || isNil(o.RotaryNumber) {
		var ret int32
		return ret
	}
	return *o.RotaryNumber
}

// GetRotaryNumberOk returns a tuple with the RotaryNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) GetRotaryNumberOk() (*int32, bool) {
	if o == nil || isNil(o.RotaryNumber) {
    return nil, false
	}
	return o.RotaryNumber, true
}

// HasRotaryNumber returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) HasRotaryNumber() bool {
	if o != nil && !isNil(o.RotaryNumber) {
		return true
	}

	return false
}

// SetRotaryNumber gets a reference to the given int32 and assigns it to the RotaryNumber field.
func (o *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) SetRotaryNumber(v int32) {
	o.RotaryNumber = &v
}

func (o OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartLineNumber) {
		toSerialize["startLineNumber"] = o.StartLineNumber
	}
	if !isNil(o.EndLineNumber) {
		toSerialize["endLineNumber"] = o.EndLineNumber
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !isNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	if !isNil(o.AccessList) {
		toSerialize["accessList"] = o.AccessList
	}
	if !isNil(o.RotaryNumber) {
		toSerialize["rotaryNumber"] = o.RotaryNumber
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty struct {
	value *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty
	isSet bool
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) Get() *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) Set(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty(val *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty {
	return &NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareVty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


