# InlineResponse200429

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200429Items**](InlineResponse200429Items.md) | List of available network enforcement targets according to the provided filters. Maximum 100 items per page. | [optional] 
**Meta** | Pointer to [**InlineResponse200283Meta**](InlineResponse200283Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200429

`func NewInlineResponse200429() *InlineResponse200429`

NewInlineResponse200429 instantiates a new InlineResponse200429 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200429WithDefaults

`func NewInlineResponse200429WithDefaults() *InlineResponse200429`

NewInlineResponse200429WithDefaults instantiates a new InlineResponse200429 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200429) GetItems() []InlineResponse200429Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200429) GetItemsOk() (*[]InlineResponse200429Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200429) SetItems(v []InlineResponse200429Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200429) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200429) GetMeta() InlineResponse200283Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200429) GetMetaOk() (*InlineResponse200283Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200429) SetMeta(v InlineResponse200283Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200429) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


