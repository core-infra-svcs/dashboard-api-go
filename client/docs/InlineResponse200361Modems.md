# InlineResponse200361Modems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | 0-based modem index | [optional] 
**Sims** | Pointer to [**[]InlineResponse200361Sims**](InlineResponse200361Sims.md) | Information about the available SIMs for this modem. Maximum 3 SIM entries. | [optional] 

## Methods

### NewInlineResponse200361Modems

`func NewInlineResponse200361Modems() *InlineResponse200361Modems`

NewInlineResponse200361Modems instantiates a new InlineResponse200361Modems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200361ModemsWithDefaults

`func NewInlineResponse200361ModemsWithDefaults() *InlineResponse200361Modems`

NewInlineResponse200361ModemsWithDefaults instantiates a new InlineResponse200361Modems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *InlineResponse200361Modems) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InlineResponse200361Modems) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InlineResponse200361Modems) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *InlineResponse200361Modems) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSims

`func (o *InlineResponse200361Modems) GetSims() []InlineResponse200361Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse200361Modems) GetSimsOk() (*[]InlineResponse200361Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse200361Modems) SetSims(v []InlineResponse200361Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse200361Modems) HasSims() bool`

HasSims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


