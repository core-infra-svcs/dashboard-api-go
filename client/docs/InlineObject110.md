# InlineObject110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your floor plan. | [optional] 
**Center** | Pointer to [**NetworksNetworkIdFloorPlansFloorPlanIdCenter**](NetworksNetworkIdFloorPlansFloorPlanIdCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomLeftCorner**](NetworksNetworkIdFloorPlansBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomRightCorner**](NetworksNetworkIdFloorPlansBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopLeftCorner**](NetworksNetworkIdFloorPlansTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopRightCorner**](NetworksNetworkIdFloorPlansTopRightCorner.md) |  | [optional] 
**FloorNumber** | Pointer to **float32** | The floor number of the floors within the building | [optional] 
**ImageContents** | Pointer to **string** | The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields (&#39;center, &#39;topLeftCorner&#39;, etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image. | [optional] 

## Methods

### NewInlineObject110

`func NewInlineObject110() *InlineObject110`

NewInlineObject110 instantiates a new InlineObject110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject110WithDefaults

`func NewInlineObject110WithDefaults() *InlineObject110`

NewInlineObject110WithDefaults instantiates a new InlineObject110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject110) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject110) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject110) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject110) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCenter

`func (o *InlineObject110) GetCenter() NetworksNetworkIdFloorPlansFloorPlanIdCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *InlineObject110) GetCenterOk() (*NetworksNetworkIdFloorPlansFloorPlanIdCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *InlineObject110) SetCenter(v NetworksNetworkIdFloorPlansFloorPlanIdCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *InlineObject110) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *InlineObject110) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *InlineObject110) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *InlineObject110) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *InlineObject110) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *InlineObject110) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *InlineObject110) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *InlineObject110) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *InlineObject110) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *InlineObject110) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *InlineObject110) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *InlineObject110) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *InlineObject110) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *InlineObject110) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *InlineObject110) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *InlineObject110) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *InlineObject110) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetFloorNumber

`func (o *InlineObject110) GetFloorNumber() float32`

GetFloorNumber returns the FloorNumber field if non-nil, zero value otherwise.

### GetFloorNumberOk

`func (o *InlineObject110) GetFloorNumberOk() (*float32, bool)`

GetFloorNumberOk returns a tuple with the FloorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorNumber

`func (o *InlineObject110) SetFloorNumber(v float32)`

SetFloorNumber sets FloorNumber field to given value.

### HasFloorNumber

`func (o *InlineObject110) HasFloorNumber() bool`

HasFloorNumber returns a boolean if a field has been set.

### GetImageContents

`func (o *InlineObject110) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *InlineObject110) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *InlineObject110) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.

### HasImageContents

`func (o *InlineObject110) HasImageContents() bool`

HasImageContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


