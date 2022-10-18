/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	// The Meraki managed device Id.
	DeviceId *string `json:"deviceId,omitempty"`
	// The numerical Meraki Id of the profile.
	Id *string `json:"id,omitempty"`
	// A boolean indicating if the profile is encrypted.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// Whether or not the profile is managed by Meraki.
	IsManaged *bool `json:"isManaged,omitempty"`
	// A string containing a JSON object with the profile data.
	ProfileData *string `json:"profileData,omitempty"`
	// The identifier of the profile.
	ProfileIdentifier *string `json:"profileIdentifier,omitempty"`
	// The name of the profile.
	Name *string `json:"name,omitempty"`
	// The verison of the profile.
	Version *string `json:"version,omitempty"`
}

// NewInlineResponse20035 instantiates a new InlineResponse20035 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035WithDefaults() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20035) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20035) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20035) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20035) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20035) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20035) SetId(v string) {
	o.Id = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *InlineResponse20035) GetIsEncrypted() bool {
	if o == nil || o.IsEncrypted == nil {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || o.IsEncrypted == nil {
		return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *InlineResponse20035) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted != nil {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *InlineResponse20035) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *InlineResponse20035) GetIsManaged() bool {
	if o == nil || o.IsManaged == nil {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetIsManagedOk() (*bool, bool) {
	if o == nil || o.IsManaged == nil {
		return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *InlineResponse20035) HasIsManaged() bool {
	if o != nil && o.IsManaged != nil {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *InlineResponse20035) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetProfileData returns the ProfileData field value if set, zero value otherwise.
func (o *InlineResponse20035) GetProfileData() string {
	if o == nil || o.ProfileData == nil {
		var ret string
		return ret
	}
	return *o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetProfileDataOk() (*string, bool) {
	if o == nil || o.ProfileData == nil {
		return nil, false
	}
	return o.ProfileData, true
}

// HasProfileData returns a boolean if a field has been set.
func (o *InlineResponse20035) HasProfileData() bool {
	if o != nil && o.ProfileData != nil {
		return true
	}

	return false
}

// SetProfileData gets a reference to the given string and assigns it to the ProfileData field.
func (o *InlineResponse20035) SetProfileData(v string) {
	o.ProfileData = &v
}

// GetProfileIdentifier returns the ProfileIdentifier field value if set, zero value otherwise.
func (o *InlineResponse20035) GetProfileIdentifier() string {
	if o == nil || o.ProfileIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ProfileIdentifier
}

// GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetProfileIdentifierOk() (*string, bool) {
	if o == nil || o.ProfileIdentifier == nil {
		return nil, false
	}
	return o.ProfileIdentifier, true
}

// HasProfileIdentifier returns a boolean if a field has been set.
func (o *InlineResponse20035) HasProfileIdentifier() bool {
	if o != nil && o.ProfileIdentifier != nil {
		return true
	}

	return false
}

// SetProfileIdentifier gets a reference to the given string and assigns it to the ProfileIdentifier field.
func (o *InlineResponse20035) SetProfileIdentifier(v string) {
	o.ProfileIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20035) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20035) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20035) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20035) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20035) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20035) SetVersion(v string) {
	o.Version = &v
}

func (o InlineResponse20035) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsEncrypted != nil {
		toSerialize["isEncrypted"] = o.IsEncrypted
	}
	if o.IsManaged != nil {
		toSerialize["isManaged"] = o.IsManaged
	}
	if o.ProfileData != nil {
		toSerialize["profileData"] = o.ProfileData
	}
	if o.ProfileIdentifier != nil {
		toSerialize["profileIdentifier"] = o.ProfileIdentifier
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035 struct {
	value *InlineResponse20035
	isSet bool
}

func (v NullableInlineResponse20035) Get() *InlineResponse20035 {
	return v.value
}

func (v *NullableInlineResponse20035) Set(val *InlineResponse20035) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035(val *InlineResponse20035) *NullableInlineResponse20035 {
	return &NullableInlineResponse20035{value: val, isSet: true}
}

func (v NullableInlineResponse20035) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


