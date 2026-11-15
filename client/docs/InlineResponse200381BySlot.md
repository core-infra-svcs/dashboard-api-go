# InlineResponse200381BySlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot identifier | [optional] 
**BySignalType** | Pointer to [**[]InlineResponse200381BySignalType**](InlineResponse200381BySignalType.md) | Signal settings for the SIM. Maximum 3 signal types per SIM slot. | [optional] 

## Methods

### NewInlineResponse200381BySlot

`func NewInlineResponse200381BySlot() *InlineResponse200381BySlot`

NewInlineResponse200381BySlot instantiates a new InlineResponse200381BySlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200381BySlotWithDefaults

`func NewInlineResponse200381BySlotWithDefaults() *InlineResponse200381BySlot`

NewInlineResponse200381BySlotWithDefaults instantiates a new InlineResponse200381BySlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200381BySlot) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200381BySlot) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200381BySlot) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200381BySlot) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetBySignalType

`func (o *InlineResponse200381BySlot) GetBySignalType() []InlineResponse200381BySignalType`

GetBySignalType returns the BySignalType field if non-nil, zero value otherwise.

### GetBySignalTypeOk

`func (o *InlineResponse200381BySlot) GetBySignalTypeOk() (*[]InlineResponse200381BySignalType, bool)`

GetBySignalTypeOk returns a tuple with the BySignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySignalType

`func (o *InlineResponse200381BySlot) SetBySignalType(v []InlineResponse200381BySignalType)`

SetBySignalType sets BySignalType field to given value.

### HasBySignalType

`func (o *InlineResponse200381BySlot) HasBySignalType() bool`

HasBySignalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


