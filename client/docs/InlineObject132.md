# InlineObject132

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of your floor plan. | [optional] 
**Center** | Pointer to [**NetworksNetworkIdFloorPlansFloorPlanIdCenter**](NetworksNetworkIdFloorPlansFloorPlanIdCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomLeftCorner1**](NetworksNetworkIdFloorPlansBottomLeftCorner1.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomRightCorner1**](NetworksNetworkIdFloorPlansBottomRightCorner1.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopLeftCorner1**](NetworksNetworkIdFloorPlansTopLeftCorner1.md) |  | [optional] 
**TopRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopRightCorner1**](NetworksNetworkIdFloorPlansTopRightCorner1.md) |  | [optional] 
**FloorNumber** | Pointer to **float32** | The floor number of the floors within the building | [optional] 
**ImageContents** | Pointer to **string** | The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields (&#39;center, &#39;topLeftCorner&#39;, etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image. | [optional] 

## Methods

### NewInlineObject132

`func NewInlineObject132() *InlineObject132`

NewInlineObject132 instantiates a new InlineObject132 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject132WithDefaults

`func NewInlineObject132WithDefaults() *InlineObject132`

NewInlineObject132WithDefaults instantiates a new InlineObject132 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject132) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject132) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject132) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject132) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCenter

`func (o *InlineObject132) GetCenter() NetworksNetworkIdFloorPlansFloorPlanIdCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *InlineObject132) GetCenterOk() (*NetworksNetworkIdFloorPlansFloorPlanIdCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *InlineObject132) SetCenter(v NetworksNetworkIdFloorPlansFloorPlanIdCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *InlineObject132) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *InlineObject132) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner1`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *InlineObject132) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner1, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *InlineObject132) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner1)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *InlineObject132) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *InlineObject132) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner1`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *InlineObject132) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner1, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *InlineObject132) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner1)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *InlineObject132) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *InlineObject132) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner1`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *InlineObject132) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner1, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *InlineObject132) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner1)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *InlineObject132) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *InlineObject132) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner1`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *InlineObject132) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner1, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *InlineObject132) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner1)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *InlineObject132) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetFloorNumber

`func (o *InlineObject132) GetFloorNumber() float32`

GetFloorNumber returns the FloorNumber field if non-nil, zero value otherwise.

### GetFloorNumberOk

`func (o *InlineObject132) GetFloorNumberOk() (*float32, bool)`

GetFloorNumberOk returns a tuple with the FloorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorNumber

`func (o *InlineObject132) SetFloorNumber(v float32)`

SetFloorNumber sets FloorNumber field to given value.

### HasFloorNumber

`func (o *InlineObject132) HasFloorNumber() bool`

HasFloorNumber returns a boolean if a field has been set.

### GetImageContents

`func (o *InlineObject132) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *InlineObject132) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *InlineObject132) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.

### HasImageContents

`func (o *InlineObject132) HasImageContents() bool`

HasImageContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


