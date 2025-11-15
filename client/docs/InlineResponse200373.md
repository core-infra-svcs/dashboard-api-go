# InlineResponse200373

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200373Items**](InlineResponse200373Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200227Meta**](InlineResponse200227Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200373

`func NewInlineResponse200373() *InlineResponse200373`

NewInlineResponse200373 instantiates a new InlineResponse200373 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200373WithDefaults

`func NewInlineResponse200373WithDefaults() *InlineResponse200373`

NewInlineResponse200373WithDefaults instantiates a new InlineResponse200373 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200373) GetItems() []InlineResponse200373Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200373) GetItemsOk() (*[]InlineResponse200373Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200373) SetItems(v []InlineResponse200373Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200373) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200373) GetMeta() InlineResponse200227Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200373) GetMetaOk() (*InlineResponse200227Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200373) SetMeta(v InlineResponse200227Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200373) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


