/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject99 struct for InlineObject99
type InlineObject99 struct {
	// The name of your floor plan.
	Name string `json:"name"`
	Center *NetworksNetworkIdFloorPlansCenter `json:"center,omitempty"`
	BottomLeftCorner *NetworksNetworkIdFloorPlansBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *NetworksNetworkIdFloorPlansBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *NetworksNetworkIdFloorPlansTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *NetworksNetworkIdFloorPlansTopRightCorner `json:"topRightCorner,omitempty"`
	// The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	ImageContents string `json:"imageContents"`
}

// NewInlineObject99 instantiates a new InlineObject99 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject99(name string, imageContents string) *InlineObject99 {
	this := InlineObject99{}
	this.Name = name
	this.ImageContents = imageContents
	return &this
}

// NewInlineObject99WithDefaults instantiates a new InlineObject99 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject99WithDefaults() *InlineObject99 {
	this := InlineObject99{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject99) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject99) SetName(v string) {
	o.Name = v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *InlineObject99) GetCenter() NetworksNetworkIdFloorPlansCenter {
	if o == nil || isNil(o.Center) {
		var ret NetworksNetworkIdFloorPlansCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *InlineObject99) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given NetworksNetworkIdFloorPlansCenter and assigns it to the Center field.
func (o *InlineObject99) SetCenter(v NetworksNetworkIdFloorPlansCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *InlineObject99) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret NetworksNetworkIdFloorPlansBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *InlineObject99) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *InlineObject99) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *InlineObject99) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret NetworksNetworkIdFloorPlansBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *InlineObject99) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *InlineObject99) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *InlineObject99) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret NetworksNetworkIdFloorPlansTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *InlineObject99) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *InlineObject99) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *InlineObject99) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret NetworksNetworkIdFloorPlansTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *InlineObject99) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given NetworksNetworkIdFloorPlansTopRightCorner and assigns it to the TopRightCorner field.
func (o *InlineObject99) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner) {
	o.TopRightCorner = &v
}

// GetImageContents returns the ImageContents field value
func (o *InlineObject99) GetImageContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value
// and a boolean to check if the value has been set.
func (o *InlineObject99) GetImageContentsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImageContents, true
}

// SetImageContents sets field value
func (o *InlineObject99) SetImageContents(v string) {
	o.ImageContents = v
}

func (o InlineObject99) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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
	if true {
		toSerialize["imageContents"] = o.ImageContents
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject99 struct {
	value *InlineObject99
	isSet bool
}

func (v NullableInlineObject99) Get() *InlineObject99 {
	return v.value
}

func (v *NullableInlineObject99) Set(val *InlineObject99) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject99) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject99) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject99(val *InlineObject99) *NullableInlineObject99 {
	return &NullableInlineObject99{value: val, isSet: true}
}

func (v NullableInlineObject99) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject99) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


