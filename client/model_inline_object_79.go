/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject79 struct for InlineObject79
type InlineObject79 struct {
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

// NewInlineObject79 instantiates a new InlineObject79 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject79() *InlineObject79 {
	this := InlineObject79{}
	return &this
}

// NewInlineObject79WithDefaults instantiates a new InlineObject79 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject79WithDefaults() *InlineObject79 {
	this := InlineObject79{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject79) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject79) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject79) SetType(v string) {
	o.Type = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *InlineObject79) GetDatasets() []string {
	if o == nil || o.Datasets == nil {
		var ret []string
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetDatasetsOk() ([]string, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *InlineObject79) HasDatasets() bool {
	if o != nil && o.Datasets != nil {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []string and assigns it to the Datasets field.
func (o *InlineObject79) SetDatasets(v []string) {
	o.Datasets = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InlineObject79) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InlineObject79) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InlineObject79) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InlineObject79) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InlineObject79) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InlineObject79) SetEmail(v string) {
	o.Email = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineObject79) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineObject79) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineObject79) SetMac(v string) {
	o.Mac = &v
}

// GetSmDeviceId returns the SmDeviceId field value if set, zero value otherwise.
func (o *InlineObject79) GetSmDeviceId() string {
	if o == nil || o.SmDeviceId == nil {
		var ret string
		return ret
	}
	return *o.SmDeviceId
}

// GetSmDeviceIdOk returns a tuple with the SmDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetSmDeviceIdOk() (*string, bool) {
	if o == nil || o.SmDeviceId == nil {
		return nil, false
	}
	return o.SmDeviceId, true
}

// HasSmDeviceId returns a boolean if a field has been set.
func (o *InlineObject79) HasSmDeviceId() bool {
	if o != nil && o.SmDeviceId != nil {
		return true
	}

	return false
}

// SetSmDeviceId gets a reference to the given string and assigns it to the SmDeviceId field.
func (o *InlineObject79) SetSmDeviceId(v string) {
	o.SmDeviceId = &v
}

// GetSmUserId returns the SmUserId field value if set, zero value otherwise.
func (o *InlineObject79) GetSmUserId() string {
	if o == nil || o.SmUserId == nil {
		var ret string
		return ret
	}
	return *o.SmUserId
}

// GetSmUserIdOk returns a tuple with the SmUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject79) GetSmUserIdOk() (*string, bool) {
	if o == nil || o.SmUserId == nil {
		return nil, false
	}
	return o.SmUserId, true
}

// HasSmUserId returns a boolean if a field has been set.
func (o *InlineObject79) HasSmUserId() bool {
	if o != nil && o.SmUserId != nil {
		return true
	}

	return false
}

// SetSmUserId gets a reference to the given string and assigns it to the SmUserId field.
func (o *InlineObject79) SetSmUserId(v string) {
	o.SmUserId = &v
}

func (o InlineObject79) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.SmDeviceId != nil {
		toSerialize["smDeviceId"] = o.SmDeviceId
	}
	if o.SmUserId != nil {
		toSerialize["smUserId"] = o.SmUserId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject79 struct {
	value *InlineObject79
	isSet bool
}

func (v NullableInlineObject79) Get() *InlineObject79 {
	return v.value
}

func (v *NullableInlineObject79) Set(val *InlineObject79) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject79) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject79) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject79(val *InlineObject79) *NullableInlineObject79 {
	return &NullableInlineObject79{value: val, isSet: true}
}

func (v NullableInlineObject79) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject79) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


