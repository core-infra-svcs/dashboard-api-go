# InlineResponse200433

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200432Items**](InlineResponse200432Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200279Meta**](InlineResponse200279Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200433

`func NewInlineResponse200433() *InlineResponse200433`

NewInlineResponse200433 instantiates a new InlineResponse200433 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200433WithDefaults

`func NewInlineResponse200433WithDefaults() *InlineResponse200433`

NewInlineResponse200433WithDefaults instantiates a new InlineResponse200433 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200433) GetItems() []InlineResponse200432Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200433) GetItemsOk() (*[]InlineResponse200432Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200433) SetItems(v []InlineResponse200432Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200433) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200433) GetMeta() InlineResponse200279Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200433) GetMetaOk() (*InlineResponse200279Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200433) SetMeta(v InlineResponse200279Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200433) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


