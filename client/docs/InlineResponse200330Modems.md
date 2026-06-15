# InlineResponse200330Modems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | 0-based modem index | [optional] 
**Sims** | Pointer to [**[]InlineResponse200330Sims**](InlineResponse200330Sims.md) | Information about the available SIMs for this modem. Maximum 3 SIM entries. | [optional] 

## Methods

### NewInlineResponse200330Modems

`func NewInlineResponse200330Modems() *InlineResponse200330Modems`

NewInlineResponse200330Modems instantiates a new InlineResponse200330Modems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200330ModemsWithDefaults

`func NewInlineResponse200330ModemsWithDefaults() *InlineResponse200330Modems`

NewInlineResponse200330ModemsWithDefaults instantiates a new InlineResponse200330Modems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *InlineResponse200330Modems) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InlineResponse200330Modems) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InlineResponse200330Modems) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *InlineResponse200330Modems) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSims

`func (o *InlineResponse200330Modems) GetSims() []InlineResponse200330Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse200330Modems) GetSimsOk() (*[]InlineResponse200330Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse200330Modems) SetSims(v []InlineResponse200330Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse200330Modems) HasSims() bool`

HasSims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


