# InlineResponse200415

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200415Items**](InlineResponse200415Items.md) | Wireless LAN controller interfaces packets statuses | [optional] 
**Meta** | Pointer to [**InlineResponse200240Meta**](InlineResponse200240Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200415

`func NewInlineResponse200415() *InlineResponse200415`

NewInlineResponse200415 instantiates a new InlineResponse200415 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200415WithDefaults

`func NewInlineResponse200415WithDefaults() *InlineResponse200415`

NewInlineResponse200415WithDefaults instantiates a new InlineResponse200415 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200415) GetItems() []InlineResponse200415Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200415) GetItemsOk() (*[]InlineResponse200415Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200415) SetItems(v []InlineResponse200415Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200415) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200415) GetMeta() InlineResponse200240Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200415) GetMetaOk() (*InlineResponse200240Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200415) SetMeta(v InlineResponse200240Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200415) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


