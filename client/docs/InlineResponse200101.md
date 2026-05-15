# InlineResponse200101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | Organization-specific number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**RouterId** | Pointer to **string** | The router ID of the appliance | [optional] 
**Neighbors** | Pointer to [**[]InlineResponse200101Neighbors**](InlineResponse200101Neighbors.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewInlineResponse200101

`func NewInlineResponse200101() *InlineResponse200101`

NewInlineResponse200101 instantiates a new InlineResponse200101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101WithDefaults

`func NewInlineResponse200101WithDefaults() *InlineResponse200101`

NewInlineResponse200101WithDefaults instantiates a new InlineResponse200101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200101) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200101) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200101) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200101) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *InlineResponse200101) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *InlineResponse200101) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *InlineResponse200101) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *InlineResponse200101) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *InlineResponse200101) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *InlineResponse200101) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *InlineResponse200101) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *InlineResponse200101) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetRouterId

`func (o *InlineResponse200101) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *InlineResponse200101) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *InlineResponse200101) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *InlineResponse200101) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetNeighbors

`func (o *InlineResponse200101) GetNeighbors() []InlineResponse200101Neighbors`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *InlineResponse200101) GetNeighborsOk() (*[]InlineResponse200101Neighbors, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *InlineResponse200101) SetNeighbors(v []InlineResponse200101Neighbors)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *InlineResponse200101) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


