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

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	// Subscription ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Unbound networks
	Networks []InlineResponse2004Networks `json:"networks,omitempty"`
	// Array of errors if failed
	Errors []string `json:"errors,omitempty"`
	// A list of entitlements required to successfully bind the networks to the subscription
	InsufficientEntitlements []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements `json:"insufficientEntitlements,omitempty"`
}

// NewInlineResponse2004 instantiates a new InlineResponse2004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004WithDefaults() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *InlineResponse2004) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *InlineResponse2004) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *InlineResponse2004) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineResponse2004) GetNetworks() []InlineResponse2004Networks {
	if o == nil || isNil(o.Networks) {
		var ret []InlineResponse2004Networks
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetNetworksOk() ([]InlineResponse2004Networks, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineResponse2004) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []InlineResponse2004Networks and assigns it to the Networks field.
func (o *InlineResponse2004) SetNetworks(v []InlineResponse2004Networks) {
	o.Networks = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse2004) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse2004) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse2004) SetErrors(v []string) {
	o.Errors = v
}

// GetInsufficientEntitlements returns the InsufficientEntitlements field value if set, zero value otherwise.
func (o *InlineResponse2004) GetInsufficientEntitlements() []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements {
	if o == nil || isNil(o.InsufficientEntitlements) {
		var ret []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements
		return ret
	}
	return o.InsufficientEntitlements
}

// GetInsufficientEntitlementsOk returns a tuple with the InsufficientEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetInsufficientEntitlementsOk() ([]AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements, bool) {
	if o == nil || isNil(o.InsufficientEntitlements) {
    return nil, false
	}
	return o.InsufficientEntitlements, true
}

// HasInsufficientEntitlements returns a boolean if a field has been set.
func (o *InlineResponse2004) HasInsufficientEntitlements() bool {
	if o != nil && !isNil(o.InsufficientEntitlements) {
		return true
	}

	return false
}

// SetInsufficientEntitlements gets a reference to the given []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements and assigns it to the InsufficientEntitlements field.
func (o *InlineResponse2004) SetInsufficientEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsMissingEntitlements) {
	o.InsufficientEntitlements = v
}

func (o InlineResponse2004) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.InsufficientEntitlements) {
		toSerialize["insufficientEntitlements"] = o.InsufficientEntitlements
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004 struct {
	value *InlineResponse2004
	isSet bool
}

func (v NullableInlineResponse2004) Get() *InlineResponse2004 {
	return v.value
}

func (v *NullableInlineResponse2004) Set(val *InlineResponse2004) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004(val *InlineResponse2004) *NullableInlineResponse2004 {
	return &NullableInlineResponse2004{value: val, isSet: true}
}

func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


