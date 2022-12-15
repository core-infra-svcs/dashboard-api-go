/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20043 struct for InlineResponse20043
type InlineResponse20043 struct {
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

// NewInlineResponse20043 instantiates a new InlineResponse20043 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20043() *InlineResponse20043 {
	this := InlineResponse20043{}
	return &this
}

// NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20043WithDefaults() *InlineResponse20043 {
	this := InlineResponse20043{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20043) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20043) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20043) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20043) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20043) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20043) SetId(v string) {
	o.Id = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *InlineResponse20043) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || isNil(o.IsEncrypted) {
    return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *InlineResponse20043) HasIsEncrypted() bool {
	if o != nil && !isNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *InlineResponse20043) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *InlineResponse20043) GetIsManaged() bool {
	if o == nil || isNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetIsManagedOk() (*bool, bool) {
	if o == nil || isNil(o.IsManaged) {
    return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *InlineResponse20043) HasIsManaged() bool {
	if o != nil && !isNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *InlineResponse20043) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetProfileData returns the ProfileData field value if set, zero value otherwise.
func (o *InlineResponse20043) GetProfileData() string {
	if o == nil || isNil(o.ProfileData) {
		var ret string
		return ret
	}
	return *o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetProfileDataOk() (*string, bool) {
	if o == nil || isNil(o.ProfileData) {
    return nil, false
	}
	return o.ProfileData, true
}

// HasProfileData returns a boolean if a field has been set.
func (o *InlineResponse20043) HasProfileData() bool {
	if o != nil && !isNil(o.ProfileData) {
		return true
	}

	return false
}

// SetProfileData gets a reference to the given string and assigns it to the ProfileData field.
func (o *InlineResponse20043) SetProfileData(v string) {
	o.ProfileData = &v
}

// GetProfileIdentifier returns the ProfileIdentifier field value if set, zero value otherwise.
func (o *InlineResponse20043) GetProfileIdentifier() string {
	if o == nil || isNil(o.ProfileIdentifier) {
		var ret string
		return ret
	}
	return *o.ProfileIdentifier
}

// GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetProfileIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ProfileIdentifier) {
    return nil, false
	}
	return o.ProfileIdentifier, true
}

// HasProfileIdentifier returns a boolean if a field has been set.
func (o *InlineResponse20043) HasProfileIdentifier() bool {
	if o != nil && !isNil(o.ProfileIdentifier) {
		return true
	}

	return false
}

// SetProfileIdentifier gets a reference to the given string and assigns it to the ProfileIdentifier field.
func (o *InlineResponse20043) SetProfileIdentifier(v string) {
	o.ProfileIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20043) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20043) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20043) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse20043) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse20043) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse20043) SetVersion(v string) {
	o.Version = &v
}

func (o InlineResponse20043) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20043 struct {
	value *InlineResponse20043
	isSet bool
}

func (v NullableInlineResponse20043) Get() *InlineResponse20043 {
	return v.value
}

func (v *NullableInlineResponse20043) Set(val *InlineResponse20043) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20043) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20043) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20043(val *InlineResponse20043) *NullableInlineResponse20043 {
	return &NullableInlineResponse20043{value: val, isSet: true}
}

func (v NullableInlineResponse20043) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20043) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


