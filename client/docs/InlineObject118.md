# InlineObject118

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | **[]string** | A list of serials of devices to claim | 
**DetailsByDevice** | Pointer to [**[]NetworksNetworkIdDevicesClaimDetailsByDevice**](NetworksNetworkIdDevicesClaimDetailsByDevice.md) | Optional details for claimed devices (currently only used for Catalyst devices) | [optional] 

## Methods

### NewInlineObject118

`func NewInlineObject118(serials []string, ) *InlineObject118`

NewInlineObject118 instantiates a new InlineObject118 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject118WithDefaults

`func NewInlineObject118WithDefaults() *InlineObject118`

NewInlineObject118WithDefaults instantiates a new InlineObject118 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *InlineObject118) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject118) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject118) SetSerials(v []string)`

SetSerials sets Serials field to given value.


### GetDetailsByDevice

`func (o *InlineObject118) GetDetailsByDevice() []NetworksNetworkIdDevicesClaimDetailsByDevice`

GetDetailsByDevice returns the DetailsByDevice field if non-nil, zero value otherwise.

### GetDetailsByDeviceOk

`func (o *InlineObject118) GetDetailsByDeviceOk() (*[]NetworksNetworkIdDevicesClaimDetailsByDevice, bool)`

GetDetailsByDeviceOk returns a tuple with the DetailsByDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsByDevice

`func (o *InlineObject118) SetDetailsByDevice(v []NetworksNetworkIdDevicesClaimDetailsByDevice)`

SetDetailsByDevice sets DetailsByDevice field to given value.

### HasDetailsByDevice

`func (o *InlineObject118) HasDetailsByDevice() bool`

HasDetailsByDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


