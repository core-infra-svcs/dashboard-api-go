# InlineResponse200448

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200448Items**](InlineResponse200448Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200263Meta**](InlineResponse200263Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200448

`func NewInlineResponse200448() *InlineResponse200448`

NewInlineResponse200448 instantiates a new InlineResponse200448 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200448WithDefaults

`func NewInlineResponse200448WithDefaults() *InlineResponse200448`

NewInlineResponse200448WithDefaults instantiates a new InlineResponse200448 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200448) GetItems() []InlineResponse200448Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200448) GetItemsOk() (*[]InlineResponse200448Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200448) SetItems(v []InlineResponse200448Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200448) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200448) GetMeta() InlineResponse200263Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200448) GetMetaOk() (*InlineResponse200263Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200448) SetMeta(v InlineResponse200263Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200448) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


