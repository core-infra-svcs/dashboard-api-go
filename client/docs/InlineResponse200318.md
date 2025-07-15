# InlineResponse200318

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200317Items**](InlineResponse200317Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200222Meta**](InlineResponse200222Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200318

`func NewInlineResponse200318() *InlineResponse200318`

NewInlineResponse200318 instantiates a new InlineResponse200318 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200318WithDefaults

`func NewInlineResponse200318WithDefaults() *InlineResponse200318`

NewInlineResponse200318WithDefaults instantiates a new InlineResponse200318 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200318) GetItems() []InlineResponse200317Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200318) GetItemsOk() (*[]InlineResponse200317Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200318) SetItems(v []InlineResponse200317Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200318) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200318) GetMeta() InlineResponse200222Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200318) GetMetaOk() (*InlineResponse200222Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200318) SetMeta(v InlineResponse200222Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200318) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


