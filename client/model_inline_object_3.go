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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// The name of a device
	Name *string `json:"name,omitempty"`
	// The list of tags of a device
	Tags []string `json:"tags,omitempty"`
	// The latitude of a device
	Lat *float32 `json:"lat,omitempty"`
	// The longitude of a device
	Lng *float32 `json:"lng,omitempty"`
	// The address of a device
	Address *string `json:"address,omitempty"`
	// The notes for the device. String. Limited to 255 characters.
	Notes *string `json:"notes,omitempty"`
	// Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	MoveMapMarker *bool `json:"moveMapMarker,omitempty"`
	// The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanId *string `json:"floorPlanId,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject3) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject3) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject3) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject3) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject3) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject3) SetTags(v []string) {
	o.Tags = v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *InlineObject3) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *InlineObject3) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *InlineObject3) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *InlineObject3) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *InlineObject3) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *InlineObject3) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineObject3) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineObject3) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineObject3) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject3) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject3) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject3) SetNotes(v string) {
	o.Notes = &v
}

// GetMoveMapMarker returns the MoveMapMarker field value if set, zero value otherwise.
func (o *InlineObject3) GetMoveMapMarker() bool {
	if o == nil || isNil(o.MoveMapMarker) {
		var ret bool
		return ret
	}
	return *o.MoveMapMarker
}

// GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetMoveMapMarkerOk() (*bool, bool) {
	if o == nil || isNil(o.MoveMapMarker) {
    return nil, false
	}
	return o.MoveMapMarker, true
}

// HasMoveMapMarker returns a boolean if a field has been set.
func (o *InlineObject3) HasMoveMapMarker() bool {
	if o != nil && !isNil(o.MoveMapMarker) {
		return true
	}

	return false
}

// SetMoveMapMarker gets a reference to the given bool and assigns it to the MoveMapMarker field.
func (o *InlineObject3) SetMoveMapMarker(v bool) {
	o.MoveMapMarker = &v
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *InlineObject3) GetSwitchProfileId() string {
	if o == nil || isNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchProfileId) {
    return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *InlineObject3) HasSwitchProfileId() bool {
	if o != nil && !isNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *InlineObject3) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *InlineObject3) GetFloorPlanId() string {
	if o == nil || isNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || isNil(o.FloorPlanId) {
    return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *InlineObject3) HasFloorPlanId() bool {
	if o != nil && !isNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *InlineObject3) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.MoveMapMarker) {
		toSerialize["moveMapMarker"] = o.MoveMapMarker
	}
	if !isNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !isNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


