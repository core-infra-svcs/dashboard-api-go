# InlineObject103

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your floor plan. | 
**Center** | Pointer to [**NetworksNetworkIdFloorPlansCenter**](NetworksNetworkIdFloorPlansCenter.md) |  | [optional] 
**BottomLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomLeftCorner**](NetworksNetworkIdFloorPlansBottomLeftCorner.md) |  | [optional] 
**BottomRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansBottomRightCorner**](NetworksNetworkIdFloorPlansBottomRightCorner.md) |  | [optional] 
**TopLeftCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopLeftCorner**](NetworksNetworkIdFloorPlansTopLeftCorner.md) |  | [optional] 
**TopRightCorner** | Pointer to [**NetworksNetworkIdFloorPlansTopRightCorner**](NetworksNetworkIdFloorPlansTopRightCorner.md) |  | [optional] 
**FloorNumber** | Pointer to **float32** | The floor number of the floors within the building | [optional] 
**ImageContents** | **string** | The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. | 

## Methods

### NewInlineObject103

`func NewInlineObject103(name string, imageContents string, ) *InlineObject103`

NewInlineObject103 instantiates a new InlineObject103 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject103WithDefaults

`func NewInlineObject103WithDefaults() *InlineObject103`

NewInlineObject103WithDefaults instantiates a new InlineObject103 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject103) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject103) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject103) SetName(v string)`

SetName sets Name field to given value.


### GetCenter

`func (o *InlineObject103) GetCenter() NetworksNetworkIdFloorPlansCenter`

GetCenter returns the Center field if non-nil, zero value otherwise.

### GetCenterOk

`func (o *InlineObject103) GetCenterOk() (*NetworksNetworkIdFloorPlansCenter, bool)`

GetCenterOk returns a tuple with the Center field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenter

`func (o *InlineObject103) SetCenter(v NetworksNetworkIdFloorPlansCenter)`

SetCenter sets Center field to given value.

### HasCenter

`func (o *InlineObject103) HasCenter() bool`

HasCenter returns a boolean if a field has been set.

### GetBottomLeftCorner

`func (o *InlineObject103) GetBottomLeftCorner() NetworksNetworkIdFloorPlansBottomLeftCorner`

GetBottomLeftCorner returns the BottomLeftCorner field if non-nil, zero value otherwise.

### GetBottomLeftCornerOk

`func (o *InlineObject103) GetBottomLeftCornerOk() (*NetworksNetworkIdFloorPlansBottomLeftCorner, bool)`

GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomLeftCorner

`func (o *InlineObject103) SetBottomLeftCorner(v NetworksNetworkIdFloorPlansBottomLeftCorner)`

SetBottomLeftCorner sets BottomLeftCorner field to given value.

### HasBottomLeftCorner

`func (o *InlineObject103) HasBottomLeftCorner() bool`

HasBottomLeftCorner returns a boolean if a field has been set.

### GetBottomRightCorner

`func (o *InlineObject103) GetBottomRightCorner() NetworksNetworkIdFloorPlansBottomRightCorner`

GetBottomRightCorner returns the BottomRightCorner field if non-nil, zero value otherwise.

### GetBottomRightCornerOk

`func (o *InlineObject103) GetBottomRightCornerOk() (*NetworksNetworkIdFloorPlansBottomRightCorner, bool)`

GetBottomRightCornerOk returns a tuple with the BottomRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomRightCorner

`func (o *InlineObject103) SetBottomRightCorner(v NetworksNetworkIdFloorPlansBottomRightCorner)`

SetBottomRightCorner sets BottomRightCorner field to given value.

### HasBottomRightCorner

`func (o *InlineObject103) HasBottomRightCorner() bool`

HasBottomRightCorner returns a boolean if a field has been set.

### GetTopLeftCorner

`func (o *InlineObject103) GetTopLeftCorner() NetworksNetworkIdFloorPlansTopLeftCorner`

GetTopLeftCorner returns the TopLeftCorner field if non-nil, zero value otherwise.

### GetTopLeftCornerOk

`func (o *InlineObject103) GetTopLeftCornerOk() (*NetworksNetworkIdFloorPlansTopLeftCorner, bool)`

GetTopLeftCornerOk returns a tuple with the TopLeftCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLeftCorner

`func (o *InlineObject103) SetTopLeftCorner(v NetworksNetworkIdFloorPlansTopLeftCorner)`

SetTopLeftCorner sets TopLeftCorner field to given value.

### HasTopLeftCorner

`func (o *InlineObject103) HasTopLeftCorner() bool`

HasTopLeftCorner returns a boolean if a field has been set.

### GetTopRightCorner

`func (o *InlineObject103) GetTopRightCorner() NetworksNetworkIdFloorPlansTopRightCorner`

GetTopRightCorner returns the TopRightCorner field if non-nil, zero value otherwise.

### GetTopRightCornerOk

`func (o *InlineObject103) GetTopRightCornerOk() (*NetworksNetworkIdFloorPlansTopRightCorner, bool)`

GetTopRightCornerOk returns a tuple with the TopRightCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRightCorner

`func (o *InlineObject103) SetTopRightCorner(v NetworksNetworkIdFloorPlansTopRightCorner)`

SetTopRightCorner sets TopRightCorner field to given value.

### HasTopRightCorner

`func (o *InlineObject103) HasTopRightCorner() bool`

HasTopRightCorner returns a boolean if a field has been set.

### GetFloorNumber

`func (o *InlineObject103) GetFloorNumber() float32`

GetFloorNumber returns the FloorNumber field if non-nil, zero value otherwise.

### GetFloorNumberOk

`func (o *InlineObject103) GetFloorNumberOk() (*float32, bool)`

GetFloorNumberOk returns a tuple with the FloorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorNumber

`func (o *InlineObject103) SetFloorNumber(v float32)`

SetFloorNumber sets FloorNumber field to given value.

### HasFloorNumber

`func (o *InlineObject103) HasFloorNumber() bool`

HasFloorNumber returns a boolean if a field has been set.

### GetImageContents

`func (o *InlineObject103) GetImageContents() string`

GetImageContents returns the ImageContents field if non-nil, zero value otherwise.

### GetImageContentsOk

`func (o *InlineObject103) GetImageContentsOk() (*string, bool)`

GetImageContentsOk returns a tuple with the ImageContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageContents

`func (o *InlineObject103) SetImageContents(v string)`

SetImageContents sets ImageContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


