/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200314 struct for InlineResponse200314
type InlineResponse200314 struct {
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Mac address of the device
	Mac *string `json:"mac,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
	Network *OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork `json:"network,omitempty"`
	Usage *OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage `json:"usage,omitempty"`
	Clients *OrganizationsOrganizationIdSummaryTopDevicesByUsageClients `json:"clients,omitempty"`
}

// NewInlineResponse200314 instantiates a new InlineResponse200314 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200314() *InlineResponse200314 {
	this := InlineResponse200314{}
	return &this
}

// NewInlineResponse200314WithDefaults instantiates a new InlineResponse200314 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200314WithDefaults() *InlineResponse200314 {
	this := InlineResponse200314{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200314) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200314) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200314) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200314) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200314) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200314) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200314) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200314) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200314) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200314) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200314) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200314) SetMac(v string) {
	o.Mac = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200314) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200314) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200314) SetProductType(v string) {
	o.ProductType = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200314) GetNetwork() OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetNetworkOk() (*OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200314) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork and assigns it to the Network field.
func (o *InlineResponse200314) SetNetwork(v OrganizationsOrganizationIdSummaryTopAppliancesByUtilizationNetwork) {
	o.Network = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200314) GetUsage() OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage {
	if o == nil || isNil(o.Usage) {
		var ret OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetUsageOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200314) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage and assigns it to the Usage field.
func (o *InlineResponse200314) SetUsage(v OrganizationsOrganizationIdSummaryTopDevicesByUsageUsage) {
	o.Usage = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *InlineResponse200314) GetClients() OrganizationsOrganizationIdSummaryTopDevicesByUsageClients {
	if o == nil || isNil(o.Clients) {
		var ret OrganizationsOrganizationIdSummaryTopDevicesByUsageClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200314) GetClientsOk() (*OrganizationsOrganizationIdSummaryTopDevicesByUsageClients, bool) {
	if o == nil || isNil(o.Clients) {
    return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *InlineResponse200314) HasClients() bool {
	if o != nil && !isNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given OrganizationsOrganizationIdSummaryTopDevicesByUsageClients and assigns it to the Clients field.
func (o *InlineResponse200314) SetClients(v OrganizationsOrganizationIdSummaryTopDevicesByUsageClients) {
	o.Clients = &v
}

func (o InlineResponse200314) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200314 struct {
	value *InlineResponse200314
	isSet bool
}

func (v NullableInlineResponse200314) Get() *InlineResponse200314 {
	return v.value
}

func (v *NullableInlineResponse200314) Set(val *InlineResponse200314) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200314) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200314) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200314(val *InlineResponse200314) *NullableInlineResponse200314 {
	return &NullableInlineResponse200314{value: val, isSet: true}
}

func (v NullableInlineResponse200314) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200314) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


