# DevicesSerialSwitchPortsFastPoe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether Fast PoE is enabled. Fast PoE requires Perpetual PoE to be enabled. This field is returned for all physical switch ports, but can only be enabled on ports that support PoE HA. | [optional] 

## Methods

### NewDevicesSerialSwitchPortsFastPoe

`func NewDevicesSerialSwitchPortsFastPoe() *DevicesSerialSwitchPortsFastPoe`

NewDevicesSerialSwitchPortsFastPoe instantiates a new DevicesSerialSwitchPortsFastPoe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchPortsFastPoeWithDefaults

`func NewDevicesSerialSwitchPortsFastPoeWithDefaults() *DevicesSerialSwitchPortsFastPoe`

NewDevicesSerialSwitchPortsFastPoeWithDefaults instantiates a new DevicesSerialSwitchPortsFastPoe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DevicesSerialSwitchPortsFastPoe) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DevicesSerialSwitchPortsFastPoe) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DevicesSerialSwitchPortsFastPoe) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DevicesSerialSwitchPortsFastPoe) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


