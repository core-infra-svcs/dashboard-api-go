/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200198 struct for InlineResponse200198
type InlineResponse200198 struct {
	// The name of the Identity PSK
	Name *string `json:"name,omitempty"`
	// The unique identifier of the Identity PSK
	Id *string `json:"id,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// The passphrase for client authentication
	Passphrase *string `json:"passphrase,omitempty"`
	// The WiFi Personal Network unique identifier
	WifiPersonalNetworkId *string `json:"wifiPersonalNetworkId,omitempty"`
	// The email associated with the System's Manager User
	Email *string `json:"email,omitempty"`
	// Timestamp for when the Identity PSK expires, or 'null' to never expire
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewInlineResponse200198 instantiates a new InlineResponse200198 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200198() *InlineResponse200198 {
	this := InlineResponse200198{}
	return &this
}

// NewInlineResponse200198WithDefaults instantiates a new InlineResponse200198 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200198WithDefaults() *InlineResponse200198 {
	this := InlineResponse200198{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200198) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200198) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200198) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200198) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200198) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200198) SetId(v string) {
	o.Id = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineResponse200198) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineResponse200198) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineResponse200198) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineResponse200198) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineResponse200198) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineResponse200198) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetWifiPersonalNetworkId returns the WifiPersonalNetworkId field value if set, zero value otherwise.
func (o *InlineResponse200198) GetWifiPersonalNetworkId() string {
	if o == nil || isNil(o.WifiPersonalNetworkId) {
		var ret string
		return ret
	}
	return *o.WifiPersonalNetworkId
}

// GetWifiPersonalNetworkIdOk returns a tuple with the WifiPersonalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetWifiPersonalNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.WifiPersonalNetworkId) {
    return nil, false
	}
	return o.WifiPersonalNetworkId, true
}

// HasWifiPersonalNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200198) HasWifiPersonalNetworkId() bool {
	if o != nil && !isNil(o.WifiPersonalNetworkId) {
		return true
	}

	return false
}

// SetWifiPersonalNetworkId gets a reference to the given string and assigns it to the WifiPersonalNetworkId field.
func (o *InlineResponse200198) SetWifiPersonalNetworkId(v string) {
	o.WifiPersonalNetworkId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200198) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200198) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200198) SetEmail(v string) {
	o.Email = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *InlineResponse200198) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200198) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InlineResponse200198) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *InlineResponse200198) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o InlineResponse200198) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !isNil(o.WifiPersonalNetworkId) {
		toSerialize["wifiPersonalNetworkId"] = o.WifiPersonalNetworkId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200198 struct {
	value *InlineResponse200198
	isSet bool
}

func (v NullableInlineResponse200198) Get() *InlineResponse200198 {
	return v.value
}

func (v *NullableInlineResponse200198) Set(val *InlineResponse200198) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200198) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200198) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200198(val *InlineResponse200198) *NullableInlineResponse200198 {
	return &NullableInlineResponse200198{value: val, isSet: true}
}

func (v NullableInlineResponse200198) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200198) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


