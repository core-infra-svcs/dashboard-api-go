# InlineResponse200321

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200320Items**](InlineResponse200320Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200224Meta**](InlineResponse200224Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200321

`func NewInlineResponse200321() *InlineResponse200321`

NewInlineResponse200321 instantiates a new InlineResponse200321 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200321WithDefaults

`func NewInlineResponse200321WithDefaults() *InlineResponse200321`

NewInlineResponse200321WithDefaults instantiates a new InlineResponse200321 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200321) GetItems() []InlineResponse200320Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200321) GetItemsOk() (*[]InlineResponse200320Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200321) SetItems(v []InlineResponse200320Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200321) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200321) GetMeta() InlineResponse200224Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200321) GetMetaOk() (*InlineResponse200224Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200321) SetMeta(v InlineResponse200224Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200321) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


