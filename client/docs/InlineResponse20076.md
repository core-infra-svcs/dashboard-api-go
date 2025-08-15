# InlineResponse20076

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | The number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**Neighbors** | Pointer to [**[]InlineResponse20076Neighbors**](InlineResponse20076Neighbors.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewInlineResponse20076

`func NewInlineResponse20076() *InlineResponse20076`

NewInlineResponse20076 instantiates a new InlineResponse20076 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20076WithDefaults

`func NewInlineResponse20076WithDefaults() *InlineResponse20076`

NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse20076) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20076) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20076) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20076) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *InlineResponse20076) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *InlineResponse20076) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *InlineResponse20076) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *InlineResponse20076) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *InlineResponse20076) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *InlineResponse20076) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *InlineResponse20076) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *InlineResponse20076) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse20076) GetNeighbors() []InlineResponse20076Neighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse20076) GetNeighborsOk() (*[]InlineResponse20076Neighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse20076) SetNeighbors(v []InlineResponse20076Neighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse20076) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


