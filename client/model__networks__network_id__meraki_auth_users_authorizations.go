/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NetworksNetworkIdMerakiAuthUsersAuthorizations struct for NetworksNetworkIdMerakiAuthUsersAuthorizations
type NetworksNetworkIdMerakiAuthUsersAuthorizations struct {
	// SSID number
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Authorized zone of the user
	AuthorizedZone *string `json:"authorizedZone,omitempty"`
	// Authorization expiration time
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// User is authorized by the account name
	AuthorizedByName *string `json:"authorizedByName,omitempty"`
	// User is authorized by the account email address
	AuthorizedByEmail *string `json:"authorizedByEmail,omitempty"`
}

// NewNetworksNetworkIdMerakiAuthUsersAuthorizations instantiates a new NetworksNetworkIdMerakiAuthUsersAuthorizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMerakiAuthUsersAuthorizations() *NetworksNetworkIdMerakiAuthUsersAuthorizations {
	this := NetworksNetworkIdMerakiAuthUsersAuthorizations{}
	return &this
}

// NewNetworksNetworkIdMerakiAuthUsersAuthorizationsWithDefaults instantiates a new NetworksNetworkIdMerakiAuthUsersAuthorizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMerakiAuthUsersAuthorizationsWithDefaults() *NetworksNetworkIdMerakiAuthUsersAuthorizations {
	this := NetworksNetworkIdMerakiAuthUsersAuthorizations{}
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetAuthorizedZone returns the AuthorizedZone field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedZone() string {
	if o == nil || isNil(o.AuthorizedZone) {
		var ret string
		return ret
	}
	return *o.AuthorizedZone
}

// GetAuthorizedZoneOk returns a tuple with the AuthorizedZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedZoneOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedZone) {
    return nil, false
	}
	return o.AuthorizedZone, true
}

// HasAuthorizedZone returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) HasAuthorizedZone() bool {
	if o != nil && !isNil(o.AuthorizedZone) {
		return true
	}

	return false
}

// SetAuthorizedZone gets a reference to the given string and assigns it to the AuthorizedZone field.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) SetAuthorizedZone(v string) {
	o.AuthorizedZone = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetAuthorizedByName returns the AuthorizedByName field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedByName() string {
	if o == nil || isNil(o.AuthorizedByName) {
		var ret string
		return ret
	}
	return *o.AuthorizedByName
}

// GetAuthorizedByNameOk returns a tuple with the AuthorizedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedByNameOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedByName) {
    return nil, false
	}
	return o.AuthorizedByName, true
}

// HasAuthorizedByName returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) HasAuthorizedByName() bool {
	if o != nil && !isNil(o.AuthorizedByName) {
		return true
	}

	return false
}

// SetAuthorizedByName gets a reference to the given string and assigns it to the AuthorizedByName field.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) SetAuthorizedByName(v string) {
	o.AuthorizedByName = &v
}

// GetAuthorizedByEmail returns the AuthorizedByEmail field value if set, zero value otherwise.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedByEmail() string {
	if o == nil || isNil(o.AuthorizedByEmail) {
		var ret string
		return ret
	}
	return *o.AuthorizedByEmail
}

// GetAuthorizedByEmailOk returns a tuple with the AuthorizedByEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) GetAuthorizedByEmailOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizedByEmail) {
    return nil, false
	}
	return o.AuthorizedByEmail, true
}

// HasAuthorizedByEmail returns a boolean if a field has been set.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) HasAuthorizedByEmail() bool {
	if o != nil && !isNil(o.AuthorizedByEmail) {
		return true
	}

	return false
}

// SetAuthorizedByEmail gets a reference to the given string and assigns it to the AuthorizedByEmail field.
func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations) SetAuthorizedByEmail(v string) {
	o.AuthorizedByEmail = &v
}

func (o NetworksNetworkIdMerakiAuthUsersAuthorizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.AuthorizedZone) {
		toSerialize["authorizedZone"] = o.AuthorizedZone
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !isNil(o.AuthorizedByName) {
		toSerialize["authorizedByName"] = o.AuthorizedByName
	}
	if !isNil(o.AuthorizedByEmail) {
		toSerialize["authorizedByEmail"] = o.AuthorizedByEmail
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMerakiAuthUsersAuthorizations struct {
	value *NetworksNetworkIdMerakiAuthUsersAuthorizations
	isSet bool
}

func (v NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) Get() *NetworksNetworkIdMerakiAuthUsersAuthorizations {
	return v.value
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) Set(val *NetworksNetworkIdMerakiAuthUsersAuthorizations) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMerakiAuthUsersAuthorizations(val *NetworksNetworkIdMerakiAuthUsersAuthorizations) *NullableNetworksNetworkIdMerakiAuthUsersAuthorizations {
	return &NullableNetworksNetworkIdMerakiAuthUsersAuthorizations{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMerakiAuthUsersAuthorizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


