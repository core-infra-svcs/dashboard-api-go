/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2017 struct for InlineResponse2017
type InlineResponse2017 struct {
	// Message related to whether or not the device was found and can be imported.
	Message *string `json:"message,omitempty"`
	// Device UDI certificate
	Udi *string `json:"udi,omitempty"`
	// Import ID from the Import operation
	DeviceId *string `json:"deviceId,omitempty"`
	// The import status of the device
	Status *string `json:"status,omitempty"`
	ConfigParams *OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams `json:"configParams,omitempty"`
}

// NewInlineResponse2017 instantiates a new InlineResponse2017 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2017() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// NewInlineResponse2017WithDefaults instantiates a new InlineResponse2017 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2017WithDefaults() *InlineResponse2017 {
	this := InlineResponse2017{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse2017) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse2017) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse2017) SetMessage(v string) {
	o.Message = &v
}

// GetUdi returns the Udi field value if set, zero value otherwise.
func (o *InlineResponse2017) GetUdi() string {
	if o == nil || isNil(o.Udi) {
		var ret string
		return ret
	}
	return *o.Udi
}

// GetUdiOk returns a tuple with the Udi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetUdiOk() (*string, bool) {
	if o == nil || isNil(o.Udi) {
    return nil, false
	}
	return o.Udi, true
}

// HasUdi returns a boolean if a field has been set.
func (o *InlineResponse2017) HasUdi() bool {
	if o != nil && !isNil(o.Udi) {
		return true
	}

	return false
}

// SetUdi gets a reference to the given string and assigns it to the Udi field.
func (o *InlineResponse2017) SetUdi(v string) {
	o.Udi = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse2017) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse2017) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse2017) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2017) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2017) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2017) SetStatus(v string) {
	o.Status = &v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise.
func (o *InlineResponse2017) GetConfigParams() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	if o == nil || isNil(o.ConfigParams) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams
		return ret
	}
	return *o.ConfigParams
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2017) GetConfigParamsOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams, bool) {
	if o == nil || isNil(o.ConfigParams) {
    return nil, false
	}
	return o.ConfigParams, true
}

// HasConfigParams returns a boolean if a field has been set.
func (o *InlineResponse2017) HasConfigParams() bool {
	if o != nil && !isNil(o.ConfigParams) {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams and assigns it to the ConfigParams field.
func (o *InlineResponse2017) SetConfigParams(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) {
	o.ConfigParams = &v
}

func (o InlineResponse2017) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Udi) {
		toSerialize["udi"] = o.Udi
	}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ConfigParams) {
		toSerialize["configParams"] = o.ConfigParams
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2017 struct {
	value *InlineResponse2017
	isSet bool
}

func (v NullableInlineResponse2017) Get() *InlineResponse2017 {
	return v.value
}

func (v *NullableInlineResponse2017) Set(val *InlineResponse2017) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2017) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2017) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2017(val *InlineResponse2017) *NullableInlineResponse2017 {
	return &NullableInlineResponse2017{value: val, isSet: true}
}

func (v NullableInlineResponse2017) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2017) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


