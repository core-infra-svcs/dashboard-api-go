# InlineResponse200460

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200460Items**](InlineResponse200460Items.md) | List of networks with SSID-level summaries of wireless clients impacted by connection failures. Results are sorted by impact count, with sort order configurable via the sortOrder parameter. | [optional] 
**Meta** | Pointer to [**InlineResponse200460Meta**](InlineResponse200460Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200460

`func NewInlineResponse200460() *InlineResponse200460`

NewInlineResponse200460 instantiates a new InlineResponse200460 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200460WithDefaults

`func NewInlineResponse200460WithDefaults() *InlineResponse200460`

NewInlineResponse200460WithDefaults instantiates a new InlineResponse200460 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200460) GetItems() []InlineResponse200460Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200460) GetItemsOk() (*[]InlineResponse200460Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200460) SetItems(v []InlineResponse200460Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200460) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200460) GetMeta() InlineResponse200460Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200460) GetMetaOk() (*InlineResponse200460Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200460) SetMeta(v InlineResponse200460Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200460) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


