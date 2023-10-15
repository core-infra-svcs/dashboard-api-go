/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject99 struct for InlineObject99
type InlineObject99 struct {
	// One of \"delete\" or \"restrict processing\"
	Type *string `json:"type,omitempty"`
	// The datasets related to the provided key that should be deleted. Only applies to \"delete\" requests. The value \"all\" will be expanded to all datasets applicable to this type. The datasets by applicable to each type are: mac (usage, events, traffic), email (users, loginAttempts), username (users, loginAttempts), bluetoothMac (client, connectivity), smDeviceId (device), smUserId (user)
	Datasets []string `json:"datasets,omitempty"`
	// The username of a network log in. Only applies to \"delete\" requests.
	Username *string `json:"username,omitempty"`
	// The email of a network user account. Only applies to \"delete\" requests.
	Email *string `json:"email,omitempty"`
	// The MAC of a network client device. Applies to both \"restrict processing\" and \"delete\" requests.
	Mac *string `json:"mac,omitempty"`
	// The sm_device_id of a Systems Manager device. The only way to \"restrict processing\" or \"delete\" a Systems Manager device. Must include \"device\" in the dataset for a \"delete\" request to destroy the device.
	SmDeviceId *string `json:"smDeviceId,omitempty"`
	// The sm_user_id of a Systems Manager user. The only way to \"restrict processing\" or \"delete\" a Systems Manager user. Must include \"user\" in the dataset for a \"delete\" request to destroy the user.
	SmUserId *string `json:"smUserId,omitempty"`
}

// NewInlineObject99 instantiates a new InlineObject99 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject99() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// NewInlineObject99WithDefaults instantiates a new InlineObject99 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject99WithDefaults() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject99) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject99) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject99) SetType(v string) {
	o.Type = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *InlineObject99) GetDatasets() []string {
	if o == nil || isNil(o.Datasets) {
		var ret []string
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetDatasetsOk() ([]string, bool) {
	if o == nil || isNil(o.Datasets) {
    return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *InlineObject99) HasDatasets() bool {
	if o != nil && !isNil(o.Datasets) {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []string and assigns it to the Datasets field.
func (o *InlineObject99) SetDatasets(v []string) {
	o.Datasets = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineObject99) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineObject99) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineObject99) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineObject99) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineObject99) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineObject99) SetEmail(v string) {
	o.Email = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineObject99) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineObject99) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineObject99) SetMac(v string) {
	o.Mac = &v
}

// GetSmDeviceId returns the SmDeviceId field value if set, zero value otherwise.
func (o *InlineObject99) GetSmDeviceId() string {
	if o == nil || isNil(o.SmDeviceId) {
		var ret string
		return ret
	}
	return *o.SmDeviceId
}

// GetSmDeviceIdOk returns a tuple with the SmDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetSmDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.SmDeviceId) {
    return nil, false
	}
	return o.SmDeviceId, true
}

// HasSmDeviceId returns a boolean if a field has been set.
func (o *InlineObject99) HasSmDeviceId() bool {
	if o != nil && !isNil(o.SmDeviceId) {
		return true
	}

	return false
}

// SetSmDeviceId gets a reference to the given string and assigns it to the SmDeviceId field.
func (o *InlineObject99) SetSmDeviceId(v string) {
	o.SmDeviceId = &v
}

// GetSmUserId returns the SmUserId field value if set, zero value otherwise.
func (o *InlineObject99) GetSmUserId() string {
	if o == nil || isNil(o.SmUserId) {
		var ret string
		return ret
	}
	return *o.SmUserId
}

// GetSmUserIdOk returns a tuple with the SmUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetSmUserIdOk() (*string, bool) {
	if o == nil || isNil(o.SmUserId) {
    return nil, false
	}
	return o.SmUserId, true
}

// HasSmUserId returns a boolean if a field has been set.
func (o *InlineObject99) HasSmUserId() bool {
	if o != nil && !isNil(o.SmUserId) {
		return true
	}

	return false
}

// SetSmUserId gets a reference to the given string and assigns it to the SmUserId field.
func (o *InlineObject99) SetSmUserId(v string) {
	o.SmUserId = &v
}

func (o InlineObject99) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Datasets) {
		toSerialize["datasets"] = o.Datasets
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.SmDeviceId) {
		toSerialize["smDeviceId"] = o.SmDeviceId
	}
	if !isNil(o.SmUserId) {
		toSerialize["smUserId"] = o.SmUserId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject99 struct {
	value *InlineObject99
	isSet bool
}

func (v NullableInlineObject99) Get() *InlineObject99 {
	return v.value
}

func (v *NullableInlineObject99) Set(val *InlineObject99) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject99) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject99) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject99(val *InlineObject99) *NullableInlineObject99 {
	return &NullableInlineObject99{value: val, isSet: true}
}

func (v NullableInlineObject99) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject99) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


