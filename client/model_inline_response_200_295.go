/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200295 struct for InlineResponse200295
type InlineResponse200295 struct {
	// Type of alert that the webhook is delivering
	AlertType *string `json:"alertType,omitempty"`
	// When the webhook log was created, in ISO8601 format
	LoggedAt *time.Time `json:"loggedAt,omitempty"`
	// Network ID for the webhook log
	NetworkId *string `json:"networkId,omitempty"`
	// ID for the webhook log's organization
	OrganizationId *string `json:"organizationId,omitempty"`
	// Response code from the webhook
	ResponseCode *int32 `json:"responseCode,omitempty"`
	// Duration of the response, in milliseconds
	ResponseDuration *int32 `json:"responseDuration,omitempty"`
	// When the webhook was sent, in ISO8601 format
	SentAt *time.Time `json:"sentAt,omitempty"`
	// URL where the webhook was sent
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse200295 instantiates a new InlineResponse200295 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200295() *InlineResponse200295 {
	this := InlineResponse200295{}
	return &this
}

// NewInlineResponse200295WithDefaults instantiates a new InlineResponse200295 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200295WithDefaults() *InlineResponse200295 {
	this := InlineResponse200295{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *InlineResponse200295) GetAlertType() string {
	if o == nil || isNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetAlertTypeOk() (*string, bool) {
	if o == nil || isNil(o.AlertType) {
    return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *InlineResponse200295) HasAlertType() bool {
	if o != nil && !isNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *InlineResponse200295) SetAlertType(v string) {
	o.AlertType = &v
}

// GetLoggedAt returns the LoggedAt field value if set, zero value otherwise.
func (o *InlineResponse200295) GetLoggedAt() time.Time {
	if o == nil || isNil(o.LoggedAt) {
		var ret time.Time
		return ret
	}
	return *o.LoggedAt
}

// GetLoggedAtOk returns a tuple with the LoggedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetLoggedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LoggedAt) {
    return nil, false
	}
	return o.LoggedAt, true
}

// HasLoggedAt returns a boolean if a field has been set.
func (o *InlineResponse200295) HasLoggedAt() bool {
	if o != nil && !isNil(o.LoggedAt) {
		return true
	}

	return false
}

// SetLoggedAt gets a reference to the given time.Time and assigns it to the LoggedAt field.
func (o *InlineResponse200295) SetLoggedAt(v time.Time) {
	o.LoggedAt = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200295) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200295) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200295) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200295) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200295) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse200295) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *InlineResponse200295) GetResponseCode() int32 {
	if o == nil || isNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetResponseCodeOk() (*int32, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *InlineResponse200295) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *InlineResponse200295) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetResponseDuration returns the ResponseDuration field value if set, zero value otherwise.
func (o *InlineResponse200295) GetResponseDuration() int32 {
	if o == nil || isNil(o.ResponseDuration) {
		var ret int32
		return ret
	}
	return *o.ResponseDuration
}

// GetResponseDurationOk returns a tuple with the ResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetResponseDurationOk() (*int32, bool) {
	if o == nil || isNil(o.ResponseDuration) {
    return nil, false
	}
	return o.ResponseDuration, true
}

// HasResponseDuration returns a boolean if a field has been set.
func (o *InlineResponse200295) HasResponseDuration() bool {
	if o != nil && !isNil(o.ResponseDuration) {
		return true
	}

	return false
}

// SetResponseDuration gets a reference to the given int32 and assigns it to the ResponseDuration field.
func (o *InlineResponse200295) SetResponseDuration(v int32) {
	o.ResponseDuration = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *InlineResponse200295) GetSentAt() time.Time {
	if o == nil || isNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetSentAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *InlineResponse200295) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *InlineResponse200295) SetSentAt(v time.Time) {
	o.SentAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200295) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200295) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200295) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200295) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse200295) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !isNil(o.LoggedAt) {
		toSerialize["loggedAt"] = o.LoggedAt
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.ResponseDuration) {
		toSerialize["responseDuration"] = o.ResponseDuration
	}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200295 struct {
	value *InlineResponse200295
	isSet bool
}

func (v NullableInlineResponse200295) Get() *InlineResponse200295 {
	return v.value
}

func (v *NullableInlineResponse200295) Set(val *InlineResponse200295) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200295) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200295) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200295(val *InlineResponse200295) *NullableInlineResponse200295 {
	return &NullableInlineResponse200295{value: val, isSet: true}
}

func (v NullableInlineResponse200295) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200295) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

