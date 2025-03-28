/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200143 struct for InlineResponse200143
type InlineResponse200143 struct {
	// device ID
	Id *string `json:"id,omitempty"`
	// SSID name
	SsidName *string `json:"ssidName,omitempty"`
	// device name
	Name *string `json:"name,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty"`
	// device tags
	Tags []string `json:"tags,omitempty"`
	// type of access period, either a static range or a dynamic period
	TimeboundType *string `json:"timeboundType,omitempty"`
	// Send Email Notifications
	SendExpirationEmails *bool `json:"sendExpirationEmails,omitempty"`
	// Time before access expiration reminder email sends
	NotifyTimeBeforeAccessEnds *int32 `json:"notifyTimeBeforeAccessEnds,omitempty"`
	// Optional email text
	AdditionalEmailText *string `json:"additionalEmailText,omitempty"`
	// time that access starts
	AccessStartAt *time.Time `json:"accessStartAt,omitempty"`
	// time that access ends
	AccessEndAt *time.Time `json:"accessEndAt,omitempty"`
}

// NewInlineResponse200143 instantiates a new InlineResponse200143 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200143() *InlineResponse200143 {
	this := InlineResponse200143{}
	return &this
}

// NewInlineResponse200143WithDefaults instantiates a new InlineResponse200143 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200143WithDefaults() *InlineResponse200143 {
	this := InlineResponse200143{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200143) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200143) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200143) SetId(v string) {
	o.Id = &v
}

// GetSsidName returns the SsidName field value if set, zero value otherwise.
func (o *InlineResponse200143) GetSsidName() string {
	if o == nil || isNil(o.SsidName) {
		var ret string
		return ret
	}
	return *o.SsidName
}

// GetSsidNameOk returns a tuple with the SsidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetSsidNameOk() (*string, bool) {
	if o == nil || isNil(o.SsidName) {
    return nil, false
	}
	return o.SsidName, true
}

// HasSsidName returns a boolean if a field has been set.
func (o *InlineResponse200143) HasSsidName() bool {
	if o != nil && !isNil(o.SsidName) {
		return true
	}

	return false
}

// SetSsidName gets a reference to the given string and assigns it to the SsidName field.
func (o *InlineResponse200143) SetSsidName(v string) {
	o.SsidName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200143) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200143) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200143) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200143) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200143) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineResponse200143) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200143) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200143) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200143) SetTags(v []string) {
	o.Tags = v
}

// GetTimeboundType returns the TimeboundType field value if set, zero value otherwise.
func (o *InlineResponse200143) GetTimeboundType() string {
	if o == nil || isNil(o.TimeboundType) {
		var ret string
		return ret
	}
	return *o.TimeboundType
}

// GetTimeboundTypeOk returns a tuple with the TimeboundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetTimeboundTypeOk() (*string, bool) {
	if o == nil || isNil(o.TimeboundType) {
    return nil, false
	}
	return o.TimeboundType, true
}

// HasTimeboundType returns a boolean if a field has been set.
func (o *InlineResponse200143) HasTimeboundType() bool {
	if o != nil && !isNil(o.TimeboundType) {
		return true
	}

	return false
}

// SetTimeboundType gets a reference to the given string and assigns it to the TimeboundType field.
func (o *InlineResponse200143) SetTimeboundType(v string) {
	o.TimeboundType = &v
}

// GetSendExpirationEmails returns the SendExpirationEmails field value if set, zero value otherwise.
func (o *InlineResponse200143) GetSendExpirationEmails() bool {
	if o == nil || isNil(o.SendExpirationEmails) {
		var ret bool
		return ret
	}
	return *o.SendExpirationEmails
}

// GetSendExpirationEmailsOk returns a tuple with the SendExpirationEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetSendExpirationEmailsOk() (*bool, bool) {
	if o == nil || isNil(o.SendExpirationEmails) {
    return nil, false
	}
	return o.SendExpirationEmails, true
}

// HasSendExpirationEmails returns a boolean if a field has been set.
func (o *InlineResponse200143) HasSendExpirationEmails() bool {
	if o != nil && !isNil(o.SendExpirationEmails) {
		return true
	}

	return false
}

// SetSendExpirationEmails gets a reference to the given bool and assigns it to the SendExpirationEmails field.
func (o *InlineResponse200143) SetSendExpirationEmails(v bool) {
	o.SendExpirationEmails = &v
}

// GetNotifyTimeBeforeAccessEnds returns the NotifyTimeBeforeAccessEnds field value if set, zero value otherwise.
func (o *InlineResponse200143) GetNotifyTimeBeforeAccessEnds() int32 {
	if o == nil || isNil(o.NotifyTimeBeforeAccessEnds) {
		var ret int32
		return ret
	}
	return *o.NotifyTimeBeforeAccessEnds
}

// GetNotifyTimeBeforeAccessEndsOk returns a tuple with the NotifyTimeBeforeAccessEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetNotifyTimeBeforeAccessEndsOk() (*int32, bool) {
	if o == nil || isNil(o.NotifyTimeBeforeAccessEnds) {
    return nil, false
	}
	return o.NotifyTimeBeforeAccessEnds, true
}

// HasNotifyTimeBeforeAccessEnds returns a boolean if a field has been set.
func (o *InlineResponse200143) HasNotifyTimeBeforeAccessEnds() bool {
	if o != nil && !isNil(o.NotifyTimeBeforeAccessEnds) {
		return true
	}

	return false
}

// SetNotifyTimeBeforeAccessEnds gets a reference to the given int32 and assigns it to the NotifyTimeBeforeAccessEnds field.
func (o *InlineResponse200143) SetNotifyTimeBeforeAccessEnds(v int32) {
	o.NotifyTimeBeforeAccessEnds = &v
}

// GetAdditionalEmailText returns the AdditionalEmailText field value if set, zero value otherwise.
func (o *InlineResponse200143) GetAdditionalEmailText() string {
	if o == nil || isNil(o.AdditionalEmailText) {
		var ret string
		return ret
	}
	return *o.AdditionalEmailText
}

// GetAdditionalEmailTextOk returns a tuple with the AdditionalEmailText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetAdditionalEmailTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalEmailText) {
    return nil, false
	}
	return o.AdditionalEmailText, true
}

// HasAdditionalEmailText returns a boolean if a field has been set.
func (o *InlineResponse200143) HasAdditionalEmailText() bool {
	if o != nil && !isNil(o.AdditionalEmailText) {
		return true
	}

	return false
}

// SetAdditionalEmailText gets a reference to the given string and assigns it to the AdditionalEmailText field.
func (o *InlineResponse200143) SetAdditionalEmailText(v string) {
	o.AdditionalEmailText = &v
}

// GetAccessStartAt returns the AccessStartAt field value if set, zero value otherwise.
func (o *InlineResponse200143) GetAccessStartAt() time.Time {
	if o == nil || isNil(o.AccessStartAt) {
		var ret time.Time
		return ret
	}
	return *o.AccessStartAt
}

// GetAccessStartAtOk returns a tuple with the AccessStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetAccessStartAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AccessStartAt) {
    return nil, false
	}
	return o.AccessStartAt, true
}

// HasAccessStartAt returns a boolean if a field has been set.
func (o *InlineResponse200143) HasAccessStartAt() bool {
	if o != nil && !isNil(o.AccessStartAt) {
		return true
	}

	return false
}

// SetAccessStartAt gets a reference to the given time.Time and assigns it to the AccessStartAt field.
func (o *InlineResponse200143) SetAccessStartAt(v time.Time) {
	o.AccessStartAt = &v
}

// GetAccessEndAt returns the AccessEndAt field value if set, zero value otherwise.
func (o *InlineResponse200143) GetAccessEndAt() time.Time {
	if o == nil || isNil(o.AccessEndAt) {
		var ret time.Time
		return ret
	}
	return *o.AccessEndAt
}

// GetAccessEndAtOk returns a tuple with the AccessEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200143) GetAccessEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AccessEndAt) {
    return nil, false
	}
	return o.AccessEndAt, true
}

// HasAccessEndAt returns a boolean if a field has been set.
func (o *InlineResponse200143) HasAccessEndAt() bool {
	if o != nil && !isNil(o.AccessEndAt) {
		return true
	}

	return false
}

// SetAccessEndAt gets a reference to the given time.Time and assigns it to the AccessEndAt field.
func (o *InlineResponse200143) SetAccessEndAt(v time.Time) {
	o.AccessEndAt = &v
}

func (o InlineResponse200143) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SsidName) {
		toSerialize["ssidName"] = o.SsidName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.TimeboundType) {
		toSerialize["timeboundType"] = o.TimeboundType
	}
	if !isNil(o.SendExpirationEmails) {
		toSerialize["sendExpirationEmails"] = o.SendExpirationEmails
	}
	if !isNil(o.NotifyTimeBeforeAccessEnds) {
		toSerialize["notifyTimeBeforeAccessEnds"] = o.NotifyTimeBeforeAccessEnds
	}
	if !isNil(o.AdditionalEmailText) {
		toSerialize["additionalEmailText"] = o.AdditionalEmailText
	}
	if !isNil(o.AccessStartAt) {
		toSerialize["accessStartAt"] = o.AccessStartAt
	}
	if !isNil(o.AccessEndAt) {
		toSerialize["accessEndAt"] = o.AccessEndAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200143 struct {
	value *InlineResponse200143
	isSet bool
}

func (v NullableInlineResponse200143) Get() *InlineResponse200143 {
	return v.value
}

func (v *NullableInlineResponse200143) Set(val *InlineResponse200143) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200143) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200143) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200143(val *InlineResponse200143) *NullableInlineResponse200143 {
	return &NullableInlineResponse200143{value: val, isSet: true}
}

func (v NullableInlineResponse200143) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200143) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


