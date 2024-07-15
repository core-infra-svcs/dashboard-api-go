/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200279 struct for InlineResponse200279
type InlineResponse200279 struct {
	Network *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork `json:"network,omitempty"`
	// Name of the appliance
	Name *string `json:"name,omitempty"`
	// Mac address of the appliance
	Mac *string `json:"mac,omitempty"`
	// Serial number of the appliance
	Serial *string `json:"serial,omitempty"`
	// Model of the appliance
	Model *string `json:"model,omitempty"`
	Utilization *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization `json:"utilization,omitempty"`
}

// NewInlineResponse200279 instantiates a new InlineResponse200279 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200279() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// NewInlineResponse200279WithDefaults instantiates a new InlineResponse200279 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200279WithDefaults() *InlineResponse200279 {
	this := InlineResponse200279{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200279) GetNetwork() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200279) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork and assigns it to the Network field.
func (o *InlineResponse200279) SetNetwork(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) {
	o.Network = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200279) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200279) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200279) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200279) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200279) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200279) SetMac(v string) {
	o.Mac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200279) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200279) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200279) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200279) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200279) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200279) SetModel(v string) {
	o.Model = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *InlineResponse200279) GetUtilization() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization {
	if o == nil || isNil(o.Utilization) {
		var ret OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200279) GetUtilizationOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization, bool) {
	if o == nil || isNil(o.Utilization) {
    return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *InlineResponse200279) HasUtilization() bool {
	if o != nil && !isNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization and assigns it to the Utilization field.
func (o *InlineResponse200279) SetUtilization(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationUtilization) {
	o.Utilization = &v
}

func (o InlineResponse200279) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200279 struct {
	value *InlineResponse200279
	isSet bool
}

func (v NullableInlineResponse200279) Get() *InlineResponse200279 {
	return v.value
}

func (v *NullableInlineResponse200279) Set(val *InlineResponse200279) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200279) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200279) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200279(val *InlineResponse200279) *NullableInlineResponse200279 {
	return &NullableInlineResponse200279{value: val, isSet: true}
}

func (v NullableInlineResponse200279) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200279) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

