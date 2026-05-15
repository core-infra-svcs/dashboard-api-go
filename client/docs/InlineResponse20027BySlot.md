# InlineResponse20027BySlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot identifier | [optional] 
**BySignalType** | Pointer to [**[]InlineResponse20027BySignalType**](InlineResponse20027BySignalType.md) | Signal settings for the SIM. Maximum 3 signal types per SIM slot. | [optional] 

## Methods

### NewInlineResponse20027BySlot

`func NewInlineResponse20027BySlot() *InlineResponse20027BySlot`

NewInlineResponse20027BySlot instantiates a new InlineResponse20027BySlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027BySlotWithDefaults

`func NewInlineResponse20027BySlotWithDefaults() *InlineResponse20027BySlot`

NewInlineResponse20027BySlotWithDefaults instantiates a new InlineResponse20027BySlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse20027BySlot) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse20027BySlot) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse20027BySlot) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse20027BySlot) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetBySignalType

`func (o *InlineResponse20027BySlot) GetBySignalType() []InlineResponse20027BySignalType`

GetBySignalType returns the BySignalType field if non-nil, zero value otherwise.

### GetBySignalTypeOk

`func (o *InlineResponse20027BySlot) GetBySignalTypeOk() (*[]InlineResponse20027BySignalType, bool)`

GetBySignalTypeOk returns a tuple with the BySignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySignalType

`func (o *InlineResponse20027BySlot) SetBySignalType(v []InlineResponse20027BySignalType)`

SetBySignalType sets BySignalType field to given value.

### HasBySignalType

`func (o *InlineResponse20027BySlot) HasBySignalType() bool`

HasBySignalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


