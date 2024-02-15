/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
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
	// Web order id
	WebOrderId *string `json:"webOrderId,omitempty"`
	// Products the subscription has entitlements for
	ProductTypes []string `json:"productTypes,omitempty"`
	// Entitlement info
	Entitlements []AdministeredLicensingSubscriptionSubscriptionsEntitlements `json:"entitlements,omitempty"`
	Counts *AdministeredLicensingSubscriptionSubscriptionsCounts `json:"counts,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
    return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *InlineResponse2001) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2001) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2001) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2001) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse2001) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse2001) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse2001) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2001) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2001) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2001) SetStatus(v string) {
	o.Status = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InlineResponse2001) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
    return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineResponse2001) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *InlineResponse2001) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InlineResponse2001) GetEndDate() time.Time {
	if o == nil || isNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndDate) {
    return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineResponse2001) HasEndDate() bool {
	if o != nil && !isNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *InlineResponse2001) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetWebOrderId returns the WebOrderId field value if set, zero value otherwise.
func (o *InlineResponse2001) GetWebOrderId() string {
	if o == nil || isNil(o.WebOrderId) {
		var ret string
		return ret
	}
	return *o.WebOrderId
}

// GetWebOrderIdOk returns a tuple with the WebOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetWebOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.WebOrderId) {
    return nil, false
	}
	return o.WebOrderId, true
}

// HasWebOrderId returns a boolean if a field has been set.
func (o *InlineResponse2001) HasWebOrderId() bool {
	if o != nil && !isNil(o.WebOrderId) {
		return true
	}

	return false
}

// SetWebOrderId gets a reference to the given string and assigns it to the WebOrderId field.
func (o *InlineResponse2001) SetWebOrderId(v string) {
	o.WebOrderId = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse2001) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse2001) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse2001) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *InlineResponse2001) GetEntitlements() []AdministeredLicensingSubscriptionSubscriptionsEntitlements {
	if o == nil || isNil(o.Entitlements) {
		var ret []AdministeredLicensingSubscriptionSubscriptionsEntitlements
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetEntitlementsOk() ([]AdministeredLicensingSubscriptionSubscriptionsEntitlements, bool) {
	if o == nil || isNil(o.Entitlements) {
    return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *InlineResponse2001) HasEntitlements() bool {
	if o != nil && !isNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []AdministeredLicensingSubscriptionSubscriptionsEntitlements and assigns it to the Entitlements field.
func (o *InlineResponse2001) SetEntitlements(v []AdministeredLicensingSubscriptionSubscriptionsEntitlements) {
	o.Entitlements = v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse2001) GetCounts() AdministeredLicensingSubscriptionSubscriptionsCounts {
	if o == nil || isNil(o.Counts) {
		var ret AdministeredLicensingSubscriptionSubscriptionsCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetCountsOk() (*AdministeredLicensingSubscriptionSubscriptionsCounts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse2001) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given AdministeredLicensingSubscriptionSubscriptionsCounts and assigns it to the Counts field.
func (o *InlineResponse2001) SetCounts(v AdministeredLicensingSubscriptionSubscriptionsCounts) {
	o.Counts = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.WebOrderId) {
		toSerialize["webOrderId"] = o.WebOrderId
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


