/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20077 struct for InlineResponse20077
type InlineResponse20077 struct {
	// Network Id
	NetworkId *string `json:"networkId,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Device model
	Model *string `json:"model,omitempty"`
	// Last reported time for the device
	LastReportedAt *time.Time `json:"lastReportedAt,omitempty"`
	// Uplinks info
	Uplinks []OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks `json:"uplinks,omitempty"`
}

// NewInlineResponse20077 instantiates a new InlineResponse20077 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20077() *InlineResponse20077 {
	this := InlineResponse20077{}
	return &this
}

// NewInlineResponse20077WithDefaults instantiates a new InlineResponse20077 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20077WithDefaults() *InlineResponse20077 {
	this := InlineResponse20077{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse20077) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse20077) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse20077) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20077) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20077) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20077) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse20077) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse20077) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse20077) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *InlineResponse20077) GetLastReportedAt() time.Time {
	if o == nil || isNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastReportedAt) {
    return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *InlineResponse20077) HasLastReportedAt() bool {
	if o != nil && !isNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *InlineResponse20077) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineResponse20077) GetUplinks() []OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20077) GetUplinksOk() ([]OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineResponse20077) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks and assigns it to the Uplinks field.
func (o *InlineResponse20077) SetUplinks(v []OrganizationsOrganizationIdCellularGatewayUplinkStatusesUplinks) {
	o.Uplinks = v
}

func (o InlineResponse20077) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20077 struct {
	value *InlineResponse20077
	isSet bool
}

func (v NullableInlineResponse20077) Get() *InlineResponse20077 {
	return v.value
}

func (v *NullableInlineResponse20077) Set(val *InlineResponse20077) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20077) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20077) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20077(val *InlineResponse20077) *NullableInlineResponse20077 {
	return &NullableInlineResponse20077{value: val, isSet: true}
}

func (v NullableInlineResponse20077) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20077) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


