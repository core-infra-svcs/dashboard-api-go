# InlineResponse200402DevicesSpare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Device ID | [optional] 
**History** | Pointer to [**[]InlineResponse200402History**](InlineResponse200402History.md) | Connectivity history for the device | [optional] 

## Methods

### NewInlineResponse200402DevicesSpare

`func NewInlineResponse200402DevicesSpare() *InlineResponse200402DevicesSpare`

NewInlineResponse200402DevicesSpare instantiates a new InlineResponse200402DevicesSpare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200402DevicesSpareWithDefaults

`func NewInlineResponse200402DevicesSpareWithDefaults() *InlineResponse200402DevicesSpare`

NewInlineResponse200402DevicesSpareWithDefaults instantiates a new InlineResponse200402DevicesSpare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200402DevicesSpare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200402DevicesSpare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200402DevicesSpare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200402DevicesSpare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200402DevicesSpare) GetHistory() []InlineResponse200402History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200402DevicesSpare) GetHistoryOk() (*[]InlineResponse200402History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200402DevicesSpare) SetHistory(v []InlineResponse200402History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200402DevicesSpare) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


