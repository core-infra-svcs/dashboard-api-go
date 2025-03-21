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

// InlineResponse200341Items struct for InlineResponse200341Items
type InlineResponse200341Items struct {
	// Unique serial number for the device
	Serial *string `json:"serial,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// List of custom tags for the device
	Tags []string `json:"tags,omitempty"`
	Network *InlineResponse200270Network `json:"network,omitempty"`
	// Events indicating power mode changes for the device
	Events []InlineResponse200341Events `json:"events,omitempty"`
}

// NewInlineResponse200341Items instantiates a new InlineResponse200341Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200341Items() *InlineResponse200341Items {
	this := InlineResponse200341Items{}
	return &this
}

// NewInlineResponse200341ItemsWithDefaults instantiates a new InlineResponse200341Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200341ItemsWithDefaults() *InlineResponse200341Items {
	this := InlineResponse200341Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200341Items) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200341Items) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200341Items) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200341Items) SetMac(v string) {
	o.Mac = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200341Items) SetTags(v []string) {
	o.Tags = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetNetwork() InlineResponse200270Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200270Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetNetworkOk() (*InlineResponse200270Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200270Network and assigns it to the Network field.
func (o *InlineResponse200341Items) SetNetwork(v InlineResponse200270Network) {
	o.Network = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *InlineResponse200341Items) GetEvents() []InlineResponse200341Events {
	if o == nil || isNil(o.Events) {
		var ret []InlineResponse200341Events
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200341Items) GetEventsOk() ([]InlineResponse200341Events, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *InlineResponse200341Items) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []InlineResponse200341Events and assigns it to the Events field.
func (o *InlineResponse200341Items) SetEvents(v []InlineResponse200341Events) {
	o.Events = v
}

func (o InlineResponse200341Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200341Items struct {
	value *InlineResponse200341Items
	isSet bool
}

func (v NullableInlineResponse200341Items) Get() *InlineResponse200341Items {
	return v.value
}

func (v *NullableInlineResponse200341Items) Set(val *InlineResponse200341Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200341Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200341Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200341Items(val *InlineResponse200341Items) *NullableInlineResponse200341Items {
	return &NullableInlineResponse200341Items{value: val, isSet: true}
}

func (v NullableInlineResponse200341Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200341Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


