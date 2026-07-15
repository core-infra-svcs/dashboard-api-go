# InlineResponse200414

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]InlineResponse200413Items**](InlineResponse200413Items.md) | Sentry Group Policies for the Organization keyed by the Network or Locale Id the Policy belongs to | [optional] 
**Meta** | Pointer to [**InlineResponse200275Meta**](InlineResponse200275Meta.md) |  | [optional] 

## Methods

### NewInlineResponse200414

`func NewInlineResponse200414() *InlineResponse200414`

NewInlineResponse200414 instantiates a new InlineResponse200414 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200414WithDefaults

`func NewInlineResponse200414WithDefaults() *InlineResponse200414`

NewInlineResponse200414WithDefaults instantiates a new InlineResponse200414 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InlineResponse200414) GetItems() []InlineResponse200413Items`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InlineResponse200414) GetItemsOk() (*[]InlineResponse200413Items, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InlineResponse200414) SetItems(v []InlineResponse200413Items)`

SetItems sets Items field to given value.

### HasItems

`func (o *InlineResponse200414) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *InlineResponse200414) GetMeta() InlineResponse200275Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse200414) GetMetaOk() (*InlineResponse200275Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse200414) SetMeta(v InlineResponse200275Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineResponse200414) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


