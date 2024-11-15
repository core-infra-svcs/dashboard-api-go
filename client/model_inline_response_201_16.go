/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20116 struct for InlineResponse20116
type InlineResponse20116 struct {
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

// NewInlineResponse20116 instantiates a new InlineResponse20116 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20116() *InlineResponse20116 {
	this := InlineResponse20116{}
	return &this
}

// NewInlineResponse20116WithDefaults instantiates a new InlineResponse20116 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20116WithDefaults() *InlineResponse20116 {
	this := InlineResponse20116{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse20116) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20116) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse20116) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse20116) SetMessage(v string) {
	o.Message = &v
}

// GetUdi returns the Udi field value if set, zero value otherwise.
func (o *InlineResponse20116) GetUdi() string {
	if o == nil || isNil(o.Udi) {
		var ret string
		return ret
	}
	return *o.Udi
}

// GetUdiOk returns a tuple with the Udi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20116) GetUdiOk() (*string, bool) {
	if o == nil || isNil(o.Udi) {
    return nil, false
	}
	return o.Udi, true
}

// HasUdi returns a boolean if a field has been set.
func (o *InlineResponse20116) HasUdi() bool {
	if o != nil && !isNil(o.Udi) {
		return true
	}

	return false
}

// SetUdi gets a reference to the given string and assigns it to the Udi field.
func (o *InlineResponse20116) SetUdi(v string) {
	o.Udi = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineResponse20116) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20116) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineResponse20116) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineResponse20116) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20116) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20116) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20116) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20116) SetStatus(v string) {
	o.Status = &v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise.
func (o *InlineResponse20116) GetConfigParams() OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams {
	if o == nil || isNil(o.ConfigParams) {
		var ret OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams
		return ret
	}
	return *o.ConfigParams
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20116) GetConfigParamsOk() (*OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams, bool) {
	if o == nil || isNil(o.ConfigParams) {
    return nil, false
	}
	return o.ConfigParams, true
}

// HasConfigParams returns a boolean if a field has been set.
func (o *InlineResponse20116) HasConfigParams() bool {
	if o != nil && !isNil(o.ConfigParams) {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams and assigns it to the ConfigParams field.
func (o *InlineResponse20116) SetConfigParams(v OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareConfigParams) {
	o.ConfigParams = &v
}

func (o InlineResponse20116) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20116 struct {
	value *InlineResponse20116
	isSet bool
}

func (v NullableInlineResponse20116) Get() *InlineResponse20116 {
	return v.value
}

func (v *NullableInlineResponse20116) Set(val *InlineResponse20116) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20116) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20116) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20116(val *InlineResponse20116) *NullableInlineResponse20116 {
	return &NullableInlineResponse20116{value: val, isSet: true}
}

func (v NullableInlineResponse20116) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20116) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

