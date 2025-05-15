# InlineResponse20077

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the warm spare enabled | [optional] 
**PrimarySerial** | Pointer to **string** | Serial number of the primary appliance | [optional] 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare appliance | [optional] 
**UplinkMode** | Pointer to **string** | Uplink mode, either virtual or public | [optional] 
**Wan1** | Pointer to [**InlineResponse20077Wan1**](InlineResponse20077Wan1.md) |  | [optional] 
**Wan2** | Pointer to [**InlineResponse20077Wan2**](InlineResponse20077Wan2.md) |  | [optional] 

## Methods

### NewInlineResponse20077

`func NewInlineResponse20077() *InlineResponse20077`

NewInlineResponse20077 instantiates a new InlineResponse20077 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20077WithDefaults

`func NewInlineResponse20077WithDefaults() *InlineResponse20077`

NewInlineResponse20077WithDefaults instantiates a new InlineResponse20077 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20077) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20077) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20077) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20077) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrimarySerial

`func (o *InlineResponse20077) GetPrimarySerial() string`

GetPrimarySerial returns the PrimarySerial field if non-nil, zero value otherwise.

### GetPrimarySerialOk

`func (o *InlineResponse20077) GetPrimarySerialOk() (*string, bool)`

GetPrimarySerialOk returns a tuple with the PrimarySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySerial

`func (o *InlineResponse20077) SetPrimarySerial(v string)`

SetPrimarySerial sets PrimarySerial field to given value.

### HasPrimarySerial

`func (o *InlineResponse20077) HasPrimarySerial() bool`

HasPrimarySerial returns a boolean if a field has been set.

### GetSpareSerial

`func (o *InlineResponse20077) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *InlineResponse20077) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *InlineResponse20077) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *InlineResponse20077) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.

### GetUplinkMode

`func (o *InlineResponse20077) GetUplinkMode() string`

GetUplinkMode returns the UplinkMode field if non-nil, zero value otherwise.

### GetUplinkModeOk

`func (o *InlineResponse20077) GetUplinkModeOk() (*string, bool)`

GetUplinkModeOk returns a tuple with the UplinkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMode

`func (o *InlineResponse20077) SetUplinkMode(v string)`

SetUplinkMode sets UplinkMode field to given value.

### HasUplinkMode

`func (o *InlineResponse20077) HasUplinkMode() bool`

HasUplinkMode returns a boolean if a field has been set.

### GetWan1

`func (o *InlineResponse20077) GetWan1() InlineResponse20077Wan1`

GetWan1 returns the Wan1 field if non-nil, zero value otherwise.

### GetWan1Ok

`func (o *InlineResponse20077) GetWan1Ok() (*InlineResponse20077Wan1, bool)`

GetWan1Ok returns a tuple with the Wan1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan1

`func (o *InlineResponse20077) SetWan1(v InlineResponse20077Wan1)`

SetWan1 sets Wan1 field to given value.

### HasWan1

`func (o *InlineResponse20077) HasWan1() bool`

HasWan1 returns a boolean if a field has been set.

### GetWan2

`func (o *InlineResponse20077) GetWan2() InlineResponse20077Wan2`

GetWan2 returns the Wan2 field if non-nil, zero value otherwise.

### GetWan2Ok

`func (o *InlineResponse20077) GetWan2Ok() (*InlineResponse20077Wan2, bool)`

GetWan2Ok returns a tuple with the Wan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan2

`func (o *InlineResponse20077) SetWan2(v InlineResponse20077Wan2)`

SetWan2 sets Wan2 field to given value.

### HasWan2

`func (o *InlineResponse20077) HasWan2() bool`

HasWan2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


