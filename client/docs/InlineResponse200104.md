# InlineResponse200104

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**ImageUrl** | Pointer to **string** | The url link for the floor plan image. | [optional] 
**ImageUrlExpiresAt** | Pointer to **string** | The time the image url link will expire. | [optional] 
**ImageExtension** | Pointer to **string** | The format type of the image. | [optional] 
**ImageMd5** | Pointer to **string** | The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields (&#39;center, &#39;topLeftCorner&#39;, etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image. | [optional] 
**Name** | Pointer to **string** | The name of your floor plan. | [optional] 
**Devices** | Pointer to [**[]InlineResponse20096**](InlineResponse20096.md) | List of devices for the floorplan | [optional] 
**Width** | Pointer to **float32** | The width of your floor plan. | [optional] 
**Height** | Pointer to **float32** | The height of your floor plan. | [optional] 
**Center** | Pointer to [**NetworksNetworkIdFloorPlansCenter**](NetworksNetworkIdFloorPlansCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomLeftCorner**](NetworksNetworkIdFloorPlansBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomRightCorner**](NetworksNetworkIdFloorPlansBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopLeftCorner**](NetworksNetworkIdFloorPlansTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopRightCorner**](NetworksNetworkIdFloorPlansTopRightCorner.md) |  | [optional] 
**FloorNumber** | Pointer to **float32** | The floor number of the floor within the building. | [optional] 

## Methods

### NewInlineResponse200104

`func NewInlineResponse200104() *InlineResponse200104`

NewInlineResponse200104 instantiates a new InlineResponse200104 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200104WithDefaults

`func NewInlineResponse200104WithDefaults() *InlineResponse200104`

NewInlineResponse200104WithDefaults instantiates a new InlineResponse200104 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloorPlanId

`func (o *InlineResponse200104) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse200104) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse200104) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse200104) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetImageUrl

`func (o *InlineResponse200104) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *InlineResponse200104) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *InlineResponse200104) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *InlineResponse200104) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImageUrlExpiresAt

`func (o *InlineResponse200104) GetImageUrlExpiresAt() string`

GetImageUrlExpiresAt returns the ImageUrlExpiresAt field if non-nil, zero value otherwise.

### GetImageUrlExpiresAtOk

`func (o *InlineResponse200104) GetImageUrlExpiresAtOk() (*string, bool)`

GetImageUrlExpiresAtOk returns a tuple with the ImageUrlExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrlExpiresAt

`func (o *InlineResponse200104) SetImageUrlExpiresAt(v string)`

SetImageUrlExpiresAt sets ImageUrlExpiresAt field to given value.

### HasImageUrlExpiresAt

`func (o *InlineResponse200104) HasImageUrlExpiresAt() bool`

HasImageUrlExpiresAt returns a boolean if a field has been set.

### GetImageExtension

`func (o *InlineResponse200104) GetImageExtension() string`

GetImageExtension returns the ImageExtension field if non-nil, zero value otherwise.

### GetImageExtensionOk

`func (o *InlineResponse200104) GetImageExtensionOk() (*string, bool)`

GetImageExtensionOk returns a tuple with the ImageExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtension

`func (o *InlineResponse200104) SetImageExtension(v string)`

SetImageExtension sets ImageExtension field to given value.

### HasImageExtension

`func (o *InlineResponse200104) HasImageExtension() bool`

HasImageExtension returns a boolean if a field has been set.

### GetImageMd5

`func (o *InlineResponse200104) GetImageMd5() string`

GetImageMd5 returns the ImageMd5 field if non-nil, zero value otherwise.

### GetImageMd5Ok

`func (o *InlineResponse200104) GetImageMd5Ok() (*string, bool)`

GetImageMd5Ok returns a tuple with the ImageMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMd5

`func (o *InlineResponse200104) SetImageMd5(v string)`

SetImageMd5 sets ImageMd5 field to given value.

### HasImageMd5

`func (o *InlineResponse200104) HasImageMd5() bool`

HasImageMd5 returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200104) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200104) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200104) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200104) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200104) GetDevices() []InlineResponse20096`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200104) GetDevicesOk() (*[]InlineResponse20096, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200104) SetDevices(v []InlineResponse20096)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200104) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetWidth

`func (o *InlineResponse200104) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *InlineResponse200104) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *InlineResponse200104) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *InlineResponse200104) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *InlineResponse200104) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *InlineResponse200104) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *InlineResponse200104) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *InlineResponse200104) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetCenter

`func (o *InlineResponse200104) GetCenter() NetworksNetworkIdFloorPlansCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *InlineResponse200104) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *InlineResponse200104) SetCenter(v NetworksNetworkIdFloorPlansCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *InlineResponse200104) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *InlineResponse200104) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *InlineResponse200104) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *InlineResponse200104) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *InlineResponse200104) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *InlineResponse200104) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *InlineResponse200104) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *InlineResponse200104) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *InlineResponse200104) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *InlineResponse200104) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *InlineResponse200104) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *InlineResponse200104) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *InlineResponse200104) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *InlineResponse200104) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *InlineResponse200104) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *InlineResponse200104) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *InlineResponse200104) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetFloorNumber

`func (o *InlineResponse200104) GetFloorNumber() float32`

GetFloorNumber returns the FloorNumber field if non-nil, zero value otherwise.

### GetFloorNumberOk

`func (o *InlineResponse200104) GetFloorNumberOk() (*float32, bool)`

GetFloorNumberOk returns a tuple with the FloorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorNumber

`func (o *InlineResponse200104) SetFloorNumber(v float32)`

SetFloorNumber sets FloorNumber field to given value.

### HasFloorNumber

`func (o *InlineResponse200104) HasFloorNumber() bool`

HasFloorNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


