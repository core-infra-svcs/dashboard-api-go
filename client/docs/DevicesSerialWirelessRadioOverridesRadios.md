# DevicesSerialWirelessRadioOverridesRadios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable radio. Setting to false will set tx power to -1. | [optional] 
**Index** | **string** | Indicates which radio to override | 
**Channel** | Pointer to **NullableInt32** | Sets a radio override channel. Can be null to use auto channel. | [optional] 
**TargetPower** | Pointer to **NullableInt32** | Set a radio override target power (dBm). Can be null to use auto power range. | [optional] 
**ChannelWidth** | Pointer to **NullableInt32** | Set a radio override channel width. Can be null to use auto channel width. | [optional] 

## Methods

### NewDevicesSerialWirelessRadioOverridesRadios

`func NewDevicesSerialWirelessRadioOverridesRadios(index string, ) *DevicesSerialWirelessRadioOverridesRadios`

NewDevicesSerialWirelessRadioOverridesRadios instantiates a new DevicesSerialWirelessRadioOverridesRadios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialWirelessRadioOverridesRadiosWithDefaults

`func NewDevicesSerialWirelessRadioOverridesRadiosWithDefaults() *DevicesSerialWirelessRadioOverridesRadios`

NewDevicesSerialWirelessRadioOverridesRadiosWithDefaults instantiates a new DevicesSerialWirelessRadioOverridesRadios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DevicesSerialWirelessRadioOverridesRadios) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIndex

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetChannel

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DevicesSerialWirelessRadioOverridesRadios) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *DevicesSerialWirelessRadioOverridesRadios) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetTargetPower

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *DevicesSerialWirelessRadioOverridesRadios) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.

### SetTargetPowerNil

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetTargetPowerNil(b bool)`

 SetTargetPowerNil sets the value for TargetPower to be an explicit nil

### UnsetTargetPower
`func (o *DevicesSerialWirelessRadioOverridesRadios) UnsetTargetPower()`

UnsetTargetPower ensures that no value is present for TargetPower, not even an explicit nil
### GetChannelWidth

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetChannelWidth() int32`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *DevicesSerialWirelessRadioOverridesRadios) GetChannelWidthOk() (*int32, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetChannelWidth(v int32)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *DevicesSerialWirelessRadioOverridesRadios) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### SetChannelWidthNil

`func (o *DevicesSerialWirelessRadioOverridesRadios) SetChannelWidthNil(b bool)`

 SetChannelWidthNil sets the value for ChannelWidth to be an explicit nil

### UnsetChannelWidth
`func (o *DevicesSerialWirelessRadioOverridesRadios) UnsetChannelWidth()`

UnsetChannelWidth ensures that no value is present for ChannelWidth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


