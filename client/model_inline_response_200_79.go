/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20079 struct for InlineResponse20079
type InlineResponse20079 struct {
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

// NewInlineResponse20079 instantiates a new InlineResponse20079 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20079() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20079WithDefaults() *InlineResponse20079 {
	this := InlineResponse20079{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20079) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20079) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20079) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20079) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20079) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20079) SetId(v string) {
	o.Id = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *InlineResponse20079) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || isNil(o.IsEncrypted) {
    return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *InlineResponse20079) HasIsEncrypted() bool {
	if o != nil && !isNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *InlineResponse20079) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *InlineResponse20079) GetIsManaged() bool {
	if o == nil || isNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetIsManagedOk() (*bool, bool) {
	if o == nil || isNil(o.IsManaged) {
    return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *InlineResponse20079) HasIsManaged() bool {
	if o != nil && !isNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *InlineResponse20079) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetProfileData returns the ProfileData field value if set, zero value otherwise.
func (o *InlineResponse20079) GetProfileData() string {
	if o == nil || isNil(o.ProfileData) {
		var ret string
		return ret
	}
	return *o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetProfileDataOk() (*string, bool) {
	if o == nil || isNil(o.ProfileData) {
    return nil, false
	}
	return o.ProfileData, true
}

// HasProfileData returns a boolean if a field has been set.
func (o *InlineResponse20079) HasProfileData() bool {
	if o != nil && !isNil(o.ProfileData) {
		return true
	}

	return false
}

// SetProfileData gets a reference to the given string and assigns it to the ProfileData field.
func (o *InlineResponse20079) SetProfileData(v string) {
	o.ProfileData = &v
}

// GetProfileIdentifier returns the ProfileIdentifier field value if set, zero value otherwise.
func (o *InlineResponse20079) GetProfileIdentifier() string {
	if o == nil || isNil(o.ProfileIdentifier) {
		var ret string
		return ret
	}
	return *o.ProfileIdentifier
}

// GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetProfileIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ProfileIdentifier) {
    return nil, false
	}
	return o.ProfileIdentifier, true
}

// HasProfileIdentifier returns a boolean if a field has been set.
func (o *InlineResponse20079) HasProfileIdentifier() bool {
	if o != nil && !isNil(o.ProfileIdentifier) {
		return true
	}

	return false
}

// SetProfileIdentifier gets a reference to the given string and assigns it to the ProfileIdentifier field.
func (o *InlineResponse20079) SetProfileIdentifier(v string) {
	o.ProfileIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20079) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20079) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20079) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20079) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20079) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20079) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20079) SetVersion(v string) {
	o.Version = &v
}

func (o InlineResponse20079) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsEncrypted) {
		toSerialize["isEncrypted"] = o.IsEncrypted
	}
	if !isNil(o.IsManaged) {
		toSerialize["isManaged"] = o.IsManaged
	}
	if !isNil(o.ProfileData) {
		toSerialize["profileData"] = o.ProfileData
	}
	if !isNil(o.ProfileIdentifier) {
		toSerialize["profileIdentifier"] = o.ProfileIdentifier
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20079 struct {
	value *InlineResponse20079
	isSet bool
}

func (v NullableInlineResponse20079) Get() *InlineResponse20079 {
	return v.value
}

func (v *NullableInlineResponse20079) Set(val *InlineResponse20079) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20079) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20079) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20079(val *InlineResponse20079) *NullableInlineResponse20079 {
	return &NullableInlineResponse20079{value: val, isSet: true}
}

func (v NullableInlineResponse20079) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20079) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


