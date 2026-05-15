# InlineResponse200327BySlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot identifier | [optional] 
**BySignalType** | Pointer to [**[]InlineResponse200327BySignalType**](InlineResponse200327BySignalType.md) | Signal settings for the SIM. Maximum 3 signal types per SIM slot. | [optional] 

## Methods

### NewInlineResponse200327BySlot

`func NewInlineResponse200327BySlot() *InlineResponse200327BySlot`

NewInlineResponse200327BySlot instantiates a new InlineResponse200327BySlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327BySlotWithDefaults

`func NewInlineResponse200327BySlotWithDefaults() *InlineResponse200327BySlot`

NewInlineResponse200327BySlotWithDefaults instantiates a new InlineResponse200327BySlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200327BySlot) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200327BySlot) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200327BySlot) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200327BySlot) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetBySignalType

`func (o *InlineResponse200327BySlot) GetBySignalType() []InlineResponse200327BySignalType`

GetBySignalType returns the BySignalType field if non-nil, zero value otherwise.

### GetBySignalTypeOk

`func (o *InlineResponse200327BySlot) GetBySignalTypeOk() (*[]InlineResponse200327BySignalType, bool)`

GetBySignalTypeOk returns a tuple with the BySignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySignalType

`func (o *InlineResponse200327BySlot) SetBySignalType(v []InlineResponse200327BySignalType)`

SetBySignalType sets BySignalType field to given value.

### HasBySignalType

`func (o *InlineResponse200327BySlot) HasBySignalType() bool`

HasBySignalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


