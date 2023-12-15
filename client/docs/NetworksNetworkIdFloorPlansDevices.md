# NetworksNetworkIdFloorPlansDevices

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
**Imei** | Pointer to **string** | IMEI of the device, if applicable | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**LanIp** | Pointer to **string** | LAN IP address of the device | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Details** | Pointer to [**[]NetworksNetworkIdFloorPlansDetails**](NetworksNetworkIdFloorPlansDetails.md) | Additional device information | [optional] 

## Methods

### NewNetworksNetworkIdFloorPlansDevices

`func NewNetworksNetworkIdFloorPlansDevices() *NetworksNetworkIdFloorPlansDevices`

NewNetworksNetworkIdFloorPlansDevices instantiates a new NetworksNetworkIdFloorPlansDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdFloorPlansDevicesWithDefaults

`func NewNetworksNetworkIdFloorPlansDevicesWithDefaults() *NetworksNetworkIdFloorPlansDevices`

NewNetworksNetworkIdFloorPlansDevicesWithDefaults instantiates a new NetworksNetworkIdFloorPlansDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworksNetworkIdFloorPlansDevices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworksNetworkIdFloorPlansDevices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworksNetworkIdFloorPlansDevices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *NetworksNetworkIdFloorPlansDevices) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *NetworksNetworkIdFloorPlansDevices) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *NetworksNetworkIdFloorPlansDevices) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *NetworksNetworkIdFloorPlansDevices) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *NetworksNetworkIdFloorPlansDevices) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *NetworksNetworkIdFloorPlansDevices) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *NetworksNetworkIdFloorPlansDevices) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworksNetworkIdFloorPlansDevices) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NetworksNetworkIdFloorPlansDevices) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *NetworksNetworkIdFloorPlansDevices) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *NetworksNetworkIdFloorPlansDevices) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *NetworksNetworkIdFloorPlansDevices) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *NetworksNetworkIdFloorPlansDevices) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworksNetworkIdFloorPlansDevices) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworksNetworkIdFloorPlansDevices) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *NetworksNetworkIdFloorPlansDevices) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *NetworksNetworkIdFloorPlansDevices) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *NetworksNetworkIdFloorPlansDevices) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *NetworksNetworkIdFloorPlansDevices) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NetworksNetworkIdFloorPlansDevices) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NetworksNetworkIdFloorPlansDevices) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *NetworksNetworkIdFloorPlansDevices) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NetworksNetworkIdFloorPlansDevices) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NetworksNetworkIdFloorPlansDevices) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetImei

`func (o *NetworksNetworkIdFloorPlansDevices) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *NetworksNetworkIdFloorPlansDevices) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *NetworksNetworkIdFloorPlansDevices) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetMac

`func (o *NetworksNetworkIdFloorPlansDevices) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NetworksNetworkIdFloorPlansDevices) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NetworksNetworkIdFloorPlansDevices) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *NetworksNetworkIdFloorPlansDevices) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *NetworksNetworkIdFloorPlansDevices) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *NetworksNetworkIdFloorPlansDevices) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *NetworksNetworkIdFloorPlansDevices) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *NetworksNetworkIdFloorPlansDevices) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *NetworksNetworkIdFloorPlansDevices) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetProductType

`func (o *NetworksNetworkIdFloorPlansDevices) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *NetworksNetworkIdFloorPlansDevices) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *NetworksNetworkIdFloorPlansDevices) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDetails

`func (o *NetworksNetworkIdFloorPlansDevices) GetDetails() []NetworksNetworkIdFloorPlansDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NetworksNetworkIdFloorPlansDevices) GetDetailsOk() (*[]NetworksNetworkIdFloorPlansDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NetworksNetworkIdFloorPlansDevices) SetDetails(v []NetworksNetworkIdFloorPlansDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NetworksNetworkIdFloorPlansDevices) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


