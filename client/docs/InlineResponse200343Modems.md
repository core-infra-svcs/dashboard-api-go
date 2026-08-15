# InlineResponse200343Modems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | 0-based modem index | [optional] 
**Sims** | Pointer to [**[]InlineResponse200343Sims**](InlineResponse200343Sims.md) | Information about the available SIMs for this modem. Maximum 3 SIM entries. | [optional] 

## Methods

### NewInlineResponse200343Modems

`func NewInlineResponse200343Modems() *InlineResponse200343Modems`

NewInlineResponse200343Modems instantiates a new InlineResponse200343Modems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200343ModemsWithDefaults

`func NewInlineResponse200343ModemsWithDefaults() *InlineResponse200343Modems`

NewInlineResponse200343ModemsWithDefaults instantiates a new InlineResponse200343Modems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *InlineResponse200343Modems) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *InlineResponse200343Modems) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *InlineResponse200343Modems) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *InlineResponse200343Modems) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSims

`func (o *InlineResponse200343Modems) GetSims() []InlineResponse200343Sims`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *InlineResponse200343Modems) GetSimsOk() (*[]InlineResponse200343Sims, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *InlineResponse200343Modems) SetSims(v []InlineResponse200343Sims)`

SetSims sets Sims field to given value.

### HasSims

`func (o *InlineResponse200343Modems) HasSims() bool`

HasSims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


