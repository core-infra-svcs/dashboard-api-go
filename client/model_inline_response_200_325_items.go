/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200325Items struct for InlineResponse200325Items
type InlineResponse200325Items struct {
	Network *InlineResponse200325Network `json:"network,omitempty"`
	// AP cloud ID
	Serial *string `json:"serial,omitempty"`
	Controller *InlineResponse200325Controller `json:"controller,omitempty"`
	// The time when AP joins wireless controller
	JoinedAt *string `json:"joinedAt,omitempty"`
	// AP model
	Model *string `json:"model,omitempty"`
	// The tags of the catalyst access point
	Tags []InlineResponse200325Tags `json:"tags,omitempty"`
	// AP mode (local, flex, etc.)
	Mode *string `json:"mode,omitempty"`
	// Country code (2 characters)
	CountryCode *string `json:"countryCode,omitempty"`
	// Catalyst access point details
	Details []InlineResponse200325Details `json:"details,omitempty"`
}

// NewInlineResponse200325Items instantiates a new InlineResponse200325Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200325Items() *InlineResponse200325Items {
	this := InlineResponse200325Items{}
	return &this
}

// NewInlineResponse200325ItemsWithDefaults instantiates a new InlineResponse200325Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200325ItemsWithDefaults() *InlineResponse200325Items {
	this := InlineResponse200325Items{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetNetwork() InlineResponse200325Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200325Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetNetworkOk() (*InlineResponse200325Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200325Network and assigns it to the Network field.
func (o *InlineResponse200325Items) SetNetwork(v InlineResponse200325Network) {
	o.Network = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200325Items) SetSerial(v string) {
	o.Serial = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetController() InlineResponse200325Controller {
	if o == nil || isNil(o.Controller) {
		var ret InlineResponse200325Controller
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetControllerOk() (*InlineResponse200325Controller, bool) {
	if o == nil || isNil(o.Controller) {
    return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasController() bool {
	if o != nil && !isNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given InlineResponse200325Controller and assigns it to the Controller field.
func (o *InlineResponse200325Items) SetController(v InlineResponse200325Controller) {
	o.Controller = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetJoinedAt() string {
	if o == nil || isNil(o.JoinedAt) {
		var ret string
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetJoinedAtOk() (*string, bool) {
	if o == nil || isNil(o.JoinedAt) {
    return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasJoinedAt() bool {
	if o != nil && !isNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given string and assigns it to the JoinedAt field.
func (o *InlineResponse200325Items) SetJoinedAt(v string) {
	o.JoinedAt = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200325Items) SetModel(v string) {
	o.Model = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetTags() []InlineResponse200325Tags {
	if o == nil || isNil(o.Tags) {
		var ret []InlineResponse200325Tags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetTagsOk() ([]InlineResponse200325Tags, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []InlineResponse200325Tags and assigns it to the Tags field.
func (o *InlineResponse200325Items) SetTags(v []InlineResponse200325Tags) {
	o.Tags = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *InlineResponse200325Items) SetMode(v string) {
	o.Mode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetCountryCode() string {
	if o == nil || isNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetCountryCodeOk() (*string, bool) {
	if o == nil || isNil(o.CountryCode) {
    return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasCountryCode() bool {
	if o != nil && !isNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *InlineResponse200325Items) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *InlineResponse200325Items) GetDetails() []InlineResponse200325Details {
	if o == nil || isNil(o.Details) {
		var ret []InlineResponse200325Details
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200325Items) GetDetailsOk() ([]InlineResponse200325Details, bool) {
	if o == nil || isNil(o.Details) {
    return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *InlineResponse200325Items) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []InlineResponse200325Details and assigns it to the Details field.
func (o *InlineResponse200325Items) SetDetails(v []InlineResponse200325Details) {
	o.Details = v
}

func (o InlineResponse200325Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !isNil(o.JoinedAt) {
		toSerialize["joinedAt"] = o.JoinedAt
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200325Items struct {
	value *InlineResponse200325Items
	isSet bool
}

func (v NullableInlineResponse200325Items) Get() *InlineResponse200325Items {
	return v.value
}

func (v *NullableInlineResponse200325Items) Set(val *InlineResponse200325Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200325Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200325Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200325Items(val *InlineResponse200325Items) *NullableInlineResponse200325Items {
	return &NullableInlineResponse200325Items{value: val, isSet: true}
}

func (v NullableInlineResponse200325Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200325Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


