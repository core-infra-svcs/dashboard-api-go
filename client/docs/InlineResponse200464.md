# InlineResponse200464

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200464Items**](InlineResponse200464Items.md) | List of networks with SSID-level summaries of wireless clients impacted by connection failures. Results are sorted by impact count, with sort order configurable via the sortOrder parameter. | [optional] 
**Meta** | Pointer to [**InlineResponse200464Meta**](InlineResponse200464Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200464

`func NewInlineResponse200464() *InlineResponse200464`

NewInlineResponse200464 instantiates a new InlineResponse200464 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200464WithDefaults

`func NewInlineResponse200464WithDefaults() *InlineResponse200464`

NewInlineResponse200464WithDefaults instantiates a new InlineResponse200464 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200464) GetItems() []InlineResponse200464Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200464) GetItemsOk() (*[]InlineResponse200464Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200464) SetItems(v []InlineResponse200464Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200464) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200464) GetMeta() InlineResponse200464Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200464) GetMetaOk() (*InlineResponse200464Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200464) SetMeta(v InlineResponse200464Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200464) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


