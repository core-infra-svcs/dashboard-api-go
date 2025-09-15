# InlineObject106

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your floor plan. | 
**Center** | Pointer to [**NetworksNetworkIdFloorPlansCenter1**](NetworksNetworkIdFloorPlansCenter1.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomLeftCorner1**](NetworksNetworkIdFloorPlansBottomLeftCorner1.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomRightCorner1**](NetworksNetworkIdFloorPlansBottomRightCorner1.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopLeftCorner1**](NetworksNetworkIdFloorPlansTopLeftCorner1.md) |  | [optional] 
**TopRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopRightCorner1**](NetworksNetworkIdFloorPlansTopRightCorner1.md) |  | [optional] 
**FloorNumber** | Pointer to **float32** | The floor number of the floors within the building | [optional] 
**ImageContents** | **string** | The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. | 

## Methods

### NewInlineObject106

`func NewInlineObject106(name string, imageContents string, ) *InlineObject106`

NewInlineObject106 instantiates a new InlineObject106 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject106WithDefaults

`func NewInlineObject106WithDefaults() *InlineObject106`

NewInlineObject106WithDefaults instantiates a new InlineObject106 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject106) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject106) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject106) SetName(v string)`

SetName sets Name field to given value.


### GetCenter

`func (o *InlineObject106) GetCenter() NetworksNetworkIdFloorPlansCenter1`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *InlineObject106) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter1, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *InlineObject106) SetCenter(v NetworksNetworkIdFloorPlansCenter1)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *InlineObject106) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *InlineObject106) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner1`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *InlineObject106) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner1, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *InlineObject106) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner1)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *InlineObject106) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *InlineObject106) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner1`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *InlineObject106) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner1, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *InlineObject106) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner1)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *InlineObject106) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *InlineObject106) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner1`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *InlineObject106) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner1, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *InlineObject106) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner1)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *InlineObject106) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *InlineObject106) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner1`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *InlineObject106) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner1, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *InlineObject106) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner1)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *InlineObject106) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetFloorNumber

`func (o *InlineObject106) GetFloorNumber() float32`

GetFloorNumber returns the FloorNumber field if non-nil, zero value otherwise.

### GetFloorNumberOk

`func (o *InlineObject106) GetFloorNumberOk() (*float32, bool)`

GetFloorNumberOk returns a tuple with the FloorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorNumber

`func (o *InlineObject106) SetFloorNumber(v float32)`

SetFloorNumber sets FloorNumber field to given value.

### HasFloorNumber

`func (o *InlineObject106) HasFloorNumber() bool`

HasFloorNumber returns a boolean if a field has been set.

### GetImageContents

`func (o *InlineObject106) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *InlineObject106) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *InlineObject106) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


