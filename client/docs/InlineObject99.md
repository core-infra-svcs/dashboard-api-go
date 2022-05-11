# InlineObject99

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]NetworksNetworkIdSwitchMtuOverrides**](NetworksNetworkIdSwitchMtuOverrides.md) | Override MTU size for individual switches or switch profiles. An empty array will clear overrides. | [optional] 

## Methods

### NewInlineObject99

`func NewInlineObject99() *InlineObject99`

NewInlineObject99 instantiates a new InlineObject99 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject99WithDefaults

`func NewInlineObject99WithDefaults() *InlineObject99`

NewInlineObject99WithDefaults instantiates a new InlineObject99 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *InlineObject99) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *InlineObject99) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *InlineObject99) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *InlineObject99) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineObject99) GetOverrides() []NetworksNetworkIdSwitchMtuOverrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineObject99) GetOverridesOk() (*[]NetworksNetworkIdSwitchMtuOverrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineObject99) SetOverrides(v []NetworksNetworkIdSwitchMtuOverrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineObject99) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


