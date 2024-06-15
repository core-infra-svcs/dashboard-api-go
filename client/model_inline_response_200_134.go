/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200134 struct for InlineResponse200134
type InlineResponse200134 struct {
	// device ID
	Id *string `json:"id,omitempty"`
	// device name
	Name *string `json:"name,omitempty"`
	// system type
	SystemType *string `json:"systemType,omitempty"`
	// mac address
	Mac *string `json:"mac,omitempty"`
	// username
	Username *string `json:"username,omitempty"`
	// user email
	Email *string `json:"email,omitempty"`
	// device tags
	Tags []string `json:"tags,omitempty"`
	// Array of trusted access configs
	TrustedAccessConnections []NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections `json:"trustedAccessConnections,omitempty"`
}

// NewInlineResponse200134 instantiates a new InlineResponse200134 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200134() *InlineResponse200134 {
	this := InlineResponse200134{}
	return &this
}

// NewInlineResponse200134WithDefaults instantiates a new InlineResponse200134 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200134WithDefaults() *InlineResponse200134 {
	this := InlineResponse200134{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200134) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200134) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200134) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200134) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200134) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200134) SetName(v string) {
	o.Name = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *InlineResponse200134) GetSystemType() string {
	if o == nil || isNil(o.SystemType) {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetSystemTypeOk() (*string, bool) {
	if o == nil || isNil(o.SystemType) {
    return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *InlineResponse200134) HasSystemType() bool {
	if o != nil && !isNil(o.SystemType) {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *InlineResponse200134) SetSystemType(v string) {
	o.SystemType = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200134) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200134) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200134) SetMac(v string) {
	o.Mac = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineResponse200134) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineResponse200134) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineResponse200134) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineResponse200134) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineResponse200134) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineResponse200134) SetEmail(v string) {
	o.Email = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200134) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200134) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200134) SetTags(v []string) {
	o.Tags = v
}

// GetTrustedAccessConnections returns the TrustedAccessConnections field value if set, zero value otherwise.
func (o *InlineResponse200134) GetTrustedAccessConnections() []NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections {
	if o == nil || isNil(o.TrustedAccessConnections) {
		var ret []NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections
		return ret
	}
	return o.TrustedAccessConnections
}

// GetTrustedAccessConnectionsOk returns a tuple with the TrustedAccessConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200134) GetTrustedAccessConnectionsOk() ([]NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections, bool) {
	if o == nil || isNil(o.TrustedAccessConnections) {
    return nil, false
	}
	return o.TrustedAccessConnections, true
}

// HasTrustedAccessConnections returns a boolean if a field has been set.
func (o *InlineResponse200134) HasTrustedAccessConnections() bool {
	if o != nil && !isNil(o.TrustedAccessConnections) {
		return true
	}

	return false
}

// SetTrustedAccessConnections gets a reference to the given []NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections and assigns it to the TrustedAccessConnections field.
func (o *InlineResponse200134) SetTrustedAccessConnections(v []NetworksNetworkIdSmUserAccessDevicesTrustedAccessConnections) {
	o.TrustedAccessConnections = v
}

func (o InlineResponse200134) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SystemType) {
		toSerialize["systemType"] = o.SystemType
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.TrustedAccessConnections) {
		toSerialize["trustedAccessConnections"] = o.TrustedAccessConnections
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200134 struct {
	value *InlineResponse200134
	isSet bool
}

func (v NullableInlineResponse200134) Get() *InlineResponse200134 {
	return v.value
}

func (v *NullableInlineResponse200134) Set(val *InlineResponse200134) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200134) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200134) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200134(val *InlineResponse200134) *NullableInlineResponse200134 {
	return &NullableInlineResponse200134{value: val, isSet: true}
}

func (v NullableInlineResponse200134) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200134) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


