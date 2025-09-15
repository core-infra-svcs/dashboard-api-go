# InlineResponse20080

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the warm spare enabled | [optional] 
**PrimarySerial** | Pointer to **string** | Serial number of the primary appliance | [optional] 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare appliance | [optional] 
**UplinkMode** | Pointer to **string** | Uplink mode, either virtual or public | [optional] 
**Wan1** | Pointer to [**InlineResponse20080Wan1**](InlineResponse20080Wan1.md) |  | [optional] 
**Wan2** | Pointer to [**InlineResponse20080Wan2**](InlineResponse20080Wan2.md) |  | [optional] 

## Methods

### NewInlineResponse20080

`func NewInlineResponse20080() *InlineResponse20080`

NewInlineResponse20080 instantiates a new InlineResponse20080 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080WithDefaults

`func NewInlineResponse20080WithDefaults() *InlineResponse20080`

NewInlineResponse20080WithDefaults instantiates a new InlineResponse20080 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20080) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20080) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20080) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20080) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrimarySerial

`func (o *InlineResponse20080) GetPrimarySerial() string`

GetPrimarySerial returns the PrimarySerial field if non-nil, zero value otherwise.

### GetPrimarySerialOk

`func (o *InlineResponse20080) GetPrimarySerialOk() (*string, bool)`

GetPrimarySerialOk returns a tuple with the PrimarySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySerial

`func (o *InlineResponse20080) SetPrimarySerial(v string)`

SetPrimarySerial sets PrimarySerial field to given value.

### HasPrimarySerial

`func (o *InlineResponse20080) HasPrimarySerial() bool`

HasPrimarySerial returns a boolean if a field has been set.

### GetSpareSerial

`func (o *InlineResponse20080) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *InlineResponse20080) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *InlineResponse20080) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *InlineResponse20080) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.

### GetUplinkMode

`func (o *InlineResponse20080) GetUplinkMode() string`

GetUplinkMode returns the UplinkMode field if non-nil, zero value otherwise.

### GetUplinkModeOk

`func (o *InlineResponse20080) GetUplinkModeOk() (*string, bool)`

GetUplinkModeOk returns a tuple with the UplinkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMode

`func (o *InlineResponse20080) SetUplinkMode(v string)`

SetUplinkMode sets UplinkMode field to given value.

### HasUplinkMode

`func (o *InlineResponse20080) HasUplinkMode() bool`

HasUplinkMode returns a boolean if a field has been set.

### GetWan1

`func (o *InlineResponse20080) GetWan1() InlineResponse20080Wan1`

GetWan1 returns the Wan1 field if non-nil, zero value otherwise.

### GetWan1Ok

`func (o *InlineResponse20080) GetWan1Ok() (*InlineResponse20080Wan1, bool)`

GetWan1Ok returns a tuple with the Wan1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan1

`func (o *InlineResponse20080) SetWan1(v InlineResponse20080Wan1)`

SetWan1 sets Wan1 field to given value.

### HasWan1

`func (o *InlineResponse20080) HasWan1() bool`

HasWan1 returns a boolean if a field has been set.

### GetWan2

`func (o *InlineResponse20080) GetWan2() InlineResponse20080Wan2`

GetWan2 returns the Wan2 field if non-nil, zero value otherwise.

### GetWan2Ok

`func (o *InlineResponse20080) GetWan2Ok() (*InlineResponse20080Wan2, bool)`

GetWan2Ok returns a tuple with the Wan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan2

`func (o *InlineResponse20080) SetWan2(v InlineResponse20080Wan2)`

SetWan2 sets Wan2 field to given value.

### HasWan2

`func (o *InlineResponse20080) HasWan2() bool`

HasWan2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


