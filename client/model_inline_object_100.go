/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject100 struct for InlineObject100
type InlineObject100 struct {
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

// NewInlineObject100 instantiates a new InlineObject100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject100(name string, imageContents string) *InlineObject100 {
	this := InlineObject100{}
	this.Name = name
	this.ImageContents = imageContents
	return &this
}

// NewInlineObject100WithDefaults instantiates a new InlineObject100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject100WithDefaults() *InlineObject100 {
	this := InlineObject100{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject100) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject100) SetName(v string) {
	o.Name = v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *InlineObject100) GetCenter() NetworksNetworkIdFloorPlansCenter {
	if o == nil || isNil(o.Center) {
		var ret NetworksNetworkIdFloorPlansCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *InlineObject100) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given NetworksNetworkIdFloorPlansCenter and assigns it to the Center field.
func (o *InlineObject100) SetCenter(v NetworksNetworkIdFloorPlansCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *InlineObject100) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret NetworksNetworkIdFloorPlansBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *InlineObject100) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *InlineObject100) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *InlineObject100) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret NetworksNetworkIdFloorPlansBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *InlineObject100) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *InlineObject100) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *InlineObject100) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret NetworksNetworkIdFloorPlansTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *InlineObject100) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *InlineObject100) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *InlineObject100) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret NetworksNetworkIdFloorPlansTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *InlineObject100) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given NetworksNetworkIdFloorPlansTopRightCorner and assigns it to the TopRightCorner field.
func (o *InlineObject100) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner) {
	o.TopRightCorner = &v
}

// GetImageContents returns the ImageContents field value
func (o *InlineObject100) GetImageContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value
// and a boolean to check if the value has been set.
func (o *InlineObject100) GetImageContentsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImageContents, true
}

// SetImageContents sets field value
func (o *InlineObject100) SetImageContents(v string) {
	o.ImageContents = v
}

func (o InlineObject100) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject100 struct {
	value *InlineObject100
	isSet bool
}

func (v NullableInlineObject100) Get() *InlineObject100 {
	return v.value
}

func (v *NullableInlineObject100) Set(val *InlineObject100) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject100) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject100(val *InlineObject100) *NullableInlineObject100 {
	return &NullableInlineObject100{value: val, isSet: true}
}

func (v NullableInlineObject100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


