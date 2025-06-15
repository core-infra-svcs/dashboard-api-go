# InlineResponse200317

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200316Items**](InlineResponse200316Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200222Meta**](InlineResponse200222Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200317

`func NewInlineResponse200317() *InlineResponse200317`

NewInlineResponse200317 instantiates a new InlineResponse200317 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200317WithDefaults

`func NewInlineResponse200317WithDefaults() *InlineResponse200317`

NewInlineResponse200317WithDefaults instantiates a new InlineResponse200317 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200317) GetItems() []InlineResponse200316Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200317) GetItemsOk() (*[]InlineResponse200316Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200317) SetItems(v []InlineResponse200316Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200317) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200317) GetMeta() InlineResponse200222Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200317) GetMetaOk() (*InlineResponse200222Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200317) SetMeta(v InlineResponse200222Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200317) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


