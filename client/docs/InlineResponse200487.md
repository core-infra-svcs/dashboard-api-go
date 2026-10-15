# InlineResponse200487

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200487Items**](InlineResponse200487Items.md) | List of networks with SSID-level summaries of wireless clients impacted by connection failures. Results are sorted by impact count, with sort order configurable via the sortOrder parameter. | [optional] 
**Meta** | Pointer to [**InlineResponse200487Meta**](InlineResponse200487Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200487

`func NewInlineResponse200487() *InlineResponse200487`

NewInlineResponse200487 instantiates a new InlineResponse200487 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200487WithDefaults

`func NewInlineResponse200487WithDefaults() *InlineResponse200487`

NewInlineResponse200487WithDefaults instantiates a new InlineResponse200487 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200487) GetItems() []InlineResponse200487Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200487) GetItemsOk() (*[]InlineResponse200487Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200487) SetItems(v []InlineResponse200487Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200487) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200487) GetMeta() InlineResponse200487Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200487) GetMetaOk() (*InlineResponse200487Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200487) SetMeta(v InlineResponse200487Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200487) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


