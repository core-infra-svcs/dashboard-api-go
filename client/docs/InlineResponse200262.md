# InlineResponse200262

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200261Items**](InlineResponse200261Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200259Meta**](InlineResponse200259Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200262

`func NewInlineResponse200262() *InlineResponse200262`

NewInlineResponse200262 instantiates a new InlineResponse200262 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200262WithDefaults

`func NewInlineResponse200262WithDefaults() *InlineResponse200262`

NewInlineResponse200262WithDefaults instantiates a new InlineResponse200262 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200262) GetItems() []InlineResponse200261Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200262) GetItemsOk() (*[]InlineResponse200261Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200262) SetItems(v []InlineResponse200261Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200262) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200262) GetMeta() InlineResponse200259Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200262) GetMetaOk() (*InlineResponse200259Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200262) SetMeta(v InlineResponse200259Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200262) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


