# InlineResponse20069

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | The number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**Neighbors** | Pointer to [**[]InlineResponse20069Neighbors**](InlineResponse20069Neighbors.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewInlineResponse20069

`func NewInlineResponse20069() *InlineResponse20069`

NewInlineResponse20069 instantiates a new InlineResponse20069 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20069WithDefaults

`func NewInlineResponse20069WithDefaults() *InlineResponse20069`

NewInlineResponse20069WithDefaults instantiates a new InlineResponse20069 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20069) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20069) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20069) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20069) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *InlineResponse20069) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *InlineResponse20069) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *InlineResponse20069) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *InlineResponse20069) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *InlineResponse20069) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *InlineResponse20069) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *InlineResponse20069) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *InlineResponse20069) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse20069) GetNeighbors() []InlineResponse20069Neighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse20069) GetNeighborsOk() (*[]InlineResponse20069Neighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse20069) SetNeighbors(v []InlineResponse20069Neighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse20069) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


