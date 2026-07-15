# InlineResponse200463

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200463Items**](InlineResponse200463Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200275Meta**](InlineResponse200275Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200463

`func NewInlineResponse200463() *InlineResponse200463`

NewInlineResponse200463 instantiates a new InlineResponse200463 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200463WithDefaults

`func NewInlineResponse200463WithDefaults() *InlineResponse200463`

NewInlineResponse200463WithDefaults instantiates a new InlineResponse200463 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200463) GetItems() []InlineResponse200463Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200463) GetItemsOk() (*[]InlineResponse200463Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200463) SetItems(v []InlineResponse200463Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200463) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200463) GetMeta() InlineResponse200275Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200463) GetMetaOk() (*InlineResponse200275Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200463) SetMeta(v InlineResponse200275Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200463) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


