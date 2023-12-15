/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20042 struct for InlineResponse20042
type InlineResponse20042 struct {
	// Floor plan ID
	FloorPlanId *string `json:"floorPlanId,omitempty"`
	// The url link for the floor plan image.
	ImageUrl *string `json:"imageUrl,omitempty"`
	// The time the image url link will expire.
	ImageUrlExpiresAt *string `json:"imageUrlExpiresAt,omitempty"`
	// The format type of the image.
	ImageExtension *string `json:"imageExtension,omitempty"`
	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	ImageMd5 *string `json:"imageMd5,omitempty"`
	// The name of your floor plan.
	Name *string `json:"name,omitempty"`
	// List of devices for the floorplan
	Devices []NetworksNetworkIdFloorPlansDevices `json:"devices,omitempty"`
	// The width of your floor plan.
	Width *float32 `json:"width,omitempty"`
	// The height of your floor plan.
	Height *float32 `json:"height,omitempty"`
	Center *NetworksNetworkIdFloorPlansCenter `json:"center,omitempty"`
	BottomLeftCorner *NetworksNetworkIdFloorPlansBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *NetworksNetworkIdFloorPlansBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *NetworksNetworkIdFloorPlansTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *NetworksNetworkIdFloorPlansTopRightCorner `json:"topRightCorner,omitempty"`
}

// NewInlineResponse20042 instantiates a new InlineResponse20042 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20042() *InlineResponse20042 {
	this := InlineResponse20042{}
	return &this
}

// NewInlineResponse20042WithDefaults instantiates a new InlineResponse20042 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20042WithDefaults() *InlineResponse20042 {
	this := InlineResponse20042{}
	return &this
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *InlineResponse20042) GetFloorPlanId() string {
	if o == nil || isNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || isNil(o.FloorPlanId) {
    return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *InlineResponse20042) HasFloorPlanId() bool {
	if o != nil && !isNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *InlineResponse20042) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *InlineResponse20042) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
    return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *InlineResponse20042) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *InlineResponse20042) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetImageUrlExpiresAt returns the ImageUrlExpiresAt field value if set, zero value otherwise.
func (o *InlineResponse20042) GetImageUrlExpiresAt() string {
	if o == nil || isNil(o.ImageUrlExpiresAt) {
		var ret string
		return ret
	}
	return *o.ImageUrlExpiresAt
}

// GetImageUrlExpiresAtOk returns a tuple with the ImageUrlExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetImageUrlExpiresAtOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrlExpiresAt) {
    return nil, false
	}
	return o.ImageUrlExpiresAt, true
}

// HasImageUrlExpiresAt returns a boolean if a field has been set.
func (o *InlineResponse20042) HasImageUrlExpiresAt() bool {
	if o != nil && !isNil(o.ImageUrlExpiresAt) {
		return true
	}

	return false
}

// SetImageUrlExpiresAt gets a reference to the given string and assigns it to the ImageUrlExpiresAt field.
func (o *InlineResponse20042) SetImageUrlExpiresAt(v string) {
	o.ImageUrlExpiresAt = &v
}

// GetImageExtension returns the ImageExtension field value if set, zero value otherwise.
func (o *InlineResponse20042) GetImageExtension() string {
	if o == nil || isNil(o.ImageExtension) {
		var ret string
		return ret
	}
	return *o.ImageExtension
}

// GetImageExtensionOk returns a tuple with the ImageExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetImageExtensionOk() (*string, bool) {
	if o == nil || isNil(o.ImageExtension) {
    return nil, false
	}
	return o.ImageExtension, true
}

// HasImageExtension returns a boolean if a field has been set.
func (o *InlineResponse20042) HasImageExtension() bool {
	if o != nil && !isNil(o.ImageExtension) {
		return true
	}

	return false
}

// SetImageExtension gets a reference to the given string and assigns it to the ImageExtension field.
func (o *InlineResponse20042) SetImageExtension(v string) {
	o.ImageExtension = &v
}

// GetImageMd5 returns the ImageMd5 field value if set, zero value otherwise.
func (o *InlineResponse20042) GetImageMd5() string {
	if o == nil || isNil(o.ImageMd5) {
		var ret string
		return ret
	}
	return *o.ImageMd5
}

// GetImageMd5Ok returns a tuple with the ImageMd5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetImageMd5Ok() (*string, bool) {
	if o == nil || isNil(o.ImageMd5) {
    return nil, false
	}
	return o.ImageMd5, true
}

// HasImageMd5 returns a boolean if a field has been set.
func (o *InlineResponse20042) HasImageMd5() bool {
	if o != nil && !isNil(o.ImageMd5) {
		return true
	}

	return false
}

// SetImageMd5 gets a reference to the given string and assigns it to the ImageMd5 field.
func (o *InlineResponse20042) SetImageMd5(v string) {
	o.ImageMd5 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20042) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20042) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20042) SetName(v string) {
	o.Name = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *InlineResponse20042) GetDevices() []NetworksNetworkIdFloorPlansDevices {
	if o == nil || isNil(o.Devices) {
		var ret []NetworksNetworkIdFloorPlansDevices
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetDevicesOk() ([]NetworksNetworkIdFloorPlansDevices, bool) {
	if o == nil || isNil(o.Devices) {
    return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *InlineResponse20042) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NetworksNetworkIdFloorPlansDevices and assigns it to the Devices field.
func (o *InlineResponse20042) SetDevices(v []NetworksNetworkIdFloorPlansDevices) {
	o.Devices = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *InlineResponse20042) GetWidth() float32 {
	if o == nil || isNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetWidthOk() (*float32, bool) {
	if o == nil || isNil(o.Width) {
    return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *InlineResponse20042) HasWidth() bool {
	if o != nil && !isNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *InlineResponse20042) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *InlineResponse20042) GetHeight() float32 {
	if o == nil || isNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetHeightOk() (*float32, bool) {
	if o == nil || isNil(o.Height) {
    return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *InlineResponse20042) HasHeight() bool {
	if o != nil && !isNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *InlineResponse20042) SetHeight(v float32) {
	o.Height = &v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *InlineResponse20042) GetCenter() NetworksNetworkIdFloorPlansCenter {
	if o == nil || isNil(o.Center) {
		var ret NetworksNetworkIdFloorPlansCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *InlineResponse20042) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given NetworksNetworkIdFloorPlansCenter and assigns it to the Center field.
func (o *InlineResponse20042) SetCenter(v NetworksNetworkIdFloorPlansCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *InlineResponse20042) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret NetworksNetworkIdFloorPlansBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *InlineResponse20042) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *InlineResponse20042) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *InlineResponse20042) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret NetworksNetworkIdFloorPlansBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *InlineResponse20042) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given NetworksNetworkIdFloorPlansBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *InlineResponse20042) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *InlineResponse20042) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret NetworksNetworkIdFloorPlansTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *InlineResponse20042) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given NetworksNetworkIdFloorPlansTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *InlineResponse20042) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *InlineResponse20042) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret NetworksNetworkIdFloorPlansTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20042) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *InlineResponse20042) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given NetworksNetworkIdFloorPlansTopRightCorner and assigns it to the TopRightCorner field.
func (o *InlineResponse20042) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner) {
	o.TopRightCorner = &v
}

func (o InlineResponse20042) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	if !isNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !isNil(o.ImageUrlExpiresAt) {
		toSerialize["imageUrlExpiresAt"] = o.ImageUrlExpiresAt
	}
	if !isNil(o.ImageExtension) {
		toSerialize["imageExtension"] = o.ImageExtension
	}
	if !isNil(o.ImageMd5) {
		toSerialize["imageMd5"] = o.ImageMd5
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !isNil(o.Height) {
		toSerialize["height"] = o.Height
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20042 struct {
	value *InlineResponse20042
	isSet bool
}

func (v NullableInlineResponse20042) Get() *InlineResponse20042 {
	return v.value
}

func (v *NullableInlineResponse20042) Set(val *InlineResponse20042) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20042) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20042) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20042(val *InlineResponse20042) *NullableInlineResponse20042 {
	return &NullableInlineResponse20042{value: val, isSet: true}
}

func (v NullableInlineResponse20042) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20042) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


