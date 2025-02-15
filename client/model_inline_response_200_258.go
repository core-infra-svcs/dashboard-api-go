/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200258 struct for InlineResponse200258
type InlineResponse200258 struct {
	// Time, in ISO8601 format, when the configuration change was made.
	Ts *time.Time `json:"ts,omitempty"`
	// The name of the admin who made the configuration change.
	AdminName *string `json:"adminName,omitempty"`
	// The email address of the admin who made the configuration change. This attribute may be null.
	AdminEmail *string `json:"adminEmail,omitempty"`
	// The ID of the admin who made the configuration change. This attribute may be null.
	AdminId *string `json:"adminId,omitempty"`
	// The name of the network that the configuration change was applied to. This attribute may be null.
	NetworkName *string `json:"networkName,omitempty"`
	// The ID of the network that the configuration change was applied to. This attribute may be null.
	NetworkId *string `json:"networkId,omitempty"`
	// The url of the network that the configuration change was applied to. This attribute may be null.
	NetworkUrl *string `json:"networkUrl,omitempty"`
	// The name of the ssid that the configuration change was applied to, if applicable. This attribute may be null.
	SsidName *string `json:"ssidName,omitempty"`
	// The ssid number that the configuration change was applied to, if applicable. This attribute may be null.
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// The name of the Meraki Dashboard page on which the configuration change was made.
	Page *string `json:"page,omitempty"`
	// Description of the configuration change.
	Label *string `json:"label,omitempty"`
	// The value of the configuration, before the change was applied.
	OldValue *string `json:"oldValue,omitempty"`
	// The value of the configuration, after the change was applied.
	NewValue *string `json:"newValue,omitempty"`
	Client *OrganizationsOrganizationIdConfigurationChangesClient `json:"client,omitempty"`
}

// NewInlineResponse200258 instantiates a new InlineResponse200258 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200258() *InlineResponse200258 {
	this := InlineResponse200258{}
	return &this
}

// NewInlineResponse200258WithDefaults instantiates a new InlineResponse200258 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200258WithDefaults() *InlineResponse200258 {
	this := InlineResponse200258{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse200258) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse200258) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse200258) SetTs(v time.Time) {
	o.Ts = &v
}

// GetAdminName returns the AdminName field value if set, zero value otherwise.
func (o *InlineResponse200258) GetAdminName() string {
	if o == nil || isNil(o.AdminName) {
		var ret string
		return ret
	}
	return *o.AdminName
}

// GetAdminNameOk returns a tuple with the AdminName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetAdminNameOk() (*string, bool) {
	if o == nil || isNil(o.AdminName) {
    return nil, false
	}
	return o.AdminName, true
}

// HasAdminName returns a boolean if a field has been set.
func (o *InlineResponse200258) HasAdminName() bool {
	if o != nil && !isNil(o.AdminName) {
		return true
	}

	return false
}

// SetAdminName gets a reference to the given string and assigns it to the AdminName field.
func (o *InlineResponse200258) SetAdminName(v string) {
	o.AdminName = &v
}

// GetAdminEmail returns the AdminEmail field value if set, zero value otherwise.
func (o *InlineResponse200258) GetAdminEmail() string {
	if o == nil || isNil(o.AdminEmail) {
		var ret string
		return ret
	}
	return *o.AdminEmail
}

// GetAdminEmailOk returns a tuple with the AdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetAdminEmailOk() (*string, bool) {
	if o == nil || isNil(o.AdminEmail) {
    return nil, false
	}
	return o.AdminEmail, true
}

// HasAdminEmail returns a boolean if a field has been set.
func (o *InlineResponse200258) HasAdminEmail() bool {
	if o != nil && !isNil(o.AdminEmail) {
		return true
	}

	return false
}

// SetAdminEmail gets a reference to the given string and assigns it to the AdminEmail field.
func (o *InlineResponse200258) SetAdminEmail(v string) {
	o.AdminEmail = &v
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *InlineResponse200258) GetAdminId() string {
	if o == nil || isNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetAdminIdOk() (*string, bool) {
	if o == nil || isNil(o.AdminId) {
    return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *InlineResponse200258) HasAdminId() bool {
	if o != nil && !isNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *InlineResponse200258) SetAdminId(v string) {
	o.AdminId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *InlineResponse200258) GetNetworkName() string {
	if o == nil || isNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetNetworkNameOk() (*string, bool) {
	if o == nil || isNil(o.NetworkName) {
    return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *InlineResponse200258) HasNetworkName() bool {
	if o != nil && !isNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *InlineResponse200258) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200258) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200258) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200258) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkUrl returns the NetworkUrl field value if set, zero value otherwise.
func (o *InlineResponse200258) GetNetworkUrl() string {
	if o == nil || isNil(o.NetworkUrl) {
		var ret string
		return ret
	}
	return *o.NetworkUrl
}

// GetNetworkUrlOk returns a tuple with the NetworkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetNetworkUrlOk() (*string, bool) {
	if o == nil || isNil(o.NetworkUrl) {
    return nil, false
	}
	return o.NetworkUrl, true
}

// HasNetworkUrl returns a boolean if a field has been set.
func (o *InlineResponse200258) HasNetworkUrl() bool {
	if o != nil && !isNil(o.NetworkUrl) {
		return true
	}

	return false
}

// SetNetworkUrl gets a reference to the given string and assigns it to the NetworkUrl field.
func (o *InlineResponse200258) SetNetworkUrl(v string) {
	o.NetworkUrl = &v
}

// GetSsidName returns the SsidName field value if set, zero value otherwise.
func (o *InlineResponse200258) GetSsidName() string {
	if o == nil || isNil(o.SsidName) {
		var ret string
		return ret
	}
	return *o.SsidName
}

// GetSsidNameOk returns a tuple with the SsidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetSsidNameOk() (*string, bool) {
	if o == nil || isNil(o.SsidName) {
    return nil, false
	}
	return o.SsidName, true
}

// HasSsidName returns a boolean if a field has been set.
func (o *InlineResponse200258) HasSsidName() bool {
	if o != nil && !isNil(o.SsidName) {
		return true
	}

	return false
}

// SetSsidName gets a reference to the given string and assigns it to the SsidName field.
func (o *InlineResponse200258) SetSsidName(v string) {
	o.SsidName = &v
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse200258) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse200258) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse200258) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *InlineResponse200258) GetPage() string {
	if o == nil || isNil(o.Page) {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetPageOk() (*string, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *InlineResponse200258) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *InlineResponse200258) SetPage(v string) {
	o.Page = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InlineResponse200258) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InlineResponse200258) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InlineResponse200258) SetLabel(v string) {
	o.Label = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *InlineResponse200258) GetOldValue() string {
	if o == nil || isNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetOldValueOk() (*string, bool) {
	if o == nil || isNil(o.OldValue) {
    return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *InlineResponse200258) HasOldValue() bool {
	if o != nil && !isNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *InlineResponse200258) SetOldValue(v string) {
	o.OldValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *InlineResponse200258) GetNewValue() string {
	if o == nil || isNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetNewValueOk() (*string, bool) {
	if o == nil || isNil(o.NewValue) {
    return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *InlineResponse200258) HasNewValue() bool {
	if o != nil && !isNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *InlineResponse200258) SetNewValue(v string) {
	o.NewValue = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *InlineResponse200258) GetClient() OrganizationsOrganizationIdConfigurationChangesClient {
	if o == nil || isNil(o.Client) {
		var ret OrganizationsOrganizationIdConfigurationChangesClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200258) GetClientOk() (*OrganizationsOrganizationIdConfigurationChangesClient, bool) {
	if o == nil || isNil(o.Client) {
    return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *InlineResponse200258) HasClient() bool {
	if o != nil && !isNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given OrganizationsOrganizationIdConfigurationChangesClient and assigns it to the Client field.
func (o *InlineResponse200258) SetClient(v OrganizationsOrganizationIdConfigurationChangesClient) {
	o.Client = &v
}

func (o InlineResponse200258) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.AdminName) {
		toSerialize["adminName"] = o.AdminName
	}
	if !isNil(o.AdminEmail) {
		toSerialize["adminEmail"] = o.AdminEmail
	}
	if !isNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	if !isNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.NetworkUrl) {
		toSerialize["networkUrl"] = o.NetworkUrl
	}
	if !isNil(o.SsidName) {
		toSerialize["ssidName"] = o.SsidName
	}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !isNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	if !isNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200258 struct {
	value *InlineResponse200258
	isSet bool
}

func (v NullableInlineResponse200258) Get() *InlineResponse200258 {
	return v.value
}

func (v *NullableInlineResponse200258) Set(val *InlineResponse200258) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200258) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200258) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200258(val *InlineResponse200258) *NullableInlineResponse200258 {
	return &NullableInlineResponse200258{value: val, isSet: true}
}

func (v NullableInlineResponse200258) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200258) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


