# InlineResponse20079

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the client | [optional] 
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Name of the client | [optional] 
**DeviceName** | Pointer to **string** | Bluetooth device name | [optional] 
**Manufacturer** | Pointer to **string** | Name of the manufacturer | [optional] 
**LastSeen** | Pointer to **int32** | Epoch timestamp of the device&#39;s last appearance | [optional] 
**SeenByDeviceMac** | Pointer to **string** | Seen by device MAC | [optional] 
**InSightAlert** | Pointer to **bool** | Device in sight alert | [optional] 
**OutOfSightAlert** | Pointer to **bool** | Device out of sight alert | [optional] 
**Tags** | Pointer to **[]string** | A list of tags applied to the device | [optional] 

## Methods

### NewInlineResponse20079

`func NewInlineResponse20079() *InlineResponse20079`

NewInlineResponse20079 instantiates a new InlineResponse20079 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079WithDefaults

`func NewInlineResponse20079WithDefaults() *InlineResponse20079`

NewInlineResponse20079WithDefaults instantiates a new InlineResponse20079 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20079) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20079) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20079) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20079) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse20079) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20079) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20079) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20079) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20079) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20079) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20079) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20079) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20079) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20079) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20079) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20079) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceName

`func (o *InlineResponse20079) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InlineResponse20079) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InlineResponse20079) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InlineResponse20079) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineResponse20079) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineResponse20079) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineResponse20079) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineResponse20079) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse20079) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse20079) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse20079) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse20079) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetSeenByDeviceMac

`func (o *InlineResponse20079) GetSeenByDeviceMac() string`

GetSeenByDeviceMac returns the SeenByDeviceMac field if non-nil, zero value otherwise.

### GetSeenByDeviceMacOk

`func (o *InlineResponse20079) GetSeenByDeviceMacOk() (*string, bool)`

GetSeenByDeviceMacOk returns a tuple with the SeenByDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenByDeviceMac

`func (o *InlineResponse20079) SetSeenByDeviceMac(v string)`

SetSeenByDeviceMac sets SeenByDeviceMac field to given value.

### HasSeenByDeviceMac

`func (o *InlineResponse20079) HasSeenByDeviceMac() bool`

HasSeenByDeviceMac returns a boolean if a field has been set.

### GetInSightAlert

`func (o *InlineResponse20079) GetInSightAlert() bool`

GetInSightAlert returns the InSightAlert field if non-nil, zero value otherwise.

### GetInSightAlertOk

`func (o *InlineResponse20079) GetInSightAlertOk() (*bool, bool)`

GetInSightAlertOk returns a tuple with the InSightAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSightAlert

`func (o *InlineResponse20079) SetInSightAlert(v bool)`

SetInSightAlert sets InSightAlert field to given value.

### HasInSightAlert

`func (o *InlineResponse20079) HasInSightAlert() bool`

HasInSightAlert returns a boolean if a field has been set.

### GetOutOfSightAlert

`func (o *InlineResponse20079) GetOutOfSightAlert() bool`

GetOutOfSightAlert returns the OutOfSightAlert field if non-nil, zero value otherwise.

### GetOutOfSightAlertOk

`func (o *InlineResponse20079) GetOutOfSightAlertOk() (*bool, bool)`

GetOutOfSightAlertOk returns a tuple with the OutOfSightAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSightAlert

`func (o *InlineResponse20079) SetOutOfSightAlert(v bool)`

SetOutOfSightAlert sets OutOfSightAlert field to given value.

### HasOutOfSightAlert

`func (o *InlineResponse20079) HasOutOfSightAlert() bool`

HasOutOfSightAlert returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20079) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20079) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20079) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20079) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


