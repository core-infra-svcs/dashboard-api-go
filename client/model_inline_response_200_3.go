/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	// Subscription's ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Subscription name
	Name *string `json:"name,omitempty"`
	// Subscription description
	Description *string `json:"description,omitempty"`
	// Subscription status
	Status *string `json:"status,omitempty"`
	// Subscription start date
	StartDate *time.Time `json:"startDate,omitempty"`
	// Subscription expiration date
	EndDate *time.Time `json:"endDate,omitempty"`
	// When the subscription was last changed
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	// Web order id
	WebOrderId *string `json:"webOrderId,omitempty"`
	// Subscription type
	Type *string `json:"type,omitempty"`
	SmartAccount *AdministeredLicensingSubscriptionSubscriptionsSmartAccount `json:"smartAccount,omitempty"`
	// Whether a renewal has been requested for the subscription
	RenewalRequested *bool `json:"renewalRequested,omitempty"`
	// Products the subscription has entitlements for
	ProductTypes []string `json:"productTypes,omitempty"`
	// Entitlement info
	Entitlements []AdministeredLicensingSubscriptionSubscriptionsEntitlements `json:"entitlements,omitempty"`
	Counts *AdministeredLicensingSubscriptionSubscriptionsCounts `json:"counts,omitempty"`
	EnterpriseAgreement *AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement `json:"enterpriseAgreement,omitempty"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *InlineResponse2003) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *InlineResponse2003) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *InlineResponse2003) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2003) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2003) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2003) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse2003) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse2003) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse2003) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2003) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2003) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2003) SetStatus(v string) {
	o.Status = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InlineResponse2003) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
    return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineResponse2003) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *InlineResponse2003) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InlineResponse2003) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDate) {
    return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineResponse2003) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *InlineResponse2003) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse2003) GetLastUpdatedAt() time.Time {
	if o == nil || isNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdatedAt) {
    return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse2003) HasLastUpdatedAt() bool {
	if o != nil && !isNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *InlineResponse2003) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetWebOrderId returns the WebOrderId field value if set, zero value otherwise.
func (o *InlineResponse2003) GetWebOrderId() string {
	if o == nil || isNil(o.WebOrderId) {
		var ret string
		return ret
	}
	return *o.WebOrderId
}

// GetWebOrderIdOk returns a tuple with the WebOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetWebOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.WebOrderId) {
    return nil, false
	}
	return o.WebOrderId, true
}

// HasWebOrderId returns a boolean if a field has been set.
func (o *InlineResponse2003) HasWebOrderId() bool {
	if o != nil && !isNil(o.WebOrderId) {
		return true
	}

	return false
}

// SetWebOrderId gets a reference to the given string and assigns it to the WebOrderId field.
func (o *InlineResponse2003) SetWebOrderId(v string) {
	o.WebOrderId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse2003) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse2003) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse2003) SetType(v string) {
	o.Type = &v
}

// GetSmartAccount returns the SmartAccount field value if set, zero value otherwise.
func (o *InlineResponse2003) GetSmartAccount() AdministeredLicensingSubscriptionSubscriptionsSmartAccount {
	if o == nil || isNil(o.SmartAccount) {
		var ret AdministeredLicensingSubscriptionSubscriptionsSmartAccount
		return ret
	}
	return *o.SmartAccount
}

// GetSmartAccountOk returns a tuple with the SmartAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetSmartAccountOk() (*AdministeredLicensingSubscriptionSubscriptionsSmartAccount, bool) {
	if o == nil || isNil(o.SmartAccount) {
    return nil, false
	}
	return o.SmartAccount, true
}

// HasSmartAccount returns a boolean if a field has been set.
func (o *InlineResponse2003) HasSmartAccount() bool {
	if o != nil && !isNil(o.SmartAccount) {
		return true
	}

	return false
}

// SetSmartAccount gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsSmartAccount and assigns it to the SmartAccount field.
func (o *InlineResponse2003) SetSmartAccount(v AdministeredLicensingSubscriptionSubscriptionsSmartAccount) {
	o.SmartAccount = &v
}

// GetRenewalRequested returns the RenewalRequested field value if set, zero value otherwise.
func (o *InlineResponse2003) GetRenewalRequested() bool {
	if o == nil || isNil(o.RenewalRequested) {
		var ret bool
		return ret
	}
	return *o.RenewalRequested
}

// GetRenewalRequestedOk returns a tuple with the RenewalRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetRenewalRequestedOk() (*bool, bool) {
	if o == nil || isNil(o.RenewalRequested) {
    return nil, false
	}
	return o.RenewalRequested, true
}

// HasRenewalRequested returns a boolean if a field has been set.
func (o *InlineResponse2003) HasRenewalRequested() bool {
	if o != nil && !isNil(o.RenewalRequested) {
		return true
	}

	return false
}

// SetRenewalRequested gets a reference to the given bool and assigns it to the RenewalRequested field.
func (o *InlineResponse2003) SetRenewalRequested(v bool) {
	o.RenewalRequested = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse2003) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse2003) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse2003) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *InlineResponse2003) GetEntitlements() []AdministeredLicensingSubscriptionSubscriptionsEntitlements {
	if o == nil || isNil(o.Entitlements) {
		var ret []AdministeredLicensingSubscriptionSubscriptionsEntitlements
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetEntitlementsOk() ([]AdministeredLicensingSubscriptionSubscriptionsEntitlements, bool) {
	if o == nil || isNil(o.Entitlements) {
    return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *InlineResponse2003) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []AdministeredLicensingSubscriptionSubscriptionsEntitlements and assigns it to the Entitlements field.
func (o *InlineResponse2003) SetEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsEntitlements) {
	o.Entitlements = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse2003) GetCounts() AdministeredLicensingSubscriptionSubscriptionsCounts {
	if o == nil || isNil(o.Counts) {
		var ret AdministeredLicensingSubscriptionSubscriptionsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetCountsOk() (*AdministeredLicensingSubscriptionSubscriptionsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse2003) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsCounts and assigns it to the Counts field.
func (o *InlineResponse2003) SetCounts(v AdministeredLicensingSubscriptionSubscriptionsCounts) {
	o.Counts = &v
}

// GetEnterpriseAgreement returns the EnterpriseAgreement field value if set, zero value otherwise.
func (o *InlineResponse2003) GetEnterpriseAgreement() AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement {
	if o == nil || isNil(o.EnterpriseAgreement) {
		var ret AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement
		return ret
	}
	return *o.EnterpriseAgreement
}

// GetEnterpriseAgreementOk returns a tuple with the EnterpriseAgreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetEnterpriseAgreementOk() (*AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement, bool) {
	if o == nil || isNil(o.EnterpriseAgreement) {
    return nil, false
	}
	return o.EnterpriseAgreement, true
}

// HasEnterpriseAgreement returns a boolean if a field has been set.
func (o *InlineResponse2003) HasEnterpriseAgreement() bool {
	if o != nil && !isNil(o.EnterpriseAgreement) {
		return true
	}

	return false
}

// SetEnterpriseAgreement gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement and assigns it to the EnterpriseAgreement field.
func (o *InlineResponse2003) SetEnterpriseAgreement(v AdministeredLicensingSubscriptionSubscriptionsEnterpriseAgreement) {
	o.EnterpriseAgreement = &v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !isNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !isNil(o.WebOrderId) {
		toSerialize["webOrderId"] = o.WebOrderId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.SmartAccount) {
		toSerialize["smartAccount"] = o.SmartAccount
	}
	if !isNil(o.RenewalRequested) {
		toSerialize["renewalRequested"] = o.RenewalRequested
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !isNil(o.EnterpriseAgreement) {
		toSerialize["enterpriseAgreement"] = o.EnterpriseAgreement
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


