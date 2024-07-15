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

// InlineResponse200274 struct for InlineResponse200274
type InlineResponse200274 struct {
	// The id of the VPP Account
	VppAccountId *string `json:"vppAccountId,omitempty"`
	// The VPP service token
	ContentToken *string `json:"contentToken,omitempty"`
	// The email address associated with the VPP account
	Email *string `json:"email,omitempty"`
	// The name of the VPP account
	Name *string `json:"name,omitempty"`
	// The allowed admins for the VPP account
	AllowedAdmins *string `json:"allowedAdmins,omitempty"`
	// The network IDs of the admins for the VPP account
	NetworkIdAdmins *string `json:"networkIdAdmins,omitempty"`
	// The assignable networks for the VPP account
	AssignableNetworks *string `json:"assignableNetworks,omitempty"`
	// The network IDs of the assignable networks for the VPP account
	AssignableNetworkIds []string `json:"assignableNetworkIds,omitempty"`
	// The VPP location ID
	VppLocationId *string `json:"vppLocationId,omitempty"`
	// The VPP location name
	VppLocationName *string `json:"vppLocationName,omitempty"`
	// The last time the VPP account was synced
	LastSyncedAt *string `json:"lastSyncedAt,omitempty"`
	// The last time the VPP account was force synced
	LastForceSyncedAt *string `json:"lastForceSyncedAt,omitempty"`
	ParsedToken *OrganizationsOrganizationIdSmVppAccountsParsedToken `json:"parsedToken,omitempty"`
	// The id of the VPP Account
	Id *string `json:"id,omitempty"`
	// The VPP Account's Service Token
	VppServiceToken *string `json:"vppServiceToken,omitempty"`
}

// NewInlineResponse200274 instantiates a new InlineResponse200274 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200274() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// NewInlineResponse200274WithDefaults instantiates a new InlineResponse200274 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200274WithDefaults() *InlineResponse200274 {
	this := InlineResponse200274{}
	return &this
}

// GetVppAccountId returns the VppAccountId field value if set, zero value otherwise.
func (o *InlineResponse200274) GetVppAccountId() string {
	if o == nil || isNil(o.VppAccountId) {
		var ret string
		return ret
	}
	return *o.VppAccountId
}

// GetVppAccountIdOk returns a tuple with the VppAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetVppAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.VppAccountId) {
    return nil, false
	}
	return o.VppAccountId, true
}

// HasVppAccountId returns a boolean if a field has been set.
func (o *InlineResponse200274) HasVppAccountId() bool {
	if o != nil && !isNil(o.VppAccountId) {
		return true
	}

	return false
}

// SetVppAccountId gets a reference to the given string and assigns it to the VppAccountId field.
func (o *InlineResponse200274) SetVppAccountId(v string) {
	o.VppAccountId = &v
}

// GetContentToken returns the ContentToken field value if set, zero value otherwise.
func (o *InlineResponse200274) GetContentToken() string {
	if o == nil || isNil(o.ContentToken) {
		var ret string
		return ret
	}
	return *o.ContentToken
}

// GetContentTokenOk returns a tuple with the ContentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetContentTokenOk() (*string, bool) {
	if o == nil || isNil(o.ContentToken) {
    return nil, false
	}
	return o.ContentToken, true
}

// HasContentToken returns a boolean if a field has been set.
func (o *InlineResponse200274) HasContentToken() bool {
	if o != nil && !isNil(o.ContentToken) {
		return true
	}

	return false
}

// SetContentToken gets a reference to the given string and assigns it to the ContentToken field.
func (o *InlineResponse200274) SetContentToken(v string) {
	o.ContentToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200274) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200274) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200274) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200274) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200274) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200274) SetName(v string) {
	o.Name = &v
}

// GetAllowedAdmins returns the AllowedAdmins field value if set, zero value otherwise.
func (o *InlineResponse200274) GetAllowedAdmins() string {
	if o == nil || isNil(o.AllowedAdmins) {
		var ret string
		return ret
	}
	return *o.AllowedAdmins
}

// GetAllowedAdminsOk returns a tuple with the AllowedAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetAllowedAdminsOk() (*string, bool) {
	if o == nil || isNil(o.AllowedAdmins) {
    return nil, false
	}
	return o.AllowedAdmins, true
}

// HasAllowedAdmins returns a boolean if a field has been set.
func (o *InlineResponse200274) HasAllowedAdmins() bool {
	if o != nil && !isNil(o.AllowedAdmins) {
		return true
	}

	return false
}

// SetAllowedAdmins gets a reference to the given string and assigns it to the AllowedAdmins field.
func (o *InlineResponse200274) SetAllowedAdmins(v string) {
	o.AllowedAdmins = &v
}

// GetNetworkIdAdmins returns the NetworkIdAdmins field value if set, zero value otherwise.
func (o *InlineResponse200274) GetNetworkIdAdmins() string {
	if o == nil || isNil(o.NetworkIdAdmins) {
		var ret string
		return ret
	}
	return *o.NetworkIdAdmins
}

// GetNetworkIdAdminsOk returns a tuple with the NetworkIdAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetNetworkIdAdminsOk() (*string, bool) {
	if o == nil || isNil(o.NetworkIdAdmins) {
    return nil, false
	}
	return o.NetworkIdAdmins, true
}

// HasNetworkIdAdmins returns a boolean if a field has been set.
func (o *InlineResponse200274) HasNetworkIdAdmins() bool {
	if o != nil && !isNil(o.NetworkIdAdmins) {
		return true
	}

	return false
}

// SetNetworkIdAdmins gets a reference to the given string and assigns it to the NetworkIdAdmins field.
func (o *InlineResponse200274) SetNetworkIdAdmins(v string) {
	o.NetworkIdAdmins = &v
}

// GetAssignableNetworks returns the AssignableNetworks field value if set, zero value otherwise.
func (o *InlineResponse200274) GetAssignableNetworks() string {
	if o == nil || isNil(o.AssignableNetworks) {
		var ret string
		return ret
	}
	return *o.AssignableNetworks
}

// GetAssignableNetworksOk returns a tuple with the AssignableNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetAssignableNetworksOk() (*string, bool) {
	if o == nil || isNil(o.AssignableNetworks) {
    return nil, false
	}
	return o.AssignableNetworks, true
}

// HasAssignableNetworks returns a boolean if a field has been set.
func (o *InlineResponse200274) HasAssignableNetworks() bool {
	if o != nil && !isNil(o.AssignableNetworks) {
		return true
	}

	return false
}

// SetAssignableNetworks gets a reference to the given string and assigns it to the AssignableNetworks field.
func (o *InlineResponse200274) SetAssignableNetworks(v string) {
	o.AssignableNetworks = &v
}

// GetAssignableNetworkIds returns the AssignableNetworkIds field value if set, zero value otherwise.
func (o *InlineResponse200274) GetAssignableNetworkIds() []string {
	if o == nil || isNil(o.AssignableNetworkIds) {
		var ret []string
		return ret
	}
	return o.AssignableNetworkIds
}

// GetAssignableNetworkIdsOk returns a tuple with the AssignableNetworkIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetAssignableNetworkIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AssignableNetworkIds) {
    return nil, false
	}
	return o.AssignableNetworkIds, true
}

// HasAssignableNetworkIds returns a boolean if a field has been set.
func (o *InlineResponse200274) HasAssignableNetworkIds() bool {
	if o != nil && !isNil(o.AssignableNetworkIds) {
		return true
	}

	return false
}

// SetAssignableNetworkIds gets a reference to the given []string and assigns it to the AssignableNetworkIds field.
func (o *InlineResponse200274) SetAssignableNetworkIds(v []string) {
	o.AssignableNetworkIds = v
}

// GetVppLocationId returns the VppLocationId field value if set, zero value otherwise.
func (o *InlineResponse200274) GetVppLocationId() string {
	if o == nil || isNil(o.VppLocationId) {
		var ret string
		return ret
	}
	return *o.VppLocationId
}

// GetVppLocationIdOk returns a tuple with the VppLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetVppLocationIdOk() (*string, bool) {
	if o == nil || isNil(o.VppLocationId) {
    return nil, false
	}
	return o.VppLocationId, true
}

// HasVppLocationId returns a boolean if a field has been set.
func (o *InlineResponse200274) HasVppLocationId() bool {
	if o != nil && !isNil(o.VppLocationId) {
		return true
	}

	return false
}

// SetVppLocationId gets a reference to the given string and assigns it to the VppLocationId field.
func (o *InlineResponse200274) SetVppLocationId(v string) {
	o.VppLocationId = &v
}

// GetVppLocationName returns the VppLocationName field value if set, zero value otherwise.
func (o *InlineResponse200274) GetVppLocationName() string {
	if o == nil || isNil(o.VppLocationName) {
		var ret string
		return ret
	}
	return *o.VppLocationName
}

// GetVppLocationNameOk returns a tuple with the VppLocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetVppLocationNameOk() (*string, bool) {
	if o == nil || isNil(o.VppLocationName) {
    return nil, false
	}
	return o.VppLocationName, true
}

// HasVppLocationName returns a boolean if a field has been set.
func (o *InlineResponse200274) HasVppLocationName() bool {
	if o != nil && !isNil(o.VppLocationName) {
		return true
	}

	return false
}

// SetVppLocationName gets a reference to the given string and assigns it to the VppLocationName field.
func (o *InlineResponse200274) SetVppLocationName(v string) {
	o.VppLocationName = &v
}

// GetLastSyncedAt returns the LastSyncedAt field value if set, zero value otherwise.
func (o *InlineResponse200274) GetLastSyncedAt() string {
	if o == nil || isNil(o.LastSyncedAt) {
		var ret string
		return ret
	}
	return *o.LastSyncedAt
}

// GetLastSyncedAtOk returns a tuple with the LastSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetLastSyncedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastSyncedAt) {
    return nil, false
	}
	return o.LastSyncedAt, true
}

// HasLastSyncedAt returns a boolean if a field has been set.
func (o *InlineResponse200274) HasLastSyncedAt() bool {
	if o != nil && !isNil(o.LastSyncedAt) {
		return true
	}

	return false
}

// SetLastSyncedAt gets a reference to the given string and assigns it to the LastSyncedAt field.
func (o *InlineResponse200274) SetLastSyncedAt(v string) {
	o.LastSyncedAt = &v
}

// GetLastForceSyncedAt returns the LastForceSyncedAt field value if set, zero value otherwise.
func (o *InlineResponse200274) GetLastForceSyncedAt() string {
	if o == nil || isNil(o.LastForceSyncedAt) {
		var ret string
		return ret
	}
	return *o.LastForceSyncedAt
}

// GetLastForceSyncedAtOk returns a tuple with the LastForceSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetLastForceSyncedAtOk() (*string, bool) {
	if o == nil || isNil(o.LastForceSyncedAt) {
    return nil, false
	}
	return o.LastForceSyncedAt, true
}

// HasLastForceSyncedAt returns a boolean if a field has been set.
func (o *InlineResponse200274) HasLastForceSyncedAt() bool {
	if o != nil && !isNil(o.LastForceSyncedAt) {
		return true
	}

	return false
}

// SetLastForceSyncedAt gets a reference to the given string and assigns it to the LastForceSyncedAt field.
func (o *InlineResponse200274) SetLastForceSyncedAt(v string) {
	o.LastForceSyncedAt = &v
}

// GetParsedToken returns the ParsedToken field value if set, zero value otherwise.
func (o *InlineResponse200274) GetParsedToken() OrganizationsOrganizationIdSmVppAccountsParsedToken {
	if o == nil || isNil(o.ParsedToken) {
		var ret OrganizationsOrganizationIdSmVppAccountsParsedToken
		return ret
	}
	return *o.ParsedToken
}

// GetParsedTokenOk returns a tuple with the ParsedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetParsedTokenOk() (*OrganizationsOrganizationIdSmVppAccountsParsedToken, bool) {
	if o == nil || isNil(o.ParsedToken) {
    return nil, false
	}
	return o.ParsedToken, true
}

// HasParsedToken returns a boolean if a field has been set.
func (o *InlineResponse200274) HasParsedToken() bool {
	if o != nil && !isNil(o.ParsedToken) {
		return true
	}

	return false
}

// SetParsedToken gets a reference to the given OrganizationsOrganizationIdSmVppAccountsParsedToken and assigns it to the ParsedToken field.
func (o *InlineResponse200274) SetParsedToken(v OrganizationsOrganizationIdSmVppAccountsParsedToken) {
	o.ParsedToken = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200274) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200274) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200274) SetId(v string) {
	o.Id = &v
}

// GetVppServiceToken returns the VppServiceToken field value if set, zero value otherwise.
func (o *InlineResponse200274) GetVppServiceToken() string {
	if o == nil || isNil(o.VppServiceToken) {
		var ret string
		return ret
	}
	return *o.VppServiceToken
}

// GetVppServiceTokenOk returns a tuple with the VppServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200274) GetVppServiceTokenOk() (*string, bool) {
	if o == nil || isNil(o.VppServiceToken) {
    return nil, false
	}
	return o.VppServiceToken, true
}

// HasVppServiceToken returns a boolean if a field has been set.
func (o *InlineResponse200274) HasVppServiceToken() bool {
	if o != nil && !isNil(o.VppServiceToken) {
		return true
	}

	return false
}

// SetVppServiceToken gets a reference to the given string and assigns it to the VppServiceToken field.
func (o *InlineResponse200274) SetVppServiceToken(v string) {
	o.VppServiceToken = &v
}

func (o InlineResponse200274) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VppAccountId) {
		toSerialize["vppAccountId"] = o.VppAccountId
	}
	if !isNil(o.ContentToken) {
		toSerialize["contentToken"] = o.ContentToken
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.AllowedAdmins) {
		toSerialize["allowedAdmins"] = o.AllowedAdmins
	}
	if !isNil(o.NetworkIdAdmins) {
		toSerialize["networkIdAdmins"] = o.NetworkIdAdmins
	}
	if !isNil(o.AssignableNetworks) {
		toSerialize["assignableNetworks"] = o.AssignableNetworks
	}
	if !isNil(o.AssignableNetworkIds) {
		toSerialize["assignableNetworkIds"] = o.AssignableNetworkIds
	}
	if !isNil(o.VppLocationId) {
		toSerialize["vppLocationId"] = o.VppLocationId
	}
	if !isNil(o.VppLocationName) {
		toSerialize["vppLocationName"] = o.VppLocationName
	}
	if !isNil(o.LastSyncedAt) {
		toSerialize["lastSyncedAt"] = o.LastSyncedAt
	}
	if !isNil(o.LastForceSyncedAt) {
		toSerialize["lastForceSyncedAt"] = o.LastForceSyncedAt
	}
	if !isNil(o.ParsedToken) {
		toSerialize["parsedToken"] = o.ParsedToken
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.VppServiceToken) {
		toSerialize["vppServiceToken"] = o.VppServiceToken
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200274 struct {
	value *InlineResponse200274
	isSet bool
}

func (v NullableInlineResponse200274) Get() *InlineResponse200274 {
	return v.value
}

func (v *NullableInlineResponse200274) Set(val *InlineResponse200274) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200274) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200274) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200274(val *InlineResponse200274) *NullableInlineResponse200274 {
	return &NullableInlineResponse200274{value: val, isSet: true}
}

func (v NullableInlineResponse200274) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200274) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


