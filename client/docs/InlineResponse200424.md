# InlineResponse200424

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200424Items**](InlineResponse200424Items.md) | Paginated list of mqtt settings by network ID. | [optional] 
**Meta** | Pointer to [**InlineResponse200250Meta**](InlineResponse200250Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200424

`func NewInlineResponse200424() *InlineResponse200424`

NewInlineResponse200424 instantiates a new InlineResponse200424 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200424WithDefaults

`func NewInlineResponse200424WithDefaults() *InlineResponse200424`

NewInlineResponse200424WithDefaults instantiates a new InlineResponse200424 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200424) GetItems() []InlineResponse200424Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200424) GetItemsOk() (*[]InlineResponse200424Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200424) SetItems(v []InlineResponse200424Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200424) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200424) GetMeta() InlineResponse200250Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200424) GetMetaOk() (*InlineResponse200250Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200424) SetMeta(v InlineResponse200250Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200424) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


