/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject85 struct for InlineObject85
type InlineObject85 struct {
	// The name of your floor plan.
	Name *string `json:"name,omitempty"`
	Center *NetworksNetworkIdFloorPlansFloorPlanIdCenter `json:"center,omitempty"`
	BottomLeftCorner *NetworksNetworkIdFloorPlansBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *NetworksNetworkIdFloorPlansBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *NetworksNetworkIdFloorPlansTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *NetworksNetworkIdFloorPlansTopRightCorner `json:"topRightCorner,omitempty"`
	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	ImageContents *string `json:"imageContents,omitempty"`
}

// NewInlineObject85 instantiates a new InlineObject85 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject85() *InlineObject85 {
	this := InlineObject85{}
	return &this
}

// NewInlineObject85WithDefaults instantiates a new InlineObject85 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject85WithDefaults() *InlineObject85 {
	this := InlineObject85{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject85) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject85) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject85) SetName(v string) {
	o.Name = &v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *InlineObject85) GetCenter() NetworksNetworkIdFloorPlansFloorPlanIdCenter {
	if o == nil || isNil(o.Center) {
		var ret NetworksNetworkIdFloorPlansFloorPlanIdCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetCenterOk() (*NetworksNetworkIdFloorPlansFloorPlanIdCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *InlineObject85) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given NetworksNetworkIdFloorPlansFloorPlanIdCenter and assigns it to the Center field.
func (o *InlineObject85) SetCenter(v NetworksNetworkIdFloorPlansFloorPlanIdCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *InlineObject85) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret NetworksNetworkIdFloorPlansBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *InlineObject85) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *InlineObject85) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *InlineObject85) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret NetworksNetworkIdFloorPlansBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *InlineObject85) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *InlineObject85) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *InlineObject85) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret NetworksNetworkIdFloorPlansTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *InlineObject85) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *InlineObject85) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *InlineObject85) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret NetworksNetworkIdFloorPlansTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *InlineObject85) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given NetworksNetworkIdFloorPlansTopRightCorner and assigns it to the TopRightCorner field.
func (o *InlineObject85) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner) {
	o.TopRightCorner = &v
}

// GetImageContents returns the ImageContents field value if set, zero value otherwise.
func (o *InlineObject85) GetImageContents() string {
	if o == nil || isNil(o.ImageContents) {
		var ret string
		return ret
	}
	return *o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject85) GetImageContentsOk() (*string, bool) {
	if o == nil || isNil(o.ImageContents) {
    return nil, false
	}
	return o.ImageContents, true
}

// HasImageContents returns a boolean if a field has been set.
func (o *InlineObject85) HasImageContents() bool {
	if o != nil && !isNil(o.ImageContents) {
		return true
	}

	return false
}

// SetImageContents gets a reference to the given string and assigns it to the ImageContents field.
func (o *InlineObject85) SetImageContents(v string) {
	o.ImageContents = &v
}

func (o InlineObject85) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ImageContents) {
		toSerialize["imageContents"] = o.ImageContents
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject85 struct {
	value *InlineObject85
	isSet bool
}

func (v NullableInlineObject85) Get() *InlineObject85 {
	return v.value
}

func (v *NullableInlineObject85) Set(val *InlineObject85) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject85) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject85) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject85(val *InlineObject85) *NullableInlineObject85 {
	return &NullableInlineObject85{value: val, isSet: true}
}

func (v NullableInlineObject85) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject85) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


