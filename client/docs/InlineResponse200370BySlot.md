# InlineResponse200370BySlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot identifier | [optional] 
**BySignalType** | Pointer to [**[]InlineResponse200370BySignalType**](InlineResponse200370BySignalType.md) | Signal settings for the SIM. Maximum 3 signal types per SIM slot. | [optional] 

## Methods

### NewInlineResponse200370BySlot

`func NewInlineResponse200370BySlot() *InlineResponse200370BySlot`

NewInlineResponse200370BySlot instantiates a new InlineResponse200370BySlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200370BySlotWithDefaults

`func NewInlineResponse200370BySlotWithDefaults() *InlineResponse200370BySlot`

NewInlineResponse200370BySlotWithDefaults instantiates a new InlineResponse200370BySlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200370BySlot) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200370BySlot) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200370BySlot) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200370BySlot) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetBySignalType

`func (o *InlineResponse200370BySlot) GetBySignalType() []InlineResponse200370BySignalType`

GetBySignalType returns the BySignalType field if non-nil, zero value otherwise.

### GetBySignalTypeOk

`func (o *InlineResponse200370BySlot) GetBySignalTypeOk() (*[]InlineResponse200370BySignalType, bool)`

GetBySignalTypeOk returns a tuple with the BySignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySignalType

`func (o *InlineResponse200370BySlot) SetBySignalType(v []InlineResponse200370BySignalType)`

SetBySignalType sets BySignalType field to given value.

### HasBySignalType

`func (o *InlineResponse200370BySlot) HasBySignalType() bool`

HasBySignalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


