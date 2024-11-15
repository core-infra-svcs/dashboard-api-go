# OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Device Name | [optional] 
**Serial** | Pointer to **string** | Device Serial Number | [optional] 
**Mac** | Pointer to **string** | MAC Address | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**Status** | Pointer to **string** | Device Status | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdFloorPlansAutoLocateDevicesNetwork**](OrganizationsOrganizationIdFloorPlansAutoLocateDevicesNetwork.md) |  | [optional] 
**FloorPlan** | Pointer to [**OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan**](OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan.md) |  | [optional] 
**Lat** | Pointer to **float32** | Latitude | [optional] 
**Lng** | Pointer to **float32** | Longitude | [optional] 
**AutoLocate** | Pointer to [**OrganizationsOrganizationIdFloorPlansAutoLocateDevicesAutoLocate**](OrganizationsOrganizationIdFloorPlansAutoLocateDevicesAutoLocate.md) |  | [optional] 
**Type** | Pointer to **string** | The type of auto locate position. Possible values: &#39;user&#39;, &#39;gnss&#39;, and &#39;calculated&#39; | [optional] 
**IsAnchor** | Pointer to **bool** | Whether or not this auto locate position is an anchor | [optional] 

## Methods

### NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems

`func NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems() *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems`

NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItemsWithDefaults

`func NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItemsWithDefaults() *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems`

NewOrganizationsOrganizationIdFloorPlansAutoLocateDevicesItemsWithDefaults instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetNetwork() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetNetworkOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetNetwork(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFloorPlan

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetFloorPlan() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan`

GetFloorPlan returns the FloorPlan field if non-nil, zero value otherwise.

### GetFloorPlanOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetFloorPlanOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan, bool)`

GetFloorPlanOk returns a tuple with the FloorPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlan

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetFloorPlan(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesFloorPlan)`

SetFloorPlan sets FloorPlan field to given value.

### HasFloorPlan

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasFloorPlan() bool`

HasFloorPlan returns a boolean if a field has been set.

### GetLat

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAutoLocate

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetAutoLocate() OrganizationsOrganizationIdFloorPlansAutoLocateDevicesAutoLocate`

GetAutoLocate returns the AutoLocate field if non-nil, zero value otherwise.

### GetAutoLocateOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetAutoLocateOk() (*OrganizationsOrganizationIdFloorPlansAutoLocateDevicesAutoLocate, bool)`

GetAutoLocateOk returns a tuple with the AutoLocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLocate

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetAutoLocate(v OrganizationsOrganizationIdFloorPlansAutoLocateDevicesAutoLocate)`

SetAutoLocate sets AutoLocate field to given value.

### HasAutoLocate

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasAutoLocate() bool`

HasAutoLocate returns a boolean if a field has been set.

### GetType

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsAnchor

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetIsAnchor() bool`

GetIsAnchor returns the IsAnchor field if non-nil, zero value otherwise.

### GetIsAnchorOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) GetIsAnchorOk() (*bool, bool)`

GetIsAnchorOk returns a tuple with the IsAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnchor

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) SetIsAnchor(v bool)`

SetIsAnchor sets IsAnchor field to given value.

### HasIsAnchor

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateDevicesItems) HasIsAnchor() bool`

HasIsAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


