# InlineResponse2006

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the device | [optional] 
**Lat** | Pointer to **float32** | Latitude of the device | [optional] 
**Lng** | Pointer to **float32** | Longitude of the device | [optional] 
**Address** | Pointer to **string** | Physical address of the device | [optional] 
**Notes** | Pointer to **string** | Notes for the device, limited to 255 characters | [optional] 
**Tags** | Pointer to **[]string** | List of tags assigned to the device | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the device belongs to | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**LanIp** | Pointer to **string** | LAN IP address of the device | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the device | [optional] 
**FloorPlanId** | Pointer to **string** | The floor plan to associate to this device. null disassociates the device from the floorplan. | [optional] 
**Details** | Pointer to [**[]InlineResponse2006Details**](InlineResponse2006Details.md) | Additional device information | [optional] 
**BeaconIdParams** | Pointer to [**InlineResponse2006BeaconIdParams**](InlineResponse2006BeaconIdParams.md) |  | [optional] 

## Methods

### NewInlineResponse2006

`func NewInlineResponse2006() *InlineResponse2006`

NewInlineResponse2006 instantiates a new InlineResponse2006 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2006WithDefaults

`func NewInlineResponse2006WithDefaults() *InlineResponse2006`

NewInlineResponse2006WithDefaults instantiates a new InlineResponse2006 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2006) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2006) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2006) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2006) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *InlineResponse2006) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *InlineResponse2006) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *InlineResponse2006) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *InlineResponse2006) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *InlineResponse2006) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *InlineResponse2006) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *InlineResponse2006) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *InlineResponse2006) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse2006) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse2006) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse2006) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse2006) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse2006) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse2006) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse2006) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse2006) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse2006) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse2006) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse2006) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse2006) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse2006) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse2006) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse2006) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse2006) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse2006) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse2006) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse2006) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse2006) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse2006) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse2006) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse2006) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse2006) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse2006) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse2006) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse2006) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse2006) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *InlineResponse2006) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *InlineResponse2006) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *InlineResponse2006) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *InlineResponse2006) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *InlineResponse2006) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *InlineResponse2006) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *InlineResponse2006) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *InlineResponse2006) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineResponse2006) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse2006) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse2006) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse2006) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse2006) GetDetails() []InlineResponse2006Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse2006) GetDetailsOk() (*[]InlineResponse2006Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse2006) SetDetails(v []InlineResponse2006Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse2006) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetBeaconIdParams

`func (o *InlineResponse2006) GetBeaconIdParams() InlineResponse2006BeaconIdParams`

GetBeaconIdParams returns the BeaconIdParams field if non-nil, zero value otherwise.

### GetBeaconIdParamsOk

`func (o *InlineResponse2006) GetBeaconIdParamsOk() (*InlineResponse2006BeaconIdParams, bool)`

GetBeaconIdParamsOk returns a tuple with the BeaconIdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconIdParams

`func (o *InlineResponse2006) SetBeaconIdParams(v InlineResponse2006BeaconIdParams)`

SetBeaconIdParams sets BeaconIdParams field to given value.

### HasBeaconIdParams

`func (o *InlineResponse2006) HasBeaconIdParams() bool`

HasBeaconIdParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


