/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject106 struct for InlineObject106
type InlineObject106 struct {
	// The name of your floor plan.
	Name *string `json:"name,omitempty"`
	Center *NetworksNetworkIdFloorPlansFloorPlanIdCenter `json:"center,omitempty"`
	BottomLeftCorner *NetworksNetworkIdFloorPlansBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *NetworksNetworkIdFloorPlansBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *NetworksNetworkIdFloorPlansTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *NetworksNetworkIdFloorPlansTopRightCorner `json:"topRightCorner,omitempty"`
	// The floor number of the floors within the building
	FloorNumber *int32 `json:"floorNumber,omitempty"`
	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	ImageContents *string `json:"imageContents,omitempty"`
}

// NewInlineObject106 instantiates a new InlineObject106 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject106() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// NewInlineObject106WithDefaults instantiates a new InlineObject106 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject106WithDefaults() *InlineObject106 {
	this := InlineObject106{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject106) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject106) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject106) SetName(v string) {
	o.Name = &v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *InlineObject106) GetCenter() NetworksNetworkIdFloorPlansFloorPlanIdCenter {
	if o == nil || isNil(o.Center) {
		var ret NetworksNetworkIdFloorPlansFloorPlanIdCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetCenterOk() (*NetworksNetworkIdFloorPlansFloorPlanIdCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *InlineObject106) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given NetworksNetworkIdFloorPlansFloorPlanIdCenter and assigns it to the Center field.
func (o *InlineObject106) SetCenter(v NetworksNetworkIdFloorPlansFloorPlanIdCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *InlineObject106) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret NetworksNetworkIdFloorPlansBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *InlineObject106) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *InlineObject106) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *InlineObject106) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret NetworksNetworkIdFloorPlansBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *InlineObject106) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *InlineObject106) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *InlineObject106) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret NetworksNetworkIdFloorPlansTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *InlineObject106) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *InlineObject106) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *InlineObject106) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret NetworksNetworkIdFloorPlansTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *InlineObject106) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given NetworksNetworkIdFloorPlansTopRightCorner and assigns it to the TopRightCorner field.
func (o *InlineObject106) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner) {
	o.TopRightCorner = &v
}

// GetFloorNumber returns the FloorNumber field value if set, zero value otherwise.
func (o *InlineObject106) GetFloorNumber() int32 {
	if o == nil || isNil(o.FloorNumber) {
		var ret int32
		return ret
	}
	return *o.FloorNumber
}

// GetFloorNumberOk returns a tuple with the FloorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetFloorNumberOk() (*int32, bool) {
	if o == nil || isNil(o.FloorNumber) {
    return nil, false
	}
	return o.FloorNumber, true
}

// HasFloorNumber returns a boolean if a field has been set.
func (o *InlineObject106) HasFloorNumber() bool {
	if o != nil && !isNil(o.FloorNumber) {
		return true
	}

	return false
}

// SetFloorNumber gets a reference to the given int32 and assigns it to the FloorNumber field.
func (o *InlineObject106) SetFloorNumber(v int32) {
	o.FloorNumber = &v
}

// GetImageContents returns the ImageContents field value if set, zero value otherwise.
func (o *InlineObject106) GetImageContents() string {
	if o == nil || isNil(o.ImageContents) {
		var ret string
		return ret
	}
	return *o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject106) GetImageContentsOk() (*string, bool) {
	if o == nil || isNil(o.ImageContents) {
    return nil, false
	}
	return o.ImageContents, true
}

// HasImageContents returns a boolean if a field has been set.
func (o *InlineObject106) HasImageContents() bool {
	if o != nil && !isNil(o.ImageContents) {
		return true
	}

	return false
}

// SetImageContents gets a reference to the given string and assigns it to the ImageContents field.
func (o *InlineObject106) SetImageContents(v string) {
	o.ImageContents = &v
}

func (o InlineObject106) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Center) {
		toSerialize["center"] = o.Center
	}
	if !isNil(o.BottomLeftCorner) {
		toSerialize["bottomLeftCorner"] = o.BottomLeftCorner
	}
	if !isNil(o.BottomRightCorner) {
		toSerialize["bottomRightCorner"] = o.BottomRightCorner
	}
	if !isNil(o.TopLeftCorner) {
		toSerialize["topLeftCorner"] = o.TopLeftCorner
	}
	if !isNil(o.TopRightCorner) {
		toSerialize["topRightCorner"] = o.TopRightCorner
	}
	if !isNil(o.FloorNumber) {
		toSerialize["floorNumber"] = o.FloorNumber
	}
	if !isNil(o.ImageContents) {
		toSerialize["imageContents"] = o.ImageContents
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject106 struct {
	value *InlineObject106
	isSet bool
}

func (v NullableInlineObject106) Get() *InlineObject106 {
	return v.value
}

func (v *NullableInlineObject106) Set(val *InlineObject106) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject106) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject106) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject106(val *InlineObject106) *NullableInlineObject106 {
	return &NullableInlineObject106{value: val, isSet: true}
}

func (v NullableInlineObject106) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject106) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


